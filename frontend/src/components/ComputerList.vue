<template>
  <div class="mac-list">
    <ul
      v-bind:key="computer.id"
      v-for="computer in computers"
      class="list-group list-group-flush"
    >
      <ComputerListItem :computer="computer" />
    </ul>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import ComputerListItem from './ComputerListItem.vue';
export default Vue.extend({
  name: 'Computer-List',
  components: { ComputerListItem: ComputerListItem },
  props: {
    computers: Array,
  },
  data() {
    return {};
  },
});
</script>
