package utils

import "strings"

// 驼峰转蛇形 snake string, XxYy to xx_yy , XxYY to xx_yy
func FormatSnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

//检测map
func SnakeGetMap(key string, data map[string]interface{}) interface{} {
	if data[key] != nil {
		return data[key]
	}
	if data[FormatCamelString(key, true)] != nil {
		return data[FormatCamelString(key, true)]
	}
	if data[FormatCamelString(key, false)] != nil {
		return data[FormatCamelString(key, false)]
	}
	return ""
}

// 蛇形转驼峰 camel string, xx_yy to XxYy  K首字母是否大小写 F是 T否
func FormatCamelString(s string, k bool) string {
	data := make([]byte, 0, len(s))
	j := false
	//k := true //首字母
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
