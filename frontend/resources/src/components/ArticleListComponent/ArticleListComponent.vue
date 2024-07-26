<script setup lang="ts">
import { ref } from "vue";
import type { Article, ArticleSearchQuery } from "frontend-types/data/article";
import type { Tag } from "frontend-types/data/tag";
import { BackendResponseStatus } from "frontend-types/app/backend-response";
import { sendHttpRequest } from "../../services/http-client";
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
const currentSearchQuery = ref({
  q: "",
  tag: [],
  sort: "createdAt",
  order: "DESC",
  take: 3,
  skip: 0,
} as ArticleSearchQuery);

const searchArticles = async (searchQuery: { q: string; tagKeys: string[] }) => {
  articles.value = [];

  const urlQuery: ArticleSearchQuery = {
    q: searchQuery.q,
    tag: searchQuery.tagKeys,
    sort: searchQuery.q ? "score" : "createdAt",
    order: "DESC",
    take: 3,
    skip: 0,
  };

  currentSearchQuery.value = urlQuery;

  const responseBody = await sendHttpRequest<ArticleSearchQuery, undefined, Article[]>(
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    (window as any).context.pageContext.apiBaseUrl,
    "/data/article-list",
    urlQuery,
  );

  if (responseBody.status !== BackendResponseStatus.OK) {
    return;
  }

  articles.value = responseBody.data;
};

const loadMoreArticles = async () => {
  const urlQuery: ArticleSearchQuery = {
    ...currentSearchQuery.value,
    skip: articles.value.length,
  };

  const responseBody = await sendHttpRequest<ArticleSearchQuery, undefined, Article[]>(
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    (window as any).context.pageContext.apiBaseUrl,
    "/data/article-list",
    urlQuery,
  );

  if (responseBody.status !== BackendResponseStatus.OK) {
    return;
  }

  articles.value = responseBody.data;
};
</script>

<template>
  <div class="article-list">
    <article-list-component-search
      v-if="props.config.showSearch"
      :tags="props.config.tags"
      @search="searchArticles"
    />

    <article-list-component-result
      :heading-level="props.config.headingLevel"
      :tags="props.config.tags"
      :article-base-url="props.config.articleBaseUrl"
      :articles="articles"
    />
  </div>
</template>
