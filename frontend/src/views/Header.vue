<template>
  <div id="nav">
    <router-link to="/">Home |</router-link>
    <router-link to="/signup" v-if="notAuthenticatedUser">新規登録 |</router-link>
    <router-link to="/login" v-if="notAuthenticatedUser">ログイン |</router-link>
    <router-link to="/signout" v-if="authenticatedUser">ログアウト |</router-link>
  </div>
</template>

<script>


export default {
  data() {
    return {
      authenticatedUser:'',
      notAuthenticatedUser:'',
    }
  },
  created(){
    var currentUserId = localStorage.getItem('userId')
    if (currentUserId != null) {
      this.authenticatedUser = true
      this.notAuthenticatedUser = false
    } else {
      this.authenticatedUser = false
      this.notAuthenticatedUser = true
    }
  },
}
</script>
