package forms

import (
	"github.com/go-playground/validator/v10"
)

type TagVaridator struct {
	Tags string `validate:"excludesall=/"`
}

func (form *TagVaridator) TagValidate() (ok bool, result map[string]string) {
	result = make(map[string]string)
	validate := validator.New()
	err := validate.Struct(*form)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		if len(errors) != 0 {
			for i := range errors {
				switch errors[i].StructField() {
				case "Tags":
					switch errors[i].Tag() {
					case "excludesall":
						result["Name"] = "*タグに/を含むことはできません。"
					}
				}
			}
		}
		return false, result
	}
	return true, result
}
