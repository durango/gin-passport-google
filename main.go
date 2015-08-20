package GinPassportGoogle

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

const KeyNamespace string = "gin-passport-google-profile"

const ProfileUrl string = "https://www.googleapis.com/oauth2/v1/userinfo?alt=json"

var passportOauth *oauth2.Options

func Routes(oauth *oauth2.Options, r *gin.RouterGroup) {
	passportOauth = oauth

	r.GET("/login", func(c *gin.Context) {
		Login(oauth, c)
	})
}

func Login(oauth *oauth2.Options, c *gin.Context) {
	url := oauth.AuthCodeURL("", "", "")
	c.Redirect(http.StatusFound, url)
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		getProfile(c)
	}
}

func GetProfile(c *gin.Context) *Profile {
	user, exists := c.Get(KeyNamespace)
	if !exists {
		return nil
	}

	return user.(*Profile)
}

func getProfile(c *gin.Context) {
	c.Request.ParseForm()

	opts := passportOauth
	code := c.Request.Form.Get("code")

	t, err := opts.NewTransportFromCode(code)
	// most likely already authenticated / all errors will return `t` as nil
	if t == nil {
		c.Redirect(301, "/")
		return
	} else if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	client := http.Client{Transport: t}

	resp, err := client.Get(ProfileUrl)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var userInformation Profile
	err = json.Unmarshal(contents, &userInformation)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Set(KeyNamespace, &userInformation)
	c.Next()
}
