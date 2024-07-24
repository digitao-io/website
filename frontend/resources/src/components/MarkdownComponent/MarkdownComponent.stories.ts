import type { Meta, StoryObj } from "@storybook/vue3";

import { default as MarkdownComponent } from "./MarkdownComponent.vue";

const meta: Meta<typeof MarkdownComponent> = {
  title: "Components/MarkdownComponent",
  component: MarkdownComponent,
};

export default meta;

type Story = StoryObj<typeof MarkdownComponent>;

const markdown = `
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

export const Default: Story = {
  args: {
    config: {
      markdown,
    },
  },
};
