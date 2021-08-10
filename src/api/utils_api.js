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
	WeatherToday: () => {
		return Get('utils/weather')
	},
	HeartBeat: () => {
		return Get('utils/heartbeat')
	},
	//设置系统通知信息已读
	SetInfoRead:id=>{
		return Get('utils/set_info_read?Iid='+id)
	}
}
