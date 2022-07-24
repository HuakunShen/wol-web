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
                :rules="[nonEmptyInputRule]"
              />
              <q-input
                class="w-72"
                v-model="confirmPassword"
                label="Confirm Password"
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
import { ref } from 'vue';
import LoginComponent from '../components/LoginComponent.vue';

const tab = ref('login');
const username = ref('');
const password = ref('');
const confirmPassword = ref('');

const nonEmptyInputRule = (val: string) =>
  (val && val.length > 0) || 'Please type something';

const onLogin = () => {
  console.log('onLogin');
};
const onSignUp = () => {
  console.log('onSignUp');
};
const onReset = () => {
  username.value = '';
  password.value = '';
  confirmPassword.value = '';
};
</script>
