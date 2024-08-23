<script setup lang="ts">
import { onBeforeMount, ref } from "vue";
import { ResponseStatus, sendHttpRequest } from "frontend-resources";
import type { Article, ArticleSearchQuery } from "frontend-resources";
import PageLayout from "@/components/Shared/PageLayout.vue";
import DataTable from "@/components/Shared/DataTable.vue";
import EditorSwitch from "@/components/Shared/EditorSwitch.vue";
import ArticleDetailsList from "@/components/Articles/ArticleDetailsList.vue";

const articles = ref<Article[]>([]);
const selectedArticleIds = ref<string[]>([]);

onBeforeMount(() => {
  loadData();
});

const loadData = async () => {
  const query: ArticleSearchQuery = {
    sort: "createdAt",
    order: "DESC",
    take: 100,
    skip: 0,
  };

  const response = await sendHttpRequest<ArticleSearchQuery, undefined, Article[]>("", "/data/article-list", query);
  if (response.status !== ResponseStatus.OK) {
    return;
  }

  articles.value = response.data;
};

const onSelect = (selected: string[]) => {
  selectedArticleIds.value = selected;
};
</script>

<template>
  <page-layout
    class="articles"
    title="Articles"
  >
    <template #table>
      <data-table
        :columns="[
          { label: 'Id', dataExtractor: (article) => article.id, align: 'left', isKey: true },
          { label: 'Title', dataExtractor: (article) => article.title, align: 'left' },
          { label: 'Created At', dataExtractor: (article) => article.createdAt, align: 'left' },
          { label: 'Updated At', dataExtractor: (article) => article.updatedAt, align: 'left' },
        ]"
        :data="articles"
        :selected="selectedArticleIds"
        @select="onSelect"
      />
    </template>
    <template #editor>
      <editor-switch
        :editors="[
          { name: 'details', label: 'Info', type: 'single', default: true },
        ]"
        :selected="selectedArticleIds"
        @select="onSelect"
      >
        <template #details>
          <article-details-list
            :article-id="selectedArticleIds[0]"
          />
        </template>
      </editor-switch>
    </template>
  </page-layout>
</template>

<style scoped>
</style>
