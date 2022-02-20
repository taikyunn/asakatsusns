import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: function () {
      return import('../views/Home.vue')
    },
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
  {
    path: '/tags/:id',
    name: 'HomeTag',
    props: true,
    component: () => import('../views/HomeTag.vue')
  },
  {
    path: '/logout',
    redirect: '/'
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
