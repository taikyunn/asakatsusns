<template>
  <div>
    <div>
      <p>{{ArticleData.Name}}さん</p>
      <p>内容:{{ArticleData.Body}}</p>
        <span v-for="result in results" :key="result">
          <button @click="registerLikes(ArticleData.ArticleId, ArticleData.UserId)" v-if="result.Count">いいね</button>
          <button v-else>いいね解除</button>
        </span>
      <p>いいね数:{{count}}</p>
      <span v-for="tag in ArticleData.Tags" :key="tag">
        {{tag}}&nbsp;
      </span>
    </div>
    <div>
      <h3>コメントを追加する</h3>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  props: {
    id: {
      type: String,
      default : '',
    },
  },
  data () {
    return {
      ArticleData: [],
      results:[],
      count:'',
    }
  },
  created() {
    const params = new URLSearchParams()
    params.append('articleId', this.id)
    axios.post('getArticleDetail', params)
    .then(response => {
      if (response.status != 200) {
        throw new Error('レスポンスエラー')
      } else {
        var ArticleData = response.data
        this.ArticleData = ArticleData[0]
      }
    })
  },
   mounted() {
     this.countFavorites()
     this.checkFavorite()
   },
   methods: {
     checkFavorite() {
       const params = new URLSearchParams()
       params.append('articleId', this.id)
       params.append('userId', localStorage.getItem('userId'))
      axios.post("checkFavoriteByArticleId", params)
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          var resultCheckFavorite = response.data
          this.results = resultCheckFavorite
        }
      })
     },
     registerLikes(articleId, userId) {
      try {
        if (localStorage.getItem('jwt') == '') {
          throw new Error('終了します');
        }
        const params = new URLSearchParams()
        params.append('articleId', articleId)
        params.append('userId', userId)
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
    countFavorites() {
      const params = new URLSearchParams()
      params.append('articleId', this.id)
      axios.post('getOneCountFavorites', params)
      .then(response => {
        var resultCountData = response.data
        this.count = resultCountData
      })
    },
  },
}
</script>
