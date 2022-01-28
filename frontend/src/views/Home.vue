<template>
  <div class="home">
    <Header></Header>
    <h1>投稿一覧</h1>
    <div class="card w-75" v-for="article in articles" :key="article">
      <div class="card-header">
        <router-link class="link" :to="{name: 'Mypage', params: {id:(Number(article.UserId))}}">{{article.Name}}さん</router-link>
      </div>
      <div class="card-body">
        <div class="card-body text-end" v-if="article.UserId == currentUserId">
          <div class="dropdown">
            <a class="dropdown-toggle" href="#" id="dropdownMenuButton" role="button" data-bs-toggle="dropdown" aria-expanded="false">
              <fa icon="ellipsis-v"/>
            </a>
            <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton" justify-content-end >
              <li>
                  <a class="dropdown-item" href="#">
                    <router-link :to="{name: 'Edit', params: {id:(Number(article.Id))}}">編集</router-link>
                  </a>
              </li>
              <li>
                  <a class="dropdown-item" href="#">
                    <button @click="deleteArticle(article)">削除</button>
                  </a>
              </li>
            </ul>
          </div>
        </div>
        <p class="card-text"><router-link class="link" :to="{name: 'Detail', params: {id:(Number(article.Id))}}">{{article.Body}}</router-link></p>
        <div v-for="tag in tags" :key="tag">
          <div v-if="article.Id == tag.Id">
            <div v-for="t in tag.Tag" :key="t">
              <span v-if="article.UserId == currentUserId">
                {{t}}
              </span>
            </div>
          </div>
        </div>
      </div>
      <div class="card-footer text-end">
        <div v-for="likesCount in likesCounts" :key="likesCount">
          <div v-if="likesCount.ArticleId == article.Id">
            <span v-for="commentCount in commentCounts" :key="commentCount">
              <span v-if="commentCount.ArticleId == article.Id"><fa icon="comment-alt" class="comment-icon"/>{{commentCount.Count}}</span>
            </span>
            <span v-for="result in results" :key="result">
              <span v-if="result.ArticleId == article.Id">
                <span @click="registerLikes(article)" v-if="result.Count">
                  <fa icon="heart" class="like-btn"/>
                  <span v-if="likesCount.ArticleId == article.Id" class="heart">
                    {{likesCount.Count}}
                  </span>
                </span>
                <span @click="deleteLikes(article)" v-else >
                  <fa icon="heart" class="unlike-btn"/>
                  <span v-if="likesCount.ArticleId == article.Id" class="heart">
                    {{likesCount.Count}}
                  </span>
                </span>
              </span>
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

<style scoped>
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

 .heart {
   margin-right: 20px;
 }
.comment-icon {
  margin-right: 5px;
}

.card {
  margin-bottom: 20px;
}

.link {
  text-decoration: none;
}
</style>
