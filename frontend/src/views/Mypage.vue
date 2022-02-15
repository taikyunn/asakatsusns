<template>
  <div>
    <Header />
    <div class="container">
      <div class="row">
          <div class="col-md-5 text-center">
            <Profile v-bind:id=propData />
          </div>
          <div class="col-md-6">
            <ul id="myTab" class="nav nav-tabs mb-3" role="tablist">
              <li class="nav-item" role="presentation">
                <button type="button" id="home-tab" class="nav-link active btn btn-outline-warning" data-bs-toggle="tab" data-bs-target="#home" role="tab" aria-controls="home" aria-selected="true">
                  投稿
                </button>
              </li>
              <li class="nav-item" role="presentation">
                <button type="button" id="profile-tab" class="nav-link btn btn-outline-warning" data-bs-toggle="tab" data-bs-target="#profile" role="tab" aria-controls="profile" aria-selected="false">
                  いいね
                </button>
              </li>
            </ul>
            <div id="myTabContent" class="tab-content">
              <div id="home" class="tab-pane active" role="tabpanel" aria-labelledby="home-tab">
                <div class="card" v-for="article in mypageArticle" :key="article">
                    <div class="card-body">
                      <span v-if="userImage">
                        <img :src="userImage" class="circle" />
                      </span>
                      <span v-else>
                        <img :src="defaultImage" class="circle" />
                      </span>
                      <span class="link">
                        {{userName.name}}
                      </span>
                      <p class="card-text">
                        {{article.body}}
                      </p>
                      <div class="card-footer text-end footer">
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
                <p v-if="likedPosts.length == 0">
                  いいねした投稿がありません。
                </p>
                <div class="card" v-for="likedPost in likedPosts" :key="likedPost" v-else>
                  <div class="card-body">
                    <span v-if="likedPost.Image">
                      <img :src="likedPost.Image" class="circle" />
                    </span>
                    <span v-else>
                      <img :src="defaultImage" class="circle" />
                    </span>
                    <span class="link">
                      {{likedPost.Name}}
                    </span>
                    <p class="card-text">
                      {{likedPost.Body}}
                    </p>
                  </div>
                  <div class="card-footer text-end footer">
                    <span v-for="commentCount in commentCounts" :key="commentCount">
                      <span v-if="likedPost.ArticleId == commentCount.ArticleId">
                        <fa icon="comment-alt" />
                        {{commentCount.Count}}
                      </span>
                    </span>
                    <span v-for="favoriteLikedPostCount in favoriteLikedPostCounts" :key="favoriteLikedPostCount">
                      <span v-if="favoriteLikedPostCount.ArticleId == likedPost.ArticleId">
                        <span v-if="favoriteLikedPostCount.Count">
                          <fa icon="heart" class="like-btn"/>
                        </span>
                        <span v-else>
                          <fa icon="heart" class="unlike-btn"/>
                        </span>
                      </span>
                    </span>
                    <span v-for="countData in likedPostCountData" :key="countData">
                      <span v-if="countData.ArticleId == likedPost.ArticleId">
                        {{countData.Count}}
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
import Profile from './Profile.vue'

export default {
  props:["id"],
  data(){
    return{
      userInfo: [],
      sleepTime: '',
      wakeUpTime: '',
      url: '',
      fileInfo: '',
      profileDataUrl: '',
      mypageArticle: [],
      countData: '',
      results: [],
      likedPosts:[],
      userName: '',
      userImage: '',
      defaultImage: require('@/images/default.png'),
      likedCommentCounts: '',
      propData: this.id,
      commentCounts: '',
      mypageCommentCounts: '',
      favoriteLikedPostCounts: '',
      likedPostCountData: '',
    }
  },
  components: { Header , Profile},
  mounted() {
    this.process()
  },
  methods: {
    async process() {
      await Promise.all([
        this.getUserData(),
        this.getMypageArticle(),
      ])
      await this.checkFavoriteMypage()
      await this.getCountFavoriteMypage()
      await this.getLikedPost()
      await this.checkFavoriteLikedPost()
      await this.getCountFavoriteLikedPost()
    },
    getUserData(){
    const params = new URLSearchParams()
      params.append('userId', this.id)
      axios.post('getUserData', params)
      .then(response => {
        var resultUser = response.data
        this.userInfo = resultUser[0]
      })
    },
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
          if (resultUserName[0].ProfileImagePath != '') {
            let url = process.env.VUE_APP_DATA_URL + resultUserName[0].ProfileImagePath
            axios.get(url,{responseType: "blob"})
            .then(response => {
              let blob = new Blob([response.data])
              this.userImage = URL.createObjectURL(blob)
            })
          }
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
          var likedPosts = response.data.favoritePostData
          for (let i = 0; i < likedPosts.length; i++) {
            if (likedPosts[i].ProfileImagePath == '') {
              continue;
            }
            let url = process.env.VUE_APP_DATA_URL + likedPosts[i].ProfileImagePath
            axios.get(url,{responseType: "blob"})
            .then(response => {
              let blob = new Blob([response.data])
              this.likedPosts.splice(i, 1, {
                Body: likedPosts[i].Body,
                Id: likedPosts[i].Id,
                Name: likedPosts[i].Name,
                ProfileImagePath: likedPosts[i].ProfileImagePath,
                Image: URL.createObjectURL(blob),
                UserId: likedPosts[i].UserId
              })
            })
          }
        }
      })
    },
    checkFavoriteLikedPost() {
      const params = new URLSearchParams()
      params.append('mypageUserId',this.id)
      params.append('visiterUserId', localStorage.getItem('userId'))
      axios.post('checkFavoriteLikedPost', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error("レスポンスエラー")
        } else {
          var resultCheckFavoriteLikedPost = response.data
          this.favoriteLikedPostCounts = resultCheckFavoriteLikedPost
        }
      })
    },
    getCountFavoriteLikedPost() {
      const params = new URLSearchParams()
      params.append('mypageUserId', this.id)
      axios.post('getCountFavoriteLikedPost', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error("レスポンスエラー")
        } else {
          var resultLikedPostCountData = response.data
          this.likedPostCountData = resultLikedPostCountData
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

.footer {
  background-color:white;
}

.link {
  margin-left: 1rem;
}

.circle {
  border-radius: 50%;
  width: 50px;
  height: 50px;
  object-fit: cover;
}
</style>
