package validator

import (
	"errors"
	"log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validator  *validator.Validate
	Translator ut.Translator
}

func NewValidator() *Validator {
	en := en.New()
	uni := ut.New(en, en)
	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}

	validate := validator.New()

	return &Validator{
		Validator:  validate,
		Translator: trans,
	}
}

func (v *Validator) Validate(i interface{}) error {
	err := v.Validator.Struct(i)

	if err != nil {
		object, _ := err.(validator.ValidationErrors)

		for _, key := range object {
			return errors.New(key.Translate(v.Translator))
		}
	}

	return nil
}
