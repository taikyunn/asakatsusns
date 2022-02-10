<template>
  <div>
    <div class="card w-75" v-for="article in articles" :key="article">
      <div class="card-header">
        {{article.Name}}
        {{article.UpdatedAt}}
      </div>
      <div class="card-body">
        <p class="card-text">
          {{article.Body}}
        </p>
      </div>
      <div class="card-footer text-end">
        <div v-for="likeCount in likeCounts" :key="likeCount">
          <div v-if="likeCount.ArticleId == article.Id">
            いいね数:{{likeCount.Count}}
          </div>
        </div>
      </div>
    </div>
    <InfiniteLoading :articles="articles" @infinite="load" />
  </div>
</template>

<script setup>
  import axios from 'axios'
  import { ref } from "vue";
  import InfiniteLoading from "v3-infinite-loading";
  import "v3-infinite-loading/lib/style.css";

  let articles = ref([]);
  let likeCounts = ref([]);
  let count = 1;
  const load = async $state => {
    const params = new URLSearchParams()
    params.append('count', count)
    try {
      const response = await axios.post('getNextArticles', params);
      const nextArticles = await response.data.nextArticles;
      const countData = await response.data.countData;
      if (nextArticles.length < 10 || response.status == 201) {
        $state.complete()
      } else {
        articles.value.push(...nextArticles);
        likeCounts.value.push(...countData);
        $state.loaded()
      }
      count ++;
    } catch (error) {
      $state.error()
    }
  }
</script>

<style>
  #app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
  }

  .result {
    display: flex;
    flex-direction: column;
    gap: 5px;
    font-weight: 300;
    width: 400px;
    padding: 10px;
    text-align: center;
    margin: 0 auto 10px auto;
    background: #eceef0;
    border-radius: 10px;
  }

.card {
  margin-bottom: 40px;
}
</style>
