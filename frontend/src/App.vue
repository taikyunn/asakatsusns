<template>
  <div>
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  mounted() {
    var routeInstance = this.$route;
    this.createTitleDesc(routeInstance);
  },
  methods : {
    createTitleDesc(routeInstance){
      if (!routeInstance.meta.title) {
        document.title = 'ASAKATSUSNS'
      }
    }
  }
}
</script>
