<template>
  <div>
    <h1>プロフィール</h1>
    <div>
      <img v-if="profileDataUrl" :src="profileDataUrl" width="100" />
      <p v-if="url">
        <span style="position:absolute" @click="deletePreview" width="1px" height="1px">X</span>
        <img :src="url" width="100" />
      </p>
      <p>
        <input type="file" ref="preview" @change="uploadFile" accept="image/jpeg, image/png">
        <button v-on:click="fileUpload()">アップロード</button>
      </p>
    </div>
    <div>
      <p>
        <router-link :to="{name: 'Follow', params: {id:(Number(this.id))}}">{{followData.FollowerCount}} フォロー</router-link>
        <router-link :to="{name: 'Follower', params: {id:(Number(this.id))}}">{{followData.FollowCount}} フォロワー</router-link>
      </p>
      <div v-if="!isMyOwnPage">
        <button v-if="!isFollowedBy" @click="registerFollow">フォローする</button>
        <button v-else @click="deleteFollow">フォロー中</button>
      </div>
    </div>
    <div>
      <p v-if="isEdit">
        <span v-if="!editName" class="border p-2  bg-light" v-on:click="doEditName" >{{userInfo.name}}*クリックで編集*</span>
        <span v-else >
          <input type="text" v-model="userInfo.name" v-on:blur="editName = false" v-focus>
          <button @click="updateName">登録</button>
        </span>
      </p>
      <p v-else>
        アカウント名:{{userInfo.name}}
      </p>
    </div>
    <div>
      <p v-if="isEdit">
        寝る時間:
        <span v-if="!editSleepTime" class="border p-2  bg-light" v-on:click="doEditSleepTime">{{userInfo.SleepTime}}*クリックで編集*</span>
        <span v-else >
          <input type="time" v-model="userInfo.SleepTime" v-on:blur="editSleepTime = false" v-focus>
          <button @click="updateSleepTime">登録</button>
        </span>
      </p>
      <p v-else>
        <span v-if="userInfo.SleepTime != null">
          寝る時間:{{userInfo.SleepTime}}
        </span>
        <span v-else>
          寝る時間:設定されていません。
        </span>
      </p>
    </div>
    <div>
      <p v-if="isEdit">
        起きる時間設定:
        <span v-if="!editWakeUpTime" class="border p-2  bg-light" v-on:click="doEditWakeUpTime">{{userInfo.WakeUpTime}}*クリックで編集*</span>
        <span v-else >
          <input type="time" v-model="userInfo.WakeUpTime" v-on:blur="editWakeUpTime = false" v-focus>
          <button @click="updateWakeUpTime">登録</button>
        </span>
      </p>
      <p v-else>
        <span v-if="userInfo.WakeUpTime != null">
          起きる時間:{{userInfo.WakeUpTime}}
        </span>
        <span v-else>
          起きる時間:設定されていません
        </span>
      </p>
    </div>
    <div>
      <h2>投稿</h2>
      <p v-for="article in mypageArticle" :key="article">
        {{article.body}}
        <span v-for="count in countData" :key="count">
          <span v-if="count.ArticleId == article.ID">いいね数:{{count.Count}}</span>
        </span>
        <span v-for="result in results" :key="result">
          <span v-if="result.ArticleId == article.ID">
            <button @click="registerLikes(article.ID)" v-if="result.Count">いいね</button>
            <button @click="deleteLikes(article.ID)" v-else>いいね解除</button>
          </span>
        </span>
      </p>
    </div>
    <div>
      <h2>いいね</h2>
      <p v-for="likedPost in likedPosts" :key="likedPost">
        {{likedPost.Name}}:
        {{likedPost.Body}}
        いいね数:{{likedPost.Count}}
      </p>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

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
      editName: false,
      editSleepTime: false,
      editWakeUpTime: false,
      isFollowedBy: false,
      isMyOwnPage: false,
      isEdit: false,
      followData: [],
      mypageArticle: [],
      countData: '',
      results: [],
      likedPosts:[],
    }
  },
  directives: {
    focus: {
      // ディレクティブ定義
      inserted: function (el) {
          el.focus();
      }
    }
  },
  mounted() {
    this.getUserProfile()
    this.checkFollow()
    this.checkMyOwnPage()
    this.checkIsEdit()
    this.getFollowData()
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
    getUserProfile() {
      const params = new URLSearchParams()
      params.append('userId', this.id)
      axios.post('getUserProfile', params, {responseType: "blob"})
      .then(response => {
        const blob = new Blob([response.data])
        this.profileDataUrl = URL.createObjectURL(blob);
      })
    },
    checkFollow() {
      const params = new URLSearchParams()
      params.append('follower_id',localStorage.getItem('userId'))
      params.append('followed_id',this.id)
      axios.post("checkFollow", params)
      .then(response => {
        var followResult = response.data
        this.isFollowedBy = followResult
      })
    },
    checkMyOwnPage() {
      if (this.id == localStorage.getItem('userId'))
      this.isMyOwnPage = true
    },
    checkIsEdit() {
      if (this.id == localStorage.getItem('userId')) {
        this.isEdit = true
      }
    },
    registerSleepTime() {
      const params = new URLSearchParams()
      params.append('userId', this.id)
      params.append('sleepTime', this.sleepTime)
      axios.post("/registerSleepTime", params)
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          alert('登録しました。')
          this.$router.go({path: this.$router.currentRoute.path, force: true})
        }
      })
    },
    registerWakeUpTime() {
      const params = new URLSearchParams()
      params.append('userId', this.id)
      params.append('wakeUpTime', this.wakeUpTime)
      axios.post("/registerWakeUpTime", params)
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          alert('登録しました。')
          this.$router.go({path: this.$router.currentRoute.path, force: true})
        }
      })
    },
    uploadFile() {
      const file = this.$refs.preview.files[0];
      this.url = URL.createObjectURL(file)
      this.fileInfo = this.$refs.preview.files[0]
      this.profileDataUrl = ""
      this.$refs.preview.value = "";
    },
    deletePreview(){
      this.url = ''
      URL.revokeObjectURL(this.url);
      this.mounted()
    },
    fileUpload() {
      const params = new FormData()
      params.append('file', this.fileInfo)
      params.append('userId', this.id)
      axios.post('/fileUpload', params, {headers: {'Content-Type': 'multipart/form-data'}})
      .then(response  => {
        if (response.status != 200) {
          throw new Error()
        } else {
          alert('登録しました。')
          this.$router.go({path: this.$router.currentRoute.path, force: true})
        }
      })
    },
    doEditName() {
      this.editName = true
    },
    updateName() {
      const params = new URLSearchParams()
      params.append('userId', this.id)
      params.append('name', this.userInfo.name)
      axios.post('editUserName', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error()
        } else {
          alert('登録しました。')
          this.$router.go({path: this.$router.currentRoute.path, force: true})
        }
      })
    },
    doEditSleepTime() {
      this.editSleepTime = true
    },
    updateSleepTime() {
      const params = new URLSearchParams()
      params.append('sleepTime', this.userInfo.SleepTime)
      params.append('userId', this.id)
      axios.post("updateSleepTime", params)
      .then(response => {
        if (response.status != 200) {
          throw new Error("レスポンスエラー")
        } else {
          alert('登録しました。')
          this.$router.go({path: this.$router.currentRoute.path, force: true})
        }
      })
    },
    doEditWakeUpTime() {
      this.editWakeUpTime = true
    },
    updateWakeUpTime(){
      const params = new URLSearchParams()
      params.append('wakeUpTime', this.userInfo.WakeUpTime)
      params.append('userId', this.id)
      axios.post("updateWakeUpTime", params)
      .then(response => {
        if (response.status != 200) {
          throw new Error("レスポンスエラー")
        } else {
          alert('登録しました。')
          this.$router.go({path: this.$router.currentRoute.path, force: true})
        }
      })
    },
    registerFollow() {
      try {
        if (localStorage.getItem('jwt') == '') {
          throw new Error('終了します');
        }
        const params = new URLSearchParams()
        params.append('follower_id', localStorage.getItem('userId'))
        params.append('followed_id',this.id)
        axios.post("registerFollow", params)
        .then(response => {
          if (response.status != 200) {
            throw new Error("レスポンスエラー")
          } else {
            this.checkFollow()
            this.getFollowData()
          }
        })
      } catch {
        alert("ログインからやり直してください。")
        this.$router.push('/login')
      }
    },
    deleteFollow() {
      try {
        if (localStorage.getItem('jwt') == '') {
          throw new Error('終了します');
        }
        const params = new URLSearchParams()
        params.append('follower_id', localStorage.getItem('userId'))
        params.append('followed_id',this.id)
        axios.post('deleteFollow', params)
        .then(response => {
          if (response.status != 200) {
            throw new Error("レスポンスエラー")
          } else {
            this.checkFollow()
            this.getFollowData()
          }
        })
      } catch {
        alert("ログインからやり直してください。")
        this.$router.push('/login')
      }
    },
    getFollowData() {
      const params = new URLSearchParams()
      params.append('userId',this.id)
      axios.post('getFollowData', params)
      .then(response => {
        var followDataResult = response.data
        this.followData = followDataResult[0]
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
          var resultMypageArticle = response.data
          this.mypageArticle = resultMypageArticle
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
          var likedPostResult = response.data
          this.likedPosts = likedPostResult
        }
      })
    },
  }
}
</script>

<style scoped>
img {
  border-radius: 50%;
  width: 180px;
  height: 180px;
}
</style>
