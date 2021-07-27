<template>
  <dv-full-screen-container class="bg" v-loading="FeiShu_login_loadin"
    v-bind:class="{ bg1: bg[0], bg2: bg[1], bg3: bg[2], bg4: bg[3],bg5: bg[4], bg6: bg[5],bg7: bg[6], bg8: bg[7],bg9: bg[8], bg10: bg[9],bg11: bg[10], bg12: bg[11]}">
    <div class="title">
      <dv-decoration-11 style="width:20%;height:10%;position:absolute;left:40%;top:5%;font-size:200%">
        <a href="/"><img src="../../public/sologon.png" style="width:100%;marginTop:10%" alt=""></a>
      </dv-decoration-11>
    </div>
    <div class="third_login">
      <el-tooltip class="item" effect="light" content="成电飞书登录" placement="right">
        <img src="../icons/FeiShuLogo.png" class="FeiShu" @click="FeiShu_authorize" alt="">
      </el-tooltip>
      <el-tooltip class="item" effect="light" content="微博登录" placement="right">
        <img src="../icons/WeiBoLogo.png" class="WeiBo" @click="FeiShu_authorize" alt="">
      </el-tooltip>
    </div>
    <!-- 学号登录/普通账号登录二选一 -->
    <el-tabs type="border-card" class="panel">
      <el-tab-pane label="用户登录">
        <el-form label-position="left" label-width="80px" :model="form">
          <el-form-item label="用户名称">
            <el-input prefix-icon="el-icon-user" v-model="form.username" placeholder="请输入用户名称"></el-input>
          </el-form-item>
          <el-form-item label="登录密码">
            <el-input prefix-icon="el-icon-lock" type="password" v-model="form.password" placeholder="请输入登录密码"
              show-password></el-input>
          </el-form-item>
          <el-button style="width:80%" type="primary" round @click="login" :loading="loading">登录</el-button>
        </el-form>
      </el-tab-pane>

      <el-tab-pane label="学号登录">
        <el-form label-position="left" label-width="80px" :model="form">
          <el-form-item label="学/工号">
            <el-input prefix-icon="el-icon-user" v-model="form.username" placeholder="请输入学/工号"
              @blur="check_if_stunum_registered(form.username)"></el-input>
          </el-form-item>
          <el-form-item label="登录密码">
            <el-input prefix-icon="el-icon-lock" type="password" v-model="form.password" placeholder="请输入登录密码"
              show-password></el-input>
          </el-form-item>
          <div v-if="stu_not_registered">
            <el-form-item label="验证码">
              <el-input prefix-icon="el-icon-mobile" v-model="form.captcha_str" placeholder="请输入验证码" style="width:60%;">
              </el-input>
              <img :src="captcha" @click="get_captcha_img" class="captchaImg" />
            </el-form-item>
          </div>
          <el-button style="width:80%" type="primary" round @click="login_StuNum" :loading="loading">学号登录
          </el-button>
        </el-form>
      </el-tab-pane>
      <el-link type="primary" @click="register" style="float:left;margin:20px 0 0 20px">新用户? 立即注册</el-link>
      <el-link type="primary" @click="visitor_login" style="float:right;margin:20px 20px 0 0">游客登录</el-link>

    </el-tabs>
    <!-- 得到登录code之后输入验证码注册用户 -->
    <el-dialog title="信息确认" :visible.sync="FeiShu_login_check" append-to-body width="40%">
      <span style="color:green;">注：飞书登录默认系统密码与学号相同</span>
      <el-form :model="stu_profile">
        <el-form-item label="学号" :label-width="formLabelWidth">
          <el-input v-model="stu_profile.stu_num" disabled style="width:60%;"></el-input>
        </el-form-item>
        <el-form-item label="姓名" :label-width="formLabelWidth">
          <el-input v-model="stu_profile.stu_name" placeholder="建议填写真实姓名~" style="width:60%;"></el-input>
        </el-form-item>
        <el-form-item label="验证码">
          <el-input prefix-icon="el-icon-mobile" v-model="stu_profile.captcha_str" placeholder="请输入验证码"
            style="width:30%;">
          </el-input>
          <img :src="captcha" @click="get_captcha_img" style="margin-left:10px" class="captchaImg" />
        </el-form-item>

      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancle_FeiShu_login">取 消</el-button>
        <el-button type="primary" @click="FeiShu_Login">确 定</el-button>
      </div>
    </el-dialog>
    <!-- <footer>
        <p>xxxxx</p>
      </footer> -->
    <!-- </dv-border-box-3> -->
  </dv-full-screen-container>
