import type { Meta, StoryObj } from "@storybook/vue3";

import { default as ArticleHeaderComponent } from "./ArticleHeaderComponent.vue";

const meta: Meta<typeof ArticleHeaderComponent> = {
  title: "Components/ArticleHeaderComponent",
  component: ArticleHeaderComponent,
};

export default meta;

type Story = StoryObj<typeof ArticleHeaderComponent>;

export const Default: Story = {
  args: {
    config: {
      title: "digiTAO: Organize the Dataflow",
      createdAt: "2024-07-24T23:46:10.249Z",
      summary: "A website as a product is soulless. It is nothing else than a pile of text, pictures ... "
      + "However, the goal of website is making the data flow among the users, the business and the world. "
      + "It is the dataflow makes the website meaningful, thus alive, thus have a soul.",
    },
  },
};
