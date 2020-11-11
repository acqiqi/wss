package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type WhereData struct {
	Key   string `json:"key"`
	Op    string `json:"op"`
	Value string `json:"value"`
}

func WhereToMap(where_data []WhereData) map[string]interface{} {
	if where_data != nil {
		cb := make(map[string]interface{})
		for _, v := range where_data {
			cb[v.Key+"__"+v.Op] = v.Value
		}
		return cb
	}
	return make(map[string]interface{})
}

func BuildWhere(maps map[string]interface{}) string {
	maps_arr := make([]string, len(maps))

	ii := 0
	for i, v := range maps {
		count_split := strings.Split(i, "__")
		where_ex := "="
		if len(count_split) > 1 {
			where_ex = count_split[1]
			i = count_split[0]
		}
		value := ""
		switch v.(type) {
		case string:
			value = "'" + v.(string) + "'"
			break
		case int:
			value = strconv.Itoa(v.(int))
			break
		case int64:
			value = strconv.FormatInt(v.(int64), 10)
		case float64:
			value = strconv.FormatFloat(v.(float64), 'f', -1, 64)
			break
		case []string:
			for k, t := range v.([]string) {
				v.([]string)[k] = "'" + t + "'"
			}
			value = "(" + strings.Join(v.([]string), ",") + ")"
			break
		case []int:
			str_arr := make([]string, len(v.([]int)))
			for k, t := range v.([]int) {
				str_arr[k] = strconv.Itoa(t)
			}
			value = "(" + strings.Join(str_arr, ",") + ")"
			break
		case []int64:
			str_arr := make([]string, len(v.([]int64)))
			for k, t := range v.([]int64) {
				str_arr[k] = strconv.FormatInt(t, 10)
			}
			value = "(" + strings.Join(str_arr, ",") + ")"
			break
		}
		log.Println(value)

		maps_arr[ii] = fmt.Sprintf("%s %s %s",
			i, where_ex, value)

		ii++
	}
	cb := strings.Join(maps_arr, " AND ")
	log.Println(cb)
	return cb
}
