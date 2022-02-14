<template>
  <div>
    <Header />
    <div class="container mt-4">
      <div class="row justify-content-center">
        <div class="col-md-8">
          <div class="card w-75">
            <div class="card-body">
              {{ArticleData.Name}}
              <span class="time">
                {{ArticleData.UpdatedAt}}
              </span>
              <p class="card-text">
                {{ArticleData.Body}}
              </p>
            </div>
            <div class="card-footer text-end footer">
              <fa icon="comment-alt" />
              {{ArticleData.Count}}
              <span v-for="result in results" :key="result">
                <span @click="registerLikes()" v-if="result.Count">
                <fa icon="heart" class="like-btn"/>
                {{count}}
                </span>
                <span @click="deleteLikes()" v-else>
                  <fa icon="heart" class="unlike-btn"/>
                  {{count}}
                </span>
              </span>
            </div>
            <span v-for="tag in ArticleData.Tags" :key="tag">
              {{tag}}&nbsp;
            </span>
          </div>
          <div class="mb-3 comment">
            <div class="card w-75" v-for="commentData in ArticleData.Comments" :key="commentData">
              <div class="card-body">
                {{commentData.Name}}
                <span class="time">
                  {{commentData.UpdatedAt}}
                </span>
                <div class="card-text card-text-comment">
                {{commentData.Comment}}
                </div>
              </div>
            </div>
            <div class="card w-75" v-if="currentUserName != null">
              <div class="card-body">
                <div class="card-text">
                  <label for="comment-area" class="form-label">
                    {{currentUserName}}
                  </label>
                  <textarea name="body" class="form-control" id="comment-area" rows="3" v-model="comment"></textarea>
                </div>
                <div class="submit">
                  <button @click='insertComment()' class="btn btn-outline-warning">コメントする</button>
                </div>
              </div>
            </div>
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
  props: {
    id: {
      type: String,
      default : '',
    },
  },
  data () {
    return {
      ArticleData: [],
      results: [],
      count: '',
      comment: '',
      currentUserName: localStorage.getItem('userName'),
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
          var resultArticle = response.data
          this.ArticleData = resultArticle[0]
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
.container {
  padding-top: 5rem;
}

.like-btn {
  width: 18px;
  height: 18px;
  font-size: 25px;
  color: #808080;
  margin-left: 20px;
  margin-right: 5px;
}

.unlike-btn {
  width: 18px;
  height: 18px;
  font-size: 25px;
  color: #e54747;
  margin-left: 20px;
  margin-right: 5px;
}

.comment {
  padding-top: 5rem;
}

.time, .submit {
  float: right;
}

.card-text-comment, .submit {
  padding-top: 1rem;
}

.footer {
  background-color:white;
}

</style>
