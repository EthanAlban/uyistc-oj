<template>
  <div>
    <el-form :inline="true" :model="formInline" style="margin:10px 0 0 20px;postion:flex;float:left">
      <el-form-item label="作业查询">
        <el-input v-model="formInline.keyword" placeholder="输入关键字"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="search_task">查询</el-button>
      </el-form-item>
    </el-form>
    <el-button type="primary" @click="dialogVisible = true" style="float: left; margin:10px 0 0 30%">
      <i class="el-icon-plus"></i>
      创建作业
    </el-button>
    <div>
      <el-table :data="tableData">
        <el-table-column label="竞赛" width="120" prop="title"></el-table-column>
        <el-table-column label="作业描述" width="220" prop="description">
          <template slot-scope="scope">
            <div v-html="scope.row.description"></div>
          </template>
        </el-table-column>
        <el-table-column label="截止时间" width="220" sortable prop="end_time">
          <template slot-scope="scope">
            <i class="el-icon-time"></i>
            <span style="margin-left: 10px">{{ scope.row.end_time | formatDate}}</span>
          </template>
        </el-table-column>
        <el-table-column width="120" label="创建人" prop="created_by.username"></el-table-column>
        <el-table-column width="120" label="可见状态" prop="visible">
          <template slot-scope="scope">
            <el-switch v-model="scope.row.visible" active-color="#13ce66"
              @change="ruleForm=scope.row,submitForm_update('ruleForm','put')" inactive-color="#ff4949"></el-switch>
          </template>
        </el-table-column>
        <el-table-column label="管理" header-align="center" width="450">
          <template slot-scope="scope">
            <el-button size="mini" icon="el-icon-view" type="primary" round @click="enter_assign_detail(scope.row.id)">
              查看</el-button>
            <el-button size="mini" icon="el-icon-edit" type="success" round @click=" update_assign(scope.row.id)">编辑
            </el-button>
            <el-button size="mini" icon="el-icon-edit" type="primary" round
              @click="show_drawer_for_annoce(scope.row.id)">
              查看公告</el-button>
            <!-- <el-button size="mini" icon="el-icon-download" type="info" round @click="enter_assign_detail(scope.row.id)">
              下载已提交作业</el-button> -->
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="page"
      :page-sizes="[10]" :page-size="100" layout="total, sizes, prev, pager, next, jumper" :total="total">
    </el-pagination>

    <el-drawer title="作业提交详情" :visible.sync="drawer" direction="rtl" size="85%" append-to-body>
      <SubmissionAys :subssion_detail="contest_submissions">
      </SubmissionAys>
    </el-drawer>

    <el-dialog title="创建作业" :visible.sync="dialogVisible" width="60%" destroy-on-close>
      <el-steps :active="active" finish-status="success" align-center>
        <el-step title="1 选择题目"></el-step>
        <el-step title="2 配置作业选项"></el-step>
      </el-steps>
      <div v-if="active===0">
        <!-- 这边题目能点进去查看题目详情 -->
        <el-table :data="problemData" size="medium" @selection-change="handleSelectionChange" ref="multipleTable"
          style="padding-left: 23%;" max-height="600">
          <el-table-column type="selection" width="55"></el-table-column>
          <el-table-column label="题目编号" sortable width="120" prop="id"></el-table-column>
          <el-table-column label="标题" width="120" prop="title"></el-table-column>
          <el-table-column label="标签" width="200">
            <template slot-scope="scope">
              <el-tag v-for="(tag,idx) in scope.row.tags" :key="tag" :type="tag_types[idx%5]" effect="dark">{{ tag }}
              </el-tag>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <!-- 创建/修改作业详情对话框 -->
      <div v-if="active===1 ">
        <el-form :model="ruleForm" ref="ruleForm" label-width="100px" class="demo-ruleForm">
          <el-form-item label="id" prop="id" required>
            <el-input v-model="ruleForm.id" clearable></el-input>
          </el-form-item>
          <el-form-item label="名称" prop="title" required>
            <el-input v-model="ruleForm.title" clearable></el-input>
          </el-form-item>
          <el-form-item label="描述" prop="description" required>
            <el-input type="textarea" v-model="ruleForm.description" clearable></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="password" required>
            <el-input type="password" v-model="ruleForm.password" clearable></el-input>
          </el-form-item>
          <el-form-item label="作业时间范围" required>
            <el-col :span="11">
              <el-form-item prop="date1">
                <el-date-picker type="date" placeholder="选择日期" v-model="ruleForm.start_time" style="width: 100%;">
                </el-date-picker>
              </el-form-item>
            </el-col>
            <el-col class="line" :span="2">-</el-col>
            <el-col :span="11">
              <el-form-item prop="date2">
                <el-date-picker typeplaceholder="选择日期" v-model="ruleForm.end_time" style="width: 100%;">
                </el-date-picker>
              </el-form-item>
            </el-col>
          </el-form-item>
        </el-form>
      </div>
      <el-button class="stepButton" type="primary" @click="active=1" v-show="active==0">下一步</el-button>
      <div style="display:flex;float:right;marginTop:-40px;marginRight:230px">
        <el-button type="primary" @click="active=0" v-show="active==1">上一步</el-button>
        <el-button type="primary" @click="submitForm('ruleForm','post')" v-show="active==1">完成</el-button>
      </div>
    </el-dialog>
    <!-- 修改作业详情对话框 -->
    <el-dialog title="创建作业" :visible.sync="dialogVisible_update" width="80%" destroy-on-close>
      <el-form :model="ruleForm" ref="ruleForm" label-width="100px" class="demo-ruleForm">
        <el-form-item label="id" prop="id" required>
          <el-input v-model="ruleForm.id" clearable></el-input>
        </el-form-item>
        <el-form-item label="名称" prop="title" required>
          <el-input v-model="ruleForm.title" clearable></el-input>
        </el-form-item>
        <el-form-item label="描述" prop="description" required>
          <el-input type="textarea" v-model="ruleForm.description" clearable></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password" required>
          <el-input type="text" v-model="ruleForm.password"></el-input>
        </el-form-item>
        <el-form-item label="作业时间" required>
          <el-col :span="11">
            <el-form-item prop="date1">
              <el-date-picker type="date" placeholder="选择日期" v-model="ruleForm.start_time" style="width: 100%;">
              </el-date-picker>
            </el-form-item>
          </el-col>
          <el-col class="line" :span="2">-</el-col>
          <el-col :span="11">
            <el-form-item prop="date2">
              <el-date-picker typeplaceholder="选择日期" v-model="ruleForm.end_time" style="width: 100%;">
              </el-date-picker>
            </el-form-item>
          </el-col>
        </el-form-item>

      </el-form>
      <el-button type="primary" @click="submitForm_update('ruleForm','put')">保存</el-button>
      <el-button type="primary" @click="edit_problem_list()">修改题目列表</el-button>
    </el-dialog>

    <!-- 新建公告 -->
    <el-dialog title="发布公告" :visible.sync="show_new_annonce_editor">
      <el-form :model="new_annonce_form">
        <el-form-item label="标题" :label-width="formLabelWidth">
          <el-input v-model="new_annonce_form.title" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="内容" :label-width="formLabelWidth">
          <textarea v-model="new_annonce_form.content" autocomplete="off" style="height:200px;width:100%"></textarea>
        </el-form-item>
        <el-form-item label="可见性" :label-width="formLabelWidth">
          <el-switch v-model="new_annonce_form.visible" active-color="#13ce66" inactive-color="#ff4949">
          </el-switch>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="show_editor = false">取 消</el-button>
        <el-button type="primary" @click="update_annocement_new(1)">确 定</el-button>
      </div>
    </el-dialog>
    <!--公告列表抽屉 -->
    <el-drawer title="公告列表" :visible.sync="contest_annon_list_drawer_visble" direction="rtl" size="50%">
      <el-table :data="contest_annon_list">
        <el-table-column property="title" label="标题" width="150"></el-table-column>
        <el-table-column property="content" label="内容" width="200"></el-table-column>
        <el-table-column prop="visible" label="可见">
          <template slot-scope="scope">
            <el-switch v-model="scope.row.visible" active-color="#13ce66" inactive-color="#ff4949"
              @change="update_annocement(scope.row.id)"></el-switch>
          </template>
        </el-table-column>
        <el-table-column property="content" label="" width="100">
          <template slot-scope="scope">
            <i class="el-icon-delete" @click="delete_annonce(scope.row.id)"></i>
          </template>
        </el-table-column>
      </el-table>
      <el-button type="primary" round @click="show_editor_for_annoce()">发布公告</el-button>
    </el-drawer>
    <!-- 修改作业的题目列表的对话框 -->
    <el-dialog title="题目列表" :visible.sync="contest_problem_list_drawer_visble" center>
      <el-transfer v-model="value" :data="data" :titles="titles" @change="problem_list_change">
      </el-transfer>
      <!-- <el-button type="primary" round @click="save_problem_list()">保存</el-button> -->
    </el-dialog>
  </div>
