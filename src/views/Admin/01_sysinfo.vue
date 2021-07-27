<template>
  <div id="main">
    <!-- 用户信息 -->
    <div id="userinfo">
      <el-card :body-style="{ padding: '0px' }">
        <div v-if="user_profile !==null">
          <img :src="user_profile.avatar" class="image">
          <div style="padding: 14px;postition:flex">
            <span>用户名：{{user_profile.user.username}}</span><br>
            <div v-if="user_profile.user.admin_type === 'Super Admin'">
              <span>超级管理员</span>
            </div>
          </div>
          <hr>
          <div>
            <span>上次登录：{{user_profile.user.last_login}}</span><br>
            <span>IP：{{session.ip}}</span>
          </div>
        </div>

      </el-card>
    </div>
    <!-- 系统信息 -->
    <div id="sys_info">
      <div class="userinfo">
        <i class="el-icon-s-custom"></i>
        <span>用户总数</span>
        {{dasgborad_info.user_count}}
      </div>
      <div class="userinfo">
        <i class="el-icon-position"></i>
        <span>今日提交数</span>
        {{dasgborad_info.today_submission_count}}
      </div>
      <div class="userinfo">
        <i class="el-icon-folder-opened"></i>
        <span>最近发布作业数</span>
        {{dasgborad_info.recent_contest_count}}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "profile",
  data () {
    return {
      msg: "",
      session: { ip: "127.0.0.1" },
      user_profile: this.$root.user_profile,
      dasgborad_info: {}
    };
  },
  beforeMount () {
    this.$axios.user_profile().then((res) => {
      this.selfLog(res);
      this.user_profile = res["data"];
      this.$root.user_profile = this.user_profile
      this.user_profile.user.last_login = this.user_profile.user.last_login.substr(0, 10)
      this.user_profile.avatar = this.$OJIP + this.user_profile.avatar;
      this.selfLog(this.user_profile);
    });
  },

  mounted () {
    //   获取登录历史
    this.$axios.login_history().then(res => {
      this.selfLog(res.data.slice(-1));
      if (res.error === null) {
        this.session = res.data.slice(-1)[0];
      } else {
        this.$message.error('登录历史获取失败');
      }
    })
    // 获取系统信息
    this.$axios.dashboard_info().then(res => {
      this.selfLog(res);
      if (res.error === null) {
        this.dasgborad_info = res.data
      } else {
        this.$message.error('登录历史获取失败');
      }
    })
  },
  methods: {},
  computed: {},
};
</script>

<style lang="scss" scoped >
#main {
  background-color: #c9dbdb;
  width: 60%;
  height: 50vh;
  margin: 2% 20% 0 20%;
  padding: 30px;
  .ruleForm {
    width: 60%;
    margin-left: 15%;
  }
  .editButton {
    width: 48%;
  }
  #userinfo {
    float: left;
    width: 30%;
    height: 30%;
    .image {
      position: relative;
      width: 30%;
      height: 30%;
      border-radius: 50%;
    }
  }
  #sys_info {
    margin-left: 10px;
    float: left;
    width: 65%;
    height: 27%;
    // border: 1px black solid;
    .userinfo {
      float: left;
      width: 40%;
      height: 27%;
      box-shadow: inset 0 -3em 3em rgba(0, 0, 0, 0.1),
        0 0 0 2px rgb(255, 255, 255), 0.3em 0.3em 1em rgba(0, 0, 0, 0.3);
      margin-left: 15px;
      margin-top: 10px;
      i {
        float: left;
        margin-top: 15px;
        margin-left: 5px;
      }
    }
  }
}
</style>