<template>
  <div class="sidebar_fixed">
    <img v-if="profileDataUrl" :src="profileDataUrl" width="100" />
    <p v-if="url">
      <span style="position:absolute" @click="deletePreview" width="1px" height="1px">X</span>
      <img :src="url" width="100" />
    </p>
    <p v-if="isMyOwnPage">
      <input type="file" ref="preview" @change="uploadFile" accept="image/jpeg, image/png">
      <button v-on:click="fileUpload()">アップロード</button>
    </p>
    <p>
      <router-link class="link" :to="{name: 'Follow', params: {id:(Number(this.id))}}">{{followData.FollowerCount}}
        フォロー
      </router-link>
      <router-link class="link" :to="{name:'Follower', params: {id:(Number(this.id))}}">{{followData.FollowCount}}
        フォロワー
      </router-link>
      <span v-if="!isMyOwnPage">
        <button v-if="!isFollowedBy" @click="registerFollow">
          フォローする
        </button>
        <button v-else @click="deleteFollow">
          フォロー中
        </button>
      </span>
    </p>
    <p v-if="isEdit">
      <span v-if="!editName" class="border p-2  bg-light" v-on:click="doEditName" >
        {{userInfo.name}}*クリックで編集*
      </span>
      <span v-else >
        <input type="text" v-model="userInfo.name" v-on:blur="editName = false" v-focus>
        <button @click="updateName">
          登録
        </button>
      </span>
    </p>
    <p v-else>
      アカウント名:{{userInfo.name}}
    </p>
    <p v-if="isEdit">
      寝る時間:
      <span v-if="!editSleepTime" class="border p-2  bg-light" v-on:click="doEditSleepTime">
        {{userInfo.SleepTime}}*クリックで編集*
      </span>
      <span v-else >
        <input type="time" v-model="userInfo.SleepTime" v-on:blur="editSleepTime = false" v-focus>
        <button @click="updateSleepTime">
          登録
        </button>
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
    <p v-if="isEdit">
      起きる時間設定:
      <span v-if="!editWakeUpTime" class="border p-2  bg-light" v-on:click="doEditWakeUpTime">
        {{userInfo.WakeUpTime}}*クリックで編集*
      </span>
      <span v-else >
        <input type="time" v-model="userInfo.WakeUpTime" v-on:blur="editWakeUpTime = false" v-focus>
        <button @click="updateWakeUpTime">
          登録
        </button>
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
</template>

<script>
import axios from 'axios'

export default {
  props:["id"],
  directives: {
    focus: {
      inserted: function (el) {
          el.focus();
      }
    }
  },
  data() {
    return {
      url:'',
      profileDataUrl: '',
      followData: [],
      isMyOwnPage: false,
      isFollowedBy: false,
      isEdit: false,
      editName: false,
      userInfo:[],
      editSleepTime: false,
      editWakeUpTime: false,
      wakeUpTime:'',
    }
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
  mounted() {
    this.checkMyOwnPage()
    this.getUserProfile()
    this.getFollowData()
    this.checkFollow()
    this.checkIsEdit()
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
    deletePreview(){
      this.url = ''
      URL.revokeObjectURL(this.url);
      this.mounted()
    },
    uploadFile() {
      const file = this.$refs.preview.files[0];
      this.url = URL.createObjectURL(file)
      this.fileInfo = this.$refs.preview.files[0]
      this.profileDataUrl = ""
      this.$refs.preview.value = "";
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
    getFollowData() {
      const params = new URLSearchParams()
      params.append('userId',this.id)
      axios.post('getFollowData', params)
      .then(response => {
        var followDataResult = response.data
        this.followData = followDataResult[0]
      })
    },
    checkMyOwnPage() {
      if (this.id == localStorage.getItem('userId'))
      this.isMyOwnPage = true
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
    checkIsEdit() {
      if (this.id == localStorage.getItem('userId')) {
        this.isEdit = true
      }
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
  }
}
</script>

<style scoped>
.sidebar_fixed {
  position: sticky;
  top: 5rem;
}

.link {
  text-decoration: none;
  text-align: left;
  color:black;
}
</style>
