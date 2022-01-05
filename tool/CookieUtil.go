package tool

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const CookieName = "cookie_user"
const CookieTimeLength = 10 * 60 //10min

func CookieAuth(context *gin.Context) (*http.Cookie, error) {
	cookie, err := context.Request.Cookie(CookieName)
	if err == nil {
		context.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		return cookie, nil
	} else {
		return nil, err
	}

}
