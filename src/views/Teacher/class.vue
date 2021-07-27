<template>
  <div>
    <span class="title">班级列表</span>
    <div id="qrCode" ref="qrCodeDiv"></div>
    <!-- <el-button type="primary" @click="creatQrCode">发起签到</el-button> -->
    <el-collapse v-model="activeNames" @change="handleChange">
      <el-collapse-item :title="item.id+'.'+item.name" :name="item.id" v-for="item in course_list" :key="item.id">
        <!-- 先展示些课程信息 -->
        <div style="float:left">学院:{{item.departname}}、开课年份:{{item.year}} </div>

        <el-table :data="item.stu_list" style="width: 100%">
          <el-table-column prop="stunum" label="学号" width="180"></el-table-column>
          <el-table-column prop="name" label="姓名" width="180"></el-table-column>
          <el-table-column prop="department" label="学院"></el-table-column>
          <el-table-column fixed="right" label="操作" width="100">
            <template slot-scope="scope">
              <el-button @click="handleClick(scope.row)" type="text" size="small">查看详情</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-collapse-item>
    </el-collapse>
  </div>
</template>

<script>
import QRCode from 'qrcodejs2'
export default {
  name: "Class",
  mounted () {
    // 拿教师信息
    this.$axios.user_profile().then(res => {
      this.user_profile = res["data"];
      this.selfLog(this.user_profile);
      let param = {
        work_id: this.user_profile.user.username
      }
      this.$axios.get_teacher_courses(param).then(res => {
        this.selfLog(res);
        this.course_list = res.data
      })
    })
  },
  data () {
    return {
      user_profile: this.$root.user_profile,
      //   所有课程
      course_list: [],
      activeNames: [],
      tableData: []
    };
  },
  methods: {
    handleChange (val) {
      this.selfLog(val);
    },
    handleClick (row) {
      this.selfLog(row);
    },
    getStudent () {
      for (let i = 1; i <= 5; i++) {
        let obj = {
          name: i,
          depart: "计蒜鸡学院"
        };
        this.tableData.push(obj);
      }
    },
    // 发布作业
    publish_task (stu_list) {
      this.$router.push('createproblem')
      //测试老师id  3204170
      this.selfLog(stu_list);
    },
    creatQrCode () {
      new QRCode(this.$refs.qrCodeDiv, {
        text: 'https://fslchz.api.larkfn.com/get_userinfo',
        width: 200,
        height: 200,
        colorDark: "#333333", //二维码颜色
        colorLight: "#ffffff", //二维码背景色
        correctLevel: QRCode.CorrectLevel.L//容错率，L/M/H
      })
    },
  },
  computed: {},
};
</script>

<style lang='scss' scoped>
.title {
  font-size: 30px;
  line-height: 100px;
}
</style>
