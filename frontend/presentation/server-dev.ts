import fs from "node:fs/promises";
import express from "express";
import { createServer } from "vite";

import type { Page } from "frontend-types/site/page";

import type { RenderResult } from "./src/entry-server";
import { sendHttpRequest } from "./src/common/http-client";
import { BackendResponseStatus } from "frontend-types/app/backend-response";
import { PageDetails } from "frontend-types/site/data-resolving";

const BASE_URL = "/";

(async () => {
  const app = express();

  const vite = await createServer({
    server: { middlewareMode: true },
    appType: "custom",
    base: BASE_URL,
  });

  app.use(vite.middlewares);

  app.use("*", async (req, res) => {
    try {
      const url = req.originalUrl.replace(BASE_URL, "");

      const originalHtml = await fs.readFile("./public/index.html", "utf-8");
      const template = await vite.transformIndexHtml(url, originalHtml);

      const pages = await sendHttpRequest<undefined, undefined, Page[]>("/site/page-list");

      if (pages.status !== BackendResponseStatus.OK) {
        throw Error("HTTP Request Broken");
      }

      const vueSsrRender: (page: PageDetails) => Promise<RenderResult>
        = (await vite.ssrLoadModule("./src/entry-server.ts")).vueSsrRender;
      const renderResult = await vueSsrRender(pages.data[0].details);

      const html = template
        .replace("$$PAGE_LANGUAGE$$", renderResult.language)
        .replace("$$PAGE_TITLE$$", renderResult.title)
        .replace("$$PAGE_HEAD$$", renderResult.head)
        .replace("$$PAGE_CONTENT$$", renderResult.content);

      res.status(200);
      res.set({ "Content-Type": "text/html" });
      res.send(html);
    } catch (e) {
      vite.ssrFixStacktrace(e as Error);

      console.log((e as Error).stack);

      res.status(500);
      res.end((e as Error).stack);
    }
  });

  app.listen(8080, () => {
    console.log("Server started at http://localhost:8080");
  });
})();
