<script setup lang="ts">
import { computed } from "vue";
import { DateTime } from "luxon";

export type ArticleHeaderComponentConfig = {
  title: string;
  createdAt: string;
  summary: string;
};

const props = defineProps<{
  config: ArticleHeaderComponentConfig;
}>();

const writtenOn = computed(() =>
  DateTime
    .fromISO(props.config.createdAt)
    .setZone("Europe/Berlin")
    .setLocale("us")
    .toLocaleString(DateTime.DATE_HUGE));
</script>

<template>
  <h1 class="title">
    {{ props.config.title }}
  </h1>

  <p class="date">
    Written on <time>{{ writtenOn }}</time>
  </p>

  <p class="summary">
    {{ props.config.summary }}
  </p>
  <hr>
</template>

<style scoped>
.title {
  margin: 0;
  font-family: var(--font-raleway);
  font-size: var(--font-size-xxxl);
  line-height: var(--line-height-xxxl);
  font-weight: bold;
}

.date {
  margin: 32px 0 0 0;
  font-size: var(--font-size-s);
  line-height: var(--line-height-s);
  font-weight: lighter;
  font-style: italic;
}

.summary {
  margin: 36px 0 36px 0;
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
