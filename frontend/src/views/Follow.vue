<template>
  <div>
    <h1>フォロー一覧</h1>
    <p v-if="followLists == 0">
      フォロしているアカウントはありません。
    </p>
    <p v-for="followList in followLists" :key="followList" v-else>
      <router-link :to="{name: 'Mypage', params: {id:(Number(followList.UserId))}}">{{followList.Name}}</router-link>
      <button v-if="isFollowedBy" @click="registerFollow(followList.UserId)">フォローする</button>
      <button v-else @click="deleteFollow(followList.UserId)">フォロー中</button>
    </p>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  props:["id"],
  data() {
    return {
      followLists: [],
      isFollowedBy: false,
    }
  },
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
