<template>
  <div>
    <h1>投稿内容</h1>
    <p v-if="apiErrors.length">
    <ul>
      <li v-for="error in apiErrors" :key="error">{{ error }}</li>
    </ul>
    </p>
    <textarea name="body" cols="70" rows="10" v-model="body"></textarea>
    <div>
      <input type="hidden" name="tags" :value="tagsJson">
      <vue-tags-input
      v-model="tag"
      :tags="tags"
      placeholder="タグを5個まで入力できます"
      :autocomplete-items="filteredItems"
      @tags-changed="newTags => tags = newTags"
      />
    </div>
    <button @click='CreateArticle'>投稿する</button>
  </div>
</template>

<script>
import axios from 'axios'
import VueTagsInput from '@sipec/vue3-tags-input'

export default {
  components: {
    VueTagsInput,
  },
  data() {
    return {
      body:'',
      apiErrors: [],
      tag: '',
      tags: [],
      autocompleteItems: [{
        text: 'Spain',
      }, {
        text: 'France',
      }, {
        text: 'USA',
      }, {
        text: 'Germany',
      }, {
        text: 'China',
      }],
    }
  },
  computed: {
    filteredItems() {
      return this.autocompleteItems.filter(i => {
        return i.text.toLowerCase().indexOf(this.tag.toLowerCase()) !== -1;
      });
    },
    tagsJson() {
      return JSON.stringify(this.tags)
    },
  },
  methods: {
    CreateArticle() {
      const params = new URLSearchParams
      params.append('body', this.body)
      params.append('name', localStorage.getItem('userName'))
      axios.post('/post/new', params, {
      headers: {'Authorization': `Bearer ${localStorage.getItem('jwt')}`}
      })
      .then(response => {
        if (response.status == 201) {
          if (response.data.Body != '') {
            this.apiErrors.push(response.data.Body)
          } else {
            alert("ログインからやり直してください。")
            this.$router.push('/login')
          }
          console.log(response.data.Body)
        } else if (response.status != 200) {
          console.log("エラーが発生しました。")
        } else {
          alert("投稿しました。")
          this.$router.push('/')
        }
      })
    }
  }
}
</script>

<style scoped>
.vue-tags-input {
  max-width: inherit;
}

.vue-tags-input .ti-tag {
  background: transparent;
  border: 1px solid #747373;
  color: #747373;
  margin-right: 4px;
  border-radius: 0px;
  font-size: 13px;
}
</style>
