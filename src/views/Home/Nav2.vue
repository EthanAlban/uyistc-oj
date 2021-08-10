<template>
  <div>
    <el-button type="primary" style="float:left;margin:10px 0 10px 10px" round @click="GetProblems">刷新题库</el-button>
    <i class="el-icon-s-management" style="float:left;margin:10px 0 10px 10px;color:green"
      @mouseover="this.show_tag_drawer"></i>
    <!-- 标签的抽屉 -->
    <el-drawer :visible.sync="tag_drawer" append-to-body direction="btt" :with-header="false">
      <el-tag v-for="(tag,idx) in tags" :key="idx" :disable-transitions="false" :type="tag_types[idx%5]" effect="dark"
        style="position:flex;float:left;margin-left:10px;margin-top:10px" @click="choose_tag_problem(tag)">{{tag.name}}
      </el-tag>
    </el-drawer>
    <el-table :data="problems_list" style="width: 100%" height="600" stripe :row-class-name="tableRowClassName">
      <el-table-column label="状态" prop="my_status">
        <template slot-scope="scope">
          <div v-if="scope.row.my_status===0" title="通过">
            <span style="color:green;">
              <i class="el-icon-success"></i>
            </span>
          </div>
          <div v-if="scope.row.my_status===null" title="未作"></div>
          <div v-if="scope.row.my_status===-2" title="编译错误">
            <span style="color:red;">
              <i class="el-icon-error"></i>
            </span>
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="Pid" label="#" width="120"></el-table-column>
      <el-table-column prop="Title" label="标题" width="180"></el-table-column>
      <el-table-column label="标签" width="180">
        <template slot-scope="scope">
          <el-tag v-for="(tag,idx) in scope.row.tags" :key="tag" size="mini" :type="tag_types[idx%5]" effect="dark"
            @click="tag_ = {name:tag} ,choose_tag_problem(tag_)" style="margin-left:5px">{{ tag }}
          </el-tag>

        </template>
      </el-table-column>
      <el-table-column prop="Level" label="难度" width="180">
        <template slot-scope="scope">
          <el-rate v-model="scope.row.Level" disabled show-score :max="max" text-color="#ff9900" score-template>
          </el-rate>
        </template>
      </el-table-column>
      <el-table-column prop="AcceptSubmissions" label="通过数量"></el-table-column>
      <el-table-column prop="TotalSubmissions" label="提交数"></el-table-column>
      <el-table-column prop="pass_rate" label="通过率">
        <template slot-scope="scope">
          <div v-if="scope.row.pass_rate===0" title="一个人都没做出来">
            <span style="color:black">{{scope.row.pass_rate}}%</span>
          </div>
          <div v-else>
            <el-progress :text-inside="true" :stroke-width="26" :percentage="scope.row.pass_rate"></el-progress>
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="Uid.UserName" label="创建者"></el-table-column>
      <el-table-column fixed="right" label="操作" width="100">
        <template slot-scope="scope">
          <el-button @click="problemDetail(scope.row)" type="success" round>
            做它
            <i class="el-icon-thumb"></i>
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="page"
      :page-sizes="[5,10, 20, 40]" :page-size="10" layout="total, sizes, prev, pager, next, jumper" :total="total"
      style="z-index:6;position:absolute;bottom:20%;left:40%">
    </el-pagination>
  </div>
</template>

<script>
export default {
  mounted () {
    this.GetProblems();
    // 获取tags
    this.$axios.get_tags().then(res => {
      this.selfLog("获取tags")
      this.selfLog(res)
      this.tags = res.data
    })
  },
  data () {
    return {
      max: 3,
      value: 1,
      problems_list: [],
      page: 1,
      offset: 0,
      limit: 10,
      total: 0,
      tag_types: ["info", "warning", "danger", "", "success"],
      tags: [],
      tag_types: ["info", "warning", "danger", "", "success"],
      tag_drawer: false
    };
  },
  methods: {
    handleSizeChange (val) {
      this.limit = val;
      this.offset = (this.page - 1) * this.limit;
      this.GetProblems();
      this.selfLog(`每页 ${val} 条`);
    },
    handleCurrentChange (val) {
      this.page = val;
      this.offset = (this.page - 1) * this.limit;
      this.GetProblems();
      this.selfLog(`当前页: ${val}`);
    },
    tableRowClassName ({ row, rowIndex }) {
      if (rowIndex === 1) {
        return "warning-row";
      } else if (rowIndex === 3) {
        return "success-row";
      }
      return "";
    },
    // 获取问题列表
    GetProblems () {
      this.$problem_axios.ProblemListPage(this.offset, this.limit).then(res => {
        this.selfLog(res);
        if (res.errcode === 200) {
          this.total = res.data.length;
          this.problems_list = res.data;
          let arange = this.limit
        }
      });
    },
    // 进入问题详情页
    problemDetail (raw) {
      this.selfLog(raw);
      this.$router.push(`/problemDetail/${raw._id}`);
    },
    // 展示tag的抽屉
    show_tag_drawer () {
      this.tag_drawer = true
    },
    // 按tag筛选题目
    choose_tag_problem (tag) {
      this.selfLog(tag)
      this.$axios.problem_list_page(this.offset, this.limit, this.page, tag.name).then(res => {
        this.selfLog(res);
        if (res.error === null) {
          this.total = res.data.total;
          this.problems_list = res.data.results;
          this.selfLog(this.problems_list[0]);
          this.selfLog(this.problems_list.length)
          let arange = this.limit
          if (this.limit > this.problems_list.length) {
            arange = this.problems_list.length
          }
          for (let i = 0; i < arange; i++) {
            // 离散化难度值
            if (this.problems_list[i].difficulty == "Low") {
              this.problems_list[i].difficulty = 1;
            } else if (this.problems_list[i].difficulty == "Mid") {
              this.problems_list[i].difficulty = 2;
            } else if (this.problems_list[i].difficulty == "High") {
              this.problems_list[i].difficulty = 3;
            }

            if (this.problems_list[i].submission_number === 0) {
              this.problems_list[i]["pass_rate"] = 0;
            } else {
              if (this.problems_list[i].accepted_number == 0) {
                this.problems_list[i]["pass_rate"] = 0;
              } else {
                let x = (
                  (this.problems_list[i].accepted_number /
                    this.problems_list[i].submission_number) *
                  100
                ).toFixed(2);
                this.problems_list[i]["pass_rate"] = parseFloat(x);
              }
              this.selfLog(this.problems_list[i]["pass_rate"]);
            }
          }
        }
        this.selfLog(this.problems_list);
      });
    },
  }
};
</script>

<style lang='scss' scoped>
#main {
  background-color: #3d84a8;
  width: 100%;
  line-height: 600px;
}
.el-table .warning-row {
  background: oldlace;
}

.el-table .success-row {
  background: #f0f9eb;
}
</style>