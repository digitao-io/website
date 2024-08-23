<script setup lang="ts">
import { computed, ref, watch } from "vue";
import { ResponseStatus, sendHttpRequest } from "frontend-resources";
import type { Article, ArticleIdentifier } from "frontend-resources";
import DetailsList from "@/components/Shared/DetailsList.vue";

const props = defineProps<{
  articleId: string;
}>();

const article = ref<Article | null>(null);

const details = computed(() => {
  if (!article.value) {
    return [];
  }

  return [
    { label: "Id", data: article.value.id },
    { label: "Title", data: article.value.title },
    { label: "Type", data: article.value.type },
    { label: "Summary", data: article.value.summary },
    { label: "Thumbnail Url", data: article.value.thumbnailUrl },
    { label: "Created At", data: article.value.createdAt },
    { label: "Updated At", data: article.value.updatedAt },
  ];
});

watch(() => props.articleId, async () => {
  const query = { id: props.articleId };
  const response = await sendHttpRequest<ArticleIdentifier, undefined, Article>("", "/data/article-get", query);
  if (response.status !== ResponseStatus.OK) {
    return;
  }

  article.value = response.data;
}, { immediate: true });
</script>

<template>
  <div class="article-details-list">
    <details-list
      :details="details"
    />

    <div class="article-details-extra-title">
      Details
    </div>
  </div>
</template>
