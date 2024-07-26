<script setup lang="ts">
import type { Article } from "frontend-types/data/article";
import type { Tag } from "frontend-types/data/tag";
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

const tagsText = function (keys: string[]) {
  let tagText = "";
  for (let i = 0; i < keys.length; i++) {
    const key = keys[i];

    for (let j = 0; j < props.tags.length; j++) {
      if (key === props.tags[j].key) {
        tagText += props.tags[j].name;
        break;
      }
    }

    if (i < keys.length - 1) {
      tagText += " / ";
    }
  }
  return tagText;
};

</script>

<template>
  <div class="articles">
    <div
      v-for="article of props.articles"
      :key="article.id"
      class="article"
    >
      <div>
        <img
          calss="thumbnail"
          :src="article.thumbnailUrl"
        >
      </div>
      <component
        :is="`h${props.headingLevel}`"
        class="title"
      >
        {{ article.title }}
      </component>
      <div class="blog-content">
        <div class="written-and-tags">
          <p>Written on <time>{{ formatWrittenOn(article.createdAt) }}</time></p>
          <p>Tags: {{ tagsText(article.tagKeys) }}</p>
        </div>

        <p class="summary">
          {{ article.summary }}
        </p>
      </div>
      <a :href="`${articleBaseUrl}/${article.id}`">[Read More 按钮]</a>
    </div>
  </div>
</template>

<style scoped>
.thumbnail {
  display: block;
  width: 100%;
  object-fit: fill;
}

.title {
  font-family: var(--font-raleway);
  font-size: var(--font-size-xxl);
  line-height: var(--line-height-xxl);
  color: var(--color-primary-s1);
  font-weight: bold;
}

.blog-content {
  font-family: var(--font-open-sans);
  font-weight: lighter;
  color: var(--color-primary-s1);
}

.written-and-tags {
  font-size: var(--font-size-s);
  line-height: var(--line-height-s);
}
.summary {
  font-size: 16px;
  line-height: auto;
}
</style>
