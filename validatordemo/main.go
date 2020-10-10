package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"unicode/utf8"
)

// User contains user information
type UserInfo struct {
	NameCheck string `validate:"checkName"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       int8   `validate:"gte=0,lte=100"`
	Email     string `validate:"required,email"`
}

// 自定义验证函数
func checkName(fl validator.FieldLevel) bool {
	count := utf8.RuneCountInString(fl.Field().String())
	fmt.Printf("length: %v \n", count)
	if count > 5 {
		return false
	}
	return true
}

func main() {
	zh_ch := zh.New()
	uni := ut.New(zh_ch)
	trans, _ := uni.GetTranslator("zh")
	//验证器
	validate := validator.New()
	//验证器注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)
	//注册自定义函数，与struct tag关联起来
	validate.RegisterValidation("checkName", checkName)
	//validate := validator.New()
	user := &UserInfo{
		NameCheck: "中国人中国人中国人中国人",
		FirstName: "Badger",
		LastName:  "Smith",
		Age:       -1,
		Email:     "",
	}
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			srrStr := err.Translate(trans)
			//fmt.Println(err)
			//翻译错误信息
			fmt.Println(srrStr)
		}
		return
	}
	fmt.Println("success")
}
