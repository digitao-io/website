<script setup lang="ts">

export type ArticleHeaderComponentConfig = {
  title: string;
  createdAt: string;
  summary: string;
};

const props = defineProps<{
  config: ArticleHeaderComponentConfig;
}>();
</script>

<template>
  <h1 class="blog-title">
    {{ props.config.title }}
  </h1>

  <p>Written on: <time class="blog-time">{{ props.config.createdAt }}</time></p>

  <p class="blog-summarize">
    {{ props.config.summary }}
  </p>
  <hr>
</template>

<style scoped>
.blog-title {
   font-family: var(--font-raleway);
   font-size: 48px;
   font-weight: bold;
   line-height: 48px;
   color: var(--color-primary);
}

.blog-title{
  font-size: var(--font-size-s);
  line-height: var(--line-height-s);
}

.blog-summarize {
  font-family: var(--font-open-sans);
  font-size: 16px;
  font-weight: lighter;
  font-style: italic;
  line-height: auto;
  color: var(--color-primary-s1);
}

@media (min-width: 500px) {
  .blog-title{
    font-size: 64px;
    line-height: 64px;
  }
}

@media (min-width: 100px) {
  .blog-title{
    font-size: 64px;
    line-height: 64px;
  }
}
</style>
