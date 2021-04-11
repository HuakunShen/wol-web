<template>
  <div id="app">
    <Navbar />
    <MessageList />
    <router-view />
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { mapActions } from 'vuex';
import Navbar from './components/Navbar.vue';
import MessageList from './components/MessageList.vue';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.min.js';

export default Vue.extend({
  data() {
    return { interval: undefined as number | undefined };
  },
  created() {
    this.loadAuth();
    this.loadComputers();
    this.interval = setInterval(() => {
      this.loadComputers();
    }, 1000 * 60);
  },
  destroyed() {
    if (this.interval) {
      clearInterval(this.interval);
    }
  },
  components: {
    Navbar: Navbar,
    MessageList: MessageList,
  },
  methods: {
    ...mapActions(['loadAuth', 'loadComputers', 'popMessage']),
    logout(): void {
      console.log('logout');
      this.$store.dispatch('logout');
    },
  },
});
</script>

<style lang="scss">
#app {
  min-height: 100vh;
  overflow: hidden;
  position: relative;
  width: 100%;
  height: 100%;
}
</style>
