package app

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

const (
	userCookieName = "_uid"
)

// CreateUidCookie create cookie for uid
func CreateUidCookie(c echo.Context) *http.Cookie {
	cookie := &http.Cookie{}
	cookie.Name = userCookieName
	cookie.Value = CreateUuid()
	cookie.Expires = time.Now().Add(2 * time.Hour)
	c.SetCookie(cookie)

	return cookie
}

// GetUidCookie uidキーのCookieを取得する
func GetUidCookie(c echo.Context) (*http.Cookie, error) {
	cookie, err := c.Cookie(userCookieName)

	return cookie, err
}
