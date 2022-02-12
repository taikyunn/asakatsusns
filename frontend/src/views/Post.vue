<template>
  <div>
    <Header></Header>
    <div class="text-center">
      <h1>投稿内容</h1>
      <div v-if="apiErrors.length">
        <p v-for="error in apiErrors" :key="error">{{ error }}</p>
      </div>
      <div class="mb-3">
        <textarea name="body" cols="70" rows="10" v-model="body"></textarea>
      </div>
      <div class="mb-3 mx-auto">
        <input type="hidden" id="tags" :value="tagsJson">
        <vue-tags-input
        v-model="tag"
        :tags="tags"
        placeholder="タグを5個まで入力できます"
        :autocomplete-items="filteredItems"
        @tags-changed="newTags => tags = newTags"
        class="mx-auto"
        />
      </div>
      <div class="mb-3">
        <button @click='createArticle'>投稿する</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import VueTagsInput from '@sipec/vue3-tags-input'
import Header from './Header.vue'

export default {
  components: {
    VueTagsInput,
    Header,
  },
  data() {
    return {
      body:'',
      apiErrors: [],
      tag: '',
      tags: [],
      autocompleteItems: [],
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
  mounted() {
    this.getAutocompleteItems()
  },
  methods: {
    getAutocompleteItems() {
      axios.get("/getAutocompleteItems")
      .then (response => {
        var resultAutocompleteItems = response.data
        if (resultAutocompleteItems != null) {
          var target = []
          for (var i = 0; i < resultAutocompleteItems.length; i++) {
            target[i] = {text: resultAutocompleteItems[i]}
          }
          const handler1 = {};
          const proxy1 = new Proxy(target, handler1);
          this.autocompleteItems = proxy1
        }
      })
    },
    createArticle() {
    try {
        if (this.body == '') {
          throw new Error('終了します');
       }
        const tags = document.getElementById("tags").value;
        const params = new FormData()
        params.append('body', this.body)
        params.append('name', localStorage.getItem('userName'))
        params.append('tags', tags)
        const config = {
          headers: {
            'Authorization': `Bearer ${localStorage.getItem('jwt')}`,
            'Content-Type': 'multipart/form-data'
          }
        }
        axios.post('/post/new', params, config)
        .then(response => {
          if (response.status == 201) {
            if (response.data.Body != '') {
              alert("ログインからやり直してください。")
              this.$router.push('/login')
            } else {
              this.apiErrors.push(response.data.Body)
            }
          } else if (response.status != 200) {
            throw new Error('レスポンスエラー')
          } else {
            if (response.data != '') {
              alert("早起き達成" + response.data + "日目!!")
              this.$router.push('/')
            } else {
              this.$router.push('/')
            }
          }
        })
      } catch(error) {
        this.apiErrors.push("投稿内容を入力してください。")
      }
    }
  }
}
</script>

<style scoped>
.vue-tags-input {
  max-width: 50%;
}

.vue-tags-input .ti-tag {
  background: transparent;
  border: 1px solid #747373;
  color: #747373;
  margin-right: 4px;
  border-radius: 0px;
  font-size: 13px;
}

.text-center {
  padding-top: 5rem;
}

</style>
