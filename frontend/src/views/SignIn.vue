<template>
  <div>
    <h1>ログインフォーム</h1>
    <label for="name">お名前:</label>
    <input type="text" name='name' v-model="name"><br />
    <label for="password">パスワード:</label>
    <input type="password" name='password' v-model="password"><br />
    <button @click="signIn">ログイン</button>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      name: "",
      password: "",
    }
  },
  methods: {
    signIn() {
      const params = new URLSearchParams()
      params.append('name', this.name)
      params.append('password', this.password)
      axios.post('/signIn', params)
      .then(response => {
        console.log(response.data)
      })
    }
  }
}
</script>
