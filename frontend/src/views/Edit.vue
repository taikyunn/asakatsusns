<template>
  <div>
    <Header />
    <div class="container mt-4">
      <div class="row justify-content-center">
        <div class="col-md-8 text-center">
          <div class="mb-3">
            <label for="exampleFormControlTextarea1" class="form-label"></label>
            <textarea name="body" class="form-control" id="exampleFormControlTextarea1" rows="3" col="10" v-model="body" ></textarea>
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
            <button @click="updateBody" class="btn btn-outline-warning">編集する</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import VueTagsInput from '@sipec/vue3-tags-input'
import Header from './Header.vue'

export default {
  props: {
    id: {
      type: String,
      default : '',
    },
  },
  components: {
    VueTagsInput,
    Header,
  },
  data() {
    return {
      body: '',
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
    this.process()
  },
  methods: {
    async process() {
      await this.getAutocompleteItems()
      await this.getOneArticle()
    },
    getOneArticle() {
      const params = new URLSearchParams()
      params.append('id', this.id)
      axios.post('getOneArticle', params)
      .then(response => {
        if (response.status !== 200) {
          throw new Error('レスポンスエラー')
        } else {
          var resultArticle = response.data
          this.body = resultArticle[0].Body
          if (resultArticle[0].Tags != null) {
            this.tags = resultArticle[0].Tags
          }
        }
      })
    },
    getAutocompleteItems() {
      axios.get("/getAutocompleteItems")
      .then (response => {
        var resultAutocompleteItems = response.data
        if (resultAutocompleteItems !== null) {
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
    updateBody(){
      const params = new FormData()
      const tags = document.getElementById("tags").value;
      params.append('body', this.body)
      params.append('id', this.id)
      params.append('tags', tags)
      const config = {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      }
      axios.post("updateArticle", params, config)
      .then(response => {
        if (response.status !== 200) {
          throw new Error("レスポンスエラー")
        } else {
          alert("編集しました。")
          this.$router.push('/')
        }
      })
    },
  }
}
</script>

<style scoped>
.container {
  padding-top: 5rem;
}

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
</style>
