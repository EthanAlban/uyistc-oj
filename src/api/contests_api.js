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
	GetAllContests: (offset, limit) => {
		let params = {
			'limit': limit,
			'offset': offset
		}
		return Get('contests/get_all_contests', params)
	},

	GetContestQuantity: () => {
		return Get('contests/get_contest_quantity')
	},
	// 获取用户访问竞赛权限  
	ContestAccess: contest_id => {
		let params = { 'contest_id': contest_id }
		return Get('contests/contest_access', params)
	},
	ContestPasswordCheck: (contest_id, psw) => {
		let params = {
			'pwd': psw,
			'contest_id': contest_id
		}
		return Get('contests/contest_password_check',params)
	},
	ContestProblemList:contest_id => {
		let params = { 'contest_id': contest_id }
		return Get('contests/contest_problem_list', params)
	},

	GetSysInfolist: (limit, offset) => {
		let params = {
			'limit': limit,
			'offset': offset
		}
		return Get('sysinfo/get_sysinfo_list', params)
	}
}
