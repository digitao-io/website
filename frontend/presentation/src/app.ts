import { createSSRApp } from "vue";
import App from "./App.vue";

import {

  // Components:
  MarkdownComponent,

  // Templates:
  DefaultTemplate,
} from "frontend-resources";
import "frontend-resources/index.css";

export function createApp() {
  const app = createSSRApp(App)

    // Components:
    .component("markdown-component", MarkdownComponent)

    // Templates:
    .component("default-template", DefaultTemplate);
  return app;
}
