<template>
  <h1 class="blog-title">
    [这里要写超大号标题]
  </h1>

  <p>Written on: <time class="blog-time">2024-07-24</time></p>

  <p class="blog-summarize">
    [这里要写一段文章的总结性文字] A website as a product is soulless. It is nothing else than a pile of text, pictures ...
    However, the goal of website is making the data flow among the users, the business and the world. It is the
    dataflow makes the website meaningful, thus alive, thus have a soul.
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
