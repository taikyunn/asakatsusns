<template>
  <div>
    <Header></Header>
    <h1>ログインフォーム</h1>
    <div v-if="errors.length">
      <p v-for="error in errors" :key="error">{{ error }}</p>
    </div>
    <div v-if="apiErrors.length">
      <p v-for="error in apiErrors" :key="error">{{ error }}</p>
    </div>
    <label for="name">お名前:</label>
    <input type="text" placeholder="name" name='name' v-model="name"><br />
    <label for="name">メールアドレス：:</label>
    <input type="email" placeholder="email" name='name' v-model="email"><br />
    <label for="password">パスワード:</label>
    <input type="password" placeholder="password" name='password' v-model="password"><br />
    <button @click="loginFirebase">ログイン</button>
    <p>You don't have an account?
        <router-link to="/signup">create account now!!</router-link>
    </p>
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
      console.log('ここが動いています')
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
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
.signin {
  margin-top: 20px;
  display: flex;
  flex-flow: column nowrap;
  justify-content: center;
  align-items: center
}
input {
  margin: 10px 0;
  padding: 10px;
}
button {
  margin: 10px 0;
  padding: 10px;
}
</style>
