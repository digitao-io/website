import { renderToString } from "vue/server-renderer";

import type { Page } from "frontend-types/site/page";

import { createApp } from "./app";

export type RenderResult = {
  language: string;
  title: string;
  head: string;
  content: string;
};

export async function vueSsrRender(page: Page): Promise<RenderResult> {
  const app = createApp();

  const ctx = { page };
  const content = await renderToString(app, ctx);

  return {
    language: "en",
    title: "Hello World",
    head: "",
    content: content,
  };
}
