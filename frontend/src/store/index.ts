import Vue from 'vue';
import Vuex from 'vuex';
import { v4 as uuidv4 } from 'uuid';
import { MacInterface, Message, State } from '../interfaces';
import VuexPersistence from 'vuex-persist';

const vuexLocal = new VuexPersistence<State>({
  storage: window.localStorage,
});

Vue.use(Vuex);

export default new Vuex.Store<State>({
  state: {
    auth: {
      isAuth: false,
      username: '',
    },
    computers: [],
    errors: [],
    messages: [],
  },
  getters: {
    allComputers: (state) => state.computers,
    isAuth: (state) => state.auth.isAuth,
    lastMsg: (state) =>
      state.messages && state.messages.length !== 0 ? state.messages[0] : null,
  },
  mutations: {
    updateAuth(state, payload: { isAuth: boolean; username: string }) {
      state.auth.username = payload.username;
      state.auth.isAuth = payload.isAuth;
    },
    updateComputers(state, payload: { computers: Array<MacInterface> }) {
      state.computers = payload.computers;
    },
    updateMessages(state, payload: { messages: Array<Message> }) {
      state.messages = payload.messages;
    },
  },
  actions: {
    async logout({ commit, dispatch }) {
      try {
        const res = await fetch('/api/users/logout', {
          method: 'POST',
          headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
          },
        });
        if (res.status === 200) {
          commit('updateAuth', { isAuth: false, username: null });
          commit('updateComputers', { computers: [] });
        } else {
          dispatch('loadAuth');
          console.error('error: logout');
        }
      } catch (error) {
        dispatch('loadAuth');
        console.error(error);
      }
    },
    loadAuth({ commit }) {
      return fetch('/api/users/user', {
        method: 'GET',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
      })
        .then((res) => {
          if (!res.ok) {
            throw new Error('User not loggedin');
          }
          return res.json();
        })
        .then((content) => {
          commit('updateAuth', { isAuth: true, username: content.username });
          return true;
        })
        .catch((err) => {
          commit('updateAuth', { isAuth: false, username: null });
          console.error(err);
          return false;
        });
    },
    async loadComputers({ commit, dispatch }) {
      fetch('/api/computers', {
        method: 'GET',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
      })
        .then((res) => {
          if (!res.ok) {
            throw new Error('Failed to get computers');
          }
          return res.json();
        })
        .then((content) => {
          commit('updateComputers', { computers: content.data });
          return true;
        })
        .catch((err) => {
          commit('updateComputers', { computers: [] });
          console.error(err);
          dispatch('loadAuth');
          return false;
        });
    },
    async login({ commit, dispatch }, payload) {
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
        dispatch('loadComputers');
        const content = await res.json();
        if (res.status < 400) {
          commit('updateAuth', { isAuth: true, username: username });
          dispatch('pushMessage', {
            message: 'Login Successfully',
            variant: 'alert-success',
          });
          return true;
        } else {
          console.error(content.message);
          dispatch('pushMessage', {
            message: content.message,
            variant: 'alert-danger',
          });
          return false;
        }
      }
    },
    async signup({ commit, dispatch }, payload): Promise<boolean> {
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
        dispatch('pushMessage', {
          message: 'Sign Up Successfully',
          variant: 'alert-success',
        });
        return true;
      } else {
        console.error(content.message);
        dispatch('pushMessage', {
          message: content.message,
          variant: 'alert-danger',
        });
        return false;
      }
    },
    async addComputer(
      { dispatch },
      payload: { name: string; mac: string; ip: string; port: string }
    ) {
      const { name, mac, ip, port } = payload;
      if (name && mac) {
        try {
          const res = await fetch('/api/computers', {
            method: 'POST',
            headers: {
              Accept: 'application/json',
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name, mac, ip, port }),
          });
          const content = await res.json();
          if (res.status < 400) {
            dispatch('pushMessage', {
              message: 'New Computer Added',
              variant: 'alert-success',
            });
            dispatch('loadComputers');
          } else {
            console.error(content.message);
            console.error(content.error);
            dispatch('pushMessage', {
              message: content.message,
              variant: 'alert-danger',
            });
            dispatch('loadAuth');
          }
        } catch (error) {
          console.error(error);
          dispatch('loadAuth');
          dispatch('pushMessage', {
            message: 'Failed to add computer',
            variant: 'alert-danger',
          });
        }
      } else {
        console.error('Invalid Input');
      }
    },
    async deleteComputer({ dispatch }, payload: { id: number }) {
      if (payload.id) {
        try {
          const res = await fetch(`/api/computers/${payload.id}`, {
            method: 'DELETE',
            headers: {
              Accept: 'application/json',
              'Content-Type': 'application/json',
            },
          });
          if (res.status < 400) {
            dispatch('loadComputers');
          } else {
            console.error('Fail to delete');
            console.error(res);
            dispatch('loadAuth');
          }
        } catch (error) {
          console.error(error);
          dispatch('loadAuth');
        }
      }
    },
    pushMessage(
      { commit, dispatch },
      payload: { message: string; variant: string }
    ) {
      const msgId = uuidv4();
      const newPayload = { id: msgId, ...payload };
      const messages = [newPayload, ...this.state.messages];
      commit('updateMessages', { messages });
      setTimeout(() => {
        dispatch('removeMessageById', { idToRemove: msgId });
      }, 3000);
    },
    popMessage({ commit }) {
      const messages = this.state.messages;
      messages.pop();
      commit('updateMessages', { messages });
    },
    removeMessageById({ commit }, payload: { idToRemove: string }) {
      const messages = this.state.messages.filter(
        (msg) => msg.id != payload.idToRemove
      );
      commit('updateMessages', { messages });
    },
  },
  modules: {},
  plugins: [vuexLocal.plugin],
});
