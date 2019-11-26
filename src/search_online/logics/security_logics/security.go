package security_logics

import (
	"errors"
	"github.com/spf13/viper"
	"go_search_online/src/search_online/databases/redis_module"
	"go_search_online/src/search_online/logics/config_logics"
	"go_search_online/src/search_online/logics/logging_logics"
	"go_search_online/src/search_online/logics/parameter_logics"
	"go_search_online/src/search_online/utils"
)

var (
	config *viper.Viper
)

func init() {
	config = config_logics.GlobalConfig
}

func SecurityChecks(params *parameter_logics.Params) (bool, error) {

}

func appCheck(appname, appid string) (bool, error) {
	var errorString string
	app_check_key := config.GetString("service.security.app_info_key" + ":" + appname + ":" + appid)
	res := redis_module.GlobalRedis.Exists(app_check_key)
	if res.Val() == 0 {
		logging_logics.ServiceLogger.Error("Failed to get the app info, may not config.")
		errorString = "Failed to get the app info， may be an invalid config."
	}
	check := utils.Int2Bool(int(res.Val()))
	return check, errors.New(errorString)
}

func signCheck(params *parameter_logics.Params) (bool, error) {

}