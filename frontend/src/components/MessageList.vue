<template>
  <div class="message-list">
    <div v-for="(message, index) in messages" :key="index">
      <MessageListItem
        :message="message.message"
        :variant="message.variant"
        :id="message.id"
      />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import MessageListItem from './MessageListItem.vue';
import { Message } from '../interfaces';
export default Vue.extend({
  props: ['messages'],
  components: {
    MessageListItem: MessageListItem,
  },
});
</script>

<style lang="css" scoped>
  .message-list {
    position: absolute;
    left: 2rem;
    bottom: 1rem;
    z-index: 1;
  }
</style>
