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
	trans, _ := tm.utrans.FindTranslator(pure.AcceptedLanguages(r)...)

	requestCtx := r.Context()
	requestCtx = context.WithValue(requestCtx, translatorKey{}, trans)
	r.WithContext(requestCtx)

	ctx.ResetRequest(r)
	ctx.Next()
}

func (tm TranslatorMiddleware) Get(ctx context.Context) ut.Translator {
	trans, found := ctx.Value(translatorKey{}).(ut.Translator)
	if !found {
		return tm.utrans.GetFallback()
	}
	return trans
}
