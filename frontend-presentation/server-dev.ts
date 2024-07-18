import fs from "node:fs/promises";
import express from "express";
import { createServer } from "vite";

const BASE_URL="/";

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
  
      const originalHtml = await fs.readFile("./index.html", "utf-8");
      const template = await vite.transformIndexHtml(url, originalHtml);
  
      const render: () => string = (await vite.ssrLoadModule("./src/entry-server.ts")).render;
      const renderResult = render();

      const html = template
        .replace("<!--PAGE_TITLE-->", "Hello SSR!")
        .replace("<!--PAGE_HEAD-->", "")
        .replace("<!--PAGE_CONTENT-->", renderResult);
  
      res.status(200);
      res.set({ "Content-Type": "text/html" });
      res.send(html);
    } catch (e) {
      vite.ssrFixStacktrace(e);
      console.log(e.stack);
      res.status(500);
      res.end(e.stack);
    }
  });
  
  app.listen(8080, () => {
    console.log("Server started at http://localhost:8080");
  });  
})();
