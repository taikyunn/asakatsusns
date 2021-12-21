<template>
  <div>
    <h1>新規登録</h1>
    {{nameError}}<br />
    <label for="name">お名前:</label>
    <input type="text" name='name' v-model="name"><br />
    {{emailError}}<br />
    <label for="email">メールアドレス:</label>
    <input type="text" name='email' v-model="email"><br />
    {{passwordError}}<br />
    <label for="password">パスワード:</label>
    <input type="password" name='password' v-model="password"><br />
    <button @click="signUp">登録する</button>
    <h2>登録者一覧</h2>
    <tr>
      <th>ID:</th>
      <th>Name:</th>
    </tr>
    <tr v-for="user in users" :key="user.name">
      <td>{{user.id}}:</td>
      <td>{{user.name}}</td>
    </tr>
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
      users:[],
      nameError: "",
      emailError: "",
      passwordError: "",
      errors: "",
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
      const params = new URLSearchParams()
      params.append('name', this.name)
      params.append('email', this.email)
      params.append('password', this.password)
      axios.post('/signUp', params)
      .then(response => {
        if (response.status == 201) {
          this.nameError = response.data.Name
          this.emailError = response.data.Email
          this.passwordError = response.data.Password
        } else{
          alert("登録しました。")
        }
        this.getAllUsers()
      })
    },
  }
}
</script>
