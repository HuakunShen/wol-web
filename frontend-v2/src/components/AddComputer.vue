<template>
  <q-card>
    <q-card-section>
      <div class="grid grid-cols-2 gap-4">
        <q-input color="grey-3" outlined v-model="data.name" label="Name" />
        <q-input
          color="grey-3"
          outlined
          v-model="data.mac"
          label="Mac Address"
        />
        <q-input color="grey-3" outlined v-model="data.ip" label="IP" />
        <q-input color="grey-3" outlined v-model="data.port" label="Port" />
      </div>
      <q-btn class="mt-3 w-full bg-[#2ecc71]" label="Add" />
    </q-card-section>
  </q-card>
</template>
<script setup lang="ts">
import { reactive } from 'vue';
const data = reactive({
  name: '',
  mac: '',
  ip: '',
  port: '',
});
</script>
