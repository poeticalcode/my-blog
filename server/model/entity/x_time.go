package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type XTime struct {
	time.Time
}

// Scan 读取数据库时会调用该方法将时间数据转换成自定义时间类型
func (t *XTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XTime{value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", value)
}

// Value 为 Xtime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t XTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// MarshalJSON 将 XTime 编码成 json 字符串
func (t XTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

// UnmarshalJSON 将  json 字符串编码成 XTime
func (t *XTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	return json.Unmarshal(b, &t)
}
