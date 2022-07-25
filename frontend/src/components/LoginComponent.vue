<template>
  <div>
    <q-input class="w-72" v-model="username" label="Username" />
    <q-input class="w-72" v-model="password" label="Password" />
  </div>
</template>
<script setup lang="ts">
import { ref } from 'vue';

const username = ref('');
const password = ref('');
</script>
