<template>
  <div>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>示例{{post.id}}</span>
        <!-- 这边删除按钮绑定函数，传出给父组件 -->
        <!-- 向父组件抛出事件 -->
        <el-button style="float: right; padding: 3px 0" type="text" @click="$emit('delete', post.id)">删除</el-button>
      </div>
      <div class="sampleBox">
        <el-input type="textarea" :rows="3" placeholder="样例输入" v-model="post.input"></el-input>
        <el-input type="textarea" :rows="3" placeholder="样例输出" v-model="post.output"></el-input>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "Sample",
  props:["post"],
  data() {
    return {
    };
  },
  methods: {},
  computed: {}
};
</script>

<style lang='scss' scoped>
.sampleBox{
    display: flex;
    justify-content: space-around;
}
</style>
