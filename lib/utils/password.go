package utils

// 加密密码
func PasswordEncode(password string) string {
	cb := EncodeMD5(EncodeMD5(password + "VK"))
	return cb
}

// 校验密码是否正确
func PasswordValidate(password, password_db string) bool {
	v_pwd := PasswordEncode(password)
	if v_pwd == password_db {
		return true
	} else {
		return false
	}
}
