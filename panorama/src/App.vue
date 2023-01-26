<template>
  <div id="app">
    <router-view />
  </div>
</template>
<script>
import { login, logout, getCodeImg } from '@/api/login';
import { getToken, setToken, removeToken } from '@/utils/auth';
export default {
  data() {
    return {
      loginForm: {
        username: "demo",
        password: "123456",
        code: "111",
        uuid: ""
      }
    }
  },
  created() {
    this.getCode();
    this.Login();
  },
  methods: {
    // 登录
    Login() {
      const username = this.loginForm.username.trim();
      const password = this.loginForm.password;
      const code = this.loginForm.code;
      const uuid = this.loginForm.uuid;
      return new Promise((resovle, reject) => {
        login(username, password, code, uuid).then(res => {
          // console.log(res);
          setToken(res.data.data.token)
          resovle()
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 退出登录
    LogOut() {
      return new Promise((resolve, reject) => {
        logout(getToken()).then(() => {
          removeToken()
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 获取验证码
    getCode() {
      getCodeImg().then(res => {
        // this.codeUrl = res.data.base64stringC;
        this.loginForm.uuid = res.data.idKeyC;
      });
    },
  },
  beforeDestroy() {
    this.LogOut();
  },
}


</script>

<style scoped>

</style>
