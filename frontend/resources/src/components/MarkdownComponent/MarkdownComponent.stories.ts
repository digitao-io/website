import type { Meta, StoryObj } from "@storybook/vue3";

import { default as MarkdownComponent } from "./MarkdownComponent.vue";

const meta: Meta<typeof MarkdownComponent> = {
  title: "Components/MarkdownComponent",
  component: MarkdownComponent,
};

export default meta;

type Story = StoryObj<typeof MarkdownComponent>;

const markdown = `
# This should be h1
This is the first paragraph

## This should be h2
This is the second paragraph

## Another headline h2

\`\`\`js
// The syntax highlighting should work:
const a = 42;
\`\`\`

Some random text:
* This is a list
  * Sub-list point 1
  * Sub-list point 2
* Another point
`;

export const Default: Story = {
  args: {
    config: {
      markdown,
    },
  },
};
