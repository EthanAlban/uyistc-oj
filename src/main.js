import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import dataV from '@jiaminghi/data-view'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import { vueBaberrage } from 'vue-baberrage'


import Api from './api/index'
import UserApi from './api/user_api'
import UtilsApi from './api/utils_api'
import CaptcahApi from './api/captcha_api'
import AnnonceApi from './api/annonce_api'
import ProblemApi from './api/problem_api'
import SubmissionApi from './api/submission_api'
import ContestsApi from './api/contests_api'


import './icons'
import VueCookies from 'vue-cookies'
import VueCodemirror from 'vue-codemirror'
import 'default-passive-events'
import { mavonEditor } from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
// 引入懒加载
import VueLazyload from 'vue-lazyload'
import loading from './assets/uestc/calendar/5.jpeg'



// import base style
// import 'codemirror/lib/codemirror.css'
// 下载到本地自定义
import './codemirror.css'

// import more codemirror resource...

// you can set default global options and events when Vue.use
Vue.use(VueCodemirror,
	/* {
  options: { theme: 'base16-dark', ... },
  events: ['scroll', ...]
} */
)

// 懒加载
Vue.use(VueLazyload, { loading: loading })

import VueQuillEditor from 'vue-quill-editor'

import 'quill/dist/quill.core.css' // import styles
import 'quill/dist/quill.snow.css' // for snow theme
import 'quill/dist/quill.bubble.css' // for bubble theme
// 导入文件导出插件js
import Blob from './views/Contest/export/Blob'
import Export2Excel from './views/Contest/export/Export2Excel'

Vue.use(VueQuillEditor, /* { default global options } */ )
Vue.use(VueCookies)

Vue.prototype.$axios = Api
Vue.prototype.$user_axios = UserApi
Vue.prototype.$utils_axios = UtilsApi
Vue.prototype.$captcha_axios = CaptcahApi
Vue.prototype.$annonce_axios = AnnonceApi
Vue.prototype.$problem_axios = ProblemApi 
Vue.prototype.$submission_axios = SubmissionApi  
Vue.prototype.$contests_axios = ContestsApi

// 全局的this.selfLog替换
// import selfLog from './utils/selfLog';
import { selfLog } from './utils'
import { decodeUtf8 } from './utils'
import { encodeUtf8 } from './utils'

Vue.prototype.selfLog = selfLog
Vue.prototype.decodeUtf8 = decodeUtf8
Vue.prototype.encodeUtf8 = encodeUtf8

Vue.prototype.$OJIP = 'http://47.94.135.51'
Vue.prototype.$tenant_access_token = ''
Vue.prototype.$redirect_uri = 'http://localhost:8080/login'

import VueClipboard from 'vue-clipboard2'
import contribution from 'vue-contribution'
import 'vue-contribution/dist/vue-contribution.css'
Vue.use(contribution)
Vue.use(VueClipboard)
Vue.use(vueBaberrage)
Vue.use(dataV)
Vue.use(ElementUI)

Vue.config.productionTip = false

import Moment from 'moment'

// 定义全局时间戳过滤器
Vue.filter('formatDate', function(value) {
	return Moment(value).format('YYYY-MM-DD HH:mm:ss')
})



new Vue({
	beforeCreate() {
		this.$cookies.config('1d')
		// 前端在启动的时候检查两台服务器的在线情况  心跳监测
		this.selfLog('项目启动')
		this.$utils_axios.HeartBeat().then(res => {
			selfLog(res)
			if (res.errcode === 200) {
				this.sys_status = 1
			} else {
				this.sys_status = 0
			}
		}).catch(err => {
			this.sys_status = -1
		})
		this.setInterval_id = setInterval(() => {
			this.selfLog('心跳检测...')
			this.$utils_axios.HeartBeat().then(res => {
				selfLog(res)
				if (res.errcode === 200) {
					this.sys_status = 1
				} else {
					this.sys_status = 0
				}
			}).catch(err => {
				this.selfLog('心跳检测失败...')
				this.sys_status = -1
			})
		}, 60000)

	},
	router,
	store,
	data: function() {
		return {
			sys_status: 1,
			setInterval_id: -1,
			user_profile: null,
		}
	},

	beforeDestroy() {
		clearInterval(this.setInterval_id)
	},

	render: h => h(App)
}).$mount('#app')