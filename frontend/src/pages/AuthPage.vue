<template>
  <q-page class="row justify-center items-center">
    <div>
      <q-card class="my-card text-white px-5 py-5">
        <q-tabs
          v-model="tab"
          dense
          class="text-grey"
          active-color="primary"
          indicator-color="primary"
          align="justify"
          narrow-indicator
        >
          <q-tab name="login" label="Login" />
          <q-tab name="signup" label="Sign Up" />
        </q-tabs>

        <q-separator />

        <q-tab-panels v-model="tab" animated>
          <q-tab-panel name="login">
            <q-form @submit="onLogin" @reset="onReset" class="q-gutter-md">
              <q-input
                class="w-72"
                v-model="username"
                label="Username"
                :rules="[nonEmptyInputRule]"
              />
              <q-input
                class="w-72"
                v-model="password"
                label="Password"
                type="password"
                :rules="[nonEmptyInputRule]"
              />
              <q-btn
                type="submit"
                color="primary"
                class="full-width mt-4"
                label="Login"
              />
              <q-btn type="reset" class="full-width mt-4" label="Reset" />
            </q-form>
          </q-tab-panel>

          <q-tab-panel name="signup">
            <q-form @submit="onSignUp" @reset="onReset" class="q-gutter-md">
              <p>
                Registration Quota: {{ userCount }} / {{ allowedUserCount }}
              </p>
              <p class="text-red" v-if="userCount === allowedUserCount">
                You cannot register
              </p>
              <q-input
                class="w-72"
                v-model="username"
                label="Username"
                :rules="[nonEmptyInputRule]"
              />
              <q-input
                class="w-72"
                v-model="password"
                label="Password"
                type="password"
                :rules="[nonEmptyInputRule]"
              />
              <q-input
                class="w-72"
                v-model="confirmPassword"
                label="Confirm Password"
                type="password"
                :rules="[nonEmptyInputRule]"
              />
              <q-btn
                type="submit"
                color="secondary"
                class="full-width mt-4"
                label="Sign Up"
              />
              <q-btn type="reset" class="full-width mt-4" label="Reset" />
            </q-form>
          </q-tab-panel>
        </q-tab-panels>
      </q-card>
    </div>
  </q-page>
</template>
<script setup lang="ts">
import { onMounted, ref } from 'vue';
import axios from 'axios';
import { useQuasar } from 'quasar';
import { getComputers, getUserCount, loadAuth, login } from 'src/util/apis';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../stores/auth';

const authStore = useAuthStore();
const router = useRouter();
const $q = useQuasar();
const tab = ref('login');
const username = ref('');
const password = ref('');
const confirmPassword = ref('');

const userCount = ref<number>(0);
const allowedUserCount = ref<number>(0);

getUserCount().then((res) => {
  userCount.value = res.data.data.user_count;
  allowedUserCount.value = res.data.data.num_user_allowed;
});

const nonEmptyInputRule = (val: string) =>
  (val && val.length > 0) || 'Please type something';

const onLogin = () =>
  login(username.value, password.value)
    .then(async (res) => {
      authStore.setAuth(true);
      $q.notify({ message: 'You are Logged In.', color: 'green' });
      router.push('/');
    })
    .catch((err) => {
      $q.notify({ message: `Error Loggin In: ${err}`, color: 'red' });
    });

const onSignUp = () => {
  if (password.value !== confirmPassword.value) {
    $q.notify({ message: "Passwords don't Match", color: 'red' });
    return;
  }
  axios
    .post('/api/users/register', {
      username: username.value,
      password: password.value,
    })
    .then((res) => {
      console.log(res);
      $q.notify({ message: 'User Created', color: 'green' });
      return onLogin();
    })
    .catch((err) => {
      $q.notify({ message: `Error Creating User: ${err}`, color: 'red' });
    });
};
const onReset = () => {
  username.value = '';
  password.value = '';
  confirmPassword.value = '';
};
</script>
