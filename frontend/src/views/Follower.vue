<template>
  <div>
    <h1>フォロワー一覧</h1>
    <p v-if="followerLists == 0">
      フォロワーはいません。
    </p>
    <p v-for="followerList in followerLists" :key="followerList">
      {{followerList.UserId}}
      {{followerList.Name}}
      <button v-if="!isFollowedBy" @click="registerFollow">フォローする</button>
      <button v-else @click="deleteFollow()">フォロー中</button>
    </p>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  props:["id"],
  data() {
    return {
      followerLists: [],
      isFollowedBy: false,
    }
  },
  created() {
    this.getFollower()
    this.checkFollow()
  },
  methods:{
    getFollower() {
      const params = new URLSearchParams()
      params.append('followed_id', this.id)
      axios.post('getFollower', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          var followerResult = response.data
          this.followerLists = followerResult
        }
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
            this.isFollowedBy = true
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
            this.isFollowedBy = false
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
