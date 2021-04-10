<template>
  <div id="app">
    <Navbar />
    <router-view />
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { mapActions } from 'vuex';
import Navbar from './components/Navbar.vue';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.min.js';
export default Vue.extend({
  created() {
    this.loadAuth();
    this.loadMacs();
  },
  components: {
    Navbar: Navbar,
  },
  methods: {
    ...mapActions(['loadAuth', 'loadMacs']),
    logout() {
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
