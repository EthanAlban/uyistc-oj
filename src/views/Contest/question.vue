<template>
  <div>
    <el-button type="primary" round class="createNew" @click="dialogFormVisible = true">创建新题目</el-button>
    <div class="problem_list">
      <el-table :data="tableData" size="medium">
        <el-table-column label="题目编号" sortable width="100" prop="id"></el-table-column>
        <el-table-column label="编号" width="80" prop="_id"></el-table-column>
        <el-table-column label="标题" width="120" prop="title"></el-table-column>
        <el-table-column label="创建人" width="100">
          <template slot-scope="scope">
            <i class="el-icon-s-custom"></i>
            <span style="margin-left: 10px">{{ scope.row.created_by.username}}</span>
          </template>
        </el-table-column>
        <el-table-column label="标签" width="250">
          <template slot-scope="scope">
            <el-tag v-for="(tag,idx) in scope.row.tags" :key="tag" :type="tag_types[idx%5]" effect="dark">{{ tag }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="创建时间" width="200" sortable prop="create_time">
          <template slot-scope="scope">
            <i class="el-icon-time"></i>
            <span style="margin-left: 10px">{{ scope.row.create_time | formatDate}}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="400" header-align="center">
          <template slot-scope="scope">
            <el-button size="mini" icon="el-icon-edit" type="success" round @click="edit_problem(scope.row)">编辑
            </el-button>
            <el-button size="mini" icon="el-icon-delete" type="danger" round @click="delete_problem(scope.row)">删除
            </el-button>
            <a
              :href="'http://47.94.135.51:9999/foo/problem/download_test_case?problem_id='+scope.row.id+'&testcase_id='+scope.row.test_case_id">
              <el-button size="mini" icon="el-icon-download" style="margin-left:5px" type="primary" round>下载用例
              </el-button>
            </a>
            <el-button size="mini" icon="el-icon-view" style="margin-left:5px" type="primary" round
              @click="show_test_case(scope.row)">
              查看用例</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="1"
        :page-sizes="[5, 10, 20]" :page-size="10" layout="total, sizes, prev, pager, next, jumper"
        :total="totalListLength"></el-pagination>
    </div>

    <el-dialog :visible.sync="dialogFormVisible" destroy-on-close>
      <p class="formTitle" style="fontSize:20px;marginTop:-40px">创建题目</p>
      <el-form :model="form" :inline="true" style="width:600px;marginLeft:-40px">
        <el-form-item label="题目编号">
          <el-input v-model="form._id" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="题目名称">
          <el-input v-model="form.title" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <!-- 使用富文本编辑器 -->
      <p class="formTitle">题目描述</p>
      <quill-editor :content="form.description" :options="editorOption" @change="onDescriptionChange($event)" />
      <p class="formTitle">输入描述</p>
      <quill-editor :content="form.input_description" :options="editorOption" @change="onInputDescChange($event)" />
      <p class="formTitle">输出描述</p>
      <quill-editor :content="form.output_description" :options="editorOption" @change="onOutDescChange($event)" />
      <p class="formTitle">提示</p>
      <quill-editor :content="form.hint" :options="editorOption" @change="onHintChange($event)" />

      <el-form :model="form" :inline="true">
        <el-form-item label="时间限制(ms)">
          <el-input-number v-model="form.time_limit"></el-input-number>
        </el-form-item>
        <el-form-item label="内存限制(MB)">
          <el-input-number v-model="form.memory_limit"></el-input-number>
        </el-form-item>
        <el-form-item label="难度">
          <el-select v-model="form.difficulty" placeholder="请选择问题难度">
            <el-option label="低" value="Low"></el-option>
            <el-option label="中等" value="Mid"></el-option>
            <el-option label="困难" value="High"></el-option>
          </el-select>
        </el-form-item>
      </el-form>

      <el-form :model="form" label-position="top">
        <el-form-item label="问题标签">
          <el-tag :key="tag" v-for="tag in form.tags" closable :disable-transitions="false" @close="handleClose(tag)">
            {{tag}}</el-tag>
          <el-input class="input-new-tag" v-if="inputVisible" v-model="inputValue" ref="saveTagInput" size="small"
            @blur="handleInputConfirm"></el-input>
          <el-button v-else class="button-new-tag" size="small" @click="showInput">+ 新标签</el-button>
        </el-form-item>
        <el-form-item label="问题可使用语言">
          <el-checkbox-group v-model="form.languages">
            <el-checkbox label="C"></el-checkbox>
            <el-checkbox label="C++"></el-checkbox>
            <el-checkbox label="Python"></el-checkbox>
            <el-checkbox label="Java"></el-checkbox>
          </el-checkbox-group>
        </el-form-item>
      </el-form>
      <!-- 显示各种语言模板 -->
      <codemirror v-if="form.languages.indexOf('C') > -1" style="textAlign:left" ref="cmEditor_C" :value="code_C"
        :options="cmOptions_C" @ready="onCmReady" @focus="onCmFocus" @input="onCmCodeChange_C" />
      <codemirror v-if="form.languages.indexOf('C++') > -1" style="textAlign:left;" ref="cmEditor_C_Plus"
        :value="code_C_Plus" :options="cmOptions_C_Plus" @ready="onCmReady" @focus="onCmFocus"
        @input="onCmCodeChange_C_Plus" />
      <codemirror v-if="form.languages.indexOf('Python') > -1" style="textAlign:left;" ref="cmEditor_Python"
        :value="code_Python" @ready="onCmReady" @focus="onCmFocus" :options="cmOptions_Python"
        @input="onCmCodeChange_Python" />
      <codemirror v-if="form.languages.indexOf('Java') > -1" style="textAlign:left;" ref="cmEditor_Java"
        :value="code_Java" @ready="onCmReady" @focus="onCmFocus" :options="cmOptions_Java"
        @input="onCmCodeChange_Java" />
      <!-- 使用v-for渲染组件 -->
      <el-divider>
        <i class="el-icon-mobile-phone"></i>
      </el-divider>
      <Sample class="sampleBox" v-for="item in sampleList" :key="item.id" :post="item" v-on:delete="deleteSample">
      </Sample>

      <el-button type="primary" style="width:60%" @click="addSample">
        <i class="el-icon-plus"></i>新建样例
      </el-button>
      <el-divider>
        <i class="el-icon-mobile-phone"></i>
      </el-divider>

      <el-upload class="upload-demo" action="string" accept=".zip" :limit="1" :http-request="uploadTestCase"
        :on-remove="handleRemove" :before-remove="beforeRemove" multiple :on-exceed="handleExceed"
        :file-list="fileList">
        <el-button size="small" type="primary">点击上传测试用例</el-button>
        <div slot="tip" class="el-upload__tip">
          <a href="../testcase" target="_blank">测试用例格式说明</a>
        </div>
      </el-upload>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="add_problem('new')">确 定</el-button>
      </div>
    </el-dialog>
    <!-- 编辑问题对话框 -->
    <el-dialog :visible.sync="dialogFormVisible_edit" destroy-on-close>
      <p class="formTitle" style="fontSize:20px;marginTop:-40px">编辑题目</p>
      <el-form :model="form" :inline="true" style="width:600px;marginLeft:-40px">
        <el-form-item label="题目编号">
          <el-input v-model="form._id" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="题目名称">
          <el-input v-model="form.title" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <!-- 使用富文本编辑器 -->
      <p class="formTitle">题目描述</p>
      <quill-editor :content="form.description" :options="editorOption" @change="onDescriptionChange($event)" />
      <p class="formTitle">输入描述</p>
      <quill-editor :content="form.input_description" :options="editorOption" @change="onInputDescChange($event)" />
      <p class="formTitle">输出描述</p>
      <quill-editor :content="form.output_description" :options="editorOption" @change="onOutDescChange($event)" />
      <p class="formTitle">提示</p>
      <quill-editor :content="form.hint" :options="editorOption" @change="onHintChange($event)" />

      <el-form :model="form" :inline="true">
        <el-form-item label="时间限制(ms)">
          <el-input-number v-model="form.time_limit"></el-input-number>
        </el-form-item>
        <el-form-item label="内存限制(MB)">
          <el-input-number v-model="form.memory_limit"></el-input-number>
        </el-form-item>
        <el-form-item label="难度">
          <el-select v-model="form.difficulty" placeholder="请选择问题难度">
            <el-option label="低" value="Low"></el-option>
            <el-option label="中等" value="Mid"></el-option>
            <el-option label="困难" value="High"></el-option>
          </el-select>
        </el-form-item>
      </el-form>

      <el-form :model="form" label-position="top">
        <el-form-item label="问题标签">
          <el-tag :key="tag" v-for="tag in form.tags" closable :disable-transitions="false" @close="handleClose(tag)">
            {{tag}}</el-tag>
          <el-input class="input-new-tag" v-if="inputVisible" v-model="inputValue" ref="saveTagInput" size="small"
            @blur="handleInputConfirm"></el-input>
          <el-button v-else class="button-new-tag" size="small" @click="showInput">+ 新标签</el-button>
        </el-form-item>
        <el-form-item label="问题可使用语言">
          <el-checkbox-group v-model="form.languages">
            <el-checkbox label="C"></el-checkbox>
            <el-checkbox label="C++"></el-checkbox>
            <el-checkbox label="Python"></el-checkbox>
            <el-checkbox label="Java"></el-checkbox>
          </el-checkbox-group>
        </el-form-item>
      </el-form>
      <!-- 显示各种语言模板 -->
      <codemirror v-if="form.languages.indexOf('C') > -1 && current_problem.template['C']!==''" style="textAlign:left"
        ref="cmEditor_C" :value="current_problem.template['C']" :options="cmOptions_C" @ready="onCmReady"
        @focus="onCmFocus" @input="onCmCodeChange_C" />
      <codemirror v-if="form.languages.indexOf('C++') > -1 && current_problem.template['C++']!==''"
        style="textAlign:left;" ref="cmEditor_C_Plus" :value="code_C_Plus" :options="cmOptions_C_Plus"
        @ready="onCmReady" @focus="onCmFocus" @input="onCmCodeChange_C_Plus" />
      <codemirror v-if="form.languages.indexOf('Python') > -1 && current_problem.template['Python']!==''"
        style="textAlign:left;" ref="cmEditor_Python" :value="code_Python" @ready="onCmReady" @focus="onCmFocus"
        :options="cmOptions_Python" @input="onCmCodeChange_Python" />
      <codemirror v-if="form.languages.indexOf('Java') > -1 && current_problem.template['Java']!==''"
        style="textAlign:left;" ref="cmEditor_Java" :value="code_Java" @ready="onCmReady" @focus="onCmFocus"
        :options="cmOptions_Java" @input="onCmCodeChange_Java" />
      <!-- 使用v-for渲染组件 -->
      <el-divider>
        <i class="el-icon-mobile-phone"></i>
      </el-divider>
      <Sample class="sampleBox" v-for="item in sampleList" :key="item.id" :post="item" v-on:delete="deleteSample">
      </Sample>

      <el-button type="primary" style="width:60%" @click="addSample">
        <i class="el-icon-plus"></i>新建样例
      </el-button>
      <el-divider>
        <i class="el-icon-mobile-phone"></i>
      </el-divider>
      <el-table :data="form.test_case_score" stripe style="width: 100%">
        <el-table-column prop="input_name" label="输入" width="180">
        </el-table-column>
        <el-table-column prop="output_name" label="输出" width="180">
        </el-table-column>
        <el-table-column prop="score" label="分值">
        </el-table-column>
      </el-table>
      <el-upload class="upload-demo" action="string" accept=".zip" :limit="1" :http-request="uploadTestCase"
        :on-remove="handleRemove" :before-remove="beforeRemove" multiple :on-exceed="handleExceed"
        :file-list="fileList">
        <el-button size="small" type="primary">点击上传测试用例</el-button>
        <div slot="tip" class="el-upload__tip">
          <a href="../testcase" target="_blank">测试用例格式说明</a>
        </div>
      </el-upload>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="add_problem('edit')">确认修改</el-button>
      </div>
    </el-dialog>
    <!-- 测试用例抽屉 -->
    <el-drawer title="" :visible.sync="case_drawer" direction="rtl" size="65%">
      <span v-html="'题目：'+current_problem.title"> </span><br>
      <el-tooltip class="item" effect="light" content="输入说明" placement="right">
        <span style="color:green" v-html="current_problem.input_description">----------</span>
      </el-tooltip>
      <el-tooltip class="item" effect="light" content="输出说明" placement="right">
        <span v-html="current_problem.output_description">----------</span>
      </el-tooltip>

      <el-button type="primary" style="position:absolute;right:5%;top:5%" @click="insert_testcase" round>新建测试用例
      </el-button>
      <el-table :data="test_cases.cases" max-height="500" v-loading="testcase_loading">
        <el-table-column type="index" width="50">
        </el-table-column>
        <el-table-column property="case" label="输入" width="200">
          <template slot-scope="scope">
            <!-- <textarea style="padding-bottom: 50%; " :value="scope.row.case" /> -->
            <el-input type="textarea" :autosize="{ minRows: 1, maxRows:6}" v-model="scope.row.case"></el-input>
          </template>
        </el-table-column>
        <el-table-column property="answer" label="答案" width="200">
          <template slot-scope="scope">
            <!-- <textarea style="padding-bottom: 50%; overflow: hidden;" :value="scope.row.answer" /> -->
            <el-input type="textarea" :autosize="{ minRows: 1, maxRows:6}" v-model="scope.row.answer"></el-input>
          </template>
        </el-table-column>
        <el-table-column property="in" label="输入文件" width="100"></el-table-column>
        <el-table-column property="out" label="输出文件" width="100"></el-table-column>
        <el-table-column property="out" label="Public?" width="120">
          <template slot-scope="scope">
            <el-switch v-model="scope.row.is_public" active-color="#13ce66" inactive-color="#ff4949"
              @change="set_case_to_public_or_private(scope.row)">公开
            </el-switch>
          </template>
        </el-table-column>
        <el-table-column property="out" label="输出文件" width="120">
          <template slot-scope="scope">
            <el-button size="mini" icon="el-icon-edit" type="success" round @click="update_testcase(scope.row)">确认修改
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-drawer>
  </div>
</template>

<script>
import Sample from "@/views/Contest/components/question_sample.vue";
import axios from "axios";
Vue.prototype.$http = axios;
import Vue from "vue";
// import language js
import "codemirror/mode/clike/clike.js";
import "codemirror/mode/go/go.js";
import "codemirror/mode/python/python.js";
// import theme style
import "codemirror/theme/monokai.css";
// 下载到本地自定义
import './codemirror.css'
export default {
  name: "Question",
  components: { Sample },
  beforeCreate () {
    this.$axios.user_profile().then(res => {
      this.selfLog(res);
      this.user_profile = res["data"];
      this.$root.user_profile = this.user_profile;
      this.user_profile.user.last_login = this.user_profile.user.last_login.substr(
        0,
        10
      );
      this.user_profile.avatar = this.$OJIP + this.user_profile.avatar;
      this.selfLog(this.user_profile);
      this.get_list();
    })
      .catch(err => {
        this.selfLog(err)
        this.selfLog("获取用户失败")
        this.user_profile = this.$root.user_profile
        this.$axios.user_profile()
      })
  },
  data () {
    return {
      user_profile: this.$root.user_profile,
      editorOption: {},
      code_C:
        '//PREPEND BEGIN\n//C模板\n#include <stdio.h>\n//PREPEND END\n//TEMPLATE BEGIN\nint add(int a, int b) {\n  // Please fill this blank\n  return ___________;\n}\n//TEMPLATE END\n//APPEND BEGIN\nint main() {\n  printf("%d", add(1, 2));\n  return 0;\n}\n//APPEND END\n',
      code_C_Plus:
        '//C++模板\n#include <iostream>\nint add(int a, int b) {\n  // Please fill this blank\n  return ___________;\n}\n\nint main() {\n  std::cout << add(1, 2);\n  return 0;\n}\n"',
      code_Java:
        '//Java模板\nimport java.util.Scanner;\npublic class Main{\n    public static void main(String[] args){\n        Scanner in=new Scanner(System.in);\n        int a=in.nextInt();\n        int b=in.nextInt();\n        System.out.println((a+b));  \n    }\n}',
      code_Python:
        '"## Python模板\ndef method(input):\n    result = 0\n    #     coding here\n    return result\n\n\nif __name__ == \"main\":\n    input = 0\n    print(method(input))"',
      cmOptions_C: {
        tabSize: 4,
        mode: "text/x-csrc",
        theme: "monokai",
        lineNumbers: true,
        line: true
      },
      cmOptions_C_Plus: {
        tabSize: 4,
        mode: "text/x-csrc",
        theme: "monokai",
        lineNumbers: true,
        line: true
      },
      cmOptions_Python: {
        tabSize: 4,
        mode: "text/x-python",
        theme: "monokai",
        lineNumbers: true,
        line: true
      },
      cmOptions_Java: {
        tabSize: 4,
        mode: "text/x-java",
        theme: "monokai",
        lineNumbers: true,
        line: true
      },
      dialogFormVisible: false,
      dialogFormVisible_edit: false,
      inputVisible: false,
      inputValue: "",
      testcase: [],
      tag_types: ["info", "warning", "danger", "", "success"],
      form: {
        description: "",
        difficulty: "",
        hint: "",
        input_description: "",
        io_mode: {
          io_mode: "Standard IO",
          input: "input.txt",
          output: "output.txt"
        },
        languages: [],
        memory_limit: 256,
        output_description: "",
        rule_type: "ACM",
        samples: [],
        test_case_id: "",
        test_case_score: [],
        share_submission: false,
        source: "",
        spj: false,
        spj_code: "",
        spj_compile_ok: false,
        spj_language: "C",
        tags: ["dp", "array", "greedy"],
        template: {},
        time_limit: 1000,
        title: "",
        visible: true,
        _id: ""
      },
      sampleList: [],
      fileList: [],
      tableData: [],
      totalListLength: 0,
      limit: 10,
      offset: 0,
      user_profile: null,
      //   测试用例
      test_cases: [],
      case_drawer: false,
      current_problem: {
        template: {
          "C": "",
          "C++": "",
          "Pyhon": "",
          "Java": ""
        }
      },
      testcase_loading: false
    };
  },
  methods: {
    onCmReady (cm) {
      this.selfLog("the editor is readied!", cm);
    },
    onCmFocus (cm) {
      this.selfLog("the editor is focused!", cm);
    },
    onCmCodeChange_C (newCode) {
      this.code_C = newCode;
    },
    onCmCodeChange_C_Plus (newCode) {
      this.code_C_Plus = newCode;
    },
    onCmCodeChange_Python (newCode) {
      this.code_Python = newCode;
    },
    onCmCodeChange_Java (newCode) {
      this.code_Java = newCode;
    },
    //显示编辑问题对话框
    edit_problem (row) {
      //教师只能编辑自己创建的问题
      if (row.created_by.username !== this.user_profile.user.username) {
        this.$message({
          message: "只能编辑自己创建的问题喔～",
          type: "warning"
        });
        return;
      }
      // 查询问题详情
      else {
        // this.selfLog(row)
        this.$axios.admin_problem_detail_by_id(row.id).then(res => {
          this.selfLog(res.data)
          row = res.data
          this.current_problem = row
          //   this.$axios.problem_detail_by_id(row)
          this.form = row
          this.sampleList = row.samples
        })
      }
      this.dialogFormVisible_edit = true;
    },

    //展示测试用例
    async show_test_case (problem) {
      await this.$axios.admin_problem_detail_by_id(problem.id)
      this.current_problem = problem
      this.case_drawer = true
      this.testcase_loading = true
      this.selfLog(problem)
      let testcase_id = problem.test_case_id
      this.selfLog(testcase_id)
      this.$axios.get_test_case(testcase_id).then(res => {
        //   this.$axios.get_test_case(id).then(res => {
        this.selfLog(res)
        this.test_cases = res
        this.current_problem = problem
        // 对比当前测试用例是不是已经向学生公开
        this.selfLog("新问题的Public")
        this.selfLog(this.current_problem.samples)
        for (let case_ of this.test_cases.cases) {
          this.selfLog("下一个测试用例")
          let sample_format = { "input": case_.case, "output": case_.answer }
          for (let sample of this.current_problem.samples) {
            this.selfLog(sample.input + "===" + sample_format.input + "&&" + sample.output + " === " + sample_format.output)
            this.selfLog(sample.input === sample_format.input && sample.output === sample_format.output)
            if (sample.input === sample_format.input && sample.output === sample_format.output) {
              case_["is_public"] = true
              break
            } else {
              case_["is_public"] = false
            }
          }
        }
        this.testcase_loading = false
      })
        .catch(err => {
          this.selfLog(err)
          this.test_cases = {}
          this.testcase_loading = false
          this.$message({
            message: '该问题测试用例过大，无法及时传输，为减小服务器压力，请下载到本地查看',
            type: 'warning'
          });
        })
    },

    // 修改或添加测试用例
    update_testcase (test_case) {
      //教师只能编辑自己创建的问题
      if (this.current_problem.created_by.username !== this.user_profile.user.username) {
        this.$message({
          message: "只能编辑自己创建的问题～",
          type: "warning"
        });
        return;
      }
      else {
        this.selfLog(test_case.case)
        let param = {
          answer: test_case.answer,
          case: test_case.case,
          in: test_case.in,
          out: test_case.out,
          testcase_id: this.current_problem.test_case_id
        }
        this.$axios.update_testcase(param).then(res => {
          this.selfLog(res)
          this.show_test_case(this.current_problem)
          this.$message({
            message: '修改完成',
            type: 'success'
          });
        })
      }
    },
    // 新增testcase
    insert_testcase () {
      if (this.current_problem.created_by.username !== this.user_profile.user.username) {
        this.$message({
          message: "只能修改自己创建问题的测试用例",
          type: "warning"
        });
        return;
      }
      else {
        let new_case = {
          case: "",
          answer: "",
          in: this.test_cases.cases.length + 1 + ".in",
          out: this.test_cases.cases.length + 1 + ".out"
        }
        this.selfLog(this.test_cases)
        this.test_cases.cases.push(new_case)
      }
    },
    // 将该测试用例内容展示给学生 形成Public(也可以隐藏为Private)
    set_case_to_public_or_private (testcase) {
      let sample_format = { "input": testcase.case, "output": testcase.answer }
      //   变成true 代表以前没有 变成了有 直接添加
      if (testcase.is_public) {
        //   判断当前有没有 没有就添加
        this.current_problem.samples.push(sample_format)
        this.selfLog("添加后的samples")
        this.selfLog(this.current_problem.samples)
      }
      // 变成false表示 以前有变成了没有  找到然后删除
      else {
        //   否则删除之
        let i = 0
        for (let sample of this.current_problem.samples) {
          if (sample.input === sample_format.input && sample.output === sample_format.output) {
            this.current_problem.samples.splice(i, 1)
          }
          i = i + 1
        }
        this.selfLog(this.current_problem.samples)
      }
      this.selfLog("更新Public 测试用例：")
      this.selfLog(this.current_problem.samples)
      //   更新问题
      this.$axios.admin_update_problem(this.current_problem).then(res => {
        this.selfLog("更新结果")
        this.selfLog(res)
        this.$axios.admin_problem_detail_by_id(this.current_problem.id).then(res => {
          this.selfLog("更新后的问题详情")
          this.selfLog(res.data)
          this.current_problem = res.data
          this.form = res.data
          // 成功就刷新测试用例列表
          this.show_test_case(this.current_problem)
        })
      })
    },
    //删除问题
    delete_problem (row) {
      //教师只能删除自己创建的问题
      if (row.created_by.username !== this.user_profile.user.username) {
        this.$message({
          message: "只能删除自己创建的问题喔～",
          type: "warning"
        });
        return;
      }
      //先弹出确认删除问题弹窗，之后删除问题
      this.$confirm(
        "确定要删除此题目？相关的提交也将被删除, 是否继续?",
        "提示",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        }
      )
        .then(() => {
          this.$axios.delete_problem(row.id).then(res => {
            if (res.error === null) {
              this.get_list();
              this.$message({
                type: "success",
                message: "删除成功!"
              });
            } else {
              this.$message({
                message: res.data,
                type: "error"
              });
            }
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除"
          });
        });
    },
    //获取问题列表，默认查第一页，每页10条
    get_list () {
      this.$axios.admin_problem_list(this.limit, this.offset).then(res => {
        this.selfLog(res);
        this.selfLog(this.user_profile);
        if (res.error === null) {
          this.tableData = res.data.results;
          this.totalListLength = res.data.total;
        } else {
          this.$message({
            message: res.data,
            type: "error"
          });
        }
      });
    },
    //新增问题前先新增问题标签
    add_tags () {
      for (let tag of this.form.tags) {
        this.$axios.add_tags(tag).then(res => {
          if (res.error === null) {
            // this.selfLog(res);
          } else {
            this.$message({
              message: res.error,
              type: "error"
            });
          }
        });
      }
    },
    //新增问题后需要刷新问题列表
    async add_problem (method) {
      await this.add_tags();
      //处理示例字段
      if (method === 'new') {
        let l = this.sampleList.length;
        for (let i = 0; i < l; i++) {
          let sample = {
            input: this.sampleList[i].input,
            output: this.sampleList[i].output
          };
          this.form.samples.push(sample);
        }
      }
      //处理测试用例字段
      if (method === 'new') {
        this.form.test_case_id = this.testcase[0].id;
        let obj = {
          input_name: this.testcase[0].info[0].input_name,
          output_name: this.testcase[0].info[0].output_name,
          score: 100
        };
        this.form.test_case_score.push(obj);
      }
      this.selfLog(this.testcase)

      //处理语言模板字段
      if (this.form.languages.indexOf('C') > -1) {
        this.form.template['C'] = this.code_C
      }
      if (this.form.languages.indexOf('C++') > -1) {
        this.form.template['C++'] = this.code_C_Plus
      }
      if (this.form.languages.indexOf('Java') > -1) {
        this.form.template['Java'] = this.code_Java
      }
      if (this.form.languages.indexOf('Python') > -1) {
        this.form.template['Python'] = this.code_Python
      }
      if (method === 'new') {
        this.$axios.admin_add_problem(this.form).then(res => {
          if (res.error === null) {
            //退出dialog，刷新列表
            this.dialogFormVisible = false;
            this.get_list();
          } else {
            this.$message({
              message: res.data,
              type: "error"
            });
          }
        });
      } else if (method === 'edit') {
        this.$axios.admin_update_problem(this.form).then(res => {
          this.selfLog(res)
          this.get_list()
          this.dialogFormVisible_edit = false;
        })
      }

    },
    onDescriptionChange ({ quill, html, text }) {
      this.form.description = html;
    },
    onInputDescChange ({ quill, html, text }) {
      this.form.input_description = html;
    },
    onOutDescChange ({ quill, html, text }) {
      this.form.output_description = html;
    },
    onHintChange ({ quill, html, text }) {
      this.form.hint = html;
    },
    uploadTestCase (param) {
      const formData = new FormData();
      formData.append("file", param.file);
      formData.append("spj", false);
      this.selfLog(formData)
      this.$axios.upload_test_case(formData).then(res => {
        if (res.error === null) {
          this.testcase.push(res.data);
        } else {
          this.$message({
            message: res.error,
            type: "error"
          });
        }
      });
    },
    handleRemove (file, fileList) {
      this.selfLog(file, fileList);
    },
    handleExceed (files, fileList) {
      this.$message.warning(
        `当前限制选择 3 个文件，本次选择了 ${files.length
        } 个文件，共选择了 ${files.length + fileList.length} 个文件`
      );
    },
    beforeRemove (file, fileList) {
      return this.$confirm(`确定移除 ${file.name}？`);
    },
    deleteSample (id) {
      let l = this.sampleList.length;
      for (let i = 0; i < l; i++) {
        if (this.sampleList[i].id === id) {
          this.sampleList.splice(i, 1);
          break;
        }
      }
      this.form.samples = this.sampleList
    },
    addSample () {
      //向sampeList里插入一个对象
      let l = this.sampleList.length;
      if (!l) {
        let obj = {
          id: 1,
          input: "",
          output: ""
        };
        this.sampleList.push(obj);
        return;
      }
      let id = this.sampleList[this.sampleList.length - 1].id + 1;
      let obj = {
        id: id,
        input: "",
        output: ""
      };
      this.sampleList.push(obj);
      this.form.samples = this.sampleList
    },
    handleClose (tag) {
      this.form.tags.splice(this.form.tags.indexOf(tag), 1);
    },
    showInput () {
      this.inputVisible = true;
      this.$nextTick(_ => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },
    handleInputConfirm () {
      let inputValue = this.inputValue;
      //新增的标签已有，则直接返回，标签不可重复
      if (this.form.tags.indexOf(inputValue) > -1) {
        this.inputValue = "";
        this.inputVisible = false;
        return;
      }
      if (inputValue) {
        this.form.tags.push(inputValue);
      }
      this.inputVisible = false;
      this.inputValue = "";
    },
    //每页条数
    handleSizeChange (val) {
      this.limit = val;
      this.get_list();
    },
    //第几页
    handleCurrentChange (val) {
      this.offset = this.limit * (val - 1);

      this.get_list();
    }
  },
  computed: {
    editor () {
      return this.$refs.myQuillEditor.quill;
    }
  }
  //   mounted () {
  //     this.get_list();
  //   }
};
</script>

<style lang='scss' scoped>
.createNew {
  position: absolute;
  top: 80px;
  left: 70px;
}
.problem_list {
  margin: 70px;
  /* width: 80%; */
}
.el-tag + .el-tag {
  margin-left: 10px;
}
.button-new-tag {
  margin-left: 10px;
  height: 32px;
  line-height: 30px;
  padding-top: 0;
  padding-bottom: 0;
}
.input-new-tag {
  width: 90px;
  margin-left: 10px;
  vertical-align: bottom;
}
.sampleBox {
  text-align: left;
  margin: 20px;
}
.formTitle {
  text-align: left;
}
</style>
