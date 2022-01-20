<template>
  <div>
    <h1>編集ページ</h1>
    <h2>投稿内容</h2>
    <textarea name="body" cols="70" rows="10" v-model="body"></textarea>
    <div>
      <input type="hidden" id="tags" :value="tagsJson">
      <vue-tags-input
      v-model="tag"
      :tags="tags"
      placeholder="タグを5個まで入力できます"
      :autocomplete-items="filteredItems"
      @tags-changed="newTags => tags = newTags"
      />
    </div>
    <button @click="updateBody">編集する</button>
    <button @click="back">戻る</button>
  </div>
</template>

<script>
import axios from 'axios'
import VueTagsInput from '@sipec/vue3-tags-input'

export default {
  props: {
    id: {
      type: String,
      default : '',
    },
  },
  components: {
    VueTagsInput,
  },
  data() {
    return {
      body: '',
      tag: '',
      tags: [],
      autocompleteItems: [{
        text: '早起き',
      }, {
        text: '早寝',
      }, {
        text: '朝勉',
      }, {
        text: 'カフェ勉',
      }, {
        text: '散歩',
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
          this.body = resultArticle[0].Body
          if (resultArticle[0].Tags != null) {
            this.tags = resultArticle[0].Tags
          }
        }
      })
    },
    updateBody(){
      const params = new FormData()
      const tags = document.getElementById("tags").value;
      console.log(tags)
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
