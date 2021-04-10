import Vue from 'vue';
import Vuex from 'vuex';
import { MacInterface, State } from '../interfaces';
Vue.use(Vuex);

export default new Vuex.Store<State>({
  state: {
    auth: {
      isAuth: false,
      username: '',
    },
    macs: [],
    errors: [],
    messages: [],
  },
  getters: {
    allMacs: (state) => state.macs,
    test: (state) => 'test',
  },
  mutations: {
    updateAuth(state, payload: { isAuth: boolean; username: string }) {
      state.auth.username = payload.username;
      state.auth.isAuth = payload.isAuth;
    },
    updateMacs(state, payload: { macs: Array<MacInterface> }) {
      state.macs = payload.macs;
    },
  },
  actions: {
    async logout({ commit }) {
      const res = await fetch('/api/users/logout', {
        method: 'POST',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
      });
      if (res.status === 200) {
        commit('updateAuth', { isAuth: false, username: null });
      } else {
        console.error('error: logout');
      }
    },
    async loadAuth({ commit }) {
      const res = await fetch('/api/users/user', {
        method: 'GET',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
      });
      const content = await res.json();
      console.log(content);
      if (res.status < 400) {
        commit('updateAuth', { isAuth: true, username: content.username });
      } else {
        commit('updateAuth', { isAuth: false, username: null });
      }
    },
    async loadMacs({ commit }) {
      const res = await fetch('/api/mac', {
        method: 'GET',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
      });
      const content = await res.json();
      if (res.status < 400) {
        commit('updateMacs', { macs: content.data });
      } else {
        commit('updateMacs', { macs: [] });
      }
    },
    async login({ commit }, payload) {
      const { username, password } = payload;
      if (username && password) {
        const res = await fetch('/api/users/login', {
          method: 'POST',
          headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            username: username,
            password: password,
          }),
        });
        const content = await res.json();
        if (res.status < 400) {
          console.log(content);
          commit('updateAuth', { isAuth: true, username: username });
        } else {
          console.error(content.message);
        }
      }
    },
    async signup({ commit }, payload) {
      const { username, password } = payload;
      const res = await fetch('/api/users/register', {
        method: 'POST',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          username: username,
          password: password,
        }),
      });
      const content = await res.json();
      if (res.status < 400) {
        console.log(content);
        commit('updateAuth', { isAuth: true, username: username });
      } else {
        console.error(content.message);
      }
    },
    async addMac(
      { dispatch },
      payload: { name: string; mac: string; ip: string; port: string }
    ) {
      const { name, mac, ip, port } = payload;
      if (name && mac) {
        const res = await fetch('/api/mac', {
          method: 'POST',
          headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ name, mac, ip, port }),
        });
        const content = await res.json();
        if (res.status < 400) {
          console.log(content);
          dispatch('loadMacs');
        } else {
          console.error(content.message);
        }
      } else {
        console.error('Invalid Input');
      }
    },
    async deleteMac({ dispatch }, payload: { id: number }) {
      console.log(payload);
      if (payload.id) {
        const res = await fetch(`/api/mac/${payload.id}`, {
          method: 'DELETE',
          headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
          },
        });
        if (res.status < 400) {
          dispatch('loadMacs');
        } else {
          console.error('Fail to delete');
          console.error(res);
        }
      }
    },
  },
  modules: {},
});
