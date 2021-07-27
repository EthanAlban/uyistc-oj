<template>
  <div id="main">
    <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
      <el-form-item label="原密码" prop="old_password">
        <el-input type="password" show-password v-model="ruleForm.old_password" placeholder="请输入原登录密码"></el-input>
      </el-form-item>
      <el-form-item label="新密码" prop="new_password">
        <el-input type="password" show-password v-model="ruleForm.new_password" placeholder="请输入新登录密码"></el-input>
      </el-form-item>
      <el-form-item label="重复新密码" prop="re_new_password">
        <el-input type="password" show-password v-model="ruleForm.re_new_password" placeholder="请再次输入新登录密码"></el-input>
      </el-form-item>
    </el-form>
    <el-button type="primary" @click="submitForm('ruleForm')" class="editButton">修改</el-button>
  </div>
</template>

<script>
export default {
  name: "ModifyPwd",
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
      } else if (value.length < 6 || value.length > 20) {
        callback(new Error("长度应在 6 到 20 个字符"));
      } else {
        callback();
      }
    };
    let validate_re_new_password = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请再次输入密码"));
      } else if (value !== this.ruleForm.new_password) {
        callback(new Error("两次输入密码不一致!"));
      } else {
        callback();
      }
    };
    return {
      user_profile: this.$root.user_profile,
      ruleForm: {
        old_password: "",
        new_password: "",
        re_new_password: ""
      },
      rules: {
        old_password: [{ validator: validate_password, trigger: "blur" }],
        new_password: [{ validator: validate_password, trigger: "blur" }],
        re_new_password: [
          { validator: validate_re_new_password, trigger: "blur" }
        ]
      }
    };
  },
  methods: {
    submitForm (formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          let form = {
            new_password: this.ruleForm.new_password,
            old_password: this.ruleForm.old_password
          };
          this.$axios.change_password(form).then(res => {
            if (res["data"] === "Invalid old password") {
              this.$message.error("原密码错误");
            } else if (res["error"] === null) {
              this.$message({
                message: "密码修改成功",
                type: "success"
              });
              //   发送系统消息
              var myDate = new Date();
              let params = {
                userId: this.user_profile.id,
                msg: myDate.toLocaleString() + ":修改密码",
                type: 1
              };
              this.$axios.send_sys_info(params);
              this.$notify({
                title: "警告",
                message: "修改密码后需重新登录",
                type: "warning"
              });
              // 到学生表里边把学生的密码同步改了
              let form2 = {
                username: this.user_profile.user.username,
                new_password: this.ruleForm.new_password,
                userId: this.user_profile.user.id
              };
              this.$axios.update_password_stu(form2).then(res => {
                this.selfLog(res);
              });
              this.$router.push("/login");
            } else if (res["error"] === "permission-denied") {
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
