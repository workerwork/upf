package node

import (
	"strconv"
	"strings"
)

func Bytes2Str(b []byte) string {
	//TODO::v4 or v6?
	return strings.Join([]string{strconv.Itoa(int(b[0])), strconv.Itoa(int(b[1])), strconv.Itoa(int(b[2])), strconv.Itoa(int(b[3]))}, ".")
}
