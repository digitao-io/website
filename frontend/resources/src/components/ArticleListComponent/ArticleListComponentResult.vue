<script setup lang="ts">
import type { Article } from "../../types/data/article";
import type { Tag } from "../../types/data/tag";
import { DateTime } from "luxon";

const props = defineProps<{
  headingLevel: number;
  tags: Tag[];
  articleBaseUrl: string;
  articles: Article[];
}>();

const formatWrittenOn = (createdAt: string) =>
  DateTime
    .fromISO(createdAt)
    .setZone("Europe/Berlin")
    .setLocale("us")
    .toLocaleString(DateTime.DATE_HUGE);

const generateTagsText = (keys: string[]): string => keys
  .map((key) => props.tags.find((t) => t.key === key)!.name)
  .sort((tag1, tag2) => tag1.localeCompare(tag2))
  .join(", ");

</script>

<template>
  <div class="articles">
    <div
      v-for="article of props.articles"
      :key="article.id"
      class="article"
    >
      <img
        class="thumbnail"
        :src="article.thumbnailUrl"
      >

      <component
        :is="`h${props.headingLevel}`"
        class="title"
      >
        {{ article.title }}
      </component>

      <p class="date">
        Written on <time>{{ formatWrittenOn(article.createdAt) }}</time>
      </p>
      <p class="tags">
        Tags: {{ generateTagsText(article.tagKeys) }}
      </p>

      <p class="summary">
        {{ article.summary }}
      </p>

      <a
        class="read-more-button"
        :href="`${articleBaseUrl}/${article.id}`"
      >READ MORE</a>
    </div>
  </div>
</template>

<style scoped>
.article {
  margin: 64px 0;
}

.thumbnail {
  display: block;
  width: 100%;
  max-height: 450px;
  object-fit: cover;
}

.title {
  margin: 16px 0 0 0;
  font-family: var(--font-raleway);
  font-size: var(--font-size-xl);
  line-height: var(--line-height-xl);
}

.date,
.tags {
  font-size: var(--font-size-s);
  line-height: var(--line-height-s);
  font-weight: lighter;
}

.date {
  margin: 16px 0 0 0;
}

.tags {
  margin: 8px 0 0 0;
}

.summary {
  margin: 16px 0 0 0;
  font-size: var(--font-size-m);
  line-height: var(--line-height-m);
  font-weight: lighter;
}

.read-more-button {
  display: inline-flex;
  align-items: center;
  margin: 16px 0 0 0;
  border: 1px solid var(--color-primary);
  padding: 0 24px;
  height: 36px;
  color: var(--color-secondary);
  background-color: var(--color-primary);
  font-family: var(--font-open-sans);
  font-weight: bold;
  font-size: var(--font-size-s);
  line-height: var(--line-height-s);
  text-decoration: none;
}
.read-more-button:hover {
  border-color: var(--color-primary-t1);
  background-color: var(--color-primary-t1);
}

@media (min-width: 500px) {
  .title {
    font-size: var(--font-size-xxl);
    line-height: var(--line-height-xxl);
  }
}
</style>
