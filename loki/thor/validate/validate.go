// Package validate contains the support for validating models.
package validate

import (
	"loki/thor"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var validate *validator.Validate

var translator ut.Translator

func init() {

	validate = validator.New()

	translator, _ = ut.New(en.New(), en.New()).GetTranslator("en")

	en_translations.RegisterDefaultTranslations(validate, translator)

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func Check(val any) error {
	if err := validate.Struct(val); err != nil {

		verrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}

		var fields thor.FieldErrors
		for _, verror := range verrors {
			field := thor.FieldError{
				Field: verror.Field(),
				Err:   verror.Translate(translator),
			}
			fields = append(fields, field)
		}

		return fields
	}

	return nil
}
