package init_viper

import "github.com/spf13/viper"

func InitViper() {
	v := viper.GetViper()
	v.MustBindEnv("devRoot", "ULLD_DEV_ROOT")
}
