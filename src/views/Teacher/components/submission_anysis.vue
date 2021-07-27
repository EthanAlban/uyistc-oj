<template>
  <div>
    <el-button type="primary" @click="update_similaryty" size="mini" round>计算重复率</el-button>
    <el-button type="primary" @click="export2Excel" size="mini" round>导出通过学生名单</el-button>
    <el-table :data="subssion_detail" stripe style="width: 100%" v-loading="loading"
      element-loading-text="正在计算各个提交之间重复率，提交数量较多，请您耐心等待~" max-height="600" element-loading-spinner="el-icon-loading">
      <el-table-column prop="create_time" label="提交时间" width="200" sortable>
        <template slot-scope="scope">
          <i class="el-icon-time"></i>
          <span style="margin-left: 10px">{{
            scope.row.create_time | formatDate
          }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="id" label="提交ID" width="170">
        <template slot-scope="scope">
          <el-tooltip :content=scope.row.id placement="top">
            <el-link class="sub_id">{{ scope.row.id }}</el-link>
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column prop="result" label="结果" width="80" sortable>
        <template slot-scope="scope">
          <el-tag v-if="scope.row.result === -2" effect="dark" type="danger">编译错误</el-tag>
          <el-tag v-if="scope.row.result === -1" effect="dark" type="danger">答案错误</el-tag>
          <el-tag v-if="scope.row.result === 0" effect="dark" type="success">正确</el-tag>
          <el-tag v-if="scope.row.result === 1" effect="dark" type="warning">计算超时</el-tag>
          <el-tag v-if="scope.row.result === 2" effect="dark" type="warning">超时</el-tag>
          <el-tag v-if="scope.row.result === 3" effect="dark" type="warning">内存超过</el-tag>
          <el-tag v-if="scope.row.result === 4" effect="dark" type="danger">运行时错误</el-tag>
          <el-tag v-if="scope.row.result === 5" effect="dark" type="danger">系统错误</el-tag>
          <el-tag v-if="scope.row.result === 6" effect="dark" type="info">传送...</el-tag>
          <el-tag v-if="scope.row.result === 7" effect="dark" type="info">判题中...</el-tag>
          <el-tag v-if="scope.row.result === 8" effect="dark">部分正确</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="problem" label="问题ID" width="100" sortable>
      </el-table-column>
      <el-table-column prop="statistic_info.time_cost" label="运行时间" width="100">
        <template slot-scope="scope">
          <span v-if="scope.row.result === 0" style="color:green">{{scope.row.statistic_info.time_cost}}ms</span>
          <span v-else style="color:red">{{scope.row.statistic_info.time_cost}}--ms</span>
        </template>
      </el-table-column>
      <el-table-column prop="statistic_info.memory_cost" label="运行内存" width="100">
        <template slot-scope="scope">
          <span v-if="scope.row.result === 0" style="color:green">{{scope.row.statistic_info.memory_cost}}</span>
          <span v-else style="color:red">{{scope.row.statistic_info.memory_cost}}--</span>
        </template>
      </el-table-column>
      <el-table-column prop="language" label="提交语言" width="80">
      </el-table-column>
      <el-table-column prop="username" label="用户" width="100">
      </el-table-column>
      <el-table-column prop="similarity" label="最高重复率" width="120" sortable>
        <template slot-scope="scope">
          <el-tag v-if="scope.row.similarity > 80" effect="dark" type="danger">{{scope.row.similarity}}%</el-tag>
          <el-tag v-else-if="scope.row.similarity > 60" effect="dark" type="warning">{{scope.row.similarity}}%</el-tag>
          <el-tag v-else-if="scope.row.similarity > 40" effect="dark" type="info">{{scope.row.similarity}}%</el-tag>
          <el-tag v-else-if="scope.row.similarity >= 0" effect="dark" type="success">{{scope.row.similarity}}%</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="similarity" label="对比代码" width="150">
        <template slot-scope="scope">
          <el-button type="primary" @click="constract_two_codes(scope.row.id)">比对</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-drawer title="代码对比" label="ltr" :append-to-body="true" :visible.sync="innerDrawer" size="50%">
      <el-tag type="success">本提交代码(学生：{{codes_compare.my_code_username}})</el-tag>
      <!-- <codemirror style="textAlign:left" ref="codes_compare.my_code" :value="codes_compare.my_code"
        :options="cmOptions_C" @ready="onCmReady" @focus="onCmFocus" /> -->
      <el-input type="textarea" :autosize="{ minRows: 6, maxRows:10}" v-model="codes_compare.my_code"></el-input>
      <el-tag type="success" style="margin-top:15px">比对提交代码(学生：{{codes_compare.copy_code_username}})
      </el-tag>
      <el-input type="textarea" :autosize="{ minRows: 6, maxRows:10}" v-model="codes_compare.copy_code"></el-input>
    </el-drawer>
  </div>
</template>

<script>
export default {
  name: "SubmissionAys",
  props: ["subssion_detail"],
  data () {
    return {
      submisson_codes: {},
      loading: false,
      copy_id_dict: {},
      codes_compare: {
        my_code: "",
        my_code_username: "",
        copy_code: "",
        copy_code_username: ""
      },
      innerDrawer: false,
      cmOptions_C: {
        tabSize: 4,
        mode: "text/x-csrc",
        theme: "monokai",
        lineNumbers: true,
        line: true
      },
    };
  },
  mounted () {
    //   挂好之后组建出一个所有提交的代码的集合
    this.assemble_submisson_codes()
  },
  methods: {
    async assemble_submisson_codes () {
      //   this.selfLog(this.subssion_detail)
      let submisson_codes = {}
      //   let requests = []
      for (let submission of this.subssion_detail) {
        // requests.push(this.$axios.submission_id(submission.id))
        await this.$axios.submission_id(submission.id).then(res => {
          submisson_codes[res.data.id] = res.data.code
        })
      }
      this.submisson_codes = submisson_codes

    },
    async update_similaryty () {
      this.loading = true
      await this.assemble_submisson_codes()
      let submisson_codes = this.submisson_codes
      for (let submission of this.subssion_detail) {
        let param = {
          new_code: submisson_codes[submission.id],
          codes: submisson_codes,
          submission_id: submission.id
        }
        //this.selfLog(param)
        this.$axios.submisson_code_similarity(param).then(res => {
          //   this.selfLog(res)
          submission["similarity"] = parseFloat(res.max_similarity.substring(0, 5))
          this.copy_id_dict[submission.id] = res.copy_submission_id
        })
      }
      this.selfLog(this.subssion_detail)
      this.loading = false
    },
    // 对比两次相似的提交的代码
    constract_two_codes (submission_id) {
      this.selfLog(this.copy_id_dict[submission_id])
      if (this.copy_id_dict[submission_id] === "0") {
        this.$message('该提交暂无相似代码');
      }
      this.selfLog("本提交：" + submission_id)
      this.selfLog("涉嫌抄袭提交：" + this.copy_id_dict[submission_id])
      this.$axios.submission_id(submission_id).then(res => {
        this.selfLog(res)
        this.codes_compare.my_code_username = res.data.username
        this.codes_compare.my_code = res.data.code
      })
      this.$axios.submission_id(this.copy_id_dict[submission_id]).then(res => {
        this.codes_compare.copy_code_username = res.data.username
        this.codes_compare.copy_code = res.data.code
      })
      this.innerDrawer = true

    },
    export2Excel () {
      require.ensure([], () => {
        const { export_json_to_excel } = require('../export/Export2Excel');
        const tHeader = ['提交时间', '提交ID', '结果', "问题ID", "提交语言", "用户", "重复率"];
        // 上面设置Excel的表格第一行的标题
        const filterVal = ['create_time', 'id', 'result', 'problem', 'language', 'username', 'similarity'];
        // 上面的index、phone_Num、school_Name是tableData里对象的属性
        const list = this.subssion_detail;  //把data里的tableData存到list
        const list2 = []
        let pass_users = []
        for (let submission of list) {
          if (submission.result === 0 && !(pass_users.includes(submission.username))) {
            list2.push(submission)
            pass_users.push(submission.username)
          }
        }
        const data = this.formatJson(filterVal, list2);
        export_json_to_excel(tHeader, data, '列表excel');
      })
    },
    formatJson (filterVal, jsonData) {
      return jsonData.map(v => filterVal.map(j => v[j]))
    },
    onCmReady (cm) {
      this.selfLog("the editor is readied!", cm);
    },
    onCmFocus (cm) {
      this.selfLog("the editor is focused!", cm);
    },
  },
  computed: {},
};
</script>

<style lang='scss' scoped>
.sub_id {
  // width: 180px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
