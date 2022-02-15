<template>
  <div>
    <Header></Header>
    <div class="text-center">
      <h1 class="title">新規登録</h1>
      <div v-if="errors.length">
        <p v-for="error in errors" :key="error">{{ error }}</p>
      </div>
      <p v-if="apiErrors.length">
      <ul>
        <li v-for="error in apiErrors" :key="error">{{ error }}</li>
      </ul>
      </p>
      <div class="mb-3">
        <label for="name">お名前:</label>
        <input type="text" placeholder="Username" name='name' v-model="name">
      </div>
      <div class="mb-3">
        <label for="email">メールアドレス:</label>
        <input type="text" placeholder="Password" name='email' v-model="email">
      </div>
      <div class="mb-3">
        <label for="password">パスワード:</label>
        <input type="password" name='password' v-model="password">
      </div>
      <div class="mb-3">
        <button class="btn btn-outline-warning" @click="signUp">登録する</button>
      </div>
      <div class="mb-3">
        <p>Do you have an account?
          <router-link to="/login">sign in now!!</router-link>
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
      email: "",
      password: "",
      users: [],
      errors: [],
      apiErrors: [],
    }
  },
  components: { Header },
  methods: {
    signUp() {
      firebase.auth().createUserWithEmailAndPassword(this.email, this.password)
      .then(response => {
        response.user.getIdToken()
        .then(idToken => {
          localStorage.setItem("jwt", idToken.toString())
        })
      })
      .catch(error => {
        switch (error.code) {
          case 'auth/email-already-in-use' :
            this.errors.push("*このメールアドレスはすでに登録されています。")
            break;
          case 'auth/weak-password':
            this.errors.push("*パスワードは6文字以上で入力してください。")
            break;
        }
      })
      const params = new URLSearchParams()
      params.append('name', this.name)
      params.append('email', this.email)
      params.append('password', this.password)
      axios.post('/signUp', params)
      .then(response => {
        if (response.status == 201) {
          this.apiErrors.push(response.data.Name)
          this.apiErrors.push(response.data.Email)
          this.apiErrors.push(response.data.Password)
        } else {
          alert("ようこそ" + response.data + 'さん')
          this.$router.push('/')
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
