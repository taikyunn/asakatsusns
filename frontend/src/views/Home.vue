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
        <div v-for="tag in tags" :key="tag">
          <div v-if="article.Id == tag.Id">
            <div v-for="t in tag.Tag" :key="t">
              <span v-if="article.UserId == currentUserId">
                {{t}}
              </span>
            </div>
          </div>
        </div>
        <td v-if="article.UserId == currentUserId">
          <router-link :to="{name: 'Edit', params: {id:(Number(article.Id))}}">編集</router-link>
        </td>
        <td v-if="article.UserId == currentUserId">
          <button @click="deleteArticle(article)">
            削除
          </button>
        </td>
        <td>
          <button @click="deleteLikes(article)" v-if="isLiked">いいね解除</button>
          <button @click="registerLikes(article)" v-else >いいね</button>
          <span v-for="count in counts" :key="count">
            <span v-if="count.ArticleId == article.Id">いいね数:{{count.Count}}</span>
          </span>
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
      tags:[],
      isLiked: false,
      counts: '',
    }
  },
  created() {
    axios.get('getAllArticles')
    .then(response => {
      if (response.status != 200) {
        throw new Error('レスポンスエラー')
      } else {
        var resultArticles = response.data.article
        this.articles = resultArticles
        var resultTags = response.data.tag
        this.tags = resultTags
      }
    })
  },
  mounted () {
    this.countFavorites()
  },
  methods:{
    deleteArticle(article) {
      confirm('削除してもよろしいですか。')
      const params = new URLSearchParams()
      params.append('articleId', article.Id)
      axios.post('deleteArticle', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          this.$router.go({path: this.$router.currentRoute.path, force: true})
          alert('削除しました。')
        }
      })
    },
    registerLikes(article) {
      const params = new URLSearchParams()
      params.append('articleId', article.Id)
      params.append('userId', article.UserId)
      axios.post('registerLikes', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          // this.isLiked = true
        }
      })
    },
    deleteLikes(article) {
      const params = new URLSearchParams()
      params.append('articleId', article.Id)
      params.append('userId', article.UserId)
      axios.post('deleteLikes', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          // this.isLiked = false
        }
      })
    },
    countFavorites() {
      axios.get('getCountFavorites')
      .then(response => {
        var resultCountData = response.data
        this.counts = resultCountData
      })
    }
  }
}
</script>
