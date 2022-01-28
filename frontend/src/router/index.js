import { createRouter, createWebHistory } from 'vue-router'
// import firebase from 'firebase/app'
// import 'firebase/app'
// import "firebase/auth"

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
    path: '/detail/:id',
    name: 'Detail',
    props: true,
    component: () => import('../views/Detail.vue')
  },
  {
    path: '/mypage/:id/follow',
    name: 'Follow',
    props: true,
    component: () => import('../views/Follow.vue')
  },
  {
    path: '/mypage/:id/follower',
    name: 'Follower',
    props: true,
    component: () => import('../views/Follower.vue')
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

// router.beforeEach((to, from, next) => {
//   const requiresAuth = to.matched.some(record => record.meta.requiresAuth)
//   if (requiresAuth) {
//     firebase.auth().onAuthStateChanged(function (user) {
//       console.log(user)
//       if (user) {
//         next()
//       } else {
//         next({ name: 'login' })
//       }
//     })
//   } else {
//     next()
//   }
// })

export default router
