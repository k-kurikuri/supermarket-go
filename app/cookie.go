package app

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

const (
	userCookieName = "_uid"
)

func SetUidCookie(c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = userCookieName
	cookie.Value = CreateUuid()
	cookie.Expires = time.Now().Add(2 * time.Hour)
	c.SetCookie(cookie)
}

// uidキーのCookieを取得する
func GetUidCookie(c echo.Context) (*http.Cookie, error) {
	cookie, err := c.Cookie(userCookieName)

	return cookie, err
}
