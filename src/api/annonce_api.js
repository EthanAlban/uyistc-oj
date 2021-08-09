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
	GetAnnocements: (limit, offset) => {
		let params = {
			'limit': limit,
			'offset': offset
		}
		return Post('annonce/get_annoncements', params)
	},
	GetSysInfolist: (limit, offset) => {
		let params = {
			'limit': limit,
			'offset': offset
		}
		return Get('sysinfo/get_sysinfo_list', params)
	}
}