</template>

<script>
import axios from 'axios'
export default {
  name: "Login",
  beforeMount () {
  },
  mounted () {
    const oScript = document.createElement('script');
    oScript.type = 'text/javascript';
    oScript.src = 'https://sf1-scmcdn-tos.pstatp.com/goofy/ee/lark/h5jssdk/lark/js_sdk/h5-js-sdk-1.5.4.js';
    document.body.appendChild(oScript);

    // 获取验证码图片
    this.get_captcha_img();
    this.$notify.info({
      title: '提示',
      message: '教师/学生都可以直接使用学工号进行登录'
    });
    let month = new Date().getMonth(); //获取当前月份(0-11,0代表1月)
    // month = 8
    this.bg[month] = true;

    let code = this.$route.query.code
    this.code = code
    if (this.code !== "" && typeof (this.code) !== "undefined") {
      this.get_stunum()
    } else {
      this.selfLog("code获取失败")
      this.selfLog(this.code)
    }

  },
  data () {
    return {
      bg: [false, false, false, false, false, false, false, false, false, false, false, false],
      loading: false,
      stu_not_registered: true,
      form: {
        username: "",
        password: "",
        captcha_str: ""
      },
      captcha: "",
      // 登录预授权码
      code: "",
      stu_profile: {
        stu_num: "",
        stu_name: "",
        work_id: "",
        captcha_str: ""
      },
      FeiShu_login_check: false,
      //   飞书登录时等待网络请求loading
      FeiShu_login_loadin: false,
      formLabelWidth: "80px",
      FeiShu: {
        tenant_access_token: "",
        app_access_token: ""
      },
      is_new_stu: false
    };
  },


  methods: {
    FeiShu_authorize () {
      //   this.selfLog(this.$axios.request_user_profile())
      //   window.location.href = 'https://open.feishu.cn/open-apis/authen/v1/user_auth_page_beta?app_id=cli_a0f809e4ca78100e&redirect_uri=http%3A%2F%2F127.0.0.1%3A8080%2Flogin&state=login';
      window.location.href = 'https://open.feishu.cn/open-apis/authen/v1/user_auth_page_beta?app_id=cli_a0f809e4ca78100e&redirect_uri=http%3A%2F%2F47.94.135.51%3A8080%2Flogin&state=login';
    },
    // 获取app_access_token
    async get_app_access_token () {
      //   this.selfLog(window)
      //   自建应用的 app_access_token 也代表了访问本租户内的相关API及资源，故 自建应用的 app_access_token 等同于 tenant_access_token。
      await this.$axios.get_tenant_access_token().then(res => {
        this.FeiShu.tenant_access_token = res.tenant_access_token
        this.FeiShu.app_access_token = res.tenant_access_token
        this.selfLog("用户access:" + this.FeiShu.app_access_token)
      })
      await this.$axios.get_FeiShu_user_profile(this.FeiShu.tenant_access_token, this.code).then(res => {
        this.selfLog("飞书用户信息：")
        this.selfLog(res)
        if (res.code === 0) {
          this.selfLog(res)
          this.stu_profile.stu_name = res.data.name
        }
      })
    },

    // 获取用户学号
    async get_stunum () {
      if (this.code === "") {
        return
      } else {
        this.FeiShu_login_loadin = true
        // await this.get_app_access_token()
        let eid = "uest5000006"
        let code = this.code
        // https://bytedance.feishu.cn/docs/doccneTa8HyYR7azip2vDtC0i6g#
        // await axios.get('http://feishumiddle.uestc.edu.cn/api/GetUserInfoByH5?eid=' + eid + '&code=' + code).then(function (response) {
        //   this.selfLog(res)
        //   if (res.code === 0) {
        //     this.selfLog("获取学工号成功")
        //     this.stu_profile.stu_num = res.data.studentId
        //     this.stu_profile.work_id = res.data.WorkId
        //   }
        // })
        await this.$axios.GetUserInfoByH5(eid, code).then(res => {
          this.selfLog(res)
          if (res.code === 0) {
            this.selfLog("获取学工号成功")
            this.stu_profile.stu_num = res.data.studentId
            this.stu_profile.work_id = res.data.WorkId
          }
        })
        let username = {
          username: this.stu_profile.stu_num
        }
        // this.selfLog("++++++++++++++++++++++++++++++")
        this.selfLog("学号：")
        this.selfLog(username)
        await this.$axios.check_if_stunum_registered(username).then(res => {
          //   学号已经注册直接登录
          this.selfLog(res)
          if (res.data === "registered") {
            this.form.username = this.stu_profile.stu_num
            this.form.password = res.password
            this.login()
          }
        })
        await this.get_app_access_token()
        this.FeiShu_login_loadin = false
        this.FeiShu_login_check = true
      }
    },
    // 取消飞书登录
    cancle_FeiShu_login () {
      this.code = ""
      this.FeiShu_login_check = false
      this.$router.push("/login")

    },
    // 确认飞书登录
    async FeiShu_Login () {
      await this.get_app_access_token()
      this.loading = true
      this.form.username = this.stu_profile.stu_num
      this.form.password = this.stu_profile.stu_num
      this.form.captcha_str = this.stu_profile.captcha_str
      let param = {
        username: this.form.username,
        password: this.form.password,
        captcha: this.form.captcha_str
      };
      this.$axios.stuLogin(param).then((res) => {
        this.selfLog("++++++++")
        this.selfLog(res);
        if (res.result === "ok" || res.result === "ok_1") {
          if (res.result === "ok_1") {
            // 没登录过而且系统数据库有这个学号，自动注册新的OJ用户
            this.form.username = res.data.stunum
            this.stu_profile.stu_name = res.data.name
            let param2 = {
              username: this.form.username,
              password: this.form.password,
              email: this.form.username + "@oj.org",
              captcha: this.form.captcha_str,
            };
            this.selfLog(param2);
            this.$axios.register_new(param2).then(res => {
              this.selfLog(res);
              if (res.error === "invalid-captcha" || res.data === "Invalid captcha") {
                // this.$router.push("/login");
                // 验证码错误要重置自己服务器的登录密码
                this.selfLog("验证码错误");
                let param3 = {
                  username: this.form.username,
                }
                this.$axios.reset_pass_to_zero(param3).then(res => {
                  this.$message({
                    message: "验证码错误",
                    type: "warning",
                  });
                  this.get_captcha_img()
                  this.loading = false
                })
              } else if (res.data === "Succeeded") {
                this.$notify({
                  title: "提示",
                  message: "使用默认密码第一次登录请修改密码，教师系统申请，请在账户管理->个人信息中提交",
                  duration: 0
                });
                this.login();
              }
            });
          }
          else {
            //   老用户密码正确
            this.stu_profile.stu_name = res.data.name
            this.form.username = res.data.stunum
            this.login()
          }
          //   this.login()
        } else if (res.result === "err_1") {
          // 学号没有但是飞书查到了 先注册用户  再插入学生 这也太复杂了 。。。
          // 没登录过而且系统数据库有这个学号，自动注册新的OJ用户
          this.stu_profile.stu_name = res.data.name
          let param2 = {
            username: this.form.username,
            password: this.form.password,
            email: this.form.username + "@oj.org",
            captcha: this.form.captcha_str,
          };
          this.selfLog(param2);
          this.$axios.register_new(param2).then(res => {
            this.selfLog(res);
            if (res.error === "invalid-captcha" || res.data === "Invalid captcha") {
              // this.$router.push("/login");
              // 验证码错误要重置自己服务器的登录密码
              this.selfLog("验证码错误");
              let param3 = {
                username: this.form.username,
              }
              this.$axios.reset_pass_to_zero(param3).then(res => {
                this.$message({
                  message: "验证码错误",
                  type: "warning",
                });
                this.get_captcha_img()
                this.loading = false
              })
            } else if (res.data === "Succeeded") {
              // 注册成功之后插入学生
              let param3 = {
                stuNum: this.stu_profile.stu_num,
                stuName: this.stu_profile.stu_name,
                userId: res.data.user.id,
                user_type: 3
              }
              this.$axios.insert_feishu_stu(param3).then(res => {
                this.$notify({
                  title: "提示",
                  message: "使用默认密码第一次登录请修改密码，教师系统申请，请在账户管理->个人信息中提交",
                  duration: 0
                });
                this.login();
              })

            }
          });

        } else if (res.result === "err_2") {
          this.loading = false
          this.selfLog(res.msg);
          this.$message({
            message: "学号已登录，密码不是默认密码，请使用账户密码登录~",
            type: "warning"
          });
        } else if (res.result === "err_3") {
          this.loading = false
          this.selfLog(res.msg);
          this.$message({
            message: "第一次登录请使用默认密码（学号）",
            type: "warning"
          });
        }
      });
    },
    get_captcha_img () {
      this.$axios.get_captcha_img()
        .then(res => {
          this.captcha = res.data;
        })
        .catch(err => {
          this.selfLog("验证码获取失败")
          this.get_captcha_img()
        })
    },
    register () {
      this.$router.push("/register");
    },
    visitor_login () {
      this.$router.push("/");
    },
    login () {
      this.is_new_stu = true
      this.loading = true
      let param = {
        username: this.form.username,
        password: this.form.password
      };
      this.$axios.userLogin(param).then(res => {
        if (res.error === null) {
          this.loading = false
          this.selfLog(res.data);
          this.$root.user_profile = res.data
          this.$message({
            message: "登陆成功",
            type: "success"
          });
          // 如果是新的学生第一次登录 把数据库里的userId更新了
          if (this.is_new_stu) {
            this.$axios.user_profile().then(res => {
              this.$axios.update_new_stu_userId(this.form.username, res.data.user.id)
            })
          }
          this.is_new_stu = false
          this.$router.push("/");
        } else {
          this.loading = false
          if (res.data === "Invalid username or password") {
            this.$message({
              message: "用户名密码不匹配或未注册",
              type: "warning"
            });
          }
        }
      });
    },
    // 学号默认密码登录
    login_StuNum () {
      this.loading = true
      let param = {
        username: this.form.username,
        password: this.form.password,
        captcha: this.form.captcha_str
      };
      this.$axios.stuLogin(param).then((res) => {
        this.selfLog(res);
        if (res.result === "ok" || res.result === "ok_1") {
          if (res.result === "ok_1") {
            // 自动注册新的OJ用户
            let param2 = {
              username: this.form.username,
              password: this.form.password,
              email: this.form.username + "@oj.org",
              captcha: this.form.captcha_str
            };
            this.selfLog(param2);
            this.$axios.register_new(param2).then(res => {
              this.selfLog(res);
              if (res.error === "invalid-captcha" || res.data === "Invalid captcha") {
                // this.$router.push("/login");
                // 验证码错误要重置自己服务器的登录密码
                this.selfLog("验证码错误");
                let param3 = {
                  username: this.form.username,
                }
                this.$axios.reset_pass_to_zero(param3).then(res => {
                  this.$message({
                    message: "验证码错误",
                    type: "warning",
                  });
                  this.get_captcha_img()
                  this.loading = false
                })
              } else if (res.data === "Succeeded") {
                this.$notify({
                  title: "提示",
                  message: "使用默认密码第一次登录请修改密码，教师系统申请，请在账户管理->个人信息中提交",
                  duration: 0
                });
                this.is_new_stu = true
                this.login();
              }
            });
          }
          else {
            //   老用户密码正确
            this.is_new_stu = true
            this.login()
          }
          //   this.login()
        } else if (res.result === "err_1") {
          this.selfLog(res.msg);
          this.loading = false
          this.$message({
            message: "学号未注册，请联系管理员",
            type: "warning"
          });
        } else if (res.result === "err_2") {
          this.loading = false
          this.selfLog(res.msg);
          this.$message({
            message: "学号已登录，密码错误",
            type: "warning"
          });
        } else if (res.result === "err_3") {
          this.loading = false
          this.selfLog(res.msg);
          this.$message({
            message: "第一次登录请使用默认密码（学号）",
            type: "warning"
          });
        }
      });
    },
    // 学校用户输入了学工号之后查一下这个学号注册了没 没有的话显示验证码  否则不再显示
    check_if_stunum_registered (username) {
      let username_ = {
        username: username
      }
      this.$axios.check_if_stunum_registered(username_).then(res => {
        if (res.data === "not_registered") {
          this.stu_not_registered = true
        } else if (res.data === "registered") {
          this.stu_not_registered = false
        } else if (res.data === "stunum_not_exists") {
          this.stu_not_registered = false
          this.$message({
            message: "输入的不是有效学/工号",
            type: "warning"
          });
        }
      })
    },
  }
};
</script>

