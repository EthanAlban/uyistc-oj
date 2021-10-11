<template>
  <div>
    <!-- 教师端,老师能看到所在班级的学生，可以发布公告，可以创建新的问题，看到自己创建的问题列表，能够发布作业 -->
    <el-container>
      <el-main style="padding:0">
        <el-menu :default-active="$route.path" mode="horizontal" background-color="#3f72af" text-color="#fdfdfd"
          active-text-color="#222831" router>
          <a href="/"><img src="../../../public/sologon.png" class="logo" alt=""></a>
          <p style="float:left;marginRight:100px;fontSize:18px;color:#fafafa">教师端</p>
          <el-menu-item index="/teacher/class" class="nav">授课班级</el-menu-item>
          <el-menu-item index="/teacher/question" class="nav">题目管理</el-menu-item>
          <el-menu-item index="/teacher/assign" class="nav">作业管理</el-menu-item>
        </el-menu>
        <router-view />
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "teacher",
  data () {
    return {};
  },
  methods: {
  },
  computed: {}
};
</script>

<style lang='scss' scoped>
.nav {
  font-size: 18px;
}
.logo {
  width: 17%;
  float: left;
}
</style>