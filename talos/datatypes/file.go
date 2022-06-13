package datatypes

import (
	"github.com/talos-systems/talos/pkg/machinery/config/types/v1alpha1"
)

// Data copies data from terraform state types to talos types.
func (planFile File) Data() (interface{}, error) {
	return &v1alpha1.MachineFile{
		FileContent:     planFile.Content.Value,
		FilePermissions: v1alpha1.FileMode(planFile.Permissions.Value),
		FilePath:        planFile.Path.Value,
		FileOp:          planFile.Op.Value,
	}, nil
}

func (planFile File) DataFunc() [](func(*v1alpha1.Config) error) {
	return [](func(*v1alpha1.Config) error){
		func(cfg *v1alpha1.Config) error {
			file := &v1alpha1.MachineFile{
				FileContent:     planFile.Content.Value,
				FilePermissions: v1alpha1.FileMode(planFile.Permissions.Value),
				FilePath:        planFile.Path.Value,
				FileOp:          planFile.Op.Value,
			}
			cfg.MachineConfig.MachineFiles = append(cfg.MachineConfig.MachineFiles, file)
			return nil
		},
	}
}

type Files = []v1alpha1.MachineFile

type TalosFiles struct {
	*Files
}
