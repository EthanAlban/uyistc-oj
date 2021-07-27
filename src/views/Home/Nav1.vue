<template>
  <div>
    <!-- message-banner -->
    <div
      v-bind:class="{ message_banner1: bg[0], message_banner2: bg[1], message_banner3: bg[2], message_banner4: bg[3]}">
      <!-- 弹幕输入框 -->
      <div class="message-container">
        <h1 class="message-title">留言板</h1>
        <dv-decoration-6 style="width:300px;height:20px;" />
        <div class="animated fadeInUp message-input-wrapper">
          <input v-model="messageContent" @click="show = true" placeholder="说点什么吧"
            style="background: transparent;z-index:1" />
          <div class="block">
            <span class="demonstration"></span>
            <el-color-picker v-model="color1" style="border-radius: 15px;z-index:1"></el-color-picker>
          </div>
          <button class="ml-3 animated bounceInLeft" style="border-radius: 15px;z-index:1" v-show="show"
            @click="addToList">
            发送
          </button>
        </div>
      </div>
      <!-- 弹幕列表 -->
      <div class="barrage-container">
        <vue-baberrage ref="babarrage" :lanesCount="10" :messageHeight="100" :loop="false" :isShow="true"
          :barrageList="barrageList" :maxWordCount="60" :hoverLanePause="true">
          <template v-slot:default="slotProps">
            <span style="color: #000; fontsize: 25px"><img style="height: 30px; width: 30px; border-radius: 15px"
                :src="slotProps.item.avatar" /><span :style="slotProps.item.color">{{ slotProps.item.nickname }}:
                {{ slotProps.item.msg }}</span></span>
          </template>
        </vue-baberrage>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  beforeMount () {
    this.userDetail();
    var myDate = new Date();
    let hour = myDate.getHours(); //获取当前小时数(0-23)
    this.selfLog(hour);
    // hour = 23
    if (hour < 10) {
      this.bg[0] = true
    } else if (hour >= 10 && hour < 14) {
      this.bg[1] = true
    } else if (hour >= 14 && hour < 19) {
      this.bg[2] = true
    } else {
      this.bg[3] = true
    }
    if (hour >= 23 || hour <= 4) {
      this.$notify.info({
        title: '夜深',
        message: '夜深了，奔波路上，疲惫的夜晚，请早早入睡，晚安~'
      });
    }
    this.selfLog(this.bg);
  },
  mounted () {
    // this.getUser()
    this.listMessage();
    currentId: 0;
  },
  data: {
    user_profile: null,
  },
  data () {
    return {
      bg: [false, false, false, false],
      color1: "",
      avatar: "",
      nickname: "",
      show: false,
      messageContent: "",
      barrageList: [],
      page: 1,
      interval_id: -1
    };
  },
  methods: {
    async addToList () {
      // 先判断用户登录没有 没登陆就弹去登陆
      this.selfLog("user:" + this.user_profile);
      if (this.user_profile === null) {
        this.$notify({
          title: "登录",
          message: "发送弹幕请先登录",
          type: "warning",
        });
        this.selfLog("未登录");
        this.$router.push("/login");
        return false;
      }
      if (this.messageContent.trim() === "") {
        this.$message.info("留言不能为空");
        return false;
      }
      this.selfLog(this.user_profile);
      // 发送存储弹幕请求
      let param = {
        uid: this.user_profile.id,
        // 应该引用全局数据 待修改
        avatar: this.$OJIP + this.user_profile.avatar,
        msg: this.messageContent,
        color: "color:" + this.color1,
        name: this.user_profile.user.username,
      };
      this.$axios.send_single_bullet_chat(param).then((res) => {
        this.selfLog(res);
        // 弹幕发送成功
        if (res.result === "ok") {
          // 放到弹幕列表第一个
          this.barrageList.unshift({
            color: "color:" + this.color1,
            id: -1,
            avatar: res.chat.avatar,
            msg: res.chat.msg,
            time: Math.floor(Math.random() * 10 + 3),
            nickname: res.chat.name,
            extraWidth: res.chat.extraWidth,
          });
          this.$message.success("新增留言成功");
          this.selfLog(res.chat);
        } else {
          // 弹幕发送失败
          this.$message.error("留言失败");
        }
      });
    },
    async listMessage () {
      let param = {
        page: this.page,
        pageSize: 30,
      };
      this.$axios.get_bullet_chats(param).then((res) => {
        this.selfLog("获取弹幕成功");
        this.selfLog(res);
        for (let bullet of res["list"]) {
          // this.selfLog(bullet.avatar);
          this.barrageList.push({
            color: bullet.color,
            id: bullet.pk,
            avatar: bullet.avatar,
            msg: bullet.msg,
            time: Math.floor(Math.random() * 15 + 3),
            nickname: bullet.name,
            extraWidth: bullet.extraWidth,
          });
        }
      })
      this.interval_id = setInterval(() => {
        let param = {
          page: this.page,
          pageSize: 30,
        };
        this.$axios.get_bullet_chats(param).then((res) => {
          //   this.selfLog("获取弹幕成功");
          this.selfLog(res);
          this.page++;
          while (this.barrageList.length > 1) {
            this.selfLog(this.barrageList)
            this.barrageList.shift();
          }
          for (let bullet of res["list"]) {
            // this.selfLog(bullet.avatar);
            this.barrageList.push({
              color: bullet.color,
              id: bullet.pk,
              avatar: bullet.avatar,
              msg: bullet.msg,
              time: Math.floor(Math.random() * 10 + 2),
              nickname: bullet.name,
              extraWidth: bullet.extraWidth,
            });
          }
        })
      }, 13500)

    },
    getUser () {
      this.user = window.sessionStorage.getItem("user");
      if (this.user != null) {
        this.nickname = JSON.parse(this.user).nickname;
        this.avatar = JSON.parse(this.user).avatar;
      }
    },
    userDetail () {
      // 创建页面的之前获取用户的信息
      var that = this;
      this.selfLog("页面创建之前获取用户详情：");
      this.$axios.user_profile().then((res) => {
        that.user_profile = res["data"];
        this.selfLog("页面创建之前获取用户详情：");
        this.selfLog(that.user_profile);
      });
    },
  },
  beforeDestroy () {
    clearInterval(this.interval_id)
  }
};
</script>

