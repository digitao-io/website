<script setup lang="ts">
import { computed } from "vue";
import { DateTime } from "luxon";
import type { Tag } from "../../types/data/tag";

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

const hasTags = computed(() => props.tagKeys.length !== 0);
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

    <p
      v-if="hasTags"
      class="tags"
    >
      Tags: {{ tagText }}
    </p>

    <img
      class="thumbnail"
      crossorigin="anonymous"
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
  margin: 32px 0 0 0;
}

.tags {
  margin: 8px 0 0 0;
}

.thumbnail {
  display: block;
  margin: 36px -16px 0 -16px;
  width: calc(100% + 32px);
  max-height: 400px;
  object-fit: cover;
}

.summary {
  margin: 36px 0;
  font-size: var(--font-size-m);
  line-height: var(--line-height-m);
  font-style: italic;
}

@media (min-width: 500px) {
  .title {
    font-size: var(--font-size-max);
    line-height: var(--line-height-max);
  }

  .thumbnail {
    margin: 36px 0 0 0;
    width: 100%;
  }
}
</style>
