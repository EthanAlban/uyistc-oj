<template>
  <div id="main">
    <div class="block">
      <el-timeline v-for="(login, index) in login_historys" :key="index">
        <el-timeline-item :timestamp="login.last_activity" placement="top">
          <el-card>
            <h4>登录IP：{{login.ip}}</h4>
            <h4>session_key：{{login.session_key}}</h4>
            <h4>登录设备：{{login.user_agent}}</h4>
          </el-card>
        </el-timeline-item>
      </el-timeline>
    </div>
  </div>
</template>

<script>
export default {
  name: "loginHistory",
  mounted () {
    this.$axios.login_history().then(res => {
      this.selfLog(res);
      if (res["error"] === null) {
        this.login_historys = res["data"]
      } else {
        this.$message.error('获取登录历史失败');
      }
    })
  },
  data () {
    return {
      login_historys: {},
      msg: "",
    };
  },
  methods: {},
  computed: {},
};
</script>

<style lang='scss' scoped>
#main {
  background-color: #3d84a8;
  width: 60%;
  margin: auto;
  // line-height: 600px;
}
</style>
