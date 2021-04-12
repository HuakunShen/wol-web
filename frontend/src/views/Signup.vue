<template>
  <div>
    <div class="container">
      <div class="signup-card">
        <h1>Sign Up</h1>
        <p>
          <strong>User Quota:</strong> {{ user_count }} /
          {{ num_user_allowed }}
        </p>
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
          <div
            v-if="lastMsg()"
            class="mt-3 alert"
            v-bind:class="lastMsg().variant"
            role="alert"
          >
            {{ lastMsg().message }}
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';
import { mapMutations, mapActions, mapGetters } from 'vuex';

export default Vue.extend({
  data() {
    return {
      username: '',
      password: '',
      passwordRepeat: '',
      num_user_allowed: 0,
      user_count: 0,
    };
  },
  async created() {
    const res = await fetch('/api/users/count', {
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
    });
    const content = await res.json();

    if (res.status < 400) {
      const { num_user_allowed, user_count } = content.data;
      this.num_user_allowed = num_user_allowed;
      this.user_count = user_count;
    } else {
      console.error(content);
      this.pushMessage({ message: content.message, variant: 'alert-danger' });
    }
  },
  watch: {
    '$store.state.auth': {
      deep: true,
      handler: function (newValue, oldValue) {
        if (newValue.isAuth === true) {
          this.$router.push({ path: '/' });
        }
      },
    },
  },
  methods: {
    ...mapGetters(['isAuth', 'lastMsg']),
    ...mapMutations(['updateAuth']),
    ...mapActions(['signup', 'pushMessage']),
    async submit(e: Event) {
      e.preventDefault();
      if (
        this.username &&
        this.password &&
        this.password === this.passwordRepeat
      ) {
        const success = await this.signup({
          username: this.username,
          password: this.password,
        });
        if (success) this.$router.push({ path: '/login' });
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
