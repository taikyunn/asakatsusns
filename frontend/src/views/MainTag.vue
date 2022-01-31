<template>
  <div class="card w-75">
    <div class="card-header text-center">
      <fa icon="tag" />
      メインタグ
    </div>
    <div class="card-body text-center" v-for="mainTag in mainTags" :key="mainTag">
      #{{mainTag.name}}
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      mainTags: [],
    }
  },
  created() {
    this.getMainTag()
  },
  methods: {
    getMainTag() {
      axios.get('getMainTag')
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          var resultMainTag = response.data
          this.mainTags = resultMainTag
        }
      })
    }
  }
}
</script>

