<template>
  <div class='terminal'>
    <h1 id='title'>404</h1>
    <div class='text' id='text_1'></div><br>
    <div class='text2' id='text_2'></div><br>
    <div class='text' id='text_3'></div>
    <div class='cursor'>_</div>
    <div class='c'>
      <a href='/' title='go to index'>
      <el-link type="primary">返回首页</el-link>
      </a>
    </div>
  </div>
</template>

<script>
import $ from 'jquery'
export default {
  mounted () {
    var text_1 = 'this.page===undifined',
      text_2 = 'warining:this page is ---^^^------------ not defined in the system',
      text_3 = '该页面不存在或者被删除，请返回首页或者联系我们.....';
    repeat(text_1, 1);

    setTimeout(function () {
      console.log(2);
      repeat(text_2, 2);
    }, text_1.length * 60 + 1000);

    setTimeout(function () {
      console.log(3);
      repeat(text_3, 3);
    }, (text_2.length + text_1.length) * 60 + 2000);

    function repeat (text, n) {
      var i = 0,
        repeatable = setInterval(function () {
          $('#text_' + n).text($('#text_' + n).text() + text[i]);
          i++;
          if (i >= text.length) {
            clearInterval(repeatable);
          }
        }, 100);
    }
    //# sourceURL=pen.js
  }
}

</script>

<style>
a {
  color: inherit;
}

.terminal {
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 60%;
  padding: 0 20%;
  background-color: #010101;
  color: #00ee00;
  font-family: consolas;
  word-wrap: break-all;
}
.terminal #title {
  text-align: center;
  width: 100%;
  font-size: 6em;
}
.terminal .text {
  display: inline-block;
}
.terminal .text2 {
  color: #ff8040;
  display: inline-block;
}
.terminal .cursor {
  display: inline;
  animation: blink 0.75s linear infinite;
}
.terminal .c {
  position: absolute;
  bottom: 50%;
  right: 48%;
}
.terminal .c .hearth {
  display: inline !important;
}

@keyframes blink {
  0% {
    opacity: 0;
  }
  33% {
    opacity: 1;
  }
  66% {
    opacity: 1;
  }
  100% {
    opacity: 0;
  }
}
</style >