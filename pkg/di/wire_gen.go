// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/int128/transerr/pkg/adaptors/inspector"
	"github.com/int128/transerr/pkg/cmd"
	"github.com/int128/transerr/pkg/usecases/dump"
	"github.com/int128/transerr/pkg/usecases/transform"
)

// Injectors from di.go:

func NewCmd() cmd.Interface {
	inspectorInspector := &inspector.Inspector{}
	useCase := transform.UseCase{
		Inspector: inspectorInspector,
	}
	cmdTransform := cmd.Transform{
		UseCase: useCase,
	}
	dumpUseCase := dump.UseCase{
		Inspector: inspectorInspector,
	}
	cmdDump := cmd.Dump{
		UseCase: dumpUseCase,
	}
	cmdCmd := &cmd.Cmd{
		Transform: cmdTransform,
		Dump:      cmdDump,
	}
	return cmdCmd
}