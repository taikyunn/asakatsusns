<template>
  <div>
    <h1>ログアウト画面</h1>
    <button @click="signOut">Sign out</button>
  </div>
</template>

<script>
import firebase from 'firebase/app'
import 'firebase/app'
import "firebase/auth"

export default{
  methods: {
    signOut() {
      firebase.auth().signOut()
      .then(() => {
        localStorage.removeItem('jwt')
        localStorage.removeItem('userName')
        localStorage.removeItem('userId')
        this.$router.push('/')
        alert("ログアウトしました。")
      })
    }
  }
}
</script>
