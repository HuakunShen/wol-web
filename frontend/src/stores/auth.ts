import { defineStore } from 'pinia';

export type StateType = {
  isAuth: boolean;
};

export const authStoreId = 'auth-store';

export const loadAuthFromLocalStorage = () => {
  const authData = localStorage.getItem(authStoreId);
  if (!authData) return null;
  const parsedAuthData = authData ? JSON.parse(authData) : null;
  return parsedAuthData as StateType;
};

export const clearAuthLocalStorage = () => {
  localStorage.removeItem(authStoreId);
};

export const useAuthStore = defineStore('auth', {
  state: () => ({
    isAuth: false,
  }),
  // getters: { isAuth: (state) => state.isAuthenticated },
  actions: {
    setAuth(isAuth: boolean) {
      console.log('isAuth', isAuth);
      this.isAuth = isAuth;
    },
    loadFromLocalStorage() {
      const storage = loadAuthFromLocalStorage();
      console.log(storage);

      if (storage) {
        this.isAuth = storage.isAuth;
      } else {
        this.setAuth(false);
      }
    },
    clearLocalStorage() {
      console.log('clearLocalStorage');

      clearAuthLocalStorage();
    },
  },
});
