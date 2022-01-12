<template>
  <div>
    <h1>編集ページ</h1>
    <h2>投稿内容：</h2>
    <textarea name="body" cols="70" rows="10" v-model="body"></textarea>
    <br />
    <button @click="updateBody">編集する</button>
    <button @click="back">戻る</button>
  </div>
</template>

<script>
import axios from 'axios'


export default {
  props:["id"],
  data() {
    return {
      body: '',
    }
  },
  created() {
    this.fetchArticle()
  },
  methods: {
    fetchArticle() {
      const params = new URLSearchParams()
      params.append('id', this.id)
      axios.post('getOneArticle', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          var resultArticle = response.data
          this.body = resultArticle[0].body
        }
      })
    },
    updateBody(){
      const params = new URLSearchParams()
      params.append('body', this.body)
      params.append('id', this.id)
      axios.post("updateArticle", params)
      .then(response => {
        if (response.status != 200) {
          throw new Error("レスポンスエラー")
        } else {
          alert("編集しました。")
          this.$router.push('/')
        }
      })
    },
    back() {
      this.$router.push('/')
    }
  }
}
</script>