<style lang="scss" scoped>
.bg {
  .register {
    float: right;
    top: 2px;
    right: 2px;
  }
  .title {
    height: 16%;
    .sys-name {
      font-size: 45px;
      color: #f5f5f5;
    }
  }
  .panel {
    width: 25%;
    height: 40%;
    margin: auto;
    background-color: white;
    opacity: 0.95;
    .captchaImg {
      width: 40%;
      height: 40px;
      margin-bottom: -15px;
    }
    .loginButton {
      width: 200px;
    }
  }
}
footer {
  position: absolute;
  bottom: 10%;
  @media screen and (max-width: 1440px) and (min-width: 1280px) {
    bottom: 16%;
  }
  display: flex;
  width: 100%;
  justify-content: center;
  p {
    font-size: 30px;
  }
}
.bg1 {
  background: url("../assets/uestc/calendar/1.jpeg") no-repeat;
  background-size: 100% 100%;
}
.bg2 {
  background: url("../assets/uestc/calendar/2.jpeg") no-repeat;
  background-size: 100% 100%;
}
.bg3 {
  background: url("../assets/uestc/calendar/3.jpeg") no-repeat;
  background-size: 100% 100%;
}
.bg4 {
  background: url("../assets/uestc/calendar/3.jpeg") no-repeat;
  background-size: 100% 100%;
}
.bg5 {
  background: url("../assets/uestc/calendar/5.jpeg") no-repeat;
  background-size: 100% 100%;
}
.bg6 {
  //   background: url("../assets/uestc/calendar/6.jpeg") no-repeat;
  background: url("../assets/bg/drangon_boat.png") no-repeat;
  background-size: 100% 90%;
}
.bg7 {
  background: url("../assets/uestc/calendar/7.jpeg") no-repeat;
  background-size: 100% 100%;
}
.bg8 {
  background: url("../assets/uestc/calendar/8.jpeg") no-repeat;
  background-size: 100% 100%;
}
.bg9 {
  background: url("../assets/uestc/calendar/9.jpeg") no-repeat;
  background-size: 100% 100%;
}
.bg10 {
  background: url("../assets/uestc/calendar/10.jpeg") no-repeat;
  background-size: 100% 100%;
}
.bg11 {
  background: url("../assets/uestc/calendar/11.jpeg") no-repeat;
  background-size: 100% 100%;
}
.bg12 {
  background: url("../assets/uestc/calendar/12.jpeg") no-repeat;
  background-size: 100% 100%;
}
.bg1 {
  background: url("../assets/uestc/calendar/1.jpeg") no-repeat;
  background-size: 100% 100%;
}
.bg2 {
  background: url("../assets/uestc/calendar/2.jpeg") no-repeat;
  background-size: 100% 100%;
}
.bg1 {
  background: url("../assets/uestc/calendar/1.jpeg") no-repeat;
  background-size: 100% 100%;
}
.bg2 {
  background: url("../assets/uestc/calendar/2.jpeg") no-repeat;
  background-size: 100% 100%;
}
.third_login {
  margin-top: 20%;
  position: absolute;
  right: 32%;
  top: -15%;
  //   border: 1px black solid;
  width: 6%;
  //   height: 50%;
  //   border: 1px black solid;
  .FeiShu {
    position: felx;
    background-image: url("../icons/FeiShuLogo.png");
    width: 51%;
    height: 51%;
    border-radius: 50%;
    box-shadow: 5px 3px 5px 0px rgba(45, 156, 161, 0.57);
    -webkit-box-shadow: 5px 3px 5px 0px rgba(45, 156, 161, 0.57);
    -moz-box-shadow: 5px 3px 5px 0px rgba(45, 156, 161, 0.57);
  }
  .WeiBo {
    position: felx;
    margin-left: 5%;
    background-image: url("../icons/FeiShuLogo.png");
    width: 51%;
    height: 51%;
    border-radius: 50%;
    box-shadow: 5px 3px 5px 0px rgba(45, 156, 161, 0.57);
    -webkit-box-shadow: 5px 3px 5px 0px rgba(45, 156, 161, 0.57);
    -moz-box-shadow: 5px 3px 5px 0px rgba(45, 156, 161, 0.57);
  }
}
</style>

