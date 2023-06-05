package app

import (
	"github.com/wispeeer/glast/internal/app/glast"
)

type service string

const (
	Cli service = "glast"
)

func (service service) Service() func() error {
	switch service {
	case Cli:
		return glast.Service
	}
	return nil
}
