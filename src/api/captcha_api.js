import {
	Get2,
	Post2
} from '@/api/request'
import {
	selfLog,
	UrlEncode
} from '../utils'
import axios from 'axios'
export default {
	NewCaptId: () => {
		return Get2('new')
	},
	// 鉴别验证码
	VerifyCaptcha: params=>{
		return Get2('verify',params)
	}
}
