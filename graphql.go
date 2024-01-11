package main

import (
	"context"
	"github.com/machinebox/graphql"
)

func getQueryUserInfo() string {
	return `query ($username: String!) { matchedUser(username: $username) { 
				username
				submitStats { 
				acSubmissionNum { 
					difficulty 
					count 
					} 
				} 
			}
		}`
}

func getQueryQntyQuestions() string {
	return `{ allQuestionsCount { 
			difficulty 
			count 
			}
		}`
}

func getUsersInfo(username string) UserProfileData {
	var requestUser UserProfileData
	client := graphql.NewClient("https://leetcode.com/graphql")
	query := getQueryQntyQuestions()
	request := graphql.NewRequest(query)
	ctx := context.Background()
	err := client.Run(ctx, request, &requestUser)
	check(err)
	query = getQueryUserInfo()
	request = graphql.NewRequest(query)
	request.Var("username", username)
	err = client.Run(ctx, request, &requestUser)
	if err != nil {
		check(err)
	}
	return requestUser
}
