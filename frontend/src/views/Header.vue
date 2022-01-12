<template>
  <div id="nav">
    <router-link to="/">Home |</router-link>
    <router-link to="/signup" v-if="notAuthenticatedUser">新規登録 |</router-link>
    <router-link to="/login" v-if="notAuthenticatedUser">ログイン |</router-link>
    <router-link to="/signout" v-if="authenticatedUser">ログアウト |</router-link>
    <router-link to="/post" v-if="authenticatedUser">投稿する</router-link>
    <router-link :to="{name: 'Mypage', params: {id:(Number(currentUserId))}}">{{ currentUserName }}さん</router-link>
  </div>
</template>

<script>
  import axios from 'axios'


export default {
  data() {
    return {
      authenticatedUser:'',
      notAuthenticatedUser:'',
      currentUserId:'',
      currentUserName:localStorage.getItem('userName'),
    }
  },
  created(){
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
      this.currentUserId =response.data
    })

  },
}
</script>
