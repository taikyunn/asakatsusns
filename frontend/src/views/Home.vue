<template>
  <div class="home">
    <Header />
    <div class="container">
      <div class="row">
        <div class="col-md-5">
          <SideBar />
        </div>
        <div class="col-md-6">
          <div class="card w-75" v-for="article in articles" :key="article">
            <div class="card-body">
              <span v-if="article.Image">
                <img :src="article.Image" class="circle" />
              </span>
              <span v-else>
                <img :src="defaultImage" class="circle" />
              </span>
              <router-link class="link" :to="{name: 'Mypage', params: {id:(Number(article.UserId))}}">
                {{article.Name}}
              </router-link>
              <span class="time">
                {{article.UpdatedAt}}
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
              </span>
              <p class="card-text">
                <router-link class="link" :to="{name: 'Detail', params: {id:(Number(article.Id))}}">
                  {{article.Body}}
                </router-link>
              </p>
              <div v-for="tag in tags" :key="tag">
                <div v-if="article.Id == tag.ArticleId">
                  <!-- <span v-if="article.UserId == currentUserId" > -->
                    <router-link class="tag border border-success rounded" :to="{name: 'HomeTag', params: {id:(Number(tag.Key))}}">
                      {{tag.Value}}
                    </router-link>
                  <!-- </span> -->
                </div>
              </div>
            </div>
            <div class="card-footer text-end footer">
              <div v-for="likesCount in likesCounts" :key="likesCount">
                <div v-if="likesCount.ArticleId == article.Id">
                  <span v-for="commentCount in commentCounts" :key="commentCount">
                    <span v-if="commentCount.ArticleId == article.Id">
                      <fa icon="comment-alt" />
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
          <InfiniteLoading />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Header from './Header.vue'
import SideBar from './SideBar.vue'
import InfiniteLoading from './InfiniteLoading.vue'

export default {
  data() {
    return {
      currentUserId: localStorage.getItem('userId'),
      articles: [],
      tags: [],
      likesCounts: [],
      results: [],
      commentCounts: [],
      defaultImage: require('@/images/default.png'),
    }
  },
  components: {
    Header,
    SideBar,
    InfiniteLoading,
  },
  mounted () {
    this.process()
  },
  methods:{
    async process() {
      await this.getAllArticles()
      await Promise.all([
        this.countFavorites(),
        this.checkFavorite(),
        this.countComments(),
      ]);
    },
    async getAllArticles() {
    const response = await axios.get('getAllArticles')
    if (response.status != 200) {
      throw new Error('レスポンスエラー')
    } else {
      var resultArticles = response.data.article
      this.articles = resultArticles
      var resultTags = response.data.tag
      this.tags = resultTags
      for (let i = 0; i < resultArticles.length; i++) {
        if (resultArticles[i].ProfileImagePath == '') {
          continue;
        }
        let url = process.env.VUE_APP_DATA_URL + resultArticles[i].ProfileImagePath
        axios.get(url,{responseType: "blob"})
        .then(response => {
          let blob = new Blob([response.data])
          resultArticles[i].Image = URL.createObjectURL(blob)
        })
      }
      this.articles = resultArticles
    }
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
          alert('削除しました。')
          this.$router.go({path: this.$router.currentRoute.path, force: true})
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

.card {
  margin-top: 40px;
}

.link {
  text-decoration: none;
  text-align: left;
  color:black;
  margin-left: 1rem;
}

.ellipsis {
  color:gray
}

.dropdown-toggle {
  color: black;
}

.tag {
  color: green;
  white-space: nowrap;
  text-decoration: none;
  padding: 2px;
  float: left;
  margin-right: 1rem;
}

.container {
  padding-top: 5rem;
}

.img {
  float: left;
}

.time {
  float: right;
}

.circle {
  border-radius: 50%;
  width: 50px;
  height: 50px;
  object-fit: cover;
}

.footer {
  background-color:white;
}

</style>
