import type { Meta, StoryObj } from "@storybook/vue3";

import { default as ArticleListComponent } from "./ArticleListComponent.vue";

const meta: Meta<typeof ArticleListComponent> = {
  title: "Components/ArticleListComponent",
  component: ArticleListComponent,
};

export default meta;

type Story = StoryObj<typeof ArticleListComponent>;

export const Default: Story = {
  args: {
    config: {
      showSearch: true,
      headingLevel: 2,
      articleBaseUrl: "https://example.org",
      tags: [
        { key: "cpp", name: "C++" },
        { key: "sdl", name: "SDL" },
        { key: "vue", name: "Vue" },
        { key: "js", name: "JavaScript" },
      ],
      articles: [
        {
          id: "00000000-0000-0000-0000-000000000000",
          title: "Goodbye Product, Hello Dataflow!",
          tagKeys: ["cpp", "sdl"],
          summary: "A website as a product is soulless. It is nothing else than a pile of text, pictures ... "
          + "However, the goal of website is making the data flow among the users, the business and the world. "
          + "It is the dataflow makes the website meaningful, thus alive, thus have a soul.",
          content: "Not important here",
          thumbnailUrl: "https://picsum.photos/640/480",
          createdAt: "2024-07-26T07:09:29.584+02:00",
          updatedAt: "2024-07-26T07:35:16.278+02:00",
        },
        {
          id: "00000000-0000-0000-0000-000000000001",
          title: "Another Title of a Blog Article",
          tagKeys: ["js", "vue"],
          summary: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam lectus nisi, aliquam a "
          + "laoreet eu, lobortis facilisis quam. Nulla facilisi. Nulla facilisi. Fusce sodales aliquam sem id "
          + "sollicitudin. Sed nulla dolor, egestas condimentum nulla eu, dapibus dictum lectus.",
          content: "Not important here",
          thumbnailUrl: "https://picsum.photos/640/480",
          createdAt: "2024-07-25T07:09:29.584+02:00",
          updatedAt: "2024-07-25T07:35:16.278+02:00",
        },
      ],
    },
  },
};
