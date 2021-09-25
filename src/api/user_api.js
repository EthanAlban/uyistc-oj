import {
	Get,
	Post,
	Put,
	Delete
} from '@/api/request'
import {
	selfLog,
	UrlEncode
} from '../utils'
import axios from 'axios'
export default {
	Login: params => {
		return Post('usr/login', params)
	},
	Logout: () => {
		return Get('usr/logout')
	},
	Register: params => {
		return Post('usr/register', params)
	},
	user_profile: () => {
		return Get('usr/user_profile')
	},
	login_history: () => {
		return Get('logs/get_cur_user_loginhistory')
	}
}
