package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	instance = validator.New()
	uni      *ut.UniversalTranslator
	trans    ut.Translator
)

func init() {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	instance.SetTagName("v")
	en_translations.RegisterDefaultTranslations(instance, trans)
}

func Validate(pointer any) error {
	return instance.Struct(pointer)
}

func Translate(err error) string {
	switch err.(type) {
	case validator.ValidationErrors:
		return err.(validator.ValidationErrors)[0].Translate(trans)
	default:
		return err.Error()
	}
}
