<template>
  <div>
    <Header></Header>
    <ul id="myTab" class="nav nav-tabs mb-3" role="tablist">
      <li class="nav-item" role="presentation">
        <button type="button" id="home-tab" class="nav-link active" data-bs-toggle="tab" data-bs-target="#home" role="tab" aria-controls="home" aria-selected="true">フォロー中</button>
      </li>
      <li class="nav-item" role="presentation">
        <button type="button" id="profile-tab" class="nav-link" data-bs-toggle="tab" data-bs-target="#profile" role="tab" aria-controls="profile" aria-selected="false">フォロワー</button>
      </li>
    </ul>
    <div id="myTabContent" class="tab-content">
      <div id="home" class="tab-pane active" role="tabpanel" aria-labelledby="home-tab">
        <p v-if="followLists == 0">
          フォロしているアカウントはありません。
        </p>
          <div class="card" v-for="followList in followLists" :key="followList" v-else>
            <div class="card-header">
              <router-link :to="{name: 'Mypage', params: {id:(Number(followList.UserId))}}">{{followList.Name}}</router-link>
              <button class="text-end" v-if="isFollowedBy" @click="registerFollow(followList.UserId)">
                フォローする
              </button>
              <button class="text-end" v-else @click="deleteFollow(followList.UserId)">
                フォロー中
              </button>
            </div>
          </div>
      </div>
      <div id="profile" class="tab-pane" role="tabpanel" aria-labelledby="profile-tab">
        <p><strong>これは、プロフィールタブに関連付けられたコンテンツのプレースホルダコンテンツ。</strong>...</p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Header from './Header.vue'

export default {
  props:["id"],
  data() {
    return {
      followLists: [],
      isFollowedBy: false,
    }
  },
  components: { Header },
  created() {
    this.getFollow()
  },
  methods:{
    getFollow() {
      const params = new URLSearchParams()
      params.append('follower_id', this.id)
      axios.post('getFollow', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          var followResult = response.data
          this.followLists = followResult
        }
      })
    },
    checkFollow() {
      const params = new URLSearchParams()
      params.append('follower_id', this.id)
      params.append('followed_id',localStorage.getItem('userId'))
      axios.post("checkFollow", params)
      .then(response => {
        var followResult = response.data
        this.isFollowedBy = followResult
      })
    },
    registerFollow(followedId) {
      try {
        if (localStorage.getItem('jwt') == '') {
          throw new Error('終了します');
        }
        const params = new URLSearchParams()
        params.append('follower_id', localStorage.getItem('userId'))
        params.append('followed_id', followedId)
        axios.post("registerFollow", params)
        .then(response => {
          if (response.status != 200) {
            throw new Error("レスポンスエラー")
          } else {
            this.isFollowedBy = false;
          }
        })
      } catch {
        alert("ログインからやり直してください。")
        this.$router.push('/login')
      }
    },
    deleteFollow(followedId) {
      try {
        if (localStorage.getItem('jwt') == '') {
          throw new Error('終了します');
        }
        const params = new URLSearchParams()
        params.append('follower_id', localStorage.getItem('userId'))
        params.append('followed_id', followedId)
        console.log(params)
        axios.post('deleteFollow', params)
        .then(response => {
          if (response.status != 200) {
            throw new Error("レスポンスエラー")
          } else {
            this.isFollowedBy = true;
          }
        })
      } catch {
        alert("ログインからやり直してください。")
        this.$router.push('/login')
      }
    },
  },
}
</script>
