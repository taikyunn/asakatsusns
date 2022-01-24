<template>
  <div>
    <div>
      <p>{{ArticleData.Name}}さん</p>
      <p>内容:{{ArticleData.Body}}</p>
        <span v-for="result in results" :key="result">
          <button v-if="result.Count">いいね</button>
          <button v-else>いいね解除</button>
        </span>
      <p>いいね数:{{ArticleData.LikeCount}}</p>
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
     }
   },
}
</script>
