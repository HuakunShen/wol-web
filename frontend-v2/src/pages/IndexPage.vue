<template>
  <q-page class="row justify-center">
    <q-list class="mt-5">
      <add-computer />
      <q-list>
        <computer-item
          class="mt-3"
          name="ubuntu"
          mac="ff:ff:ff:ff:ff:ff"
          :port="9"
          ip="10.6.6.100"
        />
        <computer-item
          class="mt-3"
          name="ubuntu"
          mac="ff:ff:ff:ff:ff:ff"
          :port="9"
          ip="10.6.6.100"
        />
      </q-list>
    </q-list>
  </q-page>
</template>

<script lang="ts">
import { Todo, Meta } from 'components/models';
import AddComputer from 'src/components/AddComputer.vue';
import { defineComponent, ref } from 'vue';
import ComputerItem from 'components/ComputerItem.vue';

export default defineComponent({
  name: 'IndexPage',
  components: { ComputerItem, AddComputer },
  setup() {
    const todos = ref<Todo[]>([
      {
        id: 1,
        content: 'ct1',
      },
      {
        id: 2,
        content: 'ct2',
      },
      {
        id: 3,
        content: 'ct3',
      },
      {
        id: 4,
        content: 'ct4',
      },
      {
        id: 5,
        content: 'ct5',
      },
    ]);
    const meta = ref<Meta>({
      totalCount: 1200,
    });
    return { todos, meta };
  },
});
</script>
