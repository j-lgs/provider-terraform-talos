package datatypes

import (
	"github.com/talos-systems/talos/pkg/machinery/config/types/v1alpha1"
)

// Data copies data from terraform state types to talos types.
func (planProxy ProxyConfig) Data() (any, error) {
	proxy := &v1alpha1.ProxyConfig{}
	if !planProxy.Image.Null {
		proxy.ContainerImage = planProxy.Image.Value
	}
	if !planProxy.Disabled.Null {
		proxy.Disabled = planProxy.Disabled.Value
	}
	if !planProxy.Mode.Null {
		proxy.ModeConfig = planProxy.Mode.Value
	}
	proxy.ExtraArgsConfig = map[string]string{}
	for arg, value := range planProxy.ExtraArgs {
		proxy.ExtraArgsConfig[arg] = value.Value
	}

	return proxy, nil
}

func (planProxy ProxyConfig) DataFunc() [](func(*v1alpha1.Config) error) {
	funs := [](func(*v1alpha1.Config) error){
		func(cfg *v1alpha1.Config) error {
			proxy, err := planProxy.Data()
			if err != nil {
				return err
			}
			cfg.ClusterConfig.ProxyConfig = proxy.(*v1alpha1.ProxyConfig)

			return nil
		},
	}
	return funs
}

type TalosProxyConfig struct {
	*v1alpha1.ProxyConfig
}
