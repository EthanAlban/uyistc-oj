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
	// 分页获取问题列表
	ProblemListPage: (limit, offset) => {
		let params = {
			'limit': limit,
			'offset': offset
		}
		return Post('problems/get_pages_problems', params)
	},
	//根据问题id获取问题的详情
	ProblemDetailById: Pid => {
		return Get('problems/get_problem_detail_by_id?pid=' + Pid)
	},
	//按问题标签进行分类
	ProblemsListByTag: TagName => {
		return Get('problems/get_problems_by_tag?tagname=' + TagName)
	},
	GetProblemTagsById: pid => {
		return Get('problems/get_problem_tags_by_id?pid=' + pid)
	},
	// 新的判题任务
	Submission: params => {
		return Post('judger/new_judge_task',params)
	}
}
