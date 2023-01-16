package utils

import "github.com/rs/xid"

// GetXid 获取唯一识别id
func GetXid() string {
	id := xid.New()
	return id.String()
}
