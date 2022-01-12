<template>
  <div>
    <h1>投稿内容</h1>
    <p v-if="apiErrors.length">
    <ul>
      <li v-for="error in apiErrors" :key="error">{{ error }}</li>
    </ul>
    </p>
    <textarea name="body" cols="70" rows="10" v-model="body"></textarea>
    <br />
    <button @click='CreateArticle'>投稿する</button>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      body:'',
      apiErrors: [],
    }
  },
  methods: {
    CreateArticle() {
      const params = new URLSearchParams
      params.append('body', this.body)
      params.append('name', localStorage.getItem('userName'))
      axios.post('/post/new', params, {
      headers: {'Authorization': `Bearer ${localStorage.getItem('jwt')}`}
      })
      .then(response => {
        if (response.status == 201) {
          if (response.data.Body != '') {
            this.apiErrors.push(response.data.Body)
          } else {
            alert("ログインからやり直してください。")
            this.$router.push('/login')
          }
          console.log(response.data.Body)
        } else if (response.status != 200) {
          console.log("エラーが発生しました。")
        } else {
          alert("投稿しました。")
          this.$router.push('/')
        }
      })
    }
  }
}
</script>

