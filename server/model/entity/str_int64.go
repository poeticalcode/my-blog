package entity

import (
	"strconv"
	"strings"
)

type DbUInt64 uint64

// UnmarshalJSON 这里主要注意反序列化时需要去掉两边的引号
func (d *DbUInt64) UnmarshalJSON(data []byte) (err error) {
	s := strings.Trim(string(data), "\"")
	intNum, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*d = DbUInt64(intNum)
	return
}

// MarshalJSON 这里需要手动加上引号，不会自动生成
func (d DbUInt64) MarshalJSON() ([]byte, error) {
	return ([]byte)("\"" + strconv.FormatInt(int64(d), 10) + "\""), nil
}
