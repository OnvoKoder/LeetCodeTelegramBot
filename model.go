package main

type Submission struct {
	Count      int    `json:"count"`
	Difficulty string `json:"difficulty"`
}

type UserProfileData struct {
	MatchedUser       MatchedUser  `json:"matchedUser"`
	AllQuestionsCount []Submission `json:"allQuestionsCount"`
}

type SubmitStats struct {
	AcSubmissionNum []Submission `json:"acSubmissionNum"`
}

type MatchedUser struct {
	Username    string      `json: username`
	SubmitStats SubmitStats `json:"submitStats"`
}
