package translator

import (
	"context"

	"github.com/go-playground/pure/v5"
	ut "github.com/go-playground/universal-translator"
	"github.com/kataras/iris/v12"
)

type translatorKey struct{}

func NewTranslatorMiddleware(utrans *ut.UniversalTranslator) *TranslatorMiddleware {
	return &TranslatorMiddleware{
		utrans: utrans,
	}
}

type TranslatorMiddleware struct {
	utrans *ut.UniversalTranslator
}

func (tm TranslatorMiddleware) Serve(ctx iris.Context) {
	r := ctx.Request()

	requestCtx := r.Context()
	requestCtx = context.WithValue(requestCtx, translatorKey{}, pure.AcceptedLanguages(r))
	r.WithContext(requestCtx)

	ctx.ResetRequest(r)
	ctx.Next()
}

func (tm TranslatorMiddleware) Get(ctx context.Context) ut.Translator {
	acceptedLanguages, found := ctx.Value(translatorKey{}).([]string)
	if !found {
		return tm.utrans.GetFallback()
	}
	trans, _ := tm.utrans.FindTranslator(acceptedLanguages...)
	return trans
}
