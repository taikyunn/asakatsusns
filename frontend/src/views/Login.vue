<template>
  <div>
    <Header></Header>
    <div class="text-center">
      <h1 class="title">ログインフォーム</h1>
      <div v-if="errors.length">
        <p v-for="error in errors" :key="error">{{ error }}</p>
      </div>
      <div v-if="apiErrors.length">
        <p v-for="error in apiErrors" :key="error">{{ error }}</p>
      </div>
        <div class="mb-3">
          <label for="name">お名前:</label>
          <input type="text" placeholder="name" name='name' v-model="name">
        </div>
        <div class="mb-3">
          <label for="name">メールアドレス:</label>
          <input type="email" placeholder="email" name='name' v-model="email">
        </div>
        <div class="mb-3">
          <label for="password">パスワード:</label>
          <input type="password" placeholder="password" name='password' v-model="password">
        </div>
        <div class="mb-3">
          <button @click="loginFirebase">ログイン</button>
      </div>
      <div class="mb-3">
          <p>You don't have an account?
            <router-link to="/signup">create account now!!</router-link>
          </p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import firebase from 'firebase/app'
import 'firebase/app'
import "firebase/auth"
import Header from './Header.vue'

export default {
  data() {
    return {
      name: "",
      email:'',
      password: "",
      errors: [],
      apiErrors: [],
      }
  },
  components: { Header },
  methods: {
    loginFirebase() {
      firebase.auth().signInWithEmailAndPassword(this.email, this.password)
      .then(response => {
        response.user.getIdToken()
        .then(idToken => {
          localStorage.setItem("jwt", idToken.toString())
          this.login()
        })
      })
      .catch(error => {
        switch (error.code) {
          case 'auth/invalid-email' :
            this.errors.push("*メールアドレスの形式が正しくありません")
            break;
          case 'auth/wrong-password' :
            this.errors.push("*メールアドレスまたはパスワードが違います")
            break;
        }
      })
    },
    login() {
      const params = new URLSearchParams
      params.append('name', this.name)
      params.append('password', this.password)
      params.append('email', this.email)
      axios.post('/login',params)
      .then(response => {
        if (response.status == 201) {
          this.apiErrors.push(response.data.dbError)
          this.apiErrors.push(response.data.Name)
        } else if (response.status != 200){
          throw new Error('レスポンスエラー')
        } else {
          localStorage.setItem('userName', response.data.name)
          localStorage.setItem('userId', response.data.userId)
          this.$router.push('/')
          alert("ようこそ" + response.data.name + "さん")
        }
      })
    }
  }
}
</script>

<style scoped>
.title {
  margin-top: 50px;
  margin-bottom: 50px;
}
</style>
