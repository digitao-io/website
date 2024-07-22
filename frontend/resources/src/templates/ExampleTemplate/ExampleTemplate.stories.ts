import type { Meta, StoryObj } from "@storybook/vue3";

import { default as ExampleTemplate } from "./ExampleTemplate.vue";

const meta: Meta<typeof ExampleTemplate> = {
  title: "Templates/ExampleTemplate",
  component: ExampleTemplate,

  render: () => ({
    components: { ExampleTemplate },
    template: `
      <example-template>Example Slot Content</example-template>
    `,
  }),
};

export default meta;

type Story = StoryObj<typeof ExampleTemplate>;

export const Default: Story = {};
