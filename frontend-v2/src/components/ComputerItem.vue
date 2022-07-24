<template>
  <q-card>
    <q-card-section>
      <div class="grid grid-cols-5 gap-4">
        <p class="col-span-2"><strong>Name: </strong>{{ props.name }}</p>
        <p class="col-span-2"><strong>IP: </strong>{{ props.ip }}</p>
        <div class="row items-center justify-center">
          <q-btn round color="green" :icon="mdiArrowUpBox" />
        </div>
        <p class="col-span-2"><strong>Mac Address: </strong>{{ props.mac }}</p>
        <p class="col-span-2"><strong>Port: </strong>{{ props.port }}</p>
        <div class="row items-center justify-center">
          <q-btn round color="red" :icon="mdiDelete" />
        </div>
      </div>
    </q-card-section>
  </q-card>
</template>
<script setup lang="ts">
import { mdiDelete, mdiArrowUpBox } from '@quasar/extras/mdi-v6';
const props = defineProps({
  name: String,
  ip: String,
  mac: String,
  port: Number,
});
const emit = defineEmits(['wake', 'delete']);
</script>
