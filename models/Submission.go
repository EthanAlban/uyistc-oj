package models

import "time"

type Submission struct {
	SubmissionId string    `orm:"column(submission_id);pk"`
	ProblemId    *Problems `orm:"column(problem_id);rel(fk)"`
	CreateTime   time.Time `orm:"column(create_time)"`
	UserId       *User     `orm:"column(user_id);rel(fk)"`
	UserName     string    `orm:"column(username)"`
	Code         string    `orm:"column(code)"`
	Result       int       `orm:"column(result)"`
	Language     *Language `orm:"column(language);rel(fk)"`
	Score        float64   `orm:"column(score)"`
	ErrInfo      string    `orm:"column(err_info)"`
}

// TableName 获取对应数据库表名.
func (sb *Submission) TableName() string {
	return "submission"
}

// TableEngine 获取数据使用的引擎.
func (sb *Submission) TableEngine() string {
	return "INNODB"
}

func NewSubmission() *Submission {
	return &Submission{}
}

func (sb *Submission) InsertNewSubmission(sub Submission) error {
	_, err := O.Insert(&sub)
	if err != nil {
		return err
	}
	return err
}

// GetSubmissionByID 根据id查询submission
func (sb *Submission) GetSubmissionByID(submissionID string) (error, Submission) {
	var submission Submission
	err := O.QueryTable("submission").Filter("SubmissionId", submissionID).One(&submission)
	return err, submission
}
