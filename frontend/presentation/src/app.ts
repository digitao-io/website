import { createSSRApp } from "vue";
import App from "./App.vue";

import "frontend-resources/index.css";
import {

  // Components:
  ArticleComponent,
  ArticleListComponent,

  // Templates:
  DefaultTemplate,
} from "frontend-resources";

export function createApp() {
  const app = createSSRApp(App)

    // Components:
    .component("article-component", ArticleComponent)
    .component("article-list-component", ArticleListComponent)

    // Templates:
    .component("default-template", DefaultTemplate);
  return app;
}
