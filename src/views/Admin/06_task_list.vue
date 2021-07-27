<template>
  <div>
    <!-- 作业模糊查询 -->
    <el-form :inline="true" :model="formInline" class="demo-form-inline"
      style="margin:10px 0 0 20px;postion:flex;float:left">
      <el-form-item label="作业查询">
        <el-input v-model="formInline.keyword" placeholder="输入关键字"></el-input>
      </el-form-item>
      <el-form-item label>
        <el-select v-model="formInline.status" placeholder="作业选项">
          <el-option label="全部" value></el-option>
          <el-option label="正在进行" value="0"></el-option>
          <el-option label="尚未开始" value="1"></el-option>
          <el-option label="已经结束" value="-1"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="search_task">查询</el-button>
      </el-form-item>
    </el-form>
    <!-- 作业列表显示 -->
    <el-table :data="tableData" style="width: 100%">
      <el-table-column type="expand">
        <template slot-scope="props">
          <el-form label-position="left" inline class="demo-table-expand">
            <el-form-item label="发布时间">
              <span>{{ props.row.create_time }}</span>
            </el-form-item>
            <el-form-item label="开始时间">
              <span>{{ props.row.start_time }}</span>
            </el-form-item>
            <el-form-item label="最近修改时间">
              <span>{{ props.row.last_update_time }}</span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column label="作业" prop="title"></el-table-column>
      <el-table-column label="作业描述" prop="description">
        <template slot-scope="scope">
          <div v-html="scope.row.description"></div>
        </template>
      </el-table-column>
      <el-table-column label="截止时间">
        <template slot-scope="scope">
          <i class="el-icon-time"></i>
          <span style="margin-left: 10px">{{ scope.row.end_time | formatDate}}</span>
        </template>
      </el-table-column>
      <el-table-column label="创建人" prop="created_by.username"></el-table-column>
      <el-table-column label="可见状态" prop="visible">
        <template slot-scope="scope">
          <el-switch v-model="scope.row.visible" @change="hide_task(scope.row)" active-color="#13ce66"
            inactive-color="#ff4949"></el-switch>
        </template>
      </el-table-column>
      <el-table-column label="管理">
        <template slot-scope="scope">
          <el-button @click="EnterContest(scope.row.id)" type="text" size="small">查看</el-button>
          <el-button type="text" size="small">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- 输入密码弹框控件 -->
    <el-dialog title="验证" :visible.sync="dialogFormVisible" append-to-body>
      <i class="el-icon-lock"></i>
      <el-form :model="form">
        <el-form-item label="作业密码" :label-width="formLabelWidth">
          <el-input v-model="form.password" type="password" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="EnterContest_password">确 定</el-button>
      </div>
    </el-dialog>
    <!-- 作业问题列表抽屉 -->
    <el-drawer title="作业题目列表" :visible.sync="table" direction="rtl" size="50%" append-to-body>
      <el-table :data="problems" style="width: 100%">
        <el-table-column type="expand">
          <template slot-scope="props">
            <el-form label-position="left" inline class="demo-table-expand">
              <el-form-item label="习题描述">
                <div style="color:blue" v-html="props.row.description"></div>
              </el-form-item>
              <el-form-item label="可使用语言">
                <div v-for="(value, index) of props.row.languages" :key="index" style="color:blue">{{ value }}、</div>
              </el-form-item>
              <el-form-item label>
                <el-button type="primary" @click="doProblem(props.row._id)" plain>完成习题</el-button>
              </el-form-item>
            </el-form>
          </template>
        </el-table-column>
        <el-table-column label="ID" prop="_id"></el-table-column>
        <el-table-column label="标题" prop="title"></el-table-column>
        <el-table-column label="已完成人数" prop="accepted_number"></el-table-column>
      </el-table>
    </el-drawer>
    <!-- 分页控件 -->
    <div class="block">
      <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="page"
        :page-sizes="[10, 20, 40, 80]" :page-size="100" layout="total, sizes, prev, pager, next, jumper" :total="total">
      </el-pagination>
    </div>
  </div>
</template>

