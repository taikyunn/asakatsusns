<template>
  <div class="home">
    <h1>{{ msg }}</h1>
    <h1>ユーザー名：{{ name }}</h1>
    <router-link to="/post">投稿する</router-link>
    <button @click="apiPublic">public</button>
    <button @click="apiPrivate">private</button>
  </div>
</template>

<script>
import firebase from 'firebase/app'
import "firebase/auth"
import axios from 'axios'

export default {
  data() {
    return {
      msg: 'Welcome to Your Vue.js App',
      email:firebase.auth().currentUser.email,
      name:localStorage.getItem('userName'),
    }
  },
  methods: {
    apiPublic: async function () {
      let res = await axios.get('http://localhost:3000/public')
      this.msg = res.data.message
    },
    apiPrivate: async function () {
      let res = await axios.get('http://localhost:3000/private', {
      headers: {'Authorization': `Bearer ${localStorage.getItem('jwt')}`}
      })
      this.msg = res.data
    },
  }
}
</script>
