import { readFileSync } from "node:fs";
import express from "express";
import sirv from "sirv";
import compression from "compression";
import helmet from "helmet";

import { ResponseStatus } from "frontend-resources";

import { render } from "./src/entry-server";
import { readConfig } from "./src/server/config-reader";
import { fetchPages } from "./src/server/page-fetcher";
import { DynamicRouter } from "./src/server/dynamic-router";
import { PageDetailsResolver } from "./src/resolving/page-details-resolver";

import { articleListValueResolver } from "./src/resolving/resolvers/article-list-value-resolver";
import { articleValueResolver } from "./src/resolving/resolvers/article-value-resolver";
import { resourceValueResolver } from "./src/resolving/resolvers/resource-value-resolver";
import { tagValueResolver } from "./src/resolving/resolvers/tag-value-resolver";

(async () => {
  const config = readConfig();
  const app = express();

  const htmlTemplate = readFileSync("./dist/client/index.html", { encoding: "utf8" });

  const pageDetailsResolver = new PageDetailsResolver()
    .with(articleListValueResolver)
    .with(articleValueResolver)
    .with(resourceValueResolver)
    .with(tagValueResolver);

  const dynamicRouter = new DynamicRouter();

  app.use(helmet({
    contentSecurityPolicy: {
      directives: {
        "script-src": ["'self'", "'unsafe-inline'"],
        "img-src": ["'self'", "data:", "https:"],
      },
    },
  }));

  app.use(compression());
  app.use("/", sirv("./dist/client", { extensions: [] }));

  app.post("/maintenance/cache-clear", (_, res) => {
    pageDetailsResolver.clearCache();

    res.status(200);
    res.json({ status: ResponseStatus.OK });
  });

  app.post("/maintenance/page-fetch", async (_, res) => {
    const pages = await fetchPages(config);

    dynamicRouter.buildRoutes(config, pages, pageDetailsResolver, render, htmlTemplate);

    res.status(200);
    res.json({ status: ResponseStatus.OK });
  });

  app.use("/", (req, res, next) => {
    dynamicRouter.getRouter()(req, res, next);
  });

  app.listen(config.port, () => {
    console.log(`Server started at http://localhost:${config.port}`);
  });
})();
