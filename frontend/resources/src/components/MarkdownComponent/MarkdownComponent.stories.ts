import type { Meta, StoryObj } from "@storybook/vue3";

import { default as MarkdownComponent } from "./MarkdownComponent.vue";

const meta: Meta<typeof MarkdownComponent> = {
  title: "Components/MarkdownComponent",
  component: MarkdownComponent,
};

export default meta;

type Story = StoryObj<typeof MarkdownComponent>;

const markdown = `
# This should be h2
This is the first paragraph

## This should be h3
This is the second paragraph

## Another headline

\`\`\`js
// The syntax highlighting should work:
const a = 42;
\`\`\`

Some random text:
* This is a list
* Another point
`;

export const Default: Story = {
  args: {
    config: {
      markdown,
    },
  },
};
