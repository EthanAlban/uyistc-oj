package structs

import "time"

type Submission struct {
	SubmissionId string `json:"SubmissionId"`
	ProblemId    struct {
		Pid               int    `json:"Pid"`
		Title             string `json:"Title"`
		Level             int    `json:"Level"`
		TotalSubmissions  int    `json:"TotalSubmissions"`
		AcceptSubmissions int    `json:"AcceptSubmissions"`
		Content           string `json:"Content"`
		TimeLimit         int    `json:"TimeLimit"`
		MemoryLimit       int    `json:"MemoryLimit"`
		Hint              string `json:"Hint"`
		ProblemType       struct {
			Tid      int    `json:"Tid"`
			TypeName string `json:"TypeName"`
		} `json:"ProblemType"`
		ProblemTags interface{} `json:"ProblemTags"`
		Uid         struct {
			UId       int    `json:"UId"`
			UserName  string `json:"UserName"`
			UserType  int    `json:"UserType"`
			RealName  string `json:"RealName"`
			Password  string `json:"Password"`
			Credit    int    `json:"Credit"`
			Email     string `json:"Email"`
			Tel       string `json:"Tel"`
			Signature string `json:"Signature"`
			Major     string `json:"Major"`
			Gitaddr   string `json:"Gitaddr"`
			Blogaddr  string `json:"Blogaddr"`
			Avatar    string `json:"Avatar"`
		} `json:"Uid"`
		InputDescription  string `json:"InputDescription"`
		OutputDescription string `json:"OutputDescription"`
		Template          string `json:"Template"`
	} `json:"ProblemId"`
	CreateTime time.Time `json:"CreateTime"`
	UserId     struct {
		UId       int    `json:"UId"`
		UserName  string `json:"UserName"`
		UserType  int    `json:"UserType"`
		RealName  string `json:"RealName"`
		Password  string `json:"Password"`
		Credit    int    `json:"Credit"`
		Email     string `json:"Email"`
		Tel       string `json:"Tel"`
		Signature string `json:"Signature"`
		Major     string `json:"Major"`
		Gitaddr   string `json:"Gitaddr"`
		Blogaddr  string `json:"Blogaddr"`
		Avatar    string `json:"Avatar"`
	} `json:"UserId"`
	UserName string `json:"UserName"`
	Code     string `json:"Code"`
	Result   int    `json:"Result"`
	Language struct {
		Lid      int    `json:"Lid"`
		Language string `json:"Language_"`
		Template string `json:"Template"`
		Suffix   string `json:"Suffix"`
	} `json:"Language"`
	Score   int    `json:"Score"`
	ErrInfo string `json:"ErrInfo"`
}
