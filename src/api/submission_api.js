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
	GetUserSubmissions: (limit, offset) => {
		let params = {
			'limit': limit,
			'offset': offset
		}
		return Get('submission/get_user_submissions', params)
	},
	GetUserSubmissionProfile:()=>{
		return Get('submission/get_user_submission_profile')
	}
}