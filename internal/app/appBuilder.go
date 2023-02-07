package app

import (
	"go.uber.org/zap"
)

type AppBuilderInt interface {
	Build() Apllication
	Logger(val *zap.SugaredLogger) AppBuilderInt
	Mode(val *Mode) AppBuilderInt
	DockerService(val DockerService) AppBuilderInt
	ErrCh(val chan error) AppBuilderInt
}

type dockerAppBuilder struct {
	log           *zap.SugaredLogger
	mode          *Mode
	dockerService DockerService
	errCh         chan error
}

func NewDockerAppBuilder() AppBuilderInt {
	errorCh := make(chan error, 1)
	return &dockerAppBuilder{errCh: errorCh}
}

func (b *dockerAppBuilder) Logger(val *zap.SugaredLogger) AppBuilderInt {
	b.log = val
	return b
}
func (b *dockerAppBuilder) Mode(val *Mode) AppBuilderInt {
	b.mode = val
	return b
}
func (b *dockerAppBuilder) DockerService(val DockerService) AppBuilderInt {
	b.dockerService = val
	return b
}

func (b *dockerAppBuilder) ErrCh(val chan error) AppBuilderInt {
	b.errCh = val
	return b
}

func (b *dockerAppBuilder) Build() Apllication {

	return &dockerStatsApp{log: b.log, mode: b.mode, errCh: b.errCh, dockerService: b.dockerService}
}
