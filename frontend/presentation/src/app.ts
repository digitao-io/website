import { createSSRApp } from "vue";
import App from "./App.vue";

import {

  // Components:
  ArticleComponent,
  ArticleListComponent,

  // Templates:
  DefaultTemplate,
} from "frontend-resources";
import "frontend-resources/index.css";

export function createApp() {
  const app = createSSRApp(App)

    // Components:
    .component("article-component", ArticleComponent)
    .component("article-list-component", ArticleListComponent)

    // Templates:
    .component("default-template", DefaultTemplate);
  return app;
}
