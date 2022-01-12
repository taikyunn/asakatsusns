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
    path: '/post',
    name: 'Post',
    component: () => import('../views/Post.vue')
  },
  {
    path: '/edit/:id',
    name: 'Edit',
    props: true,
    component: () => import('../views/Edit.vue')
  },
  {
    path: '/mypage/:id',
    name: 'Mypage',
    props: true,
    component: () => import('../views/Mypage.vue')
  },
  {
    path: '/sleep_time_edit/:id',
    name: 'SleepTimeEdit',
    props: true,
    component: () => import('../views/SleepTimeEdit.vue')
  },
  {
    path: '/wake_up_time_edit/:id',
    name: 'WakeUpTimeEdit',
    props: true,
    component: () => import('../views/WakeUpTimeEdit.vue')
  },
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
        console.log("ここが動いています。")

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
