import { readFile } from "node:fs/promises";
import express from "express";
import { createServer } from "vite";

import { readConfig } from "./src/server/config-reader";
import { fetchPages } from "./src/server/page-fetcher";

import { RenderFunction } from "./src/entry-server";
import { DynamicRouter } from "./src/server/dynamic-router";

import { PageDetailsResolver } from "./src/resolving/page-details-resolver";
import { articleListValueResolver } from "./src/resolving/resolvers/article-list-value-resolver";
import { articleValueResolver } from "./src/resolving/resolvers/article-value-resolver";
import { resourceValueResolver } from "./src/resolving/resolvers/resource-value-resolver";
import { tagValueResolver } from "./src/resolving/resolvers/tag-value-resolver";

(async () => {
  const config = readConfig();
  const app = express();
  const dynamicRouter = new DynamicRouter();
  const pageDetailsResolver = new PageDetailsResolver()
    .with(articleListValueResolver)
    .with(articleValueResolver)
    .with(resourceValueResolver)
    .with(tagValueResolver);

  const vite = await createServer({
    server: { middlewareMode: true },
    appType: "custom",
    base: config.pageBaseUrl,
  });

  app.use(vite.middlewares);

  app.use(
    "*",
    async (req, _, next) => {
      const url = req.originalUrl.replace(config.pageBaseUrl, "");

      const originalHtml = await readFile("./index.html", { encoding: "utf8" });
      const htmlTemplate = await vite.transformIndexHtml(url, originalHtml);

      const render: RenderFunction = (await vite.ssrLoadModule("./src/entry-server.ts")).render;

      const pages = await fetchPages(config);

      dynamicRouter.buildRoutes(config, pages, pageDetailsResolver, render, htmlTemplate);

      next();
    },
  );

  app.use("/", (req, res, next) => {
    dynamicRouter.getRouter()(req, res, next);
  });

  app.listen(config.port, () => {
    console.log(`Server started at http://localhost:${config.port}`);
  });
})();
