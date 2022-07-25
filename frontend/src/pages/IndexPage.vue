<template>
  <q-page class="row justify-center px-2">
    <q-list class="mt-3">
      <add-computer @refresh="refresh" />
      <q-list>
        <computer-item
          v-for="(computer, idx) in computers"
          :key="idx"
          class="mt-3"
          :name="computer.name"
          :mac="computer.mac"
          :port="computer.port"
          :ip="computer.ip"
          @wake="wake(computer.id)"
          @delete="del(computer.id)"
        />
      </q-list>
    </q-list>
  </q-page>
</template>

<script setup lang="ts">
import AddComputer from 'src/components/AddComputer.vue';
import { ref, onMounted } from 'vue';
import ComputerItem from 'components/ComputerItem.vue';
import { getComputers, wol, deleteComputer } from 'src/util/apis';
import { useQuasar } from 'quasar';
import { useAuthStore } from 'src/stores/auth';
import { useRouter } from 'vue-router';

const router = useRouter();
const authStore = useAuthStore();
const $q = useQuasar();

const computers = ref<
  { name: string; mac: string; port: number; ip: string; id: number }[]
>([]);

const refresh = () => {
  getComputers().then((res) => {
    console.log(res);
    computers.value = res.data.data.map((c) => ({
      name: c.name,
      mac: c.mac,
      ip: c.ip,
      port: c.port,
      id: c.id,
    }));
  });
};

const wake = (id: number) =>
  wol(id)
    .then(() => {
      $q.notify({ message: 'Woke Up Computer', color: 'success' });
    })
    .catch((err) => {
      $q.notify({ message: 'Error Waking Up Computer', color: 'red' });
    });

const del = (id: number) => {
  deleteComputer(id)
    .then(() => {
      refresh();
      $q.notify({ message: 'Deleted Computer', color: 'success' });
    })
    .catch((err) => {
      $q.notify({ message: `Error Deleting Computer: ${err}`, color: 'red' });
    });
};

onMounted(() => {
  // if (!authStore.isAuth) {
  //   router.push('/auth');
  //   $q.notify({ message: 'Please Login', color: 'red' });
  //   return;
  // }
  refresh();
});
</script>
