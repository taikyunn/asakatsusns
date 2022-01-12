<template>
  <div>
    <h1>寝る時間編集ページ</h1>
    <input type="time" v-model="sleepTime">
    <p>
      <button @click="updateSleepTime">編集する</button>
      <button @click="back">戻る</button>
    </p>
    {{userData.SleepTime}}
  </div>
</template>

<script>
import axios from 'axios'

export default {
  props:["id"],
  data(){
    return {
      userData:[],
      sleepTime:'',
    }
  },
  created() {
    const params = new URLSearchParams()
    params.append('userId', this.id)
    axios.post('/getSleepTimeData', params)
    .then(response => {
      if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          var resultUser = response.data
          this.userData = resultUser[0]
        }
    })
  },
  methods:{
    updateSleepTime(){
      const params = new URLSearchParams()
      params.append('sleepTime', this.sleepTime)
      params.append('userId', this.id)
      axios.post("updateSleepTime", params)
      .then(response => {
        if (response.status != 200) {
          throw new Error("レスポンスエラー")
        } else {
          alert("編集しました。")
          this.$router.push('/mypage/' + this.id)
        }
      })
    },
    back(){
      this.$router.push('/mypage/' + this.id)
    }
  }
}
</script>


