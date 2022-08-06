<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated class="bg-[#242424]">
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />
        <q-toolbar-title>
          <router-link to="/" class="text-white no-underline">
            <strong>WOL</strong>
          </router-link>
        </q-toolbar-title>
        <div>
          <q-btn
            round
            color="black"
            :icon="$q.dark.mode ? mdiWeatherNight : mdiWeatherSunny"
            @click="toggleColor()"
          />
          <router-link to="/auth" v-if="!authStore.isAuth">
            <q-btn round color="black" :icon="mdiAccount" class="ml-2" />
          </router-link>
          <q-btn
            v-if="authStore.isAuth"
            round
            color="black"
            :icon="mdiLogout"
            class="ml-2"
            @click="onLogout"
          />
        </div>
      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" show-if-above bordered>
      <q-list>
        <q-item-label header> Essential Links </q-item-label>

        <EssentialLink
          v-for="link in linksList"
          :key="link.title"
          v-bind="link"
        />
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import { defineComponent, ref } from 'vue';
import EssentialLink from 'components/EssentialLink.vue';
import { matMenu } from '@quasar/extras/material-icons';
import {
  mdiAbTesting,
  mdiWeatherSunny,
  mdiWeatherNight,
  mdiAccount,
  mdiLogout,
} from '@quasar/extras/mdi-v6';
import { fasFont } from '@quasar/extras/fontawesome-v5';
import { useQuasar } from 'quasar';
import { useDark, useToggle } from '@vueuse/core';
import { logout } from 'src/util/apis';
import { useAuthStore } from 'src/stores/auth';
import { useRouter } from 'vue-router';

const $q = useQuasar();
$q.dark.set(true);
const isDark = useDark();
const toggleDark = useToggle(isDark);
toggleDark(true);
const router = useRouter();
const leftDrawerOpen = ref(false);
const linksList = [
  {
    title: 'wol-web',
    caption: 'GitHub',
    icon: 'code',
    link: 'https://github.com/HuakunShen/wol-web',
  },
  {
    title: 'Huakun Shen',
    caption: 'Author',
    icon: 'person',
    link: 'https://huakunshen.com'
  },
  {
    title: 'Sponsor Author',
    caption: 'Sponsor Author',
    icon: 'money',
    link: 'https://github.com/sponsors/HuakunShen'
  }
];

const toggleColor = () => {
  toggleDark();
  $q.dark.toggle();
};

const authStore = useAuthStore();

const onLogout = () => {
  logout().then((res) => {
    console.log(res);
    authStore.setAuth(false);
    authStore.clearLocalStorage();
    router.push('/auth');
  });
};

const toggleLeftDrawer = () => {
  leftDrawerOpen.value = !leftDrawerOpen.value;
};
</script>
