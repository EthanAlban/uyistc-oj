import { Get, Post, Put, Delete,getCookie } from '@/api/request'
import { selfLog, UrlEncode } from '../utils'
import axios from 'axios'
export default {
	GetAnnoncements: params => {
	    return Post('annonce/get_annoncements',params)
	},
}
