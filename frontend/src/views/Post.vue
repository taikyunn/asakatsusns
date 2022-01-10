<template>
  <div>
    <h1>投稿内容</h1>
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
      .then(responce => {
        console.log(responce)
      })
    }
  }
}
</script>

