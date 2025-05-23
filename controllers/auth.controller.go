package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v72/github"
	"github.com/lubie-placki-be/configs"
	"github.com/lubie-placki-be/middlewares"
	"github.com/lubie-placki-be/services"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

var oauthConf = &oauth2.Config{
	ClientID:     configs.EnvClientId(),
	ClientSecret: configs.EnvClientSecret(),
	Endpoint:     githuboauth.Endpoint,
}

func Login(c *gin.Context) {
	url := oauthConf.AuthCodeURL(configs.EnvAuthState(), oauth2.AccessTypeOnline)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func Logout(c *gin.Context) {
	accessToken, err := c.Cookie("access-token")
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, configs.EnvClientPath())
		return
	}

	middlewares.GithubClient.Authorizations.Revoke(context.TODO(), configs.EnvClientId(), accessToken)
	c.SetCookie("access-token", accessToken, -1, "/", "localhost", true, true)

	c.Redirect(http.StatusTemporaryRedirect, configs.EnvClientPath())
}

func Callback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")

	if state != configs.EnvAuthState() {
		c.Redirect(http.StatusTemporaryRedirect, configs.EnvClientPath())
		return
	}

	token, err := oauthConf.Exchange(context.TODO(), code)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, configs.EnvClientPath())
		return
	}

	c.SetCookie("access-token", token.AccessToken, int(token.ExpiresIn), "/", "localhost", true, true)

	githubClient := github.NewClient(nil).WithAuthToken(token.AccessToken)

	if err = services.CreateUser(githubClient); err != nil {
		c.Redirect(http.StatusTemporaryRedirect, configs.EnvClientPath())
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, configs.EnvClientPath())
}

func GetMe(c *gin.Context) {
	me, err := services.GetMe()
	if err != nil {
		c.IndentedJSON(http.StatusNoContent, nil)
		return
	}

	c.Writer.Header().Set("Cache-Control", "max-age=3600, public")
	c.Writer.Header().Set("Vary", "Cookie")

	c.IndentedJSON(http.StatusOK, me)
}
