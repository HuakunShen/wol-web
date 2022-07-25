<template>
  <router-view />
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useAuthStore, authStoreId } from './stores/auth';

const authStore = useAuthStore();

onMounted(() => {
  authStore.loadFromLocalStorage();
});

authStore.$subscribe(
  () => {
    localStorage.setItem(authStoreId, JSON.stringify(authStore.$state));
  },
  { detached: true }
);
</script>
