import { Get, Post, Put, Delete,getCookie } from '@/api/request'
import { selfLog, UrlEncode } from '../utils'
import axios from 'axios'
export default {
	GetPagesProblems: params => {
	    return Get('problems/get_pages_problems',params)
	},
}
