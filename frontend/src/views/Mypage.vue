<template>
  <div>
    <h1>マイページです。</h1>
  </div>
</template>

<script>

export default {
  props:["id"],
  created() {
    this.fetchTodo()
  },
  methods: {
    fetchTodo(){
      console.log("ここが動いています",this.id)
    }
  }
}
</script>
