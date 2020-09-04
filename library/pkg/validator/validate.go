package validator

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/VulpesFerrilata/boardgame-server/library/pkg/errors"
	"github.com/VulpesFerrilata/boardgame-server/library/pkg/middleware/translator"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

type Validate interface {
	Struct(ctx context.Context, s interface{}) error
	StructExcept(ctx context.Context, s interface{}, fields ...string) error
	StructPartial(ctx context.Context, s interface{}, fields ...string) error
	Var(ctx context.Context, field interface{}, tag string) error
}

func NewValidate(utrans *ut.UniversalTranslator, translatorMiddleware *translator.TranslatorMiddleware) (Validate, error) {
	v := validator.New()
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("name"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	en := en.New()
	trans, found := utrans.GetTranslator(en.Locale())
	if !found {
		return nil, fmt.Errorf("translator not found: %v", en.Locale())
	}
	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		return nil, err
	}

	return &validate{
		Validate:             v,
		translatorMiddleware: translatorMiddleware,
	}, nil
}

type validate struct {
	*validator.Validate
	translatorMiddleware *translator.TranslatorMiddleware
}

func (v validate) Struct(ctx context.Context, s interface{}) error {
	return v.parseError(ctx, v.Validate.Struct(s))
}

func (v validate) StructExcept(ctx context.Context, s interface{}, fields ...string) error {
	return v.parseError(ctx, v.Validate.StructExcept(s, fields...))
}

func (v validate) StructPartial(ctx context.Context, s interface{}, fields ...string) error {
	return v.parseError(ctx, v.Validate.StructPartial(s, fields...))
}

func (v validate) Var(ctx context.Context, field interface{}, tag string) error {
	return v.parseError(ctx, v.Validate.Var(field, tag))
}

func (v validate) parseError(ctx context.Context, err error) error {
	trans := v.translatorMiddleware.Get(ctx)

	if err != nil {
		validationErrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}

		errs := make([]string, 0)
		for _, validationErr := range validationErrs {
			errs = append(errs, validationErr.Translate(trans))
		}
		return errors.NewValidationError(errs...)
	}

	return nil
}
