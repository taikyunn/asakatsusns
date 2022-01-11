import { createRouter, createWebHistory } from 'vue-router'
import firebase from 'firebase/app'
import 'firebase/app'
import "firebase/auth"

const routes = [
  {
    path: '/',
    name: 'Home',
    component: function () {
      return import('../views/Home.vue')
    },
    meta: { requiresAuth: true }
  },
  {
    path: '/signup',
    name: 'signup',
    component: function () {
      return import('../views/Signup.vue')
    }
  },
  {
    path: '/login',
    name: 'login',
    component: function () {
      return import('../views/Login.vue')
    }
  },
  {
    path: '/signout',
    name: 'SignOut',
    component: () => import('../views/SignOut.vue')
  },
  {
    path: '/post',
    name: 'Post',
    component: () => import('../views/Post.vue')
  },
  {
    path: '/edit/:id',
    name: 'Edit',
    component: () => import('../views/Edit.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth)
  if (requiresAuth) {
    firebase.auth().onAuthStateChanged(function (user) {
      if (user) {
        next()
      } else {
        next({ name: 'login' })
      }
    })
  } else {
    next()
  }
})

export default router
