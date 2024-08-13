<script setup lang="ts">
import { onBeforeMount, ref } from "vue";

const tagArray = ref();

onBeforeMount(() => {
  getTags();
});

async function getTags() {
  const tageUrl = "http://localhost:5173/data/tag-list";

  try {
    const response = await fetch(tageUrl, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }

    tagArray.value = await response.json();
  } catch (error) {
    console.log(error);
  }
}

</script>
<template>
  <table>
    <thead>
      <tr>
        <th>key</th>
        <th>name</th>
      </tr>
    </thead>
    <tbody>
      <tr
        v-for="tag in tagArray.data"
        :key="tag.key"
      >
        <td>{{ tag.key }}</td>
        <td>{{ tag.name }}</td>
      </tr>
    </tbody>
  </table>
</template>
