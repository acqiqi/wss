package utils

import "regexp"

func CheckMobile(mobileNum string) bool {
	reg := regexp.MustCompile("1([38][0-9]|14[57]|5[^4])\\d{8}$")
	return reg.MatchString(mobileNum)
}
