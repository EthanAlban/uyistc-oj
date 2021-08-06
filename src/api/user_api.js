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
	Register: params => {
		selfLog(params)
		return Post('usr/register', params)
	},
}
