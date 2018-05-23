// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package auth

import (
	"bytes"
	"fmt"
	"github.com/XiaoMi/galaxy-sdk-go/thrift"
	"github.com/XiaoMi/galaxy-sdk-go/rpc/errors"
	"github.com/XiaoMi/galaxy-sdk-go/rpc/common"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = errors.GoUnusedProtection__
var _ = common.GoUnusedProtection__
var SIGNATURE_SUPPORT map[UserType]bool
const HK_HOST = "Host"
const HK_TIMESTAMP = "X-Xiaomi-Timestamp"
const MI_DATE = "x-xiaomi-date"
const XIAOMI_HEADER_PREFIX = "x-xiaomi-"
const HK_CONTENT_MD5 = "X-Xiaomi-Content-MD5"
const HK_AUTHORIZATION = "Authorization"
const HK_SERVICE_ADMIN = "Service-Admin"
const HK_SERVICE_MARK = "#"
const HK_ATTACHED_INFO = "Attached-Info"
var SUGGESTED_SIGNATURE_HEADERS []string

func init() {
SIGNATURE_SUPPORT = map[UserType]bool{
    1: false,
    2: true,
    10: true,
    11: true,
    12: false,
    13: false,
}

SUGGESTED_SIGNATURE_HEADERS = []string{
  "Host",   "X-Xiaomi-Timestamp",   "X-Xiaomi-Content-MD5", }

}

