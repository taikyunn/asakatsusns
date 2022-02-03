<template>
  <div class="sidebar_fixed">
    <div class="card w-75">
      <div class="card-header text-center">
        <fa icon="tag" />
        メインタグ
      </div>
      <div class="card-body text-center" v-for="(mainTag, index) in mainTags" :key="mainTag">
        <router-link class="tag" :to="{name: 'HomeTag', params: {id:(Number(index))}}">
          #{{mainTag}}
        </router-link>
      </div>
    </div>
    <div class="card w-75 ranking">
      <div class="card-header text-center">
        早起き達成ランキンング
        <fa icon="crown" class="crown"/>
      </div>
      <div class="card-body text-center" v-for="(ranking, index) in rankings" :key="ranking">
        {{index + 1}}位:
        {{ranking.Name}}さん
        {{ranking.Count}}回
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      mainTags: [],
      rankings: [],
    }
  },
  created() {
    this.getMainTag()
    this.getWakeUpRanking()
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
    },
    getWakeUpRanking() {
      axios.get('getWakeUpRanking')
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          var rankingResult =response.data
          this.rankings = rankingResult
        }
      })
    }
  }
}
</script>

<style scoped>
.sidebar_fixed {
  position: sticky;
  top: 5rem;
}

.crown {
  color: #DAAF08;
}

.ranking {
  margin-top: 20px;
  margin-bottom: 20px;
}

.tag {
  color: green;
  white-space: nowrap;
  text-decoration: none;
}


</style>
