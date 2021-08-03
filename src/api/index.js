import { Get, Post, Put, Delete,getCookie } from '@/api/request'
import { selfLog, UrlEncode } from '../utils'
import axios from 'axios'
export default {
    testLink: () => {
        return Get('test')
    },
    GetMonthlyRemainRate:()=>{
      return Get('monthly/monthly_remain')
    },
    MonthlyOrders:()=>{
      return Get('order/monthly')
    },
    // 获取月读余额

    // =========================================================工具类==============================================
    getCookie: name => {
        return getCookie(name)
    }

}
