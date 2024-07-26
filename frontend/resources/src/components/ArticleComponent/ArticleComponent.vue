<script setup lang="ts">
import type { Article } from "frontend-types/data/article";
import type { Tag } from "frontend-types/data/tag";
import ArticleComponentHeader from "./ArticleComponentHeader.vue";
import ArticleComponentContent from "./ArticleComponentContent.vue";

type ArticleComponentConfig = {
  tags: Tag[];
  article: Article;
};

const props = defineProps<{
  config: ArticleComponentConfig;
}>();
</script>

<template>
  <div class="article-list">
    <article-component-header
      :title="props.config.article.title"
      :created-at="props.config.article.createdAt"
      :tag-keys="props.config.article.tagKeys"
      :tags="props.config.tags"
      :summary="props.config.article.summary"
      :thumbnail-url="props.config.article.thumbnailUrl"
    />

    <article-component-content
      :content="props.config.article.content"
    />
  </div>
</template>
