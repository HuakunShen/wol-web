<template>
  <div class="alert" v-bind:class="[classType]" role="alert">
    {{ msg.message }}
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
export default Vue.extend({
  props: ['msg'],
  computed: {
    classType: function () {
      return this.msg.variant;
    },
  },
});
</script>

<style lang="scss" scoped></style>
