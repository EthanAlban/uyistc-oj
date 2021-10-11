<template>
  <div id="main">
    <div v-if="user_profile!==null">
      <img class="avatar" :src="user_profile.Avatar" alt />
      <!-- <dv-border-box-8> -->
      <div class="content">
        <div>用户名：{{ user_profile.UserName }}</div>
        <div>邮箱：{{ user_profile.Email }}</div>
        <div>用户类型：{{ user_profile.UserType }}</div>
        <!-- <div>上次登录：{{ user_profile.user.last_login | formatDate}}</div> -->
        <div>博客：<a :href=user_profile.blog target="_blank">{{ user_profile.Blogaddr }}</a></div>
        <div>github：<a :href=user_profile.github target="_blank">{{ user_profile.Gitaddr }}</a></div>
        <div>专业：{{ user_profile.Major }}</div>
        <div>总分：{{user_profile.total_score}}</div>
        <!-- 若是教师则隐藏申请教师权限按钮 -->
        <el-button v-if="user_profile.UserType === 2" type="success" style="z-index:100" round
          @click="apply_to_be_a_teacher()">
          申请教师权限
        </el-button>
      </div>
      <!-- </dv-border-box-8> -->

      <el-divider direction="vertical" content-position='center' style="height:300px"></el-divider>
      <div class="content">
        <div>提交次数：{{submission}}</div>
        <div>题目总数：{{count}}</div>
        <div>做过的题目：
          <div v-for="problem in problems_id" :key="problem.id">
            <div class="tags">
              <div v-if="problem.status===0">
                <el-tooltip class="item" effect="dark" content="题目通过" placement="bottom">
                  <el-button type="success" round @click="problemDetail(problem.id)">{{problem.id}}</el-button>
                </el-tooltip>
              </div>
              <div v-else-if="problem.status===-2">
                <el-tooltip class="item" effect="dark" content="编译失败" placement="bottom">
                  <el-button type="danger" round @click="problemDetail(problem.id)">{{problem.id}}</el-button>
                </el-tooltip>
              </div>
              <div v-else>
                <el-tooltip class="item" effect="dark" content="尚未通过" placement="bottom">
                  <el-button type="warning" round @click="problemDetail(problem.id)">{{problem.id}}</el-button>
                </el-tooltip>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 登陆记录 -->
    <contribution :data="history" :rectWidth=" 12" :rectHeight="12" :fontSize="10" :click="login_history"
      monthText="zh-cn" :conditions="condition" style="width:70%" />
  </div>
</template>

