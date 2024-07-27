<script setup lang="ts">
import { computed, ref } from "vue";
import type { Tag } from "../../types/data/tag";

const props = defineProps<{
  tags: Tag[];
}>();

const emits = defineEmits<{
  search: [searchQuery: { q: string; tagKeys: string[] }];
}>();

const selectedTagKeys = ref([] as string[]);
const searchInputText = ref("");
const tagAreaExpanded = ref(false);

const selectedTags = computed(() => props.tags
  .filter((tag) => selectedTagKeys.value.includes(tag.key))
  .sort((tag1, tag2) => tag1.name.localeCompare(tag2.name)));
const unselectedTags = computed(() => props.tags
  .filter((tag) => !selectedTagKeys.value.includes(tag.key))
  .sort((tag1, tag2) => tag1.name.localeCompare(tag2.name)));

const addTag = (tagKey: string) => {
  selectedTagKeys.value.push(tagKey);
};

const removeTag = (tagKey: string) => {
  selectedTagKeys.value = selectedTagKeys.value.filter((key) => key !== tagKey);
  triggerSearch();
};

const toggleTagAreaExpanded = () => {
  tagAreaExpanded.value = !tagAreaExpanded.value;
  triggerSearch();
};

const triggerSearch = () => {
  emits("search", {
    q: searchInputText.value,
    tagKeys: selectedTagKeys.value,
  });
};
</script>

<template>
  <div class="search-area">
    <div class="input-area">
      <input
        v-model="searchInputText"
        class="search-input"
        type="text"
        placeholder="Search an article"
        size="10"
        @keyup.enter="triggerSearch()"
      >
      <button
        class="search-button"
        @click="triggerSearch()"
      >
        SEARCH
      </button>
    </div>
    <div class="tag-area">
      <button
        v-for="tag of selectedTags"
        :key="tag.key"
        class="tag tag-selected"
        @click="removeTag(tag.key)"
      >
        {{ tag.name }}
      </button>

      <button
        v-for="tag of unselectedTags"
        :key="tag.key"
        class="tag tag-unselected"
        :class="{
          'tag-expanded': tagAreaExpanded
        }"
        @click="addTag(tag.key)"
      >
        {{ tag.name }}
      </button>
    </div>

    <button
      class="tag-toggle-button"
      @click="toggleTagAreaExpanded()"
    >
      {{ tagAreaExpanded ? "SHOW LESS TAGS" : "FILTER BY TAGS" }}
    </button>
  </div>
</template>

<style scoped>
.input-area {
  display: flex;
}

.search-input {
  flex-grow: 1;
  flex-shrink: 0;
  border: 1px solid var(--color-primary);
  outline: 0;
  padding: 0 12px;
  height: 36px;
  color: var(--color-primary);
  font-family: var(--font-open-sans);
  font-size: var(--font-size-m);
  line-height: var(--line-height-m);
}

.search-button {
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
.search-button:hover {
  border-color: var(--color-primary-t1);
  background-color: var(--color-primary-t1);
}

.tag-area {
  display: flex;
  flex-wrap: wrap;
  margin: 4px 0;
}

.tag {
  display: inline-block;
  box-sizing: content-box;
  margin: 8px 0 4px 0;
  border: 1px solid var(--color-primary);
  padding: 0 16px;
  height: 36px;
  font-family: var(--font-open-sans);
  font-weight: bold;
  font-size: var(--font-size-s);
  line-height: var(--line-height-s);
  cursor: pointer;
}
.tag:not(:last-child) {
  margin-right: 12px;
}
.tag-selected {
  color: var(--color-secondary);
  background-color: var(--color-primary);
}
.tag-selected:hover {
  border-color: var(--color-primary-t1);
  background-color: var(--color-primary-t1);
}
.tag-unselected {
  display: none;
  color: var(--color-primary);
  background-color: var(--color-secondary);
}
.tag-unselected:hover {
  background-color: var(--color-secondary-s1);
}

.tag-expanded {
  display: inline-block;
}

.tag-toggle-button {
  margin: 0;
  border: none;
  padding: 0;
  color: var(--color-primary);
  background-color: transparent;
  font-family: var(--font-open-sans);
  font-weight: bold;
  font-size: var(--font-size-m);
  line-height: var(--line-height-m);
  cursor: pointer;
}
.tag-toggle-button:hover {
  color: var(--color-primary-t1);
  text-decoration: underline;
}

@media (min-width: 500px) {
  .tag-unselected {
    display: inline-block;
  }

  .tag-toggle-button {
    display: none;
  }
}
</style>
