package utils

import (
	"database/sql/driver"
	"fmt"
	"log"
	"time"
)

// 1. 创建 time.Time 类型的副本 XTime；
type Time struct {
	time.Time
}

const (
	format  = "2006-01-02 15:04:05"
	format2 = "2006-01-02"
)

// 2. 为 Xtime 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t Time) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format(format))

	return []byte(output), nil
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" {
		return nil
	}
	now, err := time.ParseInLocation(`"`+format+`"`, string(data), time.Local)
	if err != nil {
		now, err = time.ParseInLocation(`"`+format2+`"`, string(data), time.Local)
	}
	*t = Time{now}
	return
}

// 3. 为 Xtime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// 4. 为 Xtime 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// 字符串转时间戳
func StringToTimeUnix(t string) int64 {
	timeTemplate1 := "2006-01-02 15:04:05" //常规类型
	// ======= 将时间字符串转换为时间戳 =======
	stamp, _ := time.ParseInLocation(timeTemplate1, t, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	log.Println(stamp.Unix())                                      //输出：1546926630
	return stamp.Unix()
}

//时间戳转字符串
func TimeUnixToString(timeUnix int64) string {
	timeTemplate1 := "2006-01-02 15:04:05"                    //常规类型
	log.Println(time.Unix(timeUnix, 0).Format(timeTemplate1)) //输出：2019-01-08 13:50:30
	return time.Unix(timeUnix, 0).Format(timeTemplate1)
}
