package handler

import (
	"context"

	"github.com/ankorstore/yokai/config"
	"github.com/ankorstore/yokai/log"
	"github.com/ndk/test123/spec"
)

func optString(s string) *string {
	return &s
}

type apiExample struct {
	config *config.Config
}

func (a *apiExample) Translate(ctx context.Context, request spec.TranslateRequestObject) (spec.TranslateResponseObject, error) {
	log.CtxLogger(ctx).Info().Msg("kill all humans")
	return spec.Translate200JSONResponse{
		TargetText: optString("Bonjour le monde!"),
	}, nil
}

func NewAPI(config *config.Config) *apiExample {
	return &apiExample{config: config}
}
