<template>
  <div id="nav">
    <router-link to="/">Home |</router-link>
    <router-link to="/signup" v-if="notAuthenticatedUser">新規登録 |</router-link>
    <router-link to="/login" v-if="notAuthenticatedUser">ログイン |</router-link>
    <button @click="signOut" v-if="authenticatedUser">ログアウト</button>
    <router-link to="/post" v-if="authenticatedUser">投稿する</router-link>
    <router-link :to="{name: 'Mypage', params: {id:(Number(currentUserId))}}" v-if="authenticatedUser">{{ currentUserName }}さん</router-link>
  </div>
</template>

<script>
import axios from 'axios'
import firebase from 'firebase/app'
import 'firebase/app'
import "firebase/auth"


export default {
  data() {
    return {
      authenticatedUser:'',
      notAuthenticatedUser:'',
      currentUserId:'',
      currentUserName:localStorage.getItem('userName'),
    }
  },
  mounted(){
    const params = new URLSearchParams()
    params.append('id', localStorage.getItem('userId'))
    axios.post('getHeader', params)
    .then(response => {
      if (response.data != '') {
      this.authenticatedUser = true
      this.notAuthenticatedUser = false
      } else {
        this.authenticatedUser = false
        this.notAuthenticatedUser = true
      }
      this.currentUserId = response.data
    })
  },
  methods: {
    signOut() {
      confirm('ログアウトしてもよろしいですか。')
      firebase.auth().signOut()
      .then(() => {
        localStorage.removeItem('jwt')
        localStorage.removeItem('userName')
        localStorage.removeItem('userId')
        this.$router.go({path: this.$router.push('/login'), force: true})
        alert("ログアウトしました。")
      })
    }
  }
}
</script>