</template>

<script>
import * as echarts from "echarts/core";
import {
  TooltipComponent,
  LegendComponent
} from "echarts/components";
import {
  PieChart
} from "echarts/charts";
import {
  CanvasRenderer
} from "echarts/renderers";
// 引入作业提交分析组件
import SubmissionAys from "./components/submission_anysis.vue"
echarts.use([TooltipComponent, LegendComponent, PieChart, CanvasRenderer]);

export default {
  components: {
    SubmissionAys
  },
  name: "Assign",
  beforeCreate () {
    this.$axios.user_profile().then(res => {
      this.user_profile = res["data"];
      this.$root.user_profile = this.user_profile;
      this.user_profile.user.last_login = this.user_profile.user.last_login.substr(
        0,
        10
      );
      this.user_profile.user.avatar = this.$OJIP + this.user_profile.avatar;
      this.getTableData();
      this.getProblemList();
    });
  },
  data () {
    return {
      curAssign: {
        problemDetail: ""
      },
      drawer: false, //弹框
      tag_types: ["info", "warning", "danger", "", "success"],
      user_profile: "",
      ruleForm: {
        allowed_ip_ranges: [],
        description: "",
        end_time: "",
        password: "",
        real_time_rank: true,
        rule_type: "ACM",
        start_time: "",
        title: "",
        visible: true
      },
      problemData: [], //题目列表
      multipleSelection: [],
      dialogVisible: false,
      active: 0,
      activeNames: [],
      page: 1,
      classList: [],
      tableData: [], //作业列表数据
      offset: 0,
      limit: 10,
      totalListLength: 20,
      formLabelWidth: "80px",
      total: 0,
      form: {
        password: ""
      },
      formInline: {
        keyword: "",
        status: ""
      },
      problems: {},
      key: 0,
      contestid: 0,
      homework_id: 0,
      config: {
        data: [{
          name: "已提交",
          value: 60
        },
        {
          name: "未提交",
          value: 40
        }
        ],
        color: ["#00FF7F", "#f00"],
        digitalFlopStyle: {
          fill: "#000"
        },
        showOriginValue: true
      },
      option: {
        tooltip: {
          trigger: "item"
        },
        legend: {
          top: "5%",
          left: "center"
        },
        series: [{
          name: "提交分析",
          type: "pie",
          radius: ["40%", "60%"],
          avoidLabelOverlap: false,
          itemStyle: {
            borderRadius: 10,
            borderColor: "#fff",
            borderWidth: 2
          },
          label: {
            show: false,
            position: "center"
          },
          emphasis: {
            label: {
              show: true,
              fontSize: "40",
              fontWeight: "bold"
            }
          },
          labelLine: {
            show: false
          },
          data: []
        }]
      },
      new_annonce_form: {
        content: "",
        contest_id: "0",
        id: null,
        title: "",
        visible: true,
      },
      show_new_annonce_editor: false,
      contest_id_for_annon: 0,
      //   公告列表
      contest_annon_list: [],
      contest_annon_list_drawer_visble: false,
      dialogVisible_update: false,
      contest_problem_list_drawer_visble: false,
      //   问题列表 //   作业中问题列表穿梭框
      data: [],
      value: [],
      titles: ['可选题目', '已选题目'],
      //   暂存上一步的已选列表
      value2: [],
      //   作业的所有提交
      contest_submissions: []
    };
  },
  methods: {
    setEcharts () { },
    //上传题目
    upload (id) {
      for (let i = 0; i < this.multipleSelection.length; i++) {
        let params = {
          contest_id: id,
          display_id: this.multipleSelection[i]._id,
          problem_id: this.multipleSelection[i].id
        };
        this.$axios.add_problem(params).then(res => { });
      }
    },
    enter_assign_detail (contest_id) {
      this.drawer = true;
      this.$axios.search_problem_list(this.limit, this.offset, contest_id).then(res => {
        this.problemDetail = res.data;
      });
      // 获取某个作业所有的提交
      this.$axios.assign_submission_list(contest_id).then(res => {
        this.selfLog(res)
        for (let submission of res.data.results) {
          submission["similarity"] = "0"
        }
        this.contest_submissions = res.data.results
      })
    },
    handleCurrentChange (val) {
      this.page = val;
      this.offset = (this.page - 1) * this.limit;
      this.getProblemList();
    },
    handleSizeChange (val) {
      this.limit = val;
      this.offset = (this.page - 1) * this.limit;
      this.getProblemList();
    },
    //搜索函数
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
    //获取作业列表
    getProblemList () {
      this.$axios.admin_contests_list(this.offset, this.limit).then(res => {
        this.selfLog(res);
        this.tableData = []
        // this.selfLog(this.user_profile)
        if (res.error === null) {
          for (let key of res.data.results) {
            if (key.created_by.username == this.user_profile.user.username) {
              this.tableData.push(key);
            }
          }
          this.total = res.data.total;
          this.selfLog(this.tableData);
          this.selfLog(this.total);
        }
      });
    },
    //获取题目列表
    async getTableData () {
      await this.$axios.admin_problem_list(this.totalListLength, this.offset).then(res => {
        this.totalListLength = res.data.total;
      })
      this.$axios.admin_problem_list(this.totalListLength, this.offset).then(res => {
        if (res.error == null) {
          this.problemData = res.data.results;
          this.totalListLength = res.data.total;
          this.selfLog(this.problemData);
          this.selfLog(this.totalListLength);
        } else {
          this.$message({
            message: res.error,
            type: "error"
          });
        }
      });
    },
    // 显示新增公告对话框
    show_editor_for_annoce () {
      this.show_new_annonce_editor = true
    },
    // 显示公告列表抽屉
    async show_drawer_for_annoce (id) {
      this.contest_id_for_annon = id
      await this.get_annonce_for_contest(id)
      this.contest_annon_list_drawer_visble = true
    },
    // 给作业发布公告
    update_annocement_new () {
      let param = {
        content: this.new_annonce_form.content,
        contest_id: this.contest_id_for_annon,
        id: null,
        title: this.new_annonce_form.title,
        visible: this.new_annonce_form.visible,
      }
      this.selfLog(param);
      this.$axios.publish_annoncement_for_assign(param).then(res => {
        if (res.error === null) {
          this.$message({
            message: '发布公告成功',
            type: 'success'
          });
          this.get_annonce_for_contest(this.contest_id_for_annon)
          this.show_new_annonce_editor = false
        }
      })
    },
    // 查询公告列表
    get_annonce_for_contest (id) {
      this.$axios.get_annonce_for_contest(id).then(res => {
        this.selfLog(res);
        this.contest_annon_list = res.data
      })
    },
    // 删除作业公告
    delete_annonce (annonce_id) {
      //   this.selfLog(annonce_id);
      this.$axios.delete_annonce(annonce_id).then(res => {
        this.selfLog(res);
        if (res.error === null) {
          this.$message({
            message: '删除公告成功',
            type: 'success'
          });
          this.get_annonce_for_contest(this.contest_id_for_annon)
        }
      })
    },
    // 为修改作业准备数据
    async update_assign (id) {
      this.selfLog(id);
      await this.get_detail_assign_id(id)
      this.selfLog(this.ruleForm);
      this.dialogVisible_update = true
    },
    // 获取作业详情
    get_detail_assign_id (id) {
      this.$axios.get_detail_assign_id(id).then(res => {
        this.ruleForm = res.data
      })
    },
    // 创建/修改作业
    submitForm (formName, method) {
      let parms = {
        allowed_ip_ranges: [],
        description: this.ruleForm.description,
        end_time: this.ruleForm.end_time,
        password: this.ruleForm.password,
        real_time_rank: true,
        rule_type: "ACM",
        start_time: this.ruleForm.start_time,
        title: this.ruleForm.title,
        visible: true,
        id: 1
      };
      this.$axios
        .create_task(parms, method)
        .then(res => {
          this.selfLog(res);
          this.homework_id = res.data.id;
          this.dialogVisible = false;
          this.active = 0;
          this.clearable();
          this.getProblemList()
          this.upload(this.homework_id);
        })
    },
    // 修改作业
    submitForm_update (formName, method) {
      console.log(this.ruleForm);
      let parms = {
        allowed_ip_ranges: [],
        description: this.ruleForm.description,
        end_time: this.ruleForm.end_time,
        password: this.ruleForm.password,
        real_time_rank: true,
        rule_type: "ACM",
        start_time: this.ruleForm.start_time,
        title: this.ruleForm.title,
        visible: this.ruleForm.visible,
        id: this.ruleForm.id
      };
      this.selfLog(method);
      this.$axios
        .create_task(parms, method)
        .then(res => {
          this.selfLog(res);
          if (res.error === null) {
            this.$message({
              message: '修改成功',
              type: 'success'
            });
          }
          this.dialogVisible_update = false;
          this.getProblemList()
          this.ruleForm = {}
        })
    },
    // 编辑作业的问题列表
    async edit_problem_list () {
      this.data = []
      this.contest_problem_list_drawer_visble = true
      await this.getTableData()
      let list = []
      let list_objetc = []
      for (let problem of this.problemData) {
        list.push(problem.title)
        list_objetc.unshift(problem)
      }
      list.forEach((title, index) => {
        this.data.push({
          label: title,
          key: index,
          problem: list_objetc.pop()
        });
      });
      this.$axios.search_problem_list(this.limit, this.offset, this.ruleForm.id).then(res => {
        this.selfLog(res.data.results);
        for (let problem of res.data.results) {
          //   this.selfLog("全部data");
          //   this.selfLog(this.data);
          this.data.forEach((title, index) => {
            // this.selfLog("--------");
            // this.selfLog(this.data[index].problem);
            // this.selfLog(problem._id + ' === ' + this.data[index].problem.id);
            // this.selfLog(problem._id + '' === this.data[index].problem.id + '');
            // this.selfLog(this.data[index]);
            if (problem._id + '' === this.data[index].problem.id + '') {
              this.selfLog("压入" + index);
              this.value.push(index)
              this.value2 = this.value
            }
          });
        }
        this.selfLog(this.value);
      })
    },
    // 保存选择了的问题列表
    // save_problem_list () {
    //   this.selfLog(this.value);
    //   for (let id of this.value) {
    //     this.selfLog(this.data[id].problem);
    //     let params = {
    //       contest_id: this.ruleForm.id,
    //       display_id: this.data[id].problem.id,
    //       problem_id: this.data[id].problem.id
    //     };
    //     this.$axios.add_problem(params).then(res => {
    //       this.$message({
    //         message: '添加成功',
    //         type: 'success'
    //       });
    //     });
    //   }
    // },
    // 作业问题列表改变
    problem_list_change () {
      // value是右侧的现在的 value2是右侧改变之前的内容
      this.selfLog("最新value：");
      this.selfLog(this.value);
      this.selfLog("最新value2：");
      this.selfLog(this.value2);
      for (let index of this.value) {
        //   如果原来列表就有的不管他
        if (this.value2.includes(index)) { }
        // 如果是原来列表中没有的话 新增
        // else if (!(index in this.value2)) {
        else if (!this.value2.includes(index)) {
          this.selfLog("++++++++++++++++++++++++++++++++++")
          this.selfLog("this.data[index].problem");
          this.selfLog(this.data[index].problem);
          let params = {
            contest_id: this.ruleForm.id,
            display_id: this.data[index].problem.id,
            problem_id: this.data[index].problem.id
          };
          this.$axios.add_problem(params).then(res => {
            this.selfLog("---------");
            this.selfLog(res);
            this.$message({
              message: '保存' + this.data[index].problem.title + '成功',
              type: 'success'
            });
          });
        }
      }
      //   在原来的列标中确不在新的列表中表示被删除了
      for (let index of this.value2) {
        this.selfLog("indexxx:" + index)
        this.selfLog(this.value.includes(index))
        // if (!(index in this.value)) {
        if (!this.value.includes(index)) {

          let problemid_to_delte = this.data[index].problem.id + ''
          //   this.selfLog(this.data[index].problem); //id:34 _id:'jhc-test'
          this.$axios.search_problem_list(this.limit, this.offset, this.ruleForm.id).then(res => {
            for (let problem of res.data.results) {
              if (problemid_to_delte === problem._id) {
                this.selfLog(this.data[index].problem.id);
                this.selfLog(problem.id)
                this.$axios.delete_problem_from_assign(problem.id).then(res => {
                  this.selfLog(res);
                  if (res.error === null) {
                    this.$message({
                      message: '删除成功',
                      type: 'success'
                    });
                  } else if (res.data === "Can't delete the problem as it has submissions") {
                    this.$message({
                      message: '该题目已有提交记录，无法移除',
                      type: 'error'
                    });
                    this.selfLog("最新====value：");
                    this.selfLog(this.value);
                    this.selfLog("最新====value2：");
                    this.selfLog(this.value2);
                    this.edit_problem_list()
                  }

                })
              }
            }
          })
        }
      }
      this.value2 = this.value
    },
    handleSelectionChange (val) {
      this.selfLog(val);
      this.multipleSelection = val;
    },
    handleClose (done) {
      done();
    },
    handleChange (val) {
      this.selfLog(val);
    },
    handleClick (row) {
      this.selfLog(row);
    }
  },
  computed: {}
};
</script>

<style lang='scss' scoped>
.demo-ruleForm {
  width: 60%;
  margin-left: 20%;
}

.stepButton {
  margin-top: 20px;
  width: 60%;
}

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
