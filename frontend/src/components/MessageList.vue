<template>
  <div class="message-list">
    <div v-for="(msg, index) in messages" :key="index">
      <MessageListItem :msg="msg" />
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import MessageListItem from './MessageListItem.vue';
export default Vue.extend({
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
