import { renderToString } from "vue/server-renderer";

import { createApp } from "./app";
import type { ResolvedPageDetails } from "./resolving/page-details-resolver";

export type RenderResult = {
  language: string;
  title: string;
  head: string;
  content: string;
};

export type RenderFunction = (pageDetails: ResolvedPageDetails) => Promise<RenderResult>;

export async function render(pageDetails: ResolvedPageDetails): Promise<RenderResult> {
  const app = createApp();

  const ctx = { pageDetails };
  const content = await renderToString(app, ctx);

  return {
    language: pageDetails.language,
    title: pageDetails.title,
    head: `<script>window.context = ${JSON.stringify(ctx)}</script>`,
    content,
  };
}
