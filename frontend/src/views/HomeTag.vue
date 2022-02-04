<template>
  <div>
    <Header />
    <div class="container">
      <div class="row">
        <div class="col-md-5">
          <SideBar />
        </div>
        <div class="col-md-6">
          <div class="card w-75">
            <div class="card-body">
              <p class="card-text tag">
                #{{tag.name}}
              </p>
              <p class="card-text text-end">
                {{count}}件
              </p>
            </div>
          </div>
          <div class="card w-75" v-for="article in articles" :key="article">
            <div class="card-header">
              <router-link class="link" :to="{name: 'Mypage', params: {id:(Number(article.UserId))}}">
                {{article.Name}}
              </router-link>
              {{article.CreatedAt}}
              <span v-if="article.UserId == currentUserId">
                <span class="dropdown">
                  <a class="dropdown" data-bs-toggle="dropdown" aria-expanded="false">
                    <fa icon="ellipsis-v" class="ellipsis" />
                  </a>
                  <ul class="dropdown-menu" justify-content-end >
                    <li>
                        <a class="dropdown-item" href="#">
                          <router-link class="btn btn-warning" :to="{name: 'Edit', params: {id:(Number(article.Id))}}">
                            編集
                          </router-link>
                        </a>
                    </li>
                    <li>
                      <a class="dropdown-item" href="#">
                        <button class="btn btn-warning" @click="deleteArticle(article)">
                          削除
                        </button>
                      </a>
                    </li>
                  </ul>
                </span>
              </span>
            </div>
            <div class="card-body">
              <p class="card-text">
                <router-link class="link" :to="{name: 'Detail', params: {id:(Number(article.Id))}}">
                  {{article.Body}}
                </router-link>
              </p>
              <div v-for="tag in tags" :key="tag">
                <div v-if="article.Id == tag.ArticleId">
                  <span v-if="article.UserId == currentUserId" >
                    <router-link class="tag" :to="{name: 'HomeTag', params: {id:(Number(tag.Key))}}">
                      #{{tag.Value}}
                    </router-link>
                  </span>
                </div>
              </div>
            </div>
            <div class="card-footer text-end">
              <div v-for="likesCount in likesCounts" :key="likesCount">
                <div v-if="likesCount.ArticleId == article.Id">
                  <span v-for="commentCount in commentCounts" :key="commentCount">
                    <span v-if="commentCount.ArticleId == article.Id">
                      <fa icon="comment-alt" class="comment-icon" />
                      {{commentCount.Count}}
                    </span>
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
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Header from './Header.vue'
import SideBar from './SideBar.vue'

export default {
  props:["id"],
  components: { Header, SideBar},
  data() {
    return {
      currentUserId: localStorage.getItem('userId'),
      articles: [],
      tags: [],
      likesCounts: [],
      tag: '',
      count:'',
    }
  },
  mounted() {
    this.getTagArticles()
    this.countFavorites()
    this.checkFavorite()
    this.countComments()
  },
  methods: {
    getTagArticles() {
      const params = new URLSearchParams()
      params.append('tagId', this.id)
      axios.post('getTagArticles', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          var resultArticles = response.data.article
          this.articles = resultArticles
          var resultTags = response.data.tag
          this.tags = resultTags
          var resultTag = response.data.topTag
          this.tag = resultTag[0]
          var resultCount = response.data.count
          this.count = resultCount
        }
      })
    },
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
    },
  },
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
  margin-bottom: 40px;
}

.link {
  text-decoration: none;
  text-align: left;
  color:black;
}

.ellipsis {
  float: right;
}

.dropdown-toggle {
  color: black;
}

.tag {
  color: green;
  white-space: nowrap;
  text-decoration: none;
}

.container {
  padding-top: 5rem;
}

</style>
