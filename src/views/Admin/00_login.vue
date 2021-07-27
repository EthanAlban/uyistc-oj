<template>
  <dv-full-screen-container class="bg">
    <dv-border-box-3>
      <div class="title">
        <span class="sys-name">神技Online Judge 管理员登录</span>
        <dv-decoration-5 style="width:300px;height:40px; margin:auto;" :color="['#3490de', '#07689f']" />
      </div>
      <el-tabs type="border-card" class="panel">
        <el-tab-pane label="管理员登录">
          <el-form label-position="left" label-width="80px" :model="form">
            <el-form-item label="用户名称">
              <el-input prefix-icon="el-icon-user" v-model="form.username" placeholder="请输入用户名称"></el-input>
            </el-form-item>
            <el-form-item label="登录密码">
              <el-input prefix-icon="el-icon-s-custom" type="password" v-model="form.password" placeholder="请输入登录密码"
                show-password></el-input>
            </el-form-item>
            <el-button style="width:100%" type="primary" round @click="login" :loading="loading">登录</el-button>
          </el-form>
        </el-tab-pane>
      </el-tabs>

    </dv-border-box-3>
  </dv-full-screen-container>
</template>
<script>
export default {
  name: "Login",
  mounted () {

  },
  data () {
    return {
      loading: false,
      form: {
        username: "",
        password: "",
      },
    };
  },
  methods: {
    login () {
      this.loading = true
      let param = {
        username: this.form.username,
        password: this.form.password
      };
      this.$axios.userLogin(param).then(res => {
        if (res.error === null) {
          this.selfLog(res.data);
          this.$axios.user_profile().then(res => {
            this.selfLog(res)
            this.loading = false
            var myDate = new Date();
            const h = this.$createElement;
            this.$notify({
              title: '登陆成功',
              message: h('i', { style: 'color: teal' }, "管理员：" + this.form.username + "  " + myDate.toLocaleString() + ' 登录成功')
            });
            let params = {
              userId: res.data.user.id,
              msg: myDate.toLocaleString() + ":管理员后台登录",
              type: 1
            };
            this.$axios.send_sys_info(params);
            this.selfLog(res);
            if (res.data.user.admin_type === "Super Admin") {
              this.$root.user_profile = res.data
              this.$router.push("/admin/sysinfo");
            }
            else {
              this.$message({
                message: "该账户不是管理员账户",
                type: "warning"
              });
            }
          })

        } else {
          this.loading = false
          if (res.data === "Invalid username or password") {
            this.$message({
              message: "用户名密码不匹配或未注册",
              type: "warning"
            });
          }
        }
      });
    },
  }
};
</script>

<style lang="scss" scoped>
.bg {
  background-image: url("../../assets/login/bg.png");
  .register {
    float: right;
    top: 2px;
    right: 2px;
  }
  .title {
    height: 16%;
    .sys-name {
      font-size: 45px;
      color: #f5f5f5;
    }
  }
  .panel {
    width: 25%;
    height: 40%;
    margin: auto;
    background-color: white;
    opacity: 0.85;
    .captchaImg {
      width: 200px;
      height: 40px;
      margin-bottom: -15px;
    }
    .loginButton {
      width: 200px;
    }
  }
}
footer {
  position: absolute;
  bottom: 10%;
  @media screen and (max-width: 1440px) and (min-width: 1280px) {
    bottom: 16%;
  }
  display: flex;
  width: 100%;
  justify-content: center;
  p {
    font-size: 30px;
  }
}
</style>

