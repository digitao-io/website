import fs from "node:fs/promises";
import express from "express";
import { createServer } from "vite";

import type { Page } from "frontend-resources";
import { ResponseStatus } from "frontend-resources";
import { sendHttpRequest } from "frontend-resources";

import { RenderFunction } from "./src/entry-server";
import { DynamicRouter } from "./src/server/dynamic-router";
import { PageDetailsResolver } from "./src/resolving/page-details-resolver";
import { dataValueResolver } from "./src/resolving/resolvers/data-value-resolver";

const BASE_URL = "/";

(async () => {
  const app = express();
  const dynamicRouter = new DynamicRouter();
  const pageDetailsResolver = new PageDetailsResolver()
    .with(dataValueResolver);

  const vite = await createServer({
    server: { middlewareMode: true },
    appType: "custom",
    base: BASE_URL,
  });

  app.use(vite.middlewares);

  app.use(
    "*",
    async (req, _, next) => {
      const url = req.originalUrl.replace(BASE_URL, "");

      const originalHtml = await fs.readFile("./public/index.html", "utf-8");
      const htmlTemplate = await vite.transformIndexHtml(url, originalHtml);

      const pages = await sendHttpRequest<undefined, undefined, Page[]>("http://localhost:3000", "/site/page-list");

      if (pages.status !== ResponseStatus.OK) {
        throw Error("Cannot fetch pages for some unknown reason");
      }

      const render: RenderFunction = (await vite.ssrLoadModule("./src/entry-server.ts")).render;

      dynamicRouter.buildRoutes(pages.data, pageDetailsResolver, render, htmlTemplate);

      next();
    },
  );

  app.use("/", (req, res, next) => {
    dynamicRouter.getRouter()(req, res, next);
  });

  app.listen(8080, () => {
    console.log("Server started at http://localhost:8080");
  });
})();
