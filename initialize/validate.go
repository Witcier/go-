package initialize

import (
	"fmt"
	"witcier/go-api/global"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
)

var trans ut.Translator

func ValidateTrans(locale string) ut.Translator {
	validate := binding.Validator.Engine().(*validator.Validate)
	zhT := zh.New()
	enT := en.New()

	uni := ut.New(zhT, zhT, enT)
	if trans, ok := uni.GetTranslator(locale); !ok {
		fmt.Errorf("uni.GetTranslator(%s) fail", locale)

		return nil
	} else {
		enTrans.RegisterDefaultTranslations(validate, trans)

		validate.RegisterTranslation("required", trans, func(ut1 ut.Translator) error {
			return ut1.Add("required", "{0} must have a value!", true)
		}, func(ut2 ut.Translator, fe validator.FieldError) string {
			t, _ := ut2.T("required", fe.Field())

			return t
		})

		global.Log.Info("initialize validate trans success")

		return trans
	}
}
