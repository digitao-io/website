import type { Preview } from "@storybook/vue3";

import "./fonts.css";
import "../src/main.css";

(window as any).context = {
  pageContext: {
    apiBaseUrl: "http://localhost:6006/",
  }
};

const preview: Preview = {
  parameters: {
    mockAddonConfigs: {
      globalMockData: [],
      ignoreQueryParams: true,
    },
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i,
      },
    },
  },
};

export default preview;
