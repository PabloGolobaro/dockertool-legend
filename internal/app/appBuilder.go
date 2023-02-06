package app

import (
	"github.com/pablogolobaro/dockertool-legend/internal/config"
	"go.uber.org/zap"
)

type AppBuilderInt interface {
	Build() Apllication
}

type appBuilder struct {
	log  *zap.SugaredLogger
	mode Mode

	errorCh chan error
}

func NewAppBuilder() (AppBuilderInt, error) {
	development, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	sugar := development.Sugar()
	sugar.Debug("Builder get flags")
	modeFlags, err := config.GetModeFlags()
	if err != nil {
		return nil, err
	}

	sugar.Infow("Flags", modeFlags)

	mode := NewMode(modeFlags.Stream, modeFlags.Duration)

	errorCh := make(chan error, 1)

	sugar.Debug("Builder Created")
	return &appBuilder{log: sugar, mode: mode, errorCh: errorCh}, nil
}

func (b appBuilder) Logger(val *zap.SugaredLogger) {
	b.log = val
}

func (b appBuilder) Mode(val Mode) {
	b.mode = val
}

func (b appBuilder) Build() Apllication {

	return &dockerStatsApp{log: b.log, mode: b.mode, errCh: b.errorCh}
}
