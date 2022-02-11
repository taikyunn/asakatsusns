<template>
  <div>
    <div class="card w-75" v-for="article in articles" :key="article">
      <div class="card-header">
        <router-link class="link" :to="{name: 'Mypage', params: {id:(Number(article.UserId))}}">
          {{article.Name}}
        </router-link>
        {{article.UpdatedAt}}
        <span v-if="article.UserId == currentUserId">
          <span class="dropdown">
            <a class="dropdown" data-bs-toggle="dropdown" aria-expanded="false">
              <fa icon="ellipsis-v" class="ellipsis" />
            </a>
            <ul class="dropdown-menu" justify-content-end >
              <li>
                  <a class="dropdown-item" href="#">
                    <router-link class="btn btn-warning" :to="{name: 'Edit', params: {id:(Number(article.Id))}}">
                      編集
                    </router-link>
                  </a>
              </li>
              <li>
                  <a class="dropdown-item" href="#">
                    <button class="btn btn-warning" @click="deleteArticle(article.Id)">
                      削除
                    </button>
                  </a>
              </li>
            </ul>
          </span>
        </span>
      </div>
      <div class="card-body">
        <p class="card-text">
          {{article.Body}}
        </p>
        <div v-for="tag in tags" :key="tag">
          <div v-if="article.Id == tag.ArticleId">
            <router-link class="tag" :to="{name: 'HomeTag', params: {id:(Number(tag.Key))}}">
              #{{tag.Value}}
            </router-link>
          </div>
        </div>
      </div>
      <div class="card-footer text-end">
        <div v-for="likeCount in likeCounts" :key="likeCount">
          <div v-if="likeCount.ArticleId == article.Id">
            <span v-for="commentCount in commentCounts" :key="commentCount">
              <span v-if="commentCount.ArticleId == article.Id">
                <fa icon="comment-alt" />
                {{commentCount.Count}}
              </span>
            </span>
            <span v-for="result in results" :key="result">
              <span v-if="result.ArticleId == article.Id">
                <span @click="registerLikes(article.Id)" v-if="result.Count">
                  <fa icon="heart" class="like-btn"/>
                  <span v-if="likeCount.ArticleId == article.Id" class="heart">
                    {{likeCount.Count}}
                  </span>
                </span>
                <span @click="deleteLikes(article.Id)" v-else >
                  <fa icon="heart" class="unlike-btn"/>
                  <span v-if="likeCount.ArticleId == article.Id" class="heart">
                    {{likeCount.Count}}
                  </span>
                </span>
              </span>
            </span>
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
  import { useRouter } from 'vue-router'

  let articles = ref([]);
  let likeCounts = ref([]);
  let commentCounts = ref([]);
  let results = ref([]);
  let tags = ref([]);
  let count = 1;
  let currentUserId = ref(localStorage.getItem('userId'))
  const router = useRouter()

  const load = async $state => {
    const params = new URLSearchParams()
    params.append('count', count)
    params.append('userId', localStorage.getItem('userId'))
    try {
      const response = await axios.post('getNextArticles', params);
      const nextArticles = await response.data.nextArticles;
      const countData = await response.data.countData;
      const commentCount = await response.data.commentCount;
      const favoriteData = await response.data.favoriteData;
      const tagData = await response.data.tagData;
      if (nextArticles.length < 10 || response.status == 201) {
        $state.complete()
      } else {
        articles.value.push(...nextArticles);
        likeCounts.value.push(...countData);
        commentCounts.value.push(...commentCount);
        results.value.push(...favoriteData);
        tags.value.push(...tagData)
        $state.loaded()
      }
      count ++;
    } catch (error) {
      $state.error()
    }
  }

  async function registerLikes(articleId) {
    try {
      if (localStorage.getItem('jwt') == '') {
        throw new Error('終了します');
      }
      const params = new URLSearchParams()
      params.append('articleId', articleId)
      params.append('userId', localStorage.getItem('userId'))
      const config = {
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`,
        }
      }
      try {
        const response = await axios.post('/post/registerLikes', params, config);
        const count = await axios.post('getNextCountFavorites', params);
        const favorite = await axios.post("checkNextFavorite", params);
        if (response.status == 201) {
          if (response.data.Body != '') {
            alert("ログインからやり直してください。")
            router.push('/login')
          }
        } else if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          var resultLikesCount = count.data[0]
          for (var i = 0; i < likeCounts.value.length; i++) {
            if (likeCounts.value[i].ArticleId == resultLikesCount.ArticleId) {
                likeCounts.value[i] = resultLikesCount
            }
          }
          var resultFavorite = favorite.data[0]
          for (var index = 0; index < results.value.length; index++) {
            if (results.value[index].ArticleId == resultFavorite.ArticleId) {
                results.value[index] = resultFavorite
            }
          }
        }
      } catch {
        alert("ログインからやり直してください。")
        router.push('/login')
      }
    } catch {
      alert("ログインからやり直してください。")
      router.push('/login')
    }
  }

  async function deleteLikes(articleId) {
    try {
      if (localStorage.getItem('jwt') == '') {
        throw new Error('終了します');
      }
      const params = new URLSearchParams()
      params.append('articleId', articleId)
      params.append('userId', localStorage.getItem('userId'))
      const config = {
        headers: {
          'Authorization': `Bearer ${localStorage.getItem('jwt')}`,
        }
      }
      try {
        const response = await axios.post('/post/deleteLikes', params, config)
        if (response.status == 201) {
          if (response.data.Body != '') {
            alert("ログインからやり直してください。")
            router.push('/login')
          }
        } else if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          const count = await axios.post('getNextCountFavorites', params);
          const favorite = await axios.post("checkNextFavorite", params);
          var resultLikesCount = count.data[0]
          for (var i = 0; i < likeCounts.value.length; i++) {
            if (likeCounts.value[i].ArticleId == resultLikesCount.ArticleId) {
                likeCounts.value[i] = resultLikesCount
            }
          }
          var resultFavorite = favorite.data[0]
          for (var index = 0; index < results.value.length; index++) {
            if (results.value[index].ArticleId == resultFavorite.ArticleId) {
                results.value[index] = resultFavorite
            }
          }
        }
      } catch {
        alert("ログインからやり直してください。")
        router.push('/login')
      }
    } catch {
      alert("ログインからやり直してください。")
      router.push('/login')
    }
  }

  async function deleteArticle(articleId) {
    confirm('削除してもよろしいですか。')
    const params = new URLSearchParams()
    params.append('articleId', articleId)
    try {
      const response = await axios.post('deleteArticle', params)
      if (response.status != 200) {
        throw new Error('レスポンスエラー')
      } else {
        router.go({path: router.currentRoute.path, force: true})
        alert('削除しました。')
      }
    } catch {
      alert("ログインからやり直してください。")
      router.push('/login')
    }
  }
</script>

<style scoped>

.like-btn {
  width: 18px;
  height: 18px;
  font-size: 25px;
  color: #808080;
  margin-left: 20px;
  margin-right: 5px;
}

.unlike-btn {
  width: 18px;
  height: 18px;
  font-size: 25px;
  color: #e54747;
  margin-left: 20px;
  margin-right: 5px;
}

 .heart {
   margin-right: 20px;
 }

 .link {
  text-decoration: none;
  text-align: left;
  color:black;
 }

.card {
  margin-bottom: 40px;
}

.ellipsis {
  float: right;
}

.tag {
  color: green;
  white-space: nowrap;
  text-decoration: none;
}

.link {
  text-decoration: none;
  text-align: left;
  color:black;
}
</style>
