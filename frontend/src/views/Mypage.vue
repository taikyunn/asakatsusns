<template>
  <div>
    <Header></Header>
    <div class="container">
      <div class="row">
          <div class="col-md-5 text-center">
            <Profile v-bind:id=propData></Profile>
          </div>
          <div class="col-md-6">
            <ul id="myTab" class="nav nav-tabs mb-3" role="tablist">
              <li class="nav-item" role="presentation">
                <button type="button" id="home-tab" class="nav-link active" data-bs-toggle="tab" data-bs-target="#home" role="tab" aria-controls="home" aria-selected="true">
                  投稿
                </button>
              </li>
              <li class="nav-item" role="presentation">
                <button type="button" id="profile-tab" class="nav-link" data-bs-toggle="tab" data-bs-target="#profile" role="tab" aria-controls="profile" aria-selected="false">
                  いいね
                </button>
              </li>
            </ul>
            <div id="myTabContent" class="tab-content">
              <div id="home" class="tab-pane active" role="tabpanel" aria-labelledby="home-tab">
                <div class="card" v-for="article in mypageArticle" :key="article">
                  <div class="card-header">
                    {{userName.name}}
                  </div>
                    <div class="card-body">
                      <p class="card-text">
                        {{article.body}}
                      </p>
                      <div class="card-footer text-end">
                        <span v-for="mypageCommentCount in mypageCommentCounts" :key="mypageCommentCount">
                          <span v-if="mypageCommentCount.ArticleId == article.ID">
                            <fa icon="comment-alt" class="comment-icon" />
                            {{mypageCommentCount.Count}}
                          </span>
                        </span>
                        <span v-for="result in results" :key="result">
                          <span v-if="result.ArticleId == article.ID">
                            <span @click="registerLikes(article.ID)" v-if="result.Count">
                              <fa icon="heart" class="like-btn"/>
                            </span>
                            <span @click="deleteLikes(article.ID)" v-else>
                              <fa icon="heart" class="unlike-btn"/>
                            </span>
                          </span>
                        </span>
                        <span v-for="count in countData" :key="count">
                          <span v-if="count.ArticleId == article.ID">
                            {{count.Count}}
                          </span>
                        </span>
                      </div>
                    </div>
                </div>
              </div>
              <div id="profile" class="tab-pane" role="tabpanel" aria-labelledby="profile-tab" v-for="article in mypageArticle" :key="article">
                <div class="card" v-for="likedPost in likedPosts" :key="likedPost">
                  <div class="card-header">
                    {{likedPost.Name}}
                  </div>
                  <div class="card-body">
                    <p class="card-text">
                      {{likedPost.Body}}
                    </p>
                  </div>
                  <div class="card-footer text-end">
                    <span v-for="commentCount in commentCounts" :key="commentCount">
                      <span v-if="likedPost.ArticleId == commentCount.ArticleId">
                        <fa icon="comment-alt" />
                        {{commentCount.Count}}
                      </span>
                    </span>
                    いいね数{{likedPost.Count}}
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
import Profile from './Profile.vue'

export default {
  props:["id"],
  data(){
    return{
      userInfo:[],
      sleepTime:'',
      wakeUpTime:'',
      url:'',
      fileInfo:'',
      profileDataUrl: '',
      mypageArticle: [],
      countData: '',
      results: [],
      likedPosts:[],
      userName: '',
      likedCommentCounts: '',
      propData: this.id,
      commentCounts: '',
      mypageCommentCounts: ''
    }
  },
  components: { Header , Profile},
  mounted() {
    this.getMypageArticle()
    this.checkFavoriteMypage()
    this.getCountFavoriteMypage()
    this.getLikedPost()
  },
  created() {
    const params = new URLSearchParams()
    params.append('userId', this.id)
    axios.post('getUserData', params)
    .then(response => {
      var resultUser = response.data
      this.userInfo = resultUser[0]
    })
  },
  methods: {
    getMypageArticle() {
      const params = new URLSearchParams()
      params.append('userId',this.id)
      axios.post('getMypageArticle', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error("レスポンスエラー")
        } else {
          var resultMypageArticle = response.data.mypageArticle
          this.mypageArticle = resultMypageArticle
          var resultUserName = response.data.userName
          this.userName = resultUserName[0]
          var resultCommentCounts = response.data.commentCount
          this.mypageCommentCounts = resultCommentCounts
        }
      })
    },
    getCountFavoriteMypage() {
      const params = new URLSearchParams()
      params.append('userId',this.id)
      axios.post('getCountFavoriteMypage', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error("レスポンスエラー")
        } else {
          var resultCountData = response.data
          this.countData = resultCountData
        }
      })
    },
    checkFavoriteMypage() {
      const params = new URLSearchParams()
      params.append('mypageUserId',this.id)
      params.append('visiterUserId', localStorage.getItem('userId'))
      axios.post('checkFavoriteMypage', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error("レスポンスエラー")
        } else {
          var resultCheckFavorite = response.data
          this.results = resultCheckFavorite
        }
      })
    },
    registerLikes(articleId) {
      try {
        if (localStorage.getItem('jwt') == '') {
          throw new Error('終了します')
        }
        const params = new URLSearchParams()
        params.append('articleId', articleId)
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
            this.checkFavoriteMypage()
            this.getCountFavoriteMypage()
          }
        })
      } catch {
        alert("ログインからやり直してください。")
        this.$router.push('/login')
      }
    },
    deleteLikes(articleId) {
      try {
        if (localStorage.getItem('jwt') == '') {
          throw new Error('終了します');
        }
        const params = new URLSearchParams()
        params.append('articleId', articleId)
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
            this.checkFavoriteMypage()
            this.getCountFavoriteMypage()
          }
        })
      } catch {
        alert("ログインからやり直してください。")
        this.$router.push('/login')
      }
    },
    getLikedPost() {
      const params = new URLSearchParams()
      params.append('mypageUserId', this.id)
      axios.post('getLikedPost', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error("レスポンスエラー")
        } else {
          this.likedPosts = response.data.favoritePostData
          this.commentCounts = response.data.commentCount
        }
      })
    },
  }
}
</script>

<style scoped>

.row {
  margin-top: 50px;
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

 .card {
  margin-top: 50px;
  margin-bottom: 20px;
}

.container {
  padding-top: 5rem;
}

</style>
