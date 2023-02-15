package app

import (
	"github.com/pablogolobaro/dockertool-legend/internal/config"
	"go.uber.org/zap"
)

type AppBuilderInt interface {
	Build() Apllication
	Logger(val *zap.SugaredLogger) AppBuilderInt
	Mode(val *config.Mode) AppBuilderInt
	ContainerStreamer(val ContainerStreamer) AppBuilderInt
	ErrCh(val chan error) AppBuilderInt
}

type dockerAppBuilder struct {
	log      *zap.SugaredLogger
	mode     *config.Mode
	streamer ContainerStreamer
	errCh    chan error
}

func NewDockerAppBuilder() AppBuilderInt {
	errorCh := make(chan error, 1)
	return &dockerAppBuilder{errCh: errorCh}
}

func (b *dockerAppBuilder) Logger(val *zap.SugaredLogger) AppBuilderInt {
	b.log = val
	return b
}
func (b *dockerAppBuilder) Mode(val *config.Mode) AppBuilderInt {
	b.mode = val
	return b
}

func (b *dockerAppBuilder) ContainerStreamer(val ContainerStreamer) AppBuilderInt {
	b.streamer = val
	return b
}

func (b *dockerAppBuilder) ErrCh(val chan error) AppBuilderInt {
	b.errCh = val
	return b
}

func (b *dockerAppBuilder) Build() Apllication {
	return &statsApp{log: b.log, mode: b.mode, errCh: b.errCh, containerStreamer: b.streamer}
}