<script>
export default {
  beforeMount () {
    this.Get_Contests_List();
    this.$axios.user_profile().then(res => {
      this.selfLog(res);
      this.user_profile = res["data"];
      this.user_profile.avatar = this.$OJIP + this.user_profile.avatar;
      this.selfLog(this.user_profile);
    });
  },
  data () {
    return {
      user_profile: null,
      // 分页查看
      page: 1,
      offset: 0,
      limit: 10,
      total: 0,
      contest_id: 0,
      tableData: [],
      dialogFormVisible: false,
      formLabelWidth: "80px",
      form: {
        password: ""
      },
      formInline: {
        keyword: "",
        status: ""
      },
      problems: {},
      table: false
    };
  },
  methods: {
    Get_Contests_List () {
      this.$axios.admin_contests_list(this.offset, this.limit).then(res => {
        this.total = res.data.total;
        this.selfLog(res);
        this.tableData = res.data.results;
      });
    },

    // 监听分页控件
    handleSizeChange (val) {
      this.limit = val;
      this.offset = (this.page - 1) * this.limit;
      this.Get_Contests_List();
      this.selfLog(`每页 ${val} 条`);
    },
    handleCurrentChange (val) {
      this.page = val;
      this.offset = (this.page - 1) * this.limit;
      this.Get_Contests_List();
      this.selfLog(`当前页: ${val}`);
    },
    // 参加按钮监听
    EnterContest (contest_id) {
      if (this.user_profile === null) {
        this.$message({
          message: "查看习题请先登录",
          type: "warning"
        });
        return false;
      }
      this.selfLog(contest_id);
      this.contest_id = contest_id;
      this.$axios.access_contest(contest_id).then(res => {
        this.selfLog(res);
        // 没有权限就要输密码
        if (!res.data.access) {
          this.$message({
            message: "暂无权限，请输入密码",
            type: "warning"
          });
          this.dialogFormVisible = true;
        } else {
          this.$message({
            message: "拥有权限，可进入",
            type: "success"
          });
          // 查询作业的问题列表
          this.$axios.contest_problem_list(this.contest_id).then(res => {
            this.selfLog(res);
            if (res.data === "Contest has not started yet.") {
              this.$message({
                message: "作业尚未到达开始时间",
                type: "warning"
              });
            } else {
              this.problems = res.data;
              this.dialogFormVisible = false;
              this.table = true;
            }
          });
        }
      });
      // this.dialogFormVisible = false
    },
    EnterContest_password (id) {
      if (this.user_profile === null) {
        this.$message({
          message: "查看习题请先登录",
          type: "warning"
        });
        return false;
      }
      if (this.form.password === "") {
        this.$message({
          message: "密码不能为空",
          type: "warning"
        });
        return false;
      }
      let params = {
        contest_id: this.contest_id,
        password: this.form.password
      };
      this.$axios.contest_passowrd_check(params).then(res => {
        this.selfLog(res);
        if (res.data === "Wrong password or password expired") {
          this.$message({
            message: "密码错误或者密码已过期",
            type: "warning"
          });
        } else if (res.error === null) {
          // 查询作业的问题列表
          this.$axios.contest_problem_list(this.contest_id).then(res => {
            if (res.data === "Contest has not started yet.") {
              this.$message({
                message: "作业尚未到达开始时间",
                type: "warning"
              });
            } else {
              this.selfLog(res);
              this.problems = res.data;
              this.dialogFormVisible = false;
              this.table = true;
            }
          });
        }
      });
    },
    doProblem (id) {
      this.selfLog(id);
      this.$router.push(`/problemDetail_Contest/${id}/${this.contest_id}`);
    },
    search_task () {
      if (this.formInline.status === "") {
        this.$axios
          .search_task_ignore_status(
            this.offset,
            this.limit,
            this.formInline.keyword
          )
          .then(res => {
            this.total = res.data.total;
            this.tableData = res.data.results;
            this.selfLog(res);
          });
      } else {
        this.$axios
          .search_task(
            this.offset,
            this.limit,
            this.formInline.status,
            this.formInline.keyword
          )
          .then(res => {
            this.total = res.data.total;
            this.tableData = res.data.results;
            this.selfLog(res);
          });
      }
    },
    // 删除作业？
    hide_task (row) {
      let param = {
        allowed_ip_ranges: row.allowed_ip_ranges,
        contest_type: row.contest_type,
        create_time: row.create_time,
        created_by: row.created_by,
        description: row.description,
        end_time: row.end_time,
        id: row.id,
        last_update_time: row.last_update_time,
        password: row.password,
        real_time_rank: row.real_time_rank,
        rule_type: row.rule_type,
        start_time: row.start_time,
        status: row.status,
        title: row.title,
        visible: row.visible
      };
      this.$axios.hide_task(param).then(res => {
        this.selfLog(res);
        if (res.error === null) {
          this.$message({
            message: "修改成功",
            type: "success"
          });
        }
      });
    }
  },
  computed: {}
};
</script>

<style lang='scss' scoped>
.demo-table-expand {
  font-size: 0;
}
.demo-table-expand label {
  width: 90px;
  color: #99a9bf;
}
.demo-table-expand .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 50%;
}
</style>
