<template>
  <div style="margin: 70px; width:50%">
    <el-table :data="tableData" border>
      <el-table-column fixed prop="id" label="ID">
      </el-table-column>
      <el-table-column prop="tea_name" label="名字">
      </el-table-column>
      <el-table-column prop="userId" label="用户id">
      </el-table-column>
      <el-table-column prop="username" label="工号">
      </el-table-column>
      <el-table-column fixed="right" label="操作">
        <template slot-scope="scope">
          <el-button @click="pass_check(scope.row)" type="text" size="small">通过</el-button>
          <el-button type="text" @click="reject_tea_check(scope.row)" size="small">拒绝</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-divider></el-divider>
    <span style="color:red">已拒绝：</span>
    <el-table :data="reject_tableData" border style="width: 100%">
      <el-table-column fixed prop="id" label="ID">
      </el-table-column>
      <el-table-column prop="tea_name" label="名字">
        <template slot-scope="scope">
          <span style="color:red">{{scope.row.tea_name}}</span>
        </template>
      </el-table-column>
      <el-table-column prop="userId" label="用户id">
      </el-table-column>
      <el-table-column prop="username" label="工号">
      </el-table-column>
      <el-table-column fixed="right" label="操作">
        <template slot-scope="scope">
          <el-button @click="pass_check(scope.row)" type="text" size="small">通过</el-button>
          <el-button type="text" @click="reject_tea_check(scope.row)" size="small">拒绝</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  mounted () {
    this.$axios.user_profile().then(res => {
      this.user_profile = res.data
    })
    this.obtain_teacher_apply_list()
    this.obtain_teacher_reject_list()
  },
  data () {
    return {
      user_profile: this.$root.user_profile,
      tableData: [],
      reject_tableData: []
    }
  },
  methods: {
    handleClick (row) {
      this.selfLog(row);
    },
    obtain_teacher_apply_list () {
      this.$axios.obtain_teacher_apply_list().then(res => {
        this.selfLog(res);
        this.tableData = res.data
      })
    },
    // 获取已拒绝的列表
    obtain_teacher_reject_list () {
      this.$axios.obtain_teacher_reject_list().then(res => {
        this.selfLog(res);
        this.reject_tableData = res.data
      })
    },
    // 教师权限审核通过
    pass_check (msg) {
      this.selfLog(msg);
      //   首先设置成Admin
      let param = {
        admin_type: "Admin",
        id: msg.userId,
        username: msg.username,
        real_name: msg.tea_name,
        email: msg.username + "@oj.com",
        open_api: false,
        problem_permission: "All",
        real_tfa: false,
        two_factor_auth: false,
        is_disabled: false
      }
      this.$axios.alerter_user(param).then(res => {
        this.selfLog(res);
        if (res.error === null) {
          let param2 = {
            username: msg.username
          }
          // 然后向django更新审核状态
          this.$axios.pass_tea_check(param2).then(res => {
            //   发送系统消息
            var myDate = new Date();
            let params = {
              userId: msg.userId,
              msg: myDate.toLocaleString() + ":教师权限成功，教师权限已开放",
              type: 1
            };
            this.$axios.send_sys_info(params);
            this.$message({
              message: '审核修改成功',
              type: 'success'
            });
            this.obtain_teacher_apply_list()
          })
        }

      })

    },
    // 拒绝教师权限申请
    reject_tea_check (user) {
      let params = {
        username: user.username
      }
      this.selfLog(params);
      this.$axios.reject_tea_check(params).then(res => {
        if (res.data === "rejected") {
          this.$message({
            message: '已拒绝该教师的权限申请',
            type: 'success'
          });
          //   发送系统消息
          var myDate = new Date();
          let params = {
            userId: user.userId,
            msg: myDate.toLocaleString() + ":教师权限已被拒绝，请联系管理员",
            type: 1
          };
          this.$axios.send_sys_info(params);
          this.obtain_teacher_apply_list()
        } else {
          this.$message.error('否决权限失败~');
        }
      })
    },
  },
}
</script>