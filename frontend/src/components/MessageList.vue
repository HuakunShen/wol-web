<template>
  <div class="message-list">
    <div v-for="(message, index) in messages" :key="index">
      <MessageListItem :message="message" />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import MessageListItem from './MessageListItem.vue';
export default new Vue({
  computed: {
    messages: function () {
      return this.$store.state.messages;
    },
  },
  components: {
    MessageListItem: MessageListItem,
  },
});
</script>

<style lang="scss" scoped>
.message-list {
  position: absolute;
  right: 2rem;
  top: 2rem;
}
</style>
