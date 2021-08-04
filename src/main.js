import Vue from 'vue'
import App from './App'
import uView from 'uview-ui'
import Api from './api/index'
import UserApi from './api/user_api'
import AnnonceApi from './api/annonce_api'
import ProblemApi from './api/problem'

Vue.use(uView)
Vue.config.productionTip = true
Vue.prototype.$axios = Api
Vue.prototype.$user_axios = UserApi
Vue.prototype.$annonce_axios = AnnonceApi
Vue.prototype.$problem_axios = ProblemApi

import { selfLog } from './utils'
Vue.prototype.selfLog = selfLog


App.mpType = 'app'

const app = new Vue({ ...App })
app.$mount()
