
import express from "express";
import type { Router } from "express";
import type { Page } from "frontend-resources";
import type { RenderFunction } from "../entry-server";
import type { PageDetailsResolver } from "src/resolving/page-details-resolver";
import type { Config } from "./config-reader";

export class DynamicRouter {
  private router: Router;

  public constructor() {
    this.router = express.Router();
  }

  public buildRoutes(
    config: Config,
    pages: Page[],
    resolver: PageDetailsResolver,
    render: RenderFunction,
    htmlTemplate: string,
  ) {
    this.router = express.Router();

    for (const page of pages) {
      const { urlPattern, details } = page;

      this.router.get(urlPattern, async (req, res) => {
        try {
          const resolvedPageDetails = await resolver.resolve({
            apiBaseUrl: config.apiBaseUrl,
            pageBaseUrl: config.pageBaseUrl,
            pageUrlPath: req.path,
            pageUrlParams: req.params,
            pageUrlQueries: req.query as Record<string, string | string[]>,
          }, details);

          const renderResult = await render(resolvedPageDetails);

          const html = htmlTemplate
            .replace("$$PAGE_LANGUAGE$$", renderResult.language)
            .replace("$$PAGE_TITLE$$", renderResult.title)
            .replace("$$PAGE_HEAD$$", renderResult.head)
            .replace("$$PAGE_CONTENT$$", renderResult.content);

          res.status(200);
          res.set({ "Content-Type": "text/html" });
          res.send(html);
        } catch (e) {
          console.log((e as Error).stack);
          res.status(500);
          res.end((e as Error).stack);
        }
      });
    }
  }

  public getRouter(): Router {
    return this.router;
  }
}
