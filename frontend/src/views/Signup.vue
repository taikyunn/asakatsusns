<template>
  <div>
    <h1>新規登録</h1>
    <div v-if="errors.length">
      <p v-for="error in errors" :key="error">{{ error }}</p>
    </div>
    <p v-if="apiErrors.length">
    <ul>
      <li v-for="error in apiErrors" :key="error">{{ error }}</li>
    </ul>
    </p>
    <label for="name">お名前:</label>
    <input type="text" placeholder="Username" name='name' v-model="name"><br />
    <label for="email">メールアドレス:</label>
    <input type="text" placeholder="Password" name='email' v-model="email"><br />
    <label for="password">パスワード:</label>
    <input type="password" name='password' v-model="password"><br />
    <button @click="signUp">登録する</button>
    <p>Do you have an account?
      <router-link to="/login">sign in now!!</router-link>
    </p>
    <h2>登録者一覧</h2>
    <tr>
      <th>ID:</th>
      <th>Name:</th>
    </tr>
    <tr v-for="user in users" :key="user.name">
      <td>{{user.ID}}:</td>
      <td>{{user.name}}</td>
    </tr>
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
      name: "",
      email: "",
      password: "",
      users:[],
      errors: [],
      apiErrors: [],
    }
  },
  created() {
    axios.get('getAllUsers')
    .then(response => {
      if (response.status != 200) {
        throw new Error('レスポンスエラー')
      } else {
        var resultUsers = response.data
        this.users = resultUsers
      }
    })
  },
  methods: {
    getAllUsers() {
      axios.get('getAllUsers')
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          var resultUsers = response.data
          this.users = resultUsers
        }
      })
    },
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
          // エラーハンドリング
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
.signup {
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
