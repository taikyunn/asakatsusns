<template>
  <div>
    <h1>ログインフォーム</h1>
    <label for="name">お名前:</label>
    <input type="text" name='name' v-model="name"><br />
    <label for="email">メールアドレス:</label>
    <input type="text" name='email' v-model="email"><br />
    <label for="password">パスワード:</label>
    <input type="password" name='password' v-model="password"><br />
    <button @click="registerUser">登録する</button>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      name: "",
      email: "",
      password: "",
    }
  },
  methods: {
    registerUser() {
      const params = new URLSearchParams()
      params.append('name', this.name)
      params.append('email', this.email)
      params.append('password', this.password)
      axios.post('/registerUser', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          alert("登録しました。")
          this.getAllUser()
        }
      })
    },
  }
}
</script>
