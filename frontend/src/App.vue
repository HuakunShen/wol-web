<template>
  <div id="app">
    <div id="nav">
      <!-- <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link> -->
      <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
        <div class="container-fluid">
          <router-link class="nav-link active" to="/">WOL</router-link>
          <button
            class="navbar-toggler"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#navbarNavAltMarkup"
            aria-controls="navbarNavAltMarkup"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <span class="navbar-toggler-icon"></span>
          </button>
          <div class="collapse navbar-collapse" id="navbarNavAltMarkup">
            <div class="navbar-nav">
              <router-link class="nav-link active" to="/login"
                >Login</router-link
              >
              <router-link class="nav-link active" to="/signup"
                >Signup
              </router-link>
              <li class="nav-item">
                <a class="nav-link" href="#" v-on:click="logout">Logout</a>
              </li>
            </div>
          </div>
        </div>
      </nav>
    </div>
    <router-view />
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { mapActions } from 'vuex';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.min.js';
export default Vue.extend({
  created() {
    this.loadAuth();
  },
  methods: {
    ...mapActions(['loadAuth']),
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
}
</style>
