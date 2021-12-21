package forms

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidateUser struct {
	Name     string `validate:"required","unique"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type LoginValidateUser struct {
	Name     string `validate:"required"`
	Password string `validate:"required"`
}

// バリデーションチェックを行って、結果とNGの場合にエラーメッセージ群を返す
func (form *ValidateUser) Validate() (ok bool, result map[string]string) {
	result = make(map[string]string)
	// 構造体のデータをタグで定義した検証方法でチェック
	err := validator.New().Struct(*form)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		fmt.Println(errors)
		if len(errors) != 0 {
			for i := range errors {
				// フィールドごとに、検証
				switch errors[i].StructField() {
				case "Name":
					switch errors[i].Tag() {
					case "required":
						result["Name"] = "*お名前は必須入力です。"
					case "unique":
						result["Name"] = "*すでに登録されているお名前です。"
					}
				case "Email":
					switch errors[i].Tag() {
					case "required":
						result["Email"] = "*メールアドレスは必須入力です。"
					case "email":
						result["Email"] = "正しいメールアドレス形式で入力してください"
					}
				case "Password":
					switch errors[i].Tag() {
					case "required":
						result["Password"] = "*パスワードは必須入力です。"
					}
				}
			}
		}
		return false, result
	}
	return true, result
}

func (form *LoginValidateUser) LoginValidate() (ok bool, result map[string]string) {
	result = make(map[string]string)
	// 構造体のデータをタグで定義した検証方法でチェック
	err := validator.New().Struct(*form)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		fmt.Println(errors)
		if len(errors) != 0 {
			for i := range errors {
				// フィールドごとに、検証
				switch errors[i].StructField() {
				case "Name":
					switch errors[i].Tag() {
					case "required":
						result["Name"] = "*お名前は必須入力です。"
					case "unique":
						result["Name"] = "*すでに登録されているお名前です。"
					}
				case "Email":
					switch errors[i].Tag() {
					case "required":
						result["Email"] = "*メールアドレスは必須入力です。"
					case "email":
						result["Email"] = "正しいメールアドレス形式で入力してください"
					}
				case "Password":
					switch errors[i].Tag() {
					case "required":
						result["Password"] = "*パスワードは必須入力です。"
					}
				}
			}
		}
		return false, result
	}
	return true, result
}
