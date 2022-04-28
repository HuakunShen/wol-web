<template>
  <NuxtLayout name="default">
    <div>
      <h1 class="underline">hello</h1>
      <button class="border-4 rounded text-blue-500 px-5 py-2" @click="makeToast">btn</button>
      {{ filtersList }}
      {{ filtersListGetter }}
    </div>
  </NuxtLayout>
</template>
<script setup lang="ts">
import { useFiltersStore } from '~/store/filters';
import { useToast } from 'primevue/usetoast';
import { storeToRefs } from 'pinia';

definePageMeta({
  layout: 'default',
});
const store = useFiltersStore();
const { filtersList } = storeToRefs(store);

const toast = useToast();
const filtersListGetter = store.filtersListGetter;
const makeToast = () => {
  toast.add({ severity: 'info', summary: 'Hello ' });
};
</script>
