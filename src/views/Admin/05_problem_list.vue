<template>
  <div>
    <el-button type="primary" round class="createNew" @click="dialogFormVisible = true">创建新题目</el-button>
    <div class="problem_list">
      <el-table :data="tableData" size="medium">
        <el-table-column label="题目编号" sortable width="120" prop="id"></el-table-column>
        <el-table-column label="编号" width="120" prop="_id"></el-table-column>
        <el-table-column label="标题" width="120" prop="title"></el-table-column>
        <el-table-column label="创建人" width="100">
          <template slot-scope="scope">
            <i class="el-icon-s-custom"></i>
            <span style="margin-left: 10px">{{ scope.row.created_by.username}}</span>
          </template>
        </el-table-column>
        <el-table-column label="标签" width="200">
          <template slot-scope="scope">
            <el-tag v-for="(tag,idx) in scope.row.tags" :key="tag" :type="tag_types[idx%5]" effect="dark">{{ tag }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="可见状态" prop="visible">
          <template slot-scope="scope">
            <el-switch v-model="scope.row.visible" @change="set_problem_visible(scope.row)" active-color="#13ce66"
              inactive-color="#ff4949"></el-switch>
          </template>
        </el-table-column>
        <el-table-column label="创建时间" width="200" sortable>
          <template slot-scope="scope">
            <i class="el-icon-time"></i>
            <span style="margin-left: 10px">{{ scope.row.create_time | formatDate}}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="350" header-align="center">
          <template slot-scope="scope">
            <el-button size="mini" icon="el-icon-edit" type="success" round @click="edit_problem(scope.row.id)">编辑
            </el-button>
            <el-button size="mini" icon="el-icon-delete" type="danger" round @click="delete_problem(scope.row.id)">删除
            </el-button>
            <el-button size="mini" icon="el-icon-download" type="primary" round
              @click="download_test_case(scope.row.id)">下载测试用例</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="1"
        :page-sizes="[5, 10, 20]" :page-size="10" layout="total, sizes, prev, pager, next, jumper"
        :total="totalListLength"></el-pagination>
    </div>

    <el-dialog :visible.sync="dialogFormVisible" destroy-on-close>
      <p class="formTitle" style="fontSize:20px;marginTop:-40px">新建题目</p>
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
            <el-checkbox label="Java"></el-checkbox>
            <el-checkbox label="Python2"></el-checkbox>
            <el-checkbox label="Python3"></el-checkbox>
            <el-checkbox label="Golang"></el-checkbox>
          </el-checkbox-group>
        </el-form-item>
      </el-form>
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
          <a href="https://docs.onlinejudge.me/#/onlinejudge/guide/test_case" target="_blank">测试用例格式说明</a>
        </div>
      </el-upload>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="add_problem">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import Sample from "@/views/Contest/components/question_sample.vue";
import axios from "axios";
Vue.prototype.$http = axios;
import Vue from "vue";
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
    });
  },
  data () {
    return {
      user_profile: this.$root.user_profile,
      editorOption: {
      },
      dialogFormVisible: false,
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
      user_profile: null
    };
  },
  methods: {
    set_problem_visible (problem) {
      this.$axios.admin_update_problem(problem).then(res => {
        this.selfLog(res)
        this.get_list()
        this.$message({
          message: '修改完成',
          type: 'success'
        });
      })
    },
    //编辑问题
    edit_problem (id) {
      this.selfLog("编辑问题" + id);
    },
    //下载测试用例
    download_test_case (id) {
      return new Promise((resolve, reject) => {
        // this.$axios.download_test_case(id).then(resp => {
        let url = "api/admin/test_case?problem_id=" + id;
        Vue.prototype.$http
          .get(url, { responseType: "blob" })
          .then(resp => {
            this.selfLog(resp);
            let headers = resp.headers;
            this.selfLog(headers);
            if (headers["content-type"].indexOf("json") !== -1) {
              this.selfLog("+====================");
              let fr = new window.FileReader();
              if (resp.data.error) {
                //   Vue.prototype.$error(resp.data.error)
                this.$message.error(resp.data.error);
              } else {
                //   Vue.prototype.$error('Invalid file format')
                this.$message.error("Invalid file format");
              }
              fr.onload = event => {
                let data = JSON.parse(event.target.result);
                if (data.error) {
                  // Vue.prototype.$error(data.data)
                  this.$message.error(data.data);
                } else {
                  // Vue.prototype.$error('Invalid file format')
                  this.$message.error("Invalid file format");
                }
              };
              let b = new window.Blob([resp.data], {
                type: "application/json"
              });
              fr.readAsText(b);
              return;
            }
            this.selfLog(headers);
            let link = document.createElement("a");
            link.href = window.URL.createObjectURL(
              new window.Blob([resp.data], { type: headers["content-type"] })
            );
            link.download = (headers["content-disposition"] || "").split(
              "filename="
            )[1];
            document.body.appendChild(link);
            this.selfLog(link);
            link.click();
            link.remove();
            resolve();
          })
          .catch(() => { });
      });
    },
    //删除问题
    delete_problem (id) {
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
          this.$axios.delete_problem(id).then(res => {
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
    async add_problem () {
      await this.add_tags();
      //处理示例字段
      let l = this.sampleList.length;
      for (let i = 0; i < l; i++) {
        let sample = {
          input: this.sampleList[i].input,
          output: this.sampleList[i].output
        };
        this.form.samples.push(sample);
      }
      //处理测试用例字段
      this.form.test_case_id = this.testcase[0].id;
      let obj = {
        input_name: this.testcase[0].info[0].input_name,
        output_name: this.testcase[0].info[0].output_name,
        score: 100
      };
      this.form.test_case_score.push(obj);
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
