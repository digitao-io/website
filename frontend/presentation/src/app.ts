import { createSSRApp } from "vue";
import App from "./App.vue";

import { ExampleComponent, ExampleTemplate } from "frontend-resources";
import "frontend-resources/index.css";

export function createApp() {
  const app = createSSRApp(App)
    .component("example-component", ExampleComponent)
    .component("example-template", ExampleTemplate);
  return app;
}
