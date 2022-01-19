<template>
  <div class="home">
    <h1>{{ msg }}</h1>
    <h1>投稿一覧</h1>
    <tr>
      <th>投稿者:</th>
      <th>投稿:</th>
    </tr>
    <tr v-for="article in articles" :key="article">
        <td>{{article.Name}}:</td>
        <td>{{article.Body}}</td>
        <td v-if="article.Tag != null">{{article.Tag}}</td>
        <td v-if="article.UserId == currentUserId">
          <router-link :to="{name: 'Edit', params: {id:(Number(article.Id))}}">編集</router-link>
        </td>
        <td v-if="article.UserId == currentUserId">
          <button @click="deleteArticle(article)">
            削除
          </button>
        </td>
    </tr>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      msg: '朝活SNS',
      currentUserId: localStorage.getItem('userId'),
      articles:[],
    }
  },
  created() {
    axios.get('getAllArticles')
    .then(response => {
      if (response.status != 200) {
        throw new Error('レスポンスエラー')
      } else {
        console.log(response.data)
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
