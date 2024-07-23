import type { Meta, StoryObj } from "@storybook/vue3";

import { default as DefaultTemplate } from "./DefaultTemplate.vue";

const meta: Meta<typeof DefaultTemplate> = {
  title: "Templates/DefaultTemplate",
  component: DefaultTemplate,

  render: (args) => ({
    components: { DefaultTemplate },
    setup() {
      return args;
    },
    template: `
      <default-template :config="config">
        <template v-slot:["main"]>
          <h1>Hello World!</h1>
          <p>Example content of the slot.</p>
        </template>
      </default-template>
    `,
  }),
};

export default meta;

type Story = StoryObj<typeof DefaultTemplate>;

export const Default: Story = {
  args: {
    config: {
      primaryMenuEntries: [
        { label: "Home", url: "http://example.org" },
        { label: "Blog", url: "http://example.org" },
        { label: "Profile", url: "http://example.org" },
        { label: "Showcase", url: "http://example.org" },
        { label: "Contact", url: "http://example.org" },
      ],

      secondaryMenuEntries: [
        { label: "Legal Notice", url: "http://example.org" },
        { label: "Data Protection", url: "http://example.org" },
      ],
    }
  },
};
