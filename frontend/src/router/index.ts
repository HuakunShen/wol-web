import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'
import Home from '../views/Home.vue'
import store from '../store/index'

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    beforeEnter: (to, from, next) => {
      if (!store.state.auth.isAuth) next({ name: 'Login' })
      else next()
    }
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue'),
    beforeEnter: (to, from, next) => {
      if (store.state.auth.isAuth) next({ name: 'Home' })
      else next()
    }
  },
  {
    path: '/signup',
    name: 'Sign Up',
    component: () => import('../views/Signup.vue'),
    beforeEnter: (to, from, next) => {
      if (store.state.auth.isAuth) next({ name: 'Home' })
      else next()
    }
  }
]

const router = new VueRouter({
  routes
})

export default router
