<template>
  <q-card>
    <q-card-section>
      <q-form @submit="submit">
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
        <q-btn type="submit" class="mt-3 w-full bg-[#2ecc71]" label="Add" />
      </q-form>
    </q-card-section>
  </q-card>
</template>
<script setup lang="ts">
import { useQuasar } from 'quasar';
import { addComputer } from 'src/util/apis';
import { reactive } from 'vue';
const data = reactive({
  name: '',
  mac: '',
  ip: '255.255.255.255',
  port: '9',
});
const $q = useQuasar();
const emit = defineEmits(['refresh']);

const submit = () => {
  addComputer(data.name, data.mac, data.ip, data.port)
    .then((res) => {
      console.log(res);
      emit('refresh');
    })
    .catch((err) => {
      console.error(err);
      $q.notify({ message: 'Error adding computer', color: 'error' });
    });
};
</script>
