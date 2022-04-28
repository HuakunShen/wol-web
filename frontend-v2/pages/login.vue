<template>
  <NuxtLayout name="default">
    <div>
      <h1 class="underline">hello</h1>
      <button class="border-4 rounded text-blue-500 px-5 py-2" @click="makeToast">btn</button>
    </div>
  </NuxtLayout>
</template>
<script setup lang="ts">
import { useToast } from 'primevue/usetoast';
definePageMeta({
  layout: 'default',
});

const toast = useToast();

const makeToast = () => {
  toast.add({ severity: 'info', summary: 'Hello ' });
};
</script>
