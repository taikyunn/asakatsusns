<template>
  <div>
    <h1>フォロー一覧</h1>
    <p v-for="followList in followLists" :key="followList">
      {{followList.UserId}}
      {{followList.Name}}
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
    }
  },
}
</script>
