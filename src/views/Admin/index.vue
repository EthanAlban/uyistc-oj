<template>
  <div>
    <el-container>
      <el-main style="padding:0">
        <el-menu :default-active="$route.path" mode="horizontal" background-color="#3f72af" text-color="#fdfdfd"
          active-text-color="#222831" router>
          <a href="/"><img src="../../../public/sologon.png" class="logo" alt=""></a>
          <p style="float:left;marginRight:100px;fontSize:18px;color:#fafafa">管理员端</p>
          <el-menu-item index="/admin/sysinfo" class="nav">系统信息</el-menu-item>
          <el-menu-item index="/admin/annoncement" class="nav">公告管理</el-menu-item>
          <el-menu-item index="/admin/usrlist" class="nav">用户管理</el-menu-item>
          <el-menu-item index="/admin/check_tea_list" class="nav">教师审核</el-menu-item>
          <el-menu-item index="/admin/question" class="nav">题目管理</el-menu-item>
          <el-menu-item index="/admin/assign" class="nav">作业管理</el-menu-item>
        </el-menu>
        <router-view />
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "user",
  data () {
    return {
      user_profile: this.$root.user_profile,
      activeIndex: "/user/profile"
    };
  },
  beforeMount () {
    // 用户
    this.$axios.user_profile().then(res => {
      this.selfLog("登录用户");
      this.selfLog(res.data);
      this.user_profile = res["data"];
      this.user_profile.avatar = this.$OJIP + this.user_profile.avatar;
      this.selfLog(this.user_profile);
    });
  },
  methods: {
  },
};
</script>

<style lang='scss' scoped>
.nav {
  font-size: 18px;
}
.logo {
  width: 20%;
  float: left;
}
</style>
