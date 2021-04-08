import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export interface State {
  auth: {
    isAuth: boolean;
    username: string | null;
  };
  errors: Array<string> | null;
  messages: Array<string> | null;
}

export default new Vuex.Store<State>({
  state: {
    auth: {
      isAuth: false,
      username: '',
    },
    errors: [],
    messages: [],
  },
  mutations: {
    updateAuth(state, payload: { isAuth: boolean; username: string }) {
      state.auth.username = payload.username;
      state.auth.isAuth = payload.isAuth;
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
      if (res.status == 200) {
        commit('updateAuth', { isAuth: true, username: content.username });
      } else {
        commit('updateAuth', { isAuth: false, username: null });
      }
    },
  },
  modules: {},
});
