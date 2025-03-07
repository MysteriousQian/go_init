package services

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"go_server/internal/db/models"
	"regexp"
)

// 生成字符串的 MD5 哈希值
func GenerateMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// 生成一个随机盐
func GenerateSalt(size int) (string, error) {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// 用户名、密码登录
func Login(userame, password string) (user models.User, err error) {
	user, err = models.User{
		Name: userame,
	}.FindByName()
	if err != nil || user.Id == 0 {
		err = fmt.Errorf("failed to obtain user information")
		return
	}
	if user.Password != GenerateMD5(password+user.Salt) {
		err = fmt.Errorf("password error")
		return
	}
	return
}

// 获取客户列表
func GetUserList(page, size int, name string) (users []models.UserInfo, total int64, err error) {
	users, total, err = models.User{}.SelectUserList(page, size, name)
	if err != nil {
		err = fmt.Errorf("failed to obtain user information")
	}
	return
}

// 检查用户名和密码
func CheckNameAndPassword(customerName, password string) error {
	if len(customerName) < 4 || len(customerName) > 20 {
		return fmt.Errorf("username length must be between 4 and 20")
	}
	if len(password) < 8 || len(password) > 16 {
		return fmt.Errorf("password length must be between 8 and 20")
	}
	// 验证密码格式
	re := regexp.MustCompile(`^[A-Za-z0-9!@#$%^&\-\*()_+\]\[\}\{|;:,.<>?]+$`)
	if re.MatchString(customerName) {
		if re.MatchString(password) {
			return nil
		} else {
			return fmt.Errorf("password can only be composed of english, numbers, or special symbols")
		}
	} else {
		return fmt.Errorf("username can only be composed of english, numbers, or special symbols")
	}
}
