<template>
  <div>
    <h1>ログインフォーム</h1>
    <label for="name">お名前:</label>
    <input type="text" name='name' v-model="name"><br />
    <label for="password">パスワード:</label>
    <input type="password" name='password' v-model="password"><br />
    <button @click="signUp">ログイン</button>
  </div>
</template>

<script>
// import axios from 'axios'

// export default defineComponent({
//   setup() {
    
//   },
// })
</script>
