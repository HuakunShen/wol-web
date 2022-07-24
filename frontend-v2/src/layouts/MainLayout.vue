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
        <q-toolbar-title><strong>WOL</strong></q-toolbar-title>
        <div>
          <q-btn
            round
            color="black"
            :icon="$q.dark.mode ? mdiWeatherNight : mdiWeatherSunny"
            @click="toggleDark()"
          />
        </div>
      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" show-if-above bordered>
      <q-list>
        <q-item-label header> Essential Links </q-item-label>

        <EssentialLink
          v-for="link in essentialLinks"
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

<script lang="ts">
import { defineComponent, ref } from 'vue';
import EssentialLink from 'components/EssentialLink.vue';
import { matMenu } from '@quasar/extras/material-icons';
import {
  mdiAbTesting,
  mdiWeatherSunny,
  mdiWeatherNight,
} from '@quasar/extras/mdi-v6';
import { fasFont } from '@quasar/extras/fontawesome-v5';
import { useQuasar } from 'quasar';
import { useDark, useToggle } from '@vueuse/core';

const linksList = [
  {
    title: 'Docs',
    caption: 'quasar.dev',
    icon: 'school',
    link: 'https://quasar.dev',
  },
  {
    title: 'Github',
    caption: 'github.com/quasarframework',
    icon: 'code',
    link: 'https://github.com/quasarframework',
  },
  {
    title: 'Discord Chat Channel',
    caption: 'chat.quasar.dev',
    icon: 'chat',
    link: 'https://chat.quasar.dev',
  },
  {
    title: 'Forum',
    caption: 'forum.quasar.dev',
    icon: 'record_voice_over',
    link: 'https://forum.quasar.dev',
  },
  {
    title: 'Twitter',
    caption: '@quasarframework',
    icon: 'rss_feed',
    link: 'https://twitter.quasar.dev',
  },
  {
    title: 'Facebook',
    caption: '@QuasarFramework',
    icon: 'public',
    link: 'https://facebook.quasar.dev',
  },
  {
    title: 'Quasar Awesome',
    caption: 'Community Quasar projects',
    icon: 'favorite',
    link: 'https://awesome.quasar.dev',
  },
];

export default defineComponent({
  name: 'MainLayout',

  components: {
    EssentialLink,
  },
  setup() {
    const leftDrawerOpen = ref(false);
    const $q = useQuasar();
    $q.dark.set(true);
    const isDark = useDark();
    const toggleDark = useToggle(isDark);
    toggleDark(true);
    return {
      $q: $q,
      isDark,
      toggleDark: () => {
        toggleDark();
        $q.dark.toggle();
      },
      essentialLinks: linksList,
      leftDrawerOpen,
      toggleLeftDrawer() {
        leftDrawerOpen.value = !leftDrawerOpen.value;
      },
      matMenu,
      mdiAbTesting,
      fasFont,
      mdiWeatherSunny,
      mdiWeatherNight,
    };
  },
});
</script>
