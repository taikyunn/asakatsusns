<template>
  <div>
    <Header></Header>
    <div class="text-center">
      <div>
        <p>{{ArticleData.Name}}さん</p>
        <p>内容:{{ArticleData.Body}}</p>
          <span v-for="result in results" :key="result">
            <button @click="registerLikes()" v-if="result.Count">いいね</button>
            <button @click="deleteLikes()" v-else>いいね解除</button>
          </span>
        <p>いいね数:{{count}}</p>
        <span v-for="tag in ArticleData.Tags" :key="tag">
          {{tag}}&nbsp;
        </span>
      </div>
      <div>
        <h2>コメントを追加する</h2>
        <textarea name="body" cols="70" rows="3" v-model="comment"></textarea>
        <div>
          <button @click='insertComment()'>コメントする</button>
        </div>
      </div>
      <div>
        <h2>コメント一覧</h2>
        <p v-for="commentData in ArticleData.Comments" :key="commentData">
          {{commentData.Name}}さん:{{commentData.Comment}}
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Header from './Header.vue'

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
      comment:'',
    }
  },
  components: { Header },
  mounted() {
    this.process()
  },
  methods: {
    async process() {
      await Promise.all([
        this.countFavorites(),
        this.checkFavorite(),
      ])
      await this.getArticleDetail()
    },
    getArticleDetail() {
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
    registerLikes() {
      try {
        if (localStorage.getItem('jwt') == '') {
          throw new Error('終了します');
        }
        const params = new URLSearchParams()
        params.append('articleId', this.id)
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
    deleteLikes() {
      try {
        if (localStorage.getItem('jwt') == '') {
          throw new Error('終了します');
        }
        const params = new URLSearchParams()
        params.append('articleId', this.id)
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
      const params = new URLSearchParams()
      params.append('articleId', this.id)
      axios.post('getOneCountFavorites', params)
      .then(response => {
        var resultCountData = response.data
        this.count = resultCountData
      })
    },
    insertComment() {
      const params = new URLSearchParams()
      params.append('comment', this.comment)
      params.append('articleId', this.id)
      params.append('userId', localStorage.getItem('userId'))
      const config = {
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`,
        }
      }
      axios.post('/post/insertComment', params, config)
      .then(response => {
        if (response.status == 201) {
          if (response.data.Body != '') {
            alert("ログインからやり直してください。")
            this.$router.push('/login')
          }
        } else if (response.status != 200){
          throw new Error('レスポンスエラー')
        } else {
          this.$router.go({path: this.$router.currentRoute.path, force: true})
          this.getArticleDetail()
        }
      })
    }
  },
}
</script>

<style scoped>
.text-center {
  padding-top: 5rem;
}
</style>
