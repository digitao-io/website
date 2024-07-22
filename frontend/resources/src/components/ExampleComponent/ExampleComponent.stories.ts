import type { Meta, StoryObj } from "@storybook/vue3";

import { default as ExampleComponent } from "./ExampleComponent.vue";

const meta: Meta<typeof ExampleComponent> = {
  title: "Components/ExampleComponent",
  component: ExampleComponent,
};

export default meta;

type Story = StoryObj<typeof ExampleComponent>;

export const Default: Story = {};
