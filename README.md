# Gin Passport Google

Gin middleware for obtaining common google+ profile information. I don't personally use all of the permission attributes but feel free to open an issue and I can take a look into it (or open a pull request).

## Example

```go
import (
  "github.com/gin-gonic/gin"

  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"

  "github.com/durango/gin-passport-google"
)

func main() {
  opts, _ := oauth2.New(
    oauth2.Client("ClientId", "YourSecretKey"),
    oauth2.RedirectURL("Your redirect URL"),
    oauth2.Scope("https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"),
    google.Endpoint(),
  )

  router := gin.Default()

  googleAuth := router.Group("/auth/google")

  // setup the configuration and mount the "/login" route
  GinPassportGoogle.Routes(opts, googleAuth)

  // setup a customized callback url...
  googleAuth.GET("/callback", GinPassportGoogle.GetProfile(), func(c *gin.Context) {
    user, userError := c.Get(GinPassportGoogle.KeyNamespace)
    if ! userError {
      c.AbortWithStatus(500)
      return
    }

    // cast gin's middleware variable
    googleUser := user.(*GinPassportGoogle.Profile)
    c.String(200, "Got it!")
  })
}
```
