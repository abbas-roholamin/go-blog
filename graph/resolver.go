package graph

import (
	"context"
	"net/http"

	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	DB *gorm.DB
}

var customContextKey string = "CUSTOM_CONTEXT"

func CreateContext(args *Resolver, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customContext := &Resolver{
			DB: args.DB,
		}
		requestWithCtx := r.WithContext(context.WithValue(r.Context(), customContextKey, customContext))
		next.ServeHTTP(w, requestWithCtx)
	})
}

func GetContext(ctx context.Context) *Resolver {
	customContext, ok := ctx.Value(customContextKey).(*Resolver)
	if !ok {
		return nil
	}
	return customContext
}