<script>
export default {
  name: "profile",
  mounted () {
    this.$user_axios.user_profile().then(res => {
      this.selfLog(res);
      this.user_profile = res["data"];
      this.get_submission_problems()
    });
    // 获取登陆记录
    this.login_history()

  },
  data () {
    return {
      msg: "",
      user_profile: null,
      history: {},
      condition: [
        { 'condition': '<=', 'value': 0, 'color': '#ebedf0' },
        { 'condition': '<', 'value': 1, 'color': '#c6e48b' },
        { 'condition': '<', 'value': 5, 'color': '#7bc96f' },
        { 'condition': '<', 'value': 7, 'color': '#239a3b' },
        { 'condition': '>=', 'value': 7, 'color': '#196127' },
      ],
      submission: 0,
      problems: [],
      problems_id: [],
      count: 0
    };
  },
  methods: {
    apply_to_be_a_teacher () {
      this.selfLog("--------")
      let param = {
        id: this.user_profile.id,
        username: this.user_profile.user.username
      }
      if (this.user_profile.user.username.length != 7) {
        this.$message.error('请使用7位工号登录');
      } else {
        this.$axios.apply_to_be_a_teacher(param).then(res => {
          const h = this.$createElement;
          //   发送系统消息
          var myDate = new Date();
          let params = {
            userId: this.user_profile.id,
            msg: myDate.toLocaleString() + ":申请教师权限",
            type: 1
          };
          this.$axios.send_sys_info(params);
          if (res.data === "to_check") {
            this.$notify({
              title: '权限申请',
              message: h('i', { style: 'color: teal' }, '教师权限申请已发送，我们将在0-3天内审核~祝您使用愉快')
            });
          } else if (res.data === "applyed_to_check") {
            this.$notify({
              title: '权限申请',
              message: h('i', { style: 'color: teal' }, '教师权限申请已发送，请勿重复申请')
            });
          } else if (res.data === "already_passed") {
            this.$notify({
              title: '权限申请',
              message: h('i', { style: 'color: teal' }, '教师权限申请已通过，不必重复申请')
            });
          } else if (res.data === "stunum_not_exists") {
            this.$notify({
              title: '权限申请',
              message: h('i', { style: 'color: teal' }, '教师工号不在数据库中，请联系管理员')
            });
          }
        })
      }
    },
    // 登录记录点击事件
    click_history (date) {
      this.selfLog(date)
    },
    // 获取登陆记录
    login_history () {
      this.$user_axios.login_history().then(res => {
        this.selfLog(res);
        let data = {}
        if (res["errcode"] === 200) {
          let login_historys = res["data"]
          for (let his of login_historys) {
            his = his.Date.substr(0, 10)
            // this.selfLog(his.last_activity.substr(0, 10));
            let strArray = his.split('-');
            strArray = strArray.map(function (val) {
              if (val[0] == '0') {
                return val = val.slice(1);
              } else {
                return val;
              }
            })
            his = strArray.join('-')
            if (his in data) {
              data[his] = data[his] + 1
            } else {
              data[his] = 1
            }

          }
          this.history = data
          this.selfLog(data);
        } else {
          this.$message.error('获取登录历史失败');
        }
      })
        .catch(res => {
          this.login_history()
        })
    },
    // 拿自己做过的题目|提交的题目数量
    get_submission_problems () {
		this.$submission_axios.GetUserSubmissionProfile().then(res=>{
			this.selfLog(res)
			 this.submission = res.data.sub_counts
			 this.problems = res.data.problems
			 for (let key in this.problems) {
			   this.count = this.count + 1
			   var item = this.problems[key];
			   this.selfLog(item); //AA,BB,CC,DD
			   this.problems_id.push({
			     id: item['Pid'],
			     status: res.data.status[item['Pid']]
			   });
			 }
			 this.selfLog(this.problems_id);
		})
    },
    // 进入问题详情页
    problemDetail (id) {
      this.selfLog(id);
      this.$router.push(`/problemDetail/${id}`);
    }
  },
  computed: {
    user_type () {
      if (this.user_profile.UserType === 2) {
        return '普通用户'
      }
      if (this.user_profile.UserType === 1) {
        return '教师'
      }
      if (this.user_profile.UserType === 0) {
        return '管理员'
      }
    }
  }
};
</script>

<style lang='scss' scoped>
#main {
  background-color: #c9dbdb;
  width: 60%;
  height: 100%;
  margin-left: 20%;
  .avatar {
    border: 2px grey solid;
    border-radius: 50%;
    width: 20%;
    height: 20%;
    margin-top: 20px;
  }
  .content {
    border-radius: 10px;
    box-shadow: 0px 1px 4px rgba(0, 0, 0, 0.3),
      0px 0px 20px rgba(0, 0, 0, 0.1) inset;
    line-height: 30px;
    // position: flex;
    position: relative;
    z-index: 19;
    // margin-left: 2%;
    width: 28%;
    padding: 5px;
    float: right;
    margin-bottom: 7%;
    margin-top: 7%;
    margin-right: 2%;
    // border: 1px black solid;
    height: 40%;
    max-height: 300px;
    overflow: auto;
  }
  .tags {
    z-index: 20;
    position: relative;
    float: left;
    margin-top: 2%;
    margin-left: 5px;
  }
  html body {
    height: 100%;
  }
}
</style>
