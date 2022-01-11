<template>
  <div class="home">
    <h1>{{ msg }}</h1>
    <h1>ユーザー名：{{ name }}</h1>
    <router-link to="/post">投稿する</router-link>
    <h1>投稿一覧</h1>
    <tr>
      <th>ID:</th>
      <th>投稿:</th>
      <th>編集</th>
      <th>削除</th>
    </tr>
    <tr v-for="article in articles" :key="article">
      <td>{{article.ID}}:</td>
      <td>{{article.body}}</td>
      <td>
        <router-link :to="{name: 'Edit', params: {id:(Number(article.ID))},query: { id: article.ID }}">編集</router-link>
      </td>
      <td>
        <button @click="deleteArticle(article)">
          削除
        </button>
      </td>
    </tr>
  </div>
</template>

<script>
import firebase from 'firebase/app'
import "firebase/auth"
import axios from 'axios'

export default {
  data() {
    return {
      msg: 'Welcome to Your Vue.js App',
      email:firebase.auth().currentUser.email,
      name:localStorage.getItem('userName'),
      articles:[],
    }
  },
  created() {
    axios.get('getAllArticles')
    .then(response => {
      if (response.status != 200) {
        throw new Error('レスポンスエラー')
      } else {
        var resultArticles = response.data
        this.articles = resultArticles
      }
    })
  },
  methods:{
    deleteArticle(article) {
      confirm('削除してもよろしいですか。')
      const params = new URLSearchParams()
      params.append('articleId', article.ID)
      axios.post('deleteArticle', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          console.log(response)
        }
      })
    }
  }
}
</script>
