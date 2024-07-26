<script setup lang="ts">
import { ref } from "vue";
import type { Article } from "frontend-types/data/article";
import type { Tag } from "frontend-types/data/tag";
import ArticleListComponentSearch from "./ArticleListComponentSearch.vue";
import ArticleListComponentResult from "./ArticleListComponentResult.vue";

type ArticleListComponentConfig = {
  showSearch: boolean;
  headingLevel: number;
  tags: Tag[];
  articleBaseUrl: string;
  articles: Article[];
};

const props = defineProps<{
  config: ArticleListComponentConfig;
}>();

const articles = ref(props.config.articles);
</script>

<template>
  <div class="article-list">
    <article-list-component-search
      v-if="props.config.showSearch"
    />

    <article-list-component-result
      :heading-level="props.config.headingLevel"
      :tags="props.config.tags"
      :article-base-url="props.config.articleBaseUrl"
      :articles="articles"
    />
  </div>
</template>
