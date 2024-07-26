<script setup lang="ts">
import type { Article } from "frontend-types/data/article";
import type { Tag } from "frontend-types/data/tag";

const props = defineProps<{
  headingLevel: number;
  tags: Tag[];
  articleBaseUrl: string;
  articles: Article[];
}>();
</script>

<template>
  <div class="articles">
    <div
      v-for="article of props.articles"
      :key="article.id"
      class="article"
    >
      <div>[显示图片（用img标签）]</div>
      <component :is="`h${props.headingLevel}`">
        [显示标题 {{ article.title }}]
      </component>
      <p>Written on <time>[显示时间 {{ article.createdAt }}]</time></p>
      <p>Tags: [显示Tags]</p>
      <p>[显示简介（summary）]</p>
      <a :href="`${articleBaseUrl}/${article.id}`">[Read More 按钮]</a>
    </div>
  </div>
</template>
