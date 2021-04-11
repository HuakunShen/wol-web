<template>
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
            <router-link
              class="nav-link active"
              v-if="!this.isAuth()"
              to="/login"
              >Login</router-link
            >
            <router-link
              class="nav-link active"
              v-if="!this.isAuth()"
              to="/signup"
              >Signup
            </router-link>
            <li class="nav-item" v-if="this.isAuth()">
              <a class="nav-link" href="#" v-on:click="logout">Logout</a>
            </li>
          </div>
        </div>
      </div>
    </nav>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { mapActions, mapGetters } from 'vuex';
export default Vue.extend({
  created() {
    this.loadAuth();
  },
  methods: {
    ...mapGetters(['isAuth']),
    ...mapActions(['loadAuth']),
    logout() {
      this.$store.dispatch('logout');
    },
  },
});
</script>
