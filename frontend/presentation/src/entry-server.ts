import { renderToString } from "vue/server-renderer";
import { createApp } from "./app";

export type RenderResult = {
  language: string;
  title: string;
  head: string;
  content: string;
};

export async function vueSsrRender(): Promise<RenderResult> {
  const app = createApp();

  const ctx = {};
  const content = await renderToString(app, ctx);

  return {
    language: "en",
    title: "Hello World",
    head: "",
    content: content,
  };
}
