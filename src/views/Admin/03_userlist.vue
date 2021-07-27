<template>
  <div style="margin: 70px;">
    <el-table :data="tableData" border size="medium">
      <el-table-column fixed prop="id" label="ID">
      </el-table-column>
      <el-table-column fixed prop="username" label="用户名">
      </el-table-column>
      <el-table-column prop="real_name" label="真实姓名">
      </el-table-column>
      <el-table-column label="上次登录">
        <template slot-scope="scope">
          <i class="el-icon-time"></i>
          <span style="margin-left: 10px">{{ scope.row.last_login | formatDate}}</span>
        </template>
      </el-table-column>
      <el-table-column prop="email" label="邮箱">
      </el-table-column>
      <el-table-column prop="admin_type" label="用户类型">
        <template slot-scope="scope">
          <div v-if="scope.row.admin_type==='Regular User'">
            <span style="color:green">学生</span>
          </div>
          <div v-else-if="scope.row.admin_type==='Admin'">
            <span style="color:blue">教师</span>
          </div>
          <div v-else>
            <span style="color:red">超级管理员</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作">
        <template slot-scope="scope">
          <el-button type="text" size="small" @click="dialogVisible2 = true,user=scope.row">编辑</el-button>
          <el-button type="text" size="small" @click="dialogVisible = true,user=scope.row,userid=scope.row.id">删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="currentPage"
      :page-sizes="[5,10, 20, 40, 80,160]" :page-size="limit" layout="total, sizes, prev, pager, next, jumper"
      :total="total">
    </el-pagination>
    <!-- 删除用户弹窗 -->
    <el-dialog title="用户操作" :visible.sync="dialogVisible" width="30%">
      <span>确定删除该用户？删除教师，该老师将需要重新注册，申请~</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="delete_user()">确 定</el-button>
      </span>
    </el-dialog>

    <!-- 编辑用户弹窗 -->
    <el-dialog title="编辑用户" :visible.sync="dialogVisible2">
      <el-form :model="user">
        <el-form-item label="用户名" :label-width="formLabelWidth">
          <el-input v-model="user.username"></el-input>
        </el-form-item>
        <el-form-item label="邮箱" :label-width="formLabelWidth">
          <el-input v-model="user.email"></el-input>
        </el-form-item>
        <el-form-item label="禁用" :label-width="formLabelWidth">
          <el-switch v-model="user.is_disabled" active-color="#13ce66" inactive-color="#ff4949"></el-switch>
        </el-form-item>
        <el-form-item label="真实姓名" :label-width="formLabelWidth">
          <el-input v-model="user.real_name" style="height:200px;width:100%"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="show_editor2 = false">取 消</el-button>
        <el-button type="primary" @click="alerter_user()">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  beforeMount () {
    this.$axios.user_profile().then((res) => {
      this.user_profile = res["data"];
      this.$root.user_profile = this.user_profile
    });
    this.selfLog("获取用户列表");
    this.obtain_users()
  },
  data () {
    return {
      formLabelWidth: '120px',
      userid: -1,
      user: {},
      dialogVisible: false,
      dialogVisible2: false,
      currentPage: 1,
      total: 0,
      user_profile: null,
      offset: 0,
      limit: 8,
      tableData: []
    }
  },
  methods: {
    handleSizeChange (val) {
      this.limit = val
      this.obtain_users()
    },
    handleCurrentChange (val) {
      this.currentPage = val
      this.offset = (this.currentPage - 1) * this.limit
      this.obtain_users()
    },
    // 获取用户
    obtain_users () {
      this.$axios.obtain_users(this.offset, this.limit).then(res => {
        this.selfLog(res);
        this.tableData = res.data.results
        this.total = res.data.total
      })
    },
    // 删除用户
    delete_user () {
      // 自己的数据库也要清除登录密码
      let param = {
        userId: this.user.id,
        stunum: this.user.username
      }
      this.selfLog(param);
      this.$axios.delete_user(this.userid).then(res => {
        if (res.error === null) {

          this.$axios.deleteUsr(param).then(res => {
            this.$message({
              message: '用户删除成功',
              type: 'success'
            });
            this.dialogVisible = false
            this.obtain_users()
          })
        }
      })
    },
    // 更新用户信息
    alerter_user () {
      let user = this.user
      let param = {
        admin_type: user.admin_type,
        create_time: user.create_time,
        email: user.email,
        id: user.id,
        is_disabled: user.is_disabled,
        open_api: false,
        problem_permission: "None",
        real_name: user.real_name,
        real_tfa: false,
        two_factor_auth: false,
        username: user.username,
      }
      this.$axios.alerter_user(param).then(res => {
        this.selfLog(res);
        if (res.error === null) {
          this.$message({
            message: '用户修改成功',
            type: 'success'
          });
          this.dialogVisible2 = false
          this.obtain_users()
        }

      })
    },
  },


}
</script>