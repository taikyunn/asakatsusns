<template>
  <div class="sidebar_fixed">
    <div class="card w-75 main">
      <div class="card-header text-center">
        <fa icon="tag" />
        メインタグ
      </div>
      <div class="card-body text-center" v-for="mainTag in mainTags" :key="mainTag">
        <router-link class="tag" :to="{name: 'HomeTag', params: {id:(Number(mainTag.Id))}}">
          #{{mainTag.Name}}
        </router-link>
      </div>
    </div>
    <div class="card w-75 ranking">
      <div class="card-header text-center">
        <fa icon="crown" class="crown"/>
        早起き達成ランキンング
      </div>
      <div class="card-body text-center" v-for="(ranking, index) in rankings" :key="ranking">
        <span class="rank">{{index + 1}}位:</span>
        <span class="name">{{ranking.Name}}さん</span>
        <span class="count">{{ranking.Count}}回</span>
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
  mounted() {
    this.process()
  },
  methods: {
    async process() {
      await this.getMainTag()
      await this.getWakeUpRanking()
    },
    getMainTag() {
      axios.get('getMainTag')
      .then(response => {
        if (response.status !== 200) {
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
        if (response.status !==200) {
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

.main {
  margin-top: 40px;
}

.ranking {
  margin-top: 40px;
}

.tag {
  color: green;
  white-space: nowrap;
  text-decoration: none;
}

.rank {
  float: left;
}

.count {
  float: right;
}

</style>
