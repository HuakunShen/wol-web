<template>
  <div class="alert" v-bind:class="variant" role="alert">
    {{ message }}
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Message } from '../interfaces';
export default Vue.extend({
  props: { message: String, variant: String, id: String },
});
</script>

<style lang="scss" scoped></style>
