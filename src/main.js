import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import dataV from '@jiaminghi/data-view';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import { vueBaberrage } from 'vue-baberrage'
import Api from './api/index';
import './icons'
import VueCookies from 'vue-cookies'
import VueCodemirror from 'vue-codemirror'
import 'default-passive-events'
import { mavonEditor } from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
// 引入懒加载
import VueLazyload from 'vue-lazyload'
import loading from './assets/uestc/calendar/5.jpeg'

// 全局的this.selfLog替换
// import selfLog from './utils/selfLog';
import { selfLog } from './utils'
Vue.prototype.selfLog = selfLog;
// import base style
// import 'codemirror/lib/codemirror.css'
// 下载到本地自定义
import './codemirror.css'

// import more codemirror resource...

// you can set default global options and events when Vue.use
Vue.use(VueCodemirror, /* {
  options: { theme: 'base16-dark', ... },
  events: ['scroll', ...]
} */)

// 懒加载
Vue.use(VueLazyload, {
    loading: loading
})

import VueQuillEditor from 'vue-quill-editor'

import 'quill/dist/quill.core.css' // import styles
import 'quill/dist/quill.snow.css' // for snow theme
import 'quill/dist/quill.bubble.css' // for bubble theme
// 导入文件导出插件js
import Blob from './views/Teacher/export/Blob'
import Export2Excel from './views/Teacher/export/Export2Excel'

Vue.use(VueQuillEditor, /* { default global options } */)
Vue.use(VueCookies)
Vue.prototype.$axios = Api;

Vue.prototype.$OJIP = "http://47.94.135.51"
Vue.prototype.$tenant_access_token = ""
Vue.prototype.$redirect_uri = "http://localhost:8080/login"

import VueClipboard from 'vue-clipboard2'
import contribution from 'vue-contribution'
import 'vue-contribution/dist/vue-contribution.css';
Vue.use(contribution)
Vue.use(VueClipboard)
Vue.use(vueBaberrage)
Vue.use(dataV);
Vue.use(ElementUI);

Vue.config.productionTip = false

import Moment from 'moment'

// 定义全局时间戳过滤器
Vue.filter('formatDate', function (value) {
    return Moment(value).format('YYYY-MM-DD HH:mm:ss')
})

new Vue({
    beforeCreate () {
        // 前端在启动的时候检查两台服务器的在线情况  心跳监测
        this.selfLog("项目启动");
        this.setInterval_id = setInterval(() => {
            this.$axios.heartbeat().then(res => {
                this.selfLog("心跳检测.....");
                this.selfLog(res);
                if (res.err === "OJ_is_down") {
                    this.sys_status = -1
                }
                else if (res.err === null) {
                    this.sys_status = 1
                }
            }).catch(err => {
                this.sys_status = 0
            })
        }, 50000);

    },
    router,
    store,
    data: function () {
        return {
            sys_status: 1,
            setInterval_id: -1,
            user_profile: null,
        }
    },

    beforeDestroy () {
        clearInterval(this.setInterval_id);
    },
    render: h => h(App)
}).$mount('#app')
