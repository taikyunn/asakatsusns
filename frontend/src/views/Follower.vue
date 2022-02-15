<template>
  <div>
    <Header />
    <div class="text-center">
      <div class="container mt-4">
        <div class="row justify-content-center">
          <div class="col-md-8">
            <ul id="myTab" class="nav nav-tabs mb-3 justify-content-center" role="tablist">
              <li class="nav-item" role="presentation">
                <button type="button" id="home-tab" class="nav-link btn btn-outline-warning" data-bs-toggle="tab" data-bs-target="#home" role="tab" aria-controls="home" aria-selected="false">
                  フォロー中
                </button>
              </li>
              <li class="nav-item" role="presentation">
                <button type="button" id="profile-tab" class="nav-link btn btn-outline-warning" data-bs-toggle="tab" data-bs-target="#profile" role="tab" aria-controls="profile" aria-selected="true">
                  フォロワー
                </button>
              </li>
            </ul>
            <div id="myTabContent" class="tab-content">
              <div id="home" class="tab-pane" role="tabpanel" aria-labelledby="home-tab">
                <p v-if="followLists == 0">
                  フォロしているアカウントはありません。
                </p>
                <div class="card" v-for="followList in followLists" :key="followList" v-else>
                  <div class="card-body">
                    <router-link class="link" :to="{name: 'Mypage', params: {id:(Number(followList.UserId))}}">
                      {{followList.Name}}
                    </router-link>
                    <span class="follow">
                      <button class="text-end btn btn-outline-warning" v-if="isFollowedBy" @click="registerFollow(followList.UserId)">
                        フォローする
                      </button>
                      <button class="text-end btn btn-outline-warning" v-else @click="deleteFollow(followList.UserId)">
                        フォロー中
                      </button>
                    </span>
                  </div>
                </div>
              </div>
              <div id="profile" class="tab-pane active" role="tabpanel" aria-labelledby="profile-tab">
                <p v-if="followerLists == 0">
                  フォロワーはいません。
                </p>
                <div class="card" v-for="followerList in followerLists" :key="followerList">
                  <div class="card-body">
                    <span v-if="followerList.Image" class="profile-image">
                      <img :src="followerList.Image" class="circle" />
                    </span>
                    <span v-else class="profile-image">
                      <img :src="defaultImage" class="circle" />
                    </span>
                    <router-link class="link" :to="{name: 'Mypage', params: {id:(Number(followerList.UserId))}}">
                      {{followerList.Name}}
                    </router-link>
                    <button class="follow btn btn-outline-warning" v-if="!isFollowedBy" @click="registerFollow">
                      フォローする
                    </button>
                    <button class="follow btn btn-outline-warning" v-else @click="deleteFollow()">
                      フォロー中
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Header from './Header.vue'

export default {
  props:["id"],
  data() {
    return {
      followerLists: [],
      followLists: [],
      isFollowedBy: false,
      defaultImage: require('@/images/default.png'),
    }
  },
  components: { Header },
  mounted() {
    this.process()
  },
  methods:{
    async process() {
      await this.getFollower()
      await this.getFollow()
      await this.checkFollow()
    },
    getFollower() {
      const params = new URLSearchParams()
      params.append('followed_id', this.id)
      axios.post('getFollower', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          var followerResult = response.data
          this.followerLists = followerResult
          for (let i = 0; i < followerResult.length; i++) {
            if (followerResult[i].ProfileImagePath == '') {
              continue;
            }
            let url = process.env.VUE_APP_DATA_URL + followerResult[i].ProfileImagePath
            axios.get(url,{responseType: "blob"})
            .then(response => {
              let blob = new Blob([response.data])
              this.followerLists.splice(i, 1, {
                Name: followerResult[i].Name,
                ProfileImagePath: followerResult[i].ProfileImagePath,
                Image: URL.createObjectURL(blob),
              })
            })
          }
        }
      })
    },
    checkFollow() {
      const params = new URLSearchParams()
      params.append('follower_id',localStorage.getItem('userId'))
      params.append('followed_id',this.id)
      axios.post("checkFollow", params)
      .then(response => {
        var followResult = response.data
        this.isFollowedBy = followResult
      })
    },
    getFollow() {
      const params = new URLSearchParams()
      params.append('follower_id', this.id)
      axios.post('getFollow', params)
      .then(response => {
        if (response.status != 200) {
          throw new Error('レスポンスエラー')
        } else {
          var followResult = response.data
          this.followLists = followResult
        }
      })
    },
    registerFollow() {
      try {
        if (localStorage.getItem('jwt') == '') {
          throw new Error('終了します');
        }
        const params = new URLSearchParams()
        params.append('follower_id', localStorage.getItem('userId'))
        params.append('followed_id',this.id)
        axios.post("registerFollow", params)
        .then(response => {
          if (response.status != 200) {
            throw new Error("レスポンスエラー")
          } else {
            this.isFollowedBy = true
          }
        })
      } catch {
        alert("ログインからやり直してください。")
        this.$router.push('/login')
      }
    },
    deleteFollow() {
      try {
        if (localStorage.getItem('jwt') == '') {
          throw new Error('終了します');
        }
        const params = new URLSearchParams()
        params.append('follower_id', localStorage.getItem('userId'))
        params.append('followed_id',this.id)
        axios.post('deleteFollow', params)
        .then(response => {
          if (response.status != 200) {
            throw new Error("レスポンスエラー")
          } else {
            this.isFollowedBy = false
          }
        })
      } catch {
        alert("ログインからやり直してください。")
        this.$router.push('/login')
      }
    },
  },
}
</script>

<style scoped>
.text-center {
  padding-top: 5rem;
}

.card {
  margin-bottom: 1rem;
}

.link {
  text-decoration: none;
  text-align: left;
  color:black;
  float: left;
  padding-left: 1rem;
}

.follow {
  float: right;
}

.circle {
  border-radius: 50%;
  width: 50px;
  height: 50px;
  object-fit: cover;
}

.profile-image {
  float: left;
}
</style>
