package utils

//检测状态
func CheckStatusIndex(status map[int]string, index int) int {
	str := status[index]
	if str == "" {
		return 0
	} else {
		return index
	}
}
