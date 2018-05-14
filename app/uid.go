package app

import "github.com/rs/xid"

func CreateUuid() string {
	uid := xid.New()

	return uid.String()
}
