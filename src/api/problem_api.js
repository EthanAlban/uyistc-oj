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
		return Get('problems/get_pages_problems', params)
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
		return Post('judger/new_judge_task', params)
	},
	// 查询对应的submissionID的提交是否存在
	IsSubmissionExsit: submissionID => {
		return Get('submission/submision_exist?submissionID=' + submissionID)
	},
	// 查询提交的最新状态
	GetSubmissionFinal: submissionID => {
		return Get('submission/get_submission_final?submissionID=' + submissionID)
	},
	// 查询问题的提交次数
	GetAcSubTimes: pid => {
		return Get('problems/get_problem_ac_sub_times?pid=' + pid)
	},
	// 查询问题提交情况的统计信息
	GetSubmissionStaticForProblem: pid => {
		return Get('submission/get_submission_static_for_problem?pid=' + pid)
	}

}
