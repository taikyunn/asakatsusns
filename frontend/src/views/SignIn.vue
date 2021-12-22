<template>
  <div>
    <h1>ログインフォーム</h1>
    {{nameError}}<br />
    <label for="name">お名前:</label>
    <input type="text" name='name' v-model="name"><br />
     {{passwordError}}<br />
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
      emailError: "",
      passwordError: "",
    }
  },
  methods: {
    signIn() {
      const params = new URLSearchParams()
      params.append('name', this.name)
      params.append('password', this.password)
      axios.post('/signIn', params)
      .then(response => {
        if (response.status == 201) {
          this.nameError = response.data.Name
          this.passwordError = response.data.Password
        } else {
          alert("ようこそ" + response.data + "さん")
          this.$router.push({
          name:'Home',
          params: {name:response.data}
        })
        }
      })
    }
  }
}
</script>
