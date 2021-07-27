<template>
  <div style="overflow-y:auto">
    <div id="rank" style="width: 100%;height: 400px;top:40px"></div>
    <el-button type="primary" @click="show_table">排名统计</el-button>

    <el-drawer title="做题记录" append-to-body :visible.sync="showtable" direction="rtl" size="50%" max-height="700">
      <el-table :data="rank_list" stripe style="width: 100%">
        <el-table-column prop="user.username" label="用户名" width="200">
        </el-table-column>
        <el-table-column prop="passrate" sortable label="通过率" width="150">
        </el-table-column>
        <el-table-column prop="accepted_number" sortable label="通过数" width="150">
        </el-table-column>
        <el-table-column prop="submission_number" sortable label="提交数" width="150">
        </el-table-column>
      </el-table>
    </el-drawer>
  </div>
</template>

<script>
import * as echarts from "echarts";
export default {
  mounted () {
    this.getRankData()
  },
  data () {
    return {
      rank_list: [],
      limit: 10,
      page: 1,
      total: 0,
      offset: 0,
      showtable: false
    };
  },
  methods: {
    getRankData () {
        this.selfLog("获取排名信息：")
        //初始化ehcharts实例
        var myChart = echarts.init(document.getElementById("rank"));
        //指定图表的配置项和数据
        let user_list = [2018,2019,2020,2021]
        let xue_master = [152,60,90,110]
        let zhuan_master = [168,110,227,142] 
        let wangan_master = [55,13,19,20]
        var option = {
          //标题
          title: {
            text: '录取人数'
          },
          //工具箱
          //保存图片
          toolbox: {
            show: true,
            feature: {
              saveAsImage: {
                show: true
              }
            }
          },
          //图例-每一条数据的名字叫销量
          legend: {
            data: ['学硕', '专硕', '网安']
          },
          //x轴
          xAxis: {
            data: user_list
          },
          //y轴没有显式设置，根据值自动生成y轴
          yAxis: {},
          //数据-data是最终要显示的数据
          series: [{
            name: '学硕',
            type: 'line',
            data: xue_master
          },
          {
            name: '专硕',
            type: 'line',
            data: zhuan_master
          },
          {
            name: '网安',
            type: 'line',
            data: wangan_master
          }]
        };
        //使用刚刚指定的配置项和数据项显示图表
        myChart.setOption(option);
    },
    show_table () {
      this.showtable = true
    }
  }
}
</script>

<style>
</style>