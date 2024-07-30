import { ResponseStatus, sendHttpRequest, type Page } from "frontend-resources";
import type { Config } from "./config-reader";

export async function fetchPages(config: Config): Promise<Page[]> {
  try {
    const pages = await sendHttpRequest<undefined, undefined, Page[]>(
      config.apiBaseUrl,
      "/site/page-list",
    );

    if (pages.status !== ResponseStatus.OK) {
      console.log(JSON.stringify(pages));
      return [];
    }

    return pages.data;
  } catch (e) {
    console.log(e);
    return [];
  }
}
