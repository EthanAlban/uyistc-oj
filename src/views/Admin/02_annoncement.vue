<template>
  <div>
    <el-button type="primary" style="position: absolute; top: 80px; left: 70px;" round @click="edit_annoncement">创建公告
    </el-button>
    <el-button type="primary" style="position: absolute; top: 80px; left:170px;" round
      @click="admin_obtain_announcement">刷新列表</el-button>
    <div style="margin: 70px;width:60%">
      <el-table :data="announcement_list" border size="medium">
        <el-table-column fixed prop="title" label="标题"></el-table-column>
        <el-table-column prop="content" label="内容">
          <template slot-scope="scope">
            <div v-html="scope.row.content"></div>
          </template>
        </el-table-column>
        <el-table-column prop="create_time" label="创建时间"></el-table-column>
        <el-table-column prop="last_update_time" label="更新时间"></el-table-column>
        <el-table-column prop="created_by.username" label="创建人"></el-table-column>
        <el-table-column prop="visible" label="可见">
          <template slot-scope="scope">
            <el-switch v-model="scope.row.visible" active-color="#13ce66" inactive-color="#ff4949"
              @change="update_annocement(scope.row.id)"></el-switch>
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="操作">
          <template slot-scope="scope">
            <el-button type="text" size="small" @click="show_edit_panel(scope.row)">编辑</el-button>
            <el-button type="text" size="small" @click="delete_anoncement(scope.row.id)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- 新建公告 -->
    <el-dialog title="发布公告" :visible.sync="show_editor">
      <el-form :model="form">
        <el-form-item label="标题" :label-width="formLabelWidth">
          <el-input v-model="form.title" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="内容" :label-width="formLabelWidth">
          <textarea v-model="form.content" autocomplete="off" style="height:200px;width:100%"></textarea>
        </el-form-item>
        <el-form-item label="可见性" :label-width="formLabelWidth">
          <el-switch v-model="form.visible" active-color="#13ce66" inactive-color="#ff4949"></el-switch>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="show_editor = false">取 消</el-button>
        <el-button type="primary" @click="update_annocement_new(1)">确 定</el-button>
      </div>
    </el-dialog>

    <!-- 编辑公告 -->

    <el-dialog title="编辑公告" :visible.sync="show_editor2">
      <el-form :model="form2">
        <el-form-item label="标题" :label-width="formLabelWidth">
          <el-input v-model="form2.title"></el-input>
        </el-form-item>
        <el-form-item label="内容" :label-width="formLabelWidth">
          <textarea v-model="form2.content" style="height:200px;width:100%"></textarea>
        </el-form-item>
        <el-form-item label="可见性" :label-width="formLabelWidth">
          <el-switch v-model="form2.visible" active-color="#13ce66" inactive-color="#ff4949"></el-switch>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="show_editor2 = false">取 消</el-button>
        <el-button type="primary" @click="update_annocement()">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  mounted () {
    this.admin_obtain_announcement();
  },
  data () {
    return {
      formLabelWidth: "120px",
      show_editor: false,
      show_editor2: false,
      value: 1,
      announcement_list: [],
      edit_scope: null,
      form: {
        title: "",
        cotent: "",
        visible: true
      },
      form2: {
        title: "",
        cotent: "",
        visible: true
      }
    };
  },
  methods: {
    tableRowClassName ({ row, rowIndex }) {
      if (rowIndex === 1) {
        return "warning-row";
      } else if (rowIndex === 3) {
        return "success-row";
      }
      return "";
    },
    admin_obtain_announcement () {
      this.$axios.admin_obtain_announcement().then(res => {
        this.selfLog(res);
        this.announcement_list = res.data.results;
        for (let annon of this.announcement_list) {
          annon.create_time = annon.create_time.substr(0, 10);
          annon.last_update_time = annon.last_update_time.substr(0, 10);
        }
      });
    },
    show_edit_panel (scope) {
      this.edit_scope = scope;
      this.form2 = scope;
      this.selfLog("-------------");
      this.selfLog(this.form2);
      this.show_editor2 = true;
    },
    // 更新公告信息
    update_annocement () {
      let scope = this.edit_scope;
      let param = {};
      for (let anno of this.announcement_list) {
        if (anno.id === scope.id) {
          param = anno;
          break;
        }
      }
      this.selfLog(param);
      this.selfLog(this.announcement_list);
      this.$axios.update_annocement(param).then(res => {
        this.selfLog(res);
        if (res.error === null) {
          this.$message({
            message: "修改成功",
            type: "success"
          });
          this.show_editor2 = false;
        }
        this.admin_obtain_announcement();
      });
    },
    edit_annoncement () {
      this.show_editor = true;
    },
    // 删除公告
    delete_anoncement (id) {
      this.$axios.delete_anoncement(id).then(res => {
        if (res.error === null) {
          this.$message({
            message: "删除成功",
            type: "success"
          });
          this.admin_obtain_announcement();
        }
      });
    },
    // 新建公告
    update_annocement_new () {
      let param = {
        id: 1,
        title: this.form.title,
        visible: this.form.visible,
        content: this.form.content
      };
      this.selfLog(param);
      this.selfLog(this.announcement_list);
      this.$axios.update_annocement_new(param).then(res => {
        this.selfLog(res);
        if (res.error === null) {
          this.$message({
            message: "发布成功",
            type: "success"
          });
        }
        this.show_editor = false;
        this.admin_obtain_announcement();
      });
    }
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
.editor_panel {
  width: 40%;
  height: 50%;
  border: 2px black solid;
  position: absolute;
  left: 30%;
  top: 20%;
  background: silver;
}
</style>