package elasticsearch_module

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
	"go_search_online/src/search_online/logics/config_logics"
	"go_search_online/src/search_online/logics/logging_logics"
)

var (
	GlobalESClient *elastic.Client
	config *viper.Viper
)

func init() {
	fmt.Printf("Loading elasticsearch module...\n")
	config = config_logics.GlobalConfig
	GlobalESClient = initClient()
}

func initClient() *elastic.Client {
	GlobalESClient, err := elastic.NewClient(
		elastic.SetURL(config.GetString("database.elasticsearch.url")),
		elastic.SetBasicAuth(config.GetString("database.elasticsearch.username"),
			config.GetString("database.elasticsearch.password")),
		)
	if err != nil {
		logging_logics.ServiceLogger.Fatal("Failed to connect to elasticsearch cluster " + err.Error())
	}
	return GlobalESClient
}