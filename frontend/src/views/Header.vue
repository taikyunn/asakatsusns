<template>
  <div id="nav">
    <nav class="navbar navbar-expand-lg navbar-light bg-warning fixed-top">
      <div class="container-fluid">
        <a class="navbar-brand" href="#">
          <router-link class="title" to="/">
            ASAKATSUSNS
          </router-link>
        </a>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarSupportedContent">
          <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          </ul>
          <label class="nav-link" v-if="authenticatedUser">
            <fa icon="pen" class="pen" />
            <router-link to="/post" class="post" v-if="authenticatedUser">
              投稿する
            </router-link>
          </label>
          <div class="justify-content-end">
            <ul class="navbar-nav me-auto mb-2 mb-lg-0">
              <li class="nav-item dropdown">
                <a class="nav-link dropdown" href="#" id="navbarDropdown" role="button" data-bs-toggle="dropdown" aria-expanded="false">
                  <fa icon="user-circle" />
                  <router-link class="user-name" :to="{name: 'Mypage', params: {id:(Number(currentUserId))}}" v-if="authenticatedUser">
                    {{ currentUserName }}さん
                  </router-link>
                </a>
                <ul class="dropdown-menu" aria-labelledby="navbarDropdown" justify-content-end >
                  <li>
                    <a class="dropdown-item" href="#">
                      <router-link class="btn btn-warning" :to="{name: 'Mypage', params: {id:(Number(currentUserId))}}" v-if="authenticatedUser">
                        マイページ
                      </router-link>
                    </a>
                  </li>
                  <li>
                    <a class="dropdown-item" href="#">
                      <router-link class="btn btn-warning" to="/signup" v-if="notAuthenticatedUser">
                        新規登録
                      </router-link>
                    </a>
                  </li>
                  <li>
                    <a class="dropdown-item" href="#">
                      <router-link class="btn btn-warning" to="/login" v-if="notAuthenticatedUser">
                        ログイン
                      </router-link>
                    </a>
                  </li>
                  <li>
                    <a class="dropdown-item" href="#">
                      <router-link class="btn btn-warning"  to="/logout" @click="signOut" v-if="authenticatedUser">
                        ログアウト
                      </router-link>
                    </a>
                  </li>
                </ul>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </nav>
  </div>
</template>

<script>
import axios from 'axios'
import firebase from 'firebase/app'
import 'firebase/app'
import "firebase/auth"

export default {
  data() {
    return {
      authenticatedUser: '',
      notAuthenticatedUser: '',
      currentUserId: '',
      currentUserName: localStorage.getItem('userName'),
    }
  },
  mounted(){
    this.getHeader()
  },
  methods: {
    getHeader() {
      const params = new URLSearchParams()
      params.append('id', localStorage.getItem('userId'))
      axios.post('getHeader', params)
      .then(response => {
        if (response.data != '') {
        this.authenticatedUser = true
        this.notAuthenticatedUser = false
        } else {
          this.authenticatedUser = false
          this.notAuthenticatedUser = true
        }
        this.currentUserId = response.data
      })
    },
    signOut() {
      confirm('ログアウトしてもよろしいですか。')
      firebase.auth().signOut()
      .then(() => {
        localStorage.removeItem('jwt')
        localStorage.removeItem('userName')
        localStorage.removeItem('userId')
        this.$router.go({path: this.$router.push('/login'), force: true})
        alert("ログアウトしました。")
      })
    }
  }
}
</script>

<style scoped>
.title, .post, .user-name {
  text-decoration: none;
  color: black;
  font-weight: bold;
}
img {
  border-radius: 50%;
  width:3%;
  height:3%;
}
.pen {
  color: black;
}

</style>
