<template>
  <div id="main">
    <el-form :model="user_profile" :rules="rules" ref="ruleForm" label-width="100px" class="ruleForm">
      <el-form-item label="头像">
        <el-upload class="upload-demo" drag action="string" :http-request="uploadAvatar" multiple>
          <i class="el-icon-upload"></i>
          <div class="el-upload__text">
            将文件拖到此处，或
            <em>点击上传</em>
          </div>
          <div class="el-upload__tip" slot="tip">只能上传jpg/png文件，且不超过500kb</div>
        </el-upload>
      </el-form-item>
      <el-form-item label="学校" prop="school">
        <el-input v-model="user_profile.school"></el-input>
      </el-form-item>
      <el-form-item label="专业" prop="major">
        <el-input v-model="user_profile.major"></el-input>
      </el-form-item>
      <el-form-item label="github地址" prop="github">
        <el-input v-model="user_profile.github"></el-input>
      </el-form-item>
      <el-form-item label="博客地址" prop="blog">
        <el-input v-model="user_profile.blog"></el-input>
      </el-form-item>
      <el-form-item label="个性签名" prop="mood">
        <el-input v-model="user_profile.mood"></el-input>
      </el-form-item>
    </el-form>
    <el-button type="primary" @click="update_profile" class="editButton">确认</el-button>
  </div>
</template>

<script>
export default {
  name: "updateInfo",
  beforeMount () {
    if (this.user_profile === null) {
      this.$axios.user_profile().then(res => {
        this.user_profile = res["data"];
      });
    }
  },
  data () {
    let isUrl = s => {
      return /^(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$/.test(
        s
      );
    };
    let validate_url = (rule, value, callback) => {
      if (!value) return;
      if (!isUrl(value)) {
        callback(new Error("网址格式错误"));
      }
    };
    return {
      user_profile: this.$root.user_profile,
      uploadAction: this.$OJIP + "/api/upload_avatar",
      rules: {
        blog: [{ validator: validate_url, trigger: "blur" }],
        github: [{ validator: validate_url, trigger: "blur" }]
      }
    };
  },
  methods: {
    update_profile () {
      //加一个语言进去
      let params = this.user_profile;
      params.language = "简体中文"
      this.$axios.updateInfo(params).then(res => {
        if (res["error"] === null) {
          this.$root.user_profile = res["data"];
          this.user_profile = res["data"];
          this.$message({
            message: "更新成功",
            type: "success"
          });
        } else if (res["error" === "invalid-blog"]) {
          this.$message.error("博客地址不正确，更新失败~");
        } else {
          this.$message.error("更新失败~");
        }
        this.resetForm();
      });
    },
    resetForm () {
      this.$refs["ruleForm"].resetFields();
    },
    uploadAvatar (param) {
      const formData = new FormData();
      formData.append("image", param.file);
      this.$axios.uploadAvatar(formData).then(res => {
        if (res["data"] === "Succeeded") {
          this.$axios.user_profile().then(res => {
            this.selfLog(res);
            this.user_profile = res["data"];
            let updata_avatar = {
              updata_avatar: this.user_profile.avatar
            };
          });
          this.$message({
            message: "头像更新成功",
            type: "success"
          });
          //   发送系统消息
          var myDate = new Date();
          let params = {
            userId: this.user_profile.id,
            msg: myDate.toLocaleString() + ":修改头像",
            type: 1
          };
          this.$axios.send_sys_info(params);
        } else {
          this.$message.error("头像更新失败~");
        }
      });
    }
  }
};
</script>

<style lang='scss' scoped>
#main {
  background-color: #c9dbdb;
  width: 50%;
  margin: auto;
  padding: 20px;
  .ruleForm {
    width: 60%;
    margin-left: 15%;
  }
  .editButton {
    width: 48%;
  }
}
</style>
