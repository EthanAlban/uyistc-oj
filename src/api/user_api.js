import { Get, Post, Put, Delete,getCookie } from '@/api/request'
import { selfLog, UrlEncode } from '../utils'
import axios from 'axios'
export default {
	GetUserByUsernamePwd: () => {
	    return Get('usr/getuser_by_username_pwd')
	},
}
