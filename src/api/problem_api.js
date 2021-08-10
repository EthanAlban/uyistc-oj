import {
	Get,
	Post
} from '@/api/request'
import {
	selfLog,
	UrlEncode
} from '../utils'
import axios from 'axios'
export default {
	ProblemListPage: (limit, offset) => {
		let params = {
			'limit': limit,
			'offset': offset
		}
		return Post('problems/get_pages_problems',params)
	},
}
