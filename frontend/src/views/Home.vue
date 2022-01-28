<template>
  <div class="home">
    <Header></Header>
    <h1>投稿一覧</h1>
    <div class="card w-75" v-for="article in articles" :key="article">
      <div class="card-header">
        <router-link :to="{name: 'Mypage', params: {id:(Number(article.UserId))}}">{{article.Name}}さん</router-link>
      </div>
      <div class="card-body">
        <p class="card-text"><router-link :to="{name: 'Detail', params: {id:(Number(article.Id))}}">{{article.Body}}</router-link></p>
        <div v-for="tag in tags" :key="tag">
          <div v-if="article.Id == tag.Id">
            <div v-for="t in tag.Tag" :key="t">
              <span v-if="article.UserId == currentUserId">
                {{t}}
              </span>
            </div>
          </div>
        </div>
        <div v-if="article.UserId == currentUserId">
          <router-link :to="{name: 'Edit', params: {id:(Number(article.Id))}}">編集</router-link>
          <button @click="deleteArticle(article)">削除</button>
        </div>
        <div v-for="result in results" :key="result">
          <div v-if="result.ArticleId == article.Id">
            <button @click="registerLikes(article)" v-if="result.Count">いいね</button>
            <button @click="deleteLikes(article)" v-else >いいね解除</button>
          </div>
        </div>
      </div>
      <div class="card-footer text-end">
        <div v-for="likesCount in likesCounts" :key="likesCount">
          <div v-if="likesCount.ArticleId == article.Id">
            <span v-if="likesCount.ArticleId == article.Id">いいね数:{{likesCount.Count}}</span>
            <span v-for="commentCount in commentCounts" :key="commentCount">
              <span v-if="commentCount.ArticleId == article.Id">コメント数:{{commentCount.Count}}</span>
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Header from './Header.vue'

export default {
  data() {
    return {
      msg: '朝活SNS',
      currentUserId: localStorage.getItem('userId'),
      articles:[],
      tags:[],
      likesCounts: [],
      results:[],
      commentCounts:[],
    }
  },
  components: { Header },
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
    this.checkFavorite()
    this.countComments()
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
      try {
        if (localStorage.getItem('jwt') == '') {
          throw new Error('終了します');
        }
        const params = new URLSearchParams()
        params.append('articleId', article.Id)
        params.append('userId', localStorage.getItem('userId'))
        const config = {
          headers: {
            'Authorization': `Bearer ${localStorage.getItem('jwt')}`,
          }
        }
        axios.post('/post/registerLikes', params, config)
        .then(response => {
          if (response.status == 201) {
            if (response.data.Body != '') {
              alert("ログインからやり直してください。")
              this.$router.push('/login')
            }
          } else if (response.status != 200) {
            throw new Error('レスポンスエラー')
          } else {
            this.countFavorites()
            this.checkFavorite()
          }
        })
      } catch {
        alert("ログインからやり直してください。")
        this.$router.push('/login')
      }
    },
    deleteLikes(article) {
      try {
        if (localStorage.getItem('jwt') == '') {
          throw new Error('終了します');
        }
        const params = new URLSearchParams()
        params.append('articleId', article.Id)
        params.append('userId', localStorage.getItem('userId'))
        const config = {
          headers: {
            'Authorization': `Bearer ${localStorage.getItem('jwt')}`,
          }
        }
        axios.post('/post/deleteLikes', params, config)
        .then(response => {
          if (response.status == 201) {
            if (response.data.Body != '') {
              alert("ログインからやり直してください。")
              this.$router.push('/login')
            }
          } else if (response.status != 200) {
            throw new Error('レスポンスエラー')
          } else {
            this.countFavorites()
            this.checkFavorite()
          }
        })
      } catch {
        alert("ログインからやり直してください。")
        this.$router.push('/login')
      }
    },
    countFavorites() {
      axios.get('getCountFavorites')
      .then(response => {
        var resultLikesCount = response.data
        this.likesCounts = resultLikesCount
      })
    },
    checkFavorite() {
      const params = new URLSearchParams()
      params.append("userId", localStorage.getItem('userId'))
      axios.post("checkFavorite", params)
      .then(response => {
        var resultCheckFavorite = response.data
        this.results = resultCheckFavorite
      })
    },
    countComments() {
      axios.get('getCountComments')
      .then(response => {
        var resultCommentCount = response.data
        this.commentCounts = resultCommentCount
      })
    }
  }
}
</script>

