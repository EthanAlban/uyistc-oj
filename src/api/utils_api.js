import {
	Get,
	Post,
	Put,
	Delete,
	getCookie
} from '@/api/request'
import {
	selfLog,
	UrlEncode
} from '../utils'
import axios from 'axios'
export default {
	WeatherToday: () => {
		return Get('utils/weather')
	},
	HeartBeat: () => {
		return Get('utils/heartbeat')
	},
	//设置系统通知信息已读
	SetInfoRead: id => {
		return Get('sysinfo/set_info_read?Iid=' + id)
	},
	// 得到系统中的问题标签
	GetTags: () => {
		return Get('utils/get_tags')
	},
	// 获得cookie
	GetCookie: name => {
		return getCookie(name)
	}
}
