import type { Meta, StoryObj } from "@storybook/vue3";

import { default as ArticleComponent } from "./ArticleComponent.vue";

const content = `

# H1 Shouldn't be Rendered
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam lectus nisi, aliquam a laoreet eu,
lobortis facilisis quam. Nulla facilisi. Nulla facilisi. Fusce sodales aliquam sem id sollicitudin.

## The Dataflow
Nam et nulla urna. Duis ultricies mauris eu lorem ultricies sodales. Maecenas at nibh ipsum. Etiam
eget nulla vulputate, euismod mi id, aliquam lacus. Pellentesque congue dolor in hendrerit dapibus.
Aliquam nec orci vitae erat convallis interdum sed eget nibh. Nam mi neque, tempor a nulla rhoncus,
faucibus placerat elit.

## Another Headline H2
### Another Headline H3
#### Another Headline H4
##### H5 Shouldn't be Rendered
###### H6 Shouldn't be Rendered

\`\`\`js
// The syntax highlighting should work:
const a = 42;

funciont sayHello() {
  console.log("hello world");
}
\`\`\`

Maecenas non tincidunt enim, et interdum ex. Curabitur molestie nulla est, quis faucibus lacus
convallis ac. Proin tempus ex nec aliquam lacinia. Pellentesque malesuada lacus eu neque imperdiet
auctor:

* This is a list
  * Sub-list point 1
  * Sub-list point 2
* Another point
`;

const meta: Meta<typeof ArticleComponent> = {
  title: "Components/ArticleComponent",
  component: ArticleComponent,

  render: (args) => ({
    components: { ArticleComponent },
    setup() {
      return args;
    },
    template: `
      <div style="max-width: 800px;">
        <article-component :config="config" />
      </div>
    `,
  }),
};

export default meta;

type Story = StoryObj<typeof ArticleComponent>;

export const Default: Story = {
  args: {
    config: {
      article: {
        id: "00000000-0000-0000-0000-000000000000",
        title: "Goodbye Product, Hello Dataflow!",
        type: "page",
        tagKeys: ["cpp", "sdl"],
        summary: "A website as a product is soulless. It is nothing else than a pile of text, pictures ... "
        + "However, the goal of website is making the data flow among the users, the business and the world. "
        + "It is the dataflow makes the website meaningful, thus alive, thus have a soul.",
        content,
        thumbnailUrl: "https://picsum.photos/1600/1200",
        createdAt: "2024-07-26T07:09:29.584+02:00",
        updatedAt: "2024-07-26T07:35:16.278+02:00",
      },
      tags: [
        { key: "cpp", name: "C++" },
        { key: "sdl", name: "SDL" },
        { key: "vue", name: "Vue" },
        { key: "js", name: "JavaScript" },
      ],
    },
  },
};
