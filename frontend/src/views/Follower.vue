<template>
  <div>
    <h1>フォロワー一覧</h1>
    <p v-for="followerList in followerLists" :key="followerList">
      {{followerList.UserId}}
      {{followerList.Name}}
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
    }
  },
  created() {
    this.getFollower()
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
    }
  },
}
</script>
