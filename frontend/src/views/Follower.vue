<template>
  <div>
    <h1>フォロワー一覧</h1>
  </div>
</template>

<script>
export default{
  props:["id"],
  created() {
    this.getFollower()
  },
  methods:{
    getFollower() {
      console.log(this.id)
    }
  },
}
</script>
