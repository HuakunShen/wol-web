<template>
  <div class="mac-list">
    <ul
      v-bind:key="mac.id"
      v-for="mac in macs"
      class="list-group list-group-flush"
    >
      <MacListItem :mac="mac" />
    </ul>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import MacListItem from './MacListItem.vue';
export default Vue.extend({
  name: 'Mac-List',
  components: { MacListItem: MacListItem },
  props: {
    macs: Array,
  },
  data() {
    return {};
  },
});
</script>
