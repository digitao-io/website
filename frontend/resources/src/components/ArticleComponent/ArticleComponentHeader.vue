<script setup lang="ts">
import { computed } from "vue";
import { DateTime } from "luxon";
import type { Tag } from "frontend-types/data/tag";

const props = defineProps<{
  tags: Tag[];
  title: string;
  createdAt: string;
  tagKeys: string[];
  summary: string;
  thumbnailUrl: string;
}>();

const writtenOn = computed(() =>
  DateTime
    .fromISO(props.createdAt)
    .setZone("Europe/Berlin")
    .setLocale("us")
    .toLocaleString(DateTime.DATE_HUGE));

const tagText = computed(() =>
  props.tagKeys
    .map((tagKey) => props.tags.find((tag) => tag.key === tagKey)!.name)
    .join(", "));
</script>

<template>
  <div class="article-header">
    <h1 class="title">
      {{ props.title }}
    </h1>

    <p class="date">
      Written on <time>{{ writtenOn }}</time>
    </p>

    <p class="tags">
      Tags: {{ tagText }}
    </p>

    <img
      class="thumbnail"
      :src="thumbnailUrl"
    >

    <p class="summary">
      {{ props.summary }}
    </p>
    <hr>
  </div>
</template>

<style scoped>
.title {
  margin: 0;
  font-family: var(--font-raleway);
  font-size: var(--font-size-xxxl);
  line-height: var(--line-height-xxxl);
  font-weight: bold;
}

.date,
.tags {
  font-size: var(--font-size-s);
  line-height: var(--line-height-s);
  font-weight: lighter;
  font-style: italic;
}

.date {
  margin-top: 32px;
}

.tags {
  margin-top: 8px;
}

.thumbnail {
  display: block;
  margin-top: 36px;
  width: 100%;
  max-height: 450px;
  object-fit: cover;
}

.summary {
  margin-top: 36px;
  margin-bottom: 36px;
  color: var(--color-primary-s1);
  font-family: var(--font-open-sans);
  font-size: var(--font-size-m);
  line-height: var(--line-height-m);
  font-weight: lighter;
  font-style: italic;
}

@media (min-width: 500px) {
  .title {
    font-size: var(--font-size-max);
    line-height: var(--line-height-max);
  }
}
</style>
