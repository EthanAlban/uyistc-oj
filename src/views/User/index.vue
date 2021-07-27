<template>
  <div id="main">
    <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
      <!-- logo -->
      <el-col :span="4">
        <a href="/"><img src="../../../public/sologon.png" style="width:100%;margin-left:-15px" alt=""></a>
      </el-col>
      <el-menu-item index="/home/nav1">首页</el-menu-item>
      <el-menu-item index="/user/profile">个人信息</el-menu-item>
      <el-submenu index="workstation">
        <template slot="title">信息工作台</template>
        <el-menu-item index="/user/updateInfo">修改个人信息</el-menu-item>
        <el-menu-item index="/user/ModifyPwd">修改密码</el-menu-item>
        <el-menu-item index="/user/ModifyEmail">修改邮箱</el-menu-item>
        <el-menu-item index="/user/loginHistory">登录记录</el-menu-item>
      </el-submenu>
    </el-menu>
    <el-main>
      <div id="main">
        <router-view />
      </div>
    </el-main>
  </div>
</template>

<script>
export default {
  name: "user",
  data () {
    return {
      activeIndex: "/user/profile",
      user_profile: this.$root.user_profile,

    };
  },
  methods: {
    // 路由跳转
    handleSelect (key, keyPath) {
      //相同路由不能重复定向bug fix
      if (this.$route.path === key) {
        return
      }
      this.selfLog(key, keyPath);
      if (keyPath[0] === "/home/nav1") {
        this.$router.push("/home/nav1");
      } else if (keyPath[0] === "/user/profile") {
        this.$router.push("/user/profile");
      } else if (keyPath[0] === "workstation") {
        this.$router.push(keyPath[1]);
      }
    },
  },
};
</script>

<style lang='scss' scoped>
</style>
