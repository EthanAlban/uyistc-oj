package models

/*
	ACM赛制：每道题提交之后都有反馈，可以看到“通过”、“运行错误”、“答案错误”等等结果，
	但看不到错误的测试样例（leetcode周赛可以看到），每道题都有多个测试点，
	每道题必须通过了所有的测试点才算通过。每道题不限制提交次数，但没通过的话会有罚时，仅以最后一次提交为准。
	比赛过程中一般可以看到实时排名，通过题数相同的情况下按照答题时间+罚时来排名。
*/

// SubmissionacmTable1  用于存储Acm比赛相关的数据库
type SubmissionacmTable1 struct {
	Submission
	Penalty float64 `orm:"column(penalty)"`
}
