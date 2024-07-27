import type { Meta, StoryObj } from "@storybook/vue3";
import type { Article } from "frontend-types/data/article";
import { default as ArticleListComponent } from "./ArticleListComponent.vue";

const generateArticles = (): Article[] => [
  {
    id: `00000000-0000-0000-0000-00000000${Math.floor(1000 + Math.random() * 9000)}`,
    title: "Goodbye Product, Hello Dataflow!",
    tagKeys: ["cpp", "sdl"],
    summary: "A website as a product is soulless. It is nothing else than a pile of text, pictures ... "
    + "However, the goal of website is making the data flow among the users, the business and the world. "
    + "It is the dataflow makes the website meaningful, thus alive, thus have a soul.",
    content: "Not important here",
    thumbnailUrl: `https://picsum.photos/640/480?random=${Math.floor(100 + Math.random() * 900)}`,
    createdAt: "2024-07-26T07:09:29.584+02:00",
    updatedAt: "2024-07-26T07:35:16.278+02:00",
  },
  {
    id: `00000000-0000-0000-0000-00000000${Math.floor(1000 + Math.random() * 9000)}`,
    title: "Another Title of a Blog Article",
    tagKeys: ["js", "vue"],
    summary: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam lectus nisi, aliquam a "
    + "laoreet eu, lobortis facilisis quam. Nulla facilisi. Nulla facilisi. Fusce sodales aliquam sem id "
    + "sollicitudin. Sed nulla dolor, egestas condimentum nulla eu, dapibus dictum lectus.",
    content: "Not important here",
    thumbnailUrl: `https://picsum.photos/640/480?random=${Math.floor(100 + Math.random() * 900)}`,
    createdAt: "2024-07-25T07:09:29.584+02:00",
    updatedAt: "2024-07-25T07:35:16.278+02:00",
  },
  {
    id: `00000000-0000-0000-0000-00000000${Math.floor(1000 + Math.random() * 9000)}`,
    title: "Another Title of a Blog Article",
    tagKeys: ["js", "vue"],
    summary: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam lectus nisi, aliquam a "
    + "laoreet eu, lobortis facilisis quam. Nulla facilisi. Nulla facilisi. Fusce sodales aliquam sem id "
    + "sollicitudin. Sed nulla dolor, egestas condimentum nulla eu, dapibus dictum lectus.",
    content: "Not important here",
    thumbnailUrl: `https://picsum.photos/640/480?random=${Math.floor(100 + Math.random() * 900)}`,
    createdAt: "2024-07-25T07:09:29.584+02:00",
    updatedAt: "2024-07-25T07:35:16.278+02:00",
  },
];

const meta: Meta<typeof ArticleListComponent> = {
  title: "Components/ArticleListComponent",
  component: ArticleListComponent,

  parameters: {
    mockData: [
      {
        url: "http://localhost:6006/data/article-list?*",
        method: "POST",
        status: 200,
        response: () => ({
          status: "OK",
          data: generateArticles(),
        }),
      },
    ],
  },

  render: (args) => ({
    components: { ArticleListComponent },
    setup() {
      return args;
    },
    template: `
      <div style="max-width: 800px;">
        <article-list-component :config="config" />
      </div>
    `,
  }),
};

export default meta;

type Story = StoryObj<typeof ArticleListComponent>;

export const Default: Story = {
  args: {
    config: {
      showSearch: true,
      headingLevel: 2,
      numberPerLoad: 3,
      articleBaseUrl: "https://example.org",
      tags: [
        { key: "cpp", name: "C++" },
        { key: "sdl", name: "SDL" },
        { key: "vue", name: "Vue" },
        { key: "js", name: "JavaScript" },
      ],
      articles: generateArticles(),
    },
  },
};
