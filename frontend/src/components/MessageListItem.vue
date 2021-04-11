<template>
  <div class="alert alert-danger" role="alert">
    {{ message.message }}
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { Message } from '../interfaces';
export default Vue.extend({
  props: ['message'],
  computed: {},
});
</script>

<style lang="scss" scoped></style>
