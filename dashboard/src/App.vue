<script setup lang="ts">
import Header from "@/components/Header.vue";
import { DialogWrapper } from "vue3-promise-dialog";
</script>

<template>
  <Suspense>
    <div class="d-flex flex-column min-vh-100">
      <Header />
      <main class="flex-grow-1">
        <router-view></router-view>
      </main>
      <DialogWrapper />
    </div>
    <template #fallback>Loading...</template>
  </Suspense>
</template>