<style scoped>
.message_banner1 {
  position: absolute;
  top: 0px;
  left: 9%;
  right: 0;
  height: 100%;
  width: 100%;
  background: #49b1f5 url(../../assets/uestc/1.webp) no-repeat;
  animation: header-effect 1s;
  background-size: 100% 100%;
}
.message_banner2 {
  position: absolute;
  top: 0px;
  left: 9%;
  right: 0;
  height: 100%;
  width: 100%;
  background: #49b1f5 url(../../assets/uestc/2.webp) no-repeat;
  animation: header-effect 1s;
  background-size: 100% 100%;
}
.message_banner3 {
  position: absolute;
  top: 0px;
  left: 9%;
  right: 0;
  height: 100%;
  width: 100%;
  background: #49b1f5 url(../../assets/uestc/3.webp) no-repeat;
  animation: header-effect 1s;
  background-size: 100% 100%;
}
.message_banner4 {
  position: absolute;
  top: 0px;
  left: 9%;
  right: 0;
  height: 100%;
  width: 100%;
  background: #49b1f5 url(../../assets/uestc/4.webp) no-repeat;
  animation: header-effect 1s;
  background-size: 100% 100%;
}
.message-title {
  color: #393e46;
  animation: title-scale 1s;
  /* font-size: 38px; */
}
.message-container {
  position: absolute;
  width: 360px;
  top: 35%;
  left: 0;
  right: 0;
  text-align: center;
  margin: 0 auto;
  color: #222831;
}
.message-input-wrapper {
  display: flex;
  justify-content: center;
  height: 2.5rem;
  margin-top: 2rem;
}
.message-input-wrapper input {
  outline: none;
  width: 70%;
  border-radius: 20px;
  height: 100%;
  padding: 0 1.25rem;
  color: #222831;
  border: #222831 1px solid;
}
.message-input-wrapper input::-webkit-input-placeholder {
  color: #222831;
}
.message-input-wrapper button {
  width: 120px;
  outline: none;
  border-radius: 20px;
  height: 100%;
  padding: 0 1.25rem;
  border: #222831 1px solid;
}
.barrage-container {
  position: absolute;
  top: 50px;
  left: 0;
  right: 0;
  bottom: 0;
  height: calc(100% -50px);
  width: 100%;
}
.barrage-items {
  background: rgb(0, 0, 0, 0.7);
  border-radius: 100px;
  color: #fff;
  padding: 5px 10px 5px 5px;
  align-items: center;
  display: flex;
}
.el-carousel__item h3 {
  color: #475669;
  font-size: 18px;
  opacity: 0.75;
  line-height: 300px;
  margin: 0;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
</style>