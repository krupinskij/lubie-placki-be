package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v72/github"
)

func Github() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access-token")
		if err != nil {
			GithubClient = github.NewClient(nil)
			c.Next()
		}

		GithubClient = github.NewClient(nil).WithAuthToken(accessToken)
	}
}

var GithubClient *github.Client
