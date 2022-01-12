<template>
  <div>
    <h1>起きる時間編集ページ</h1>
    <input type="time" v-model="wakeUpTime">
    <p>
      <button @click="updateWakeUpTime">編集する</button>
      <button @click="back">戻る</button>
    </p>
    {{userData.WakeUpTime}}
  </div>
</template>

<script>
import axios from 'axios'

export default {
  props:["id"],
  data(){
    return {
      userData:[],
      wakeUpTime:'',
    }
  },
  created() {
    const params = new URLSearchParams()
    params.append('userId', this.id)
    axios.post('/getWakeUpTimeData', params)
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
    updateWakeUpTime(){
      const params = new URLSearchParams()
      params.append('wakeUpTime', this.wakeUpTime)
      params.append('userId', this.id)
      axios.post("updateWakeUpTime", params)
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
