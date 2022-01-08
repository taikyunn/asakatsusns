<template>
  <div id="nav">
    <router-link to="/">Home |</router-link>
    <router-link to="/signup" v-if="notAuthenticatedUser">新規登録 |</router-link>
    <router-link to="/login" v-if="notAuthenticatedUser">ログイン |</router-link>
    <router-link to="/signout" v-if="authenticatedUser">ログアウト |</router-link>
    <router-view></router-view>
  </div>
</template>

<script>

export default {
  data() {
    return {
      notAuthenticatedUser:'',
      authenticatedUser: ''
    }
  },
  mounted(){
    if (localStorage.getItem('jwt') != null) {
      this.authenticatedUser = true;
      this.notAuthenticatedUser = false;
    } else {
      this.authenticatedUser = false;
      this.notAuthenticatedUser = true;
    }
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
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
