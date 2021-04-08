<template>
  <div>
    <div class="container">
      <div class="login-card">
        <h1>Login</h1>
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
          <button type="submit" class="btn btn-primary">Submit</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { mapMutations } from 'vuex';
export default Vue.extend({
  data() {
    return {
      username: '',
      password: '',
    };
  },
  created() {
    console.log(this.$store.state);
  },
  methods: {
    ...mapMutations(['updateAuth']),
    async submit(e) {
      e.preventDefault();
      if (this.username && this.password) {
        const res = await fetch('/api/users/login', {
          method: 'POST',
          headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            username: this.username,
            password: this.password,
          }),
        });
        const content = await res.json();
        if (res.status == 200) {
          console.log(content);
          this.updateAuth({ isAuth: true, username: this.username });
        } else {
          console.error(content.message);
        }
        console.log(content);
      } else {
        console.error('error invalid input');
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
  .login-card {
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
