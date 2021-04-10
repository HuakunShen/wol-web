<template>
  <div>
    <div class="container">
      <div class="signup-card">
        <h1>Sign Up</h1>
        <form @submit="submit">
          <div class="mb-3">
            <label for="username" class="form-label">Username</label>
            <input
              v-model="username"
              type="text"
              class="form-control"
              id="username"
              aria-describedby="emailHelp"
              required
            />
          </div>
          <div class="mb-3">
            <label for="password" class="form-label">Password</label>
            <input
              v-model="password"
              type="password"
              class="form-control"
              id="password"
              required
            />
          </div>
          <div class="mb-3">
            <label for="repeat-password" class="form-label"
              >Repeat Password</label
            >
            <input
              v-model="passwordRepeat"
              type="password"
              class="form-control"
              id="repeat-password"
              required
            />
          </div>
          <button type="submit" class="btn btn-primary">Submit</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { mapMutations, mapActions } from 'vuex';

export default Vue.extend({
  data() {
    return {
      username: '',
      password: '',
      passwordRepeat: '',
    };
  },
  created() {
    console.log(this.$store.state);
  },
  methods: {
    ...mapMutations(['updateAuth']),
    ...mapActions(['signup']),
    async submit(e: Event) {
      e.preventDefault();
      if (
        this.username &&
        this.password &&
        this.password === this.passwordRepeat
      ) {
        this.signup({ username: this.username, password: this.password });
      } else {
        console.error('invalid input');
      }
    },
  },
});
</script>
<style lang="scss" scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100%;
  box-sizing: border-box;
  .signup-card {
    background-color: #eee;
    padding: 2rem 2rem 2rem 2rem;
    border-radius: 10px;
    transform: translate(0, 10vh);
    button {
      width: 100%;
    }
  }
}
</style>
