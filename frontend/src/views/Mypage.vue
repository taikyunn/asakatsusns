<template>
  <div>
    <h1>プロフィール</h1>
    <p>お名前：{{userInfo.name}}</p>
    <p>
      寝る時間：{{userInfo.SleepTime}}
      <router-link :to="{name: 'SleepTimeEdit', params: {id:(Number(userInfo.ID))}}">編集する</router-link>
      <span v-if="userInfo.SleepTime == null">
        <input type="time" name='sleepTime' v-model="sleepTime">
        <button @click="registerSleepTime">登録</button>
      </span>
    </p>
    <p>起きる時間設定：{{userInfo.WakeUpTime}}
      <router-link :to="{name: 'WakeUpTimeEdit', params: {id:(Number(userInfo.ID))}}">編集する</router-link>
      <span v-if="userInfo.WakeUpTime == null">
        <input type="time" name='wakeUpTime' v-model="wakeUpTime">
        <button @click="registerWakeUpTime">登録</button>
      </span>
    </p>
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
  methods: {
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
    }
  }
}
</script>
