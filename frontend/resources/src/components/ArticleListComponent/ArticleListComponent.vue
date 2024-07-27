<script setup lang="ts">
import { ref } from "vue";
import { ResponseStatus } from "../../types/app/response";
import type { Article, ArticleSearchQuery } from "../../types/data/article";
import type { Tag } from "../../types/data/tag";
import { sendHttpRequest } from "../../services/http-client";
import ArticleListComponentSearch from "./ArticleListComponentSearch.vue";
import ArticleListComponentResult from "./ArticleListComponentResult.vue";

type ArticleListComponentConfig = {
  showSearch: boolean;
  headingLevel: number;
  numberPerLoad: number;
  tags: Tag[];
  articleBaseUrl: string;
  articles: Article[];
};

const props = defineProps<{
  config: ArticleListComponentConfig;
}>();

const articles = ref(props.config.articles);
const currentSearchQuery = ref({
  type: "blog",
  q: "",
  tag: [],
  sort: "createdAt",
  order: "DESC",
  take: props.config.numberPerLoad,
  skip: 0,
} as ArticleSearchQuery);
const hasMore = ref(true);

const searchArticles = async (searchQuery: { q: string; tagKeys: string[] }) => {
  articles.value = [];

  const urlQuery: ArticleSearchQuery = {
    type: "blog",
    q: searchQuery.q,
    tag: searchQuery.tagKeys,
    sort: searchQuery.q ? "score" : "createdAt",
    order: "DESC",
    take: props.config.numberPerLoad,
    skip: 0,
  };

  currentSearchQuery.value = urlQuery;

  const responseBody = await sendHttpRequest<ArticleSearchQuery, undefined, Article[]>(
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    (window as any).context.pageContext.apiBaseUrl,
    "/data/article-list",
    urlQuery,
  );

  if (responseBody.status !== ResponseStatus.OK) {
    return;
  }

  articles.value = responseBody.data;

  if (responseBody.data.length < props.config.numberPerLoad) {
    hasMore.value = false;
  }
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

  if (responseBody.status !== ResponseStatus.OK) {
    return;
  }

  articles.value = articles.value.concat(responseBody.data);

  if (responseBody.data.length < props.config.numberPerLoad) {
    hasMore.value = false;
  }
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

    <div class="footer-area">
      <button
        class="load-more-button"
        @click="loadMoreArticles()"
      >
        LOAD MORE
      </button>
    </div>
  </div>
</template>

<style scoped>
.footer-area {
  display: flex;
  justify-content: center;
}

.load-more-button {
  box-sizing: content-box;
  border: 1px solid var(--color-primary);
  padding: 0 24px;
  height: 36px;
  color: var(--color-secondary);
  background-color: var(--color-primary);
  font-family: var(--font-open-sans);
  font-weight: bold;
  font-size: var(--font-size-s);
  line-height: var(--line-height-s);
  cursor: pointer;
}
.load-more-button:hover {
  border-color: var(--color-primary-t1);
  background-color: var(--color-primary-t1);
}
</style>
