package utils

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

// 数组切片求和
func ArraySum(input interface{}) (sum float64, err error) {
	list := reflect.ValueOf(input)
	switch reflect.TypeOf(input).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < list.Len(); i++ {
			val := list.Index(i)
			v, err := toFloat64(val.Interface())
			if err != nil {
				return 0, err
			}
			sum += v
		}
	default:
		return 0, errors.New("input must be slice or array")
	}
	return
}

func toFloat64(in interface{}) (f64 float64, err error) {
	switch val := in.(type) {
	case float64:
		return val, nil
	case float32:
		return float64(val), nil
	case uint8:
		return float64(val), nil
	case uint16:
		return float64(val), nil
	case uint32:
		return float64(val), nil
	case uint64:
		return float64(val), nil
	case uint:
		if strconv.IntSize == 32 || strconv.IntSize == 64 {
			return float64(val), nil
		}
		return 0, errors.New("convert uint to float64 failed")
	case int8:
		return float64(val), nil
	case int16:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case int:
		if strconv.IntSize == 32 || strconv.IntSize == 64 {
			return float64(val), nil
		}
		return 0, errors.New("convert int to float64 failed")
	case bool:
		if val {
			f64 = 1
		}
		return
	case string:
		f64, err = strconv.ParseFloat(val, 64)
		if err == nil {
			return
		}
		return 0, errors.New("convert string to float64 failed")
	default:
		return 0, errors.New("convert to float64 failed")
	}
}

// GenerateRangeNum 生成一个区间范围的随机数
func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min)
	randNum = randNum + min
	fmt.Printf("rand is %v\n", randNum)
	return randNum
}

// 获取抽奖结果 用于大转盘之类的功能
// 数组内是概率 返回item key
func GetRand(proArr []int) int {
	//result := ""

	sum, err := ArraySum(proArr)
	proSum := int(sum)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(proSum)
	result := 0
	for k, v := range proArr {
		randNum := GenerateRangeNum(1, int(proSum))
		if randNum <= v {
			result = k
			break
		} else {
			proSum = proSum - v
		}
	}
	return result
}

// 保留两位小数点
func Decimal(value float64) float64 {
	return math.Trunc(value*1e2+0.5) * 1e-2
}

//生成区间随机数
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}
