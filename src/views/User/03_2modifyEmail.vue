<template>
  <div id="main">
    <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
      <el-form-item label="密码" prop="password">
        <el-input v-model="ruleForm.password" placeholder="请先输入账户密码" type="password" show-password></el-input>
      </el-form-item>
      <el-form-item label="原邮箱" prop="old_email">
        <el-input v-model="ruleForm.old_email" placeholder="请先输入原邮箱"></el-input>
      </el-form-item>
      <el-form-item label="新邮箱" prop="new_email">
        <el-input v-model="ruleForm.new_email" placeholder="请输入新邮箱"></el-input>
      </el-form-item>
    </el-form>
    <el-button type="primary" @click="submitForm_email('ruleForm')" class="editButton">修改</el-button>
  </div>
</template>

<script>
export default {
  name: "ModifyEmail",
  beforeMount () {
    this.selfLog(this.user_profile);
    if (this.user_profile === null) {
      this.$axios.user_profile().then(res => {
        this.user_profile = res["data"];
        this.selfLog(this.user_profile);
      });
    }
  },
  data () {
    let validate_password = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请输入密码"));
      }
      if (value.length < 6 || value.length > 20) {
        callback(new Error("长度应在 6 到 20 个字符"));
      }
    };
    let isEmail = s => {
      return /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((.[a-zA-Z0-9_-]{2,3}){1,2})$/.test(
        s
      );
    };
    let validate_email = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("邮箱不能为空"));
      }
      if (!isEmail(value)) {
        callback(new Error("邮箱格式错误"));
      }
    };
    let validate_new_email = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("邮箱不能为空"));
      }
      if (!isEmail(value)) {
        callback(new Error("邮箱格式错误"));
      }
      if (value === this.ruleForm.old_email) {
        callback(new Error("新邮箱不能与旧邮箱一样"));
      }
    };
    return {
      user_profile: this.$root.user_profile,
      enable: true,
      ruleForm: {
        password: "",
        old_email: "",
        new_email: ""
      },
      rules: {
        password: [{ validator: validate_password, trigger: "blur" }],
        old_email: [{ validator: validate_email, trigger: "blur" }],
        new_email: [{ validator: validate_new_email, trigger: "blur" }]
      }
    };
  },
  methods: {
    // 修改邮件
    submitForm_email (formName) {
      this.ruleForm.old_email = this.user_profile.user.email;
      this.$refs[formName].validate(valid => {
        if (valid) {
          let form = {
            new_email: this.ruleForm.new_email,
            old_email: this.ruleForm.old_email,
            password: this.ruleForm.old_password
          };
          this.selfLog(form);
          this.$axios.change_email(form).then(res => {
            this.selfLog(res);
            if (res["data"] === "Wrong password") {
              this.$message.error("原密码错误");
            } else if (res["data"] === "The email is owned by other account") {
              this.$message({
                message: "该邮件地址已被注册",
                type: "warning"
              });
            } else if (res["error"] === null) {
              this.$message({
                message: "邮件修改成功",
                type: "success"
              });
              //   发送系统消息
              var myDate = new Date();
              let params = {
                userId: this.user_profile.id,
                msg: myDate.toLocaleString() + ":修改绑定邮箱地址",
                type: 1
              };
              this.$axios.send_sys_info(params);
            } else if (res["error"] === "permission-denied") {
              this.$message({
                message: "修改邮件需要登录",
                type: "warning"
              });
              this.$router.push("/login");
            }
          });
        } else {
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
  padding: 15px;
  .demo-ruleForm {
    width: 50%;
    margin: auto;
  }
  .editButton {
    width: 40%;
    margin-left: 80px;
  }
}
</style>
