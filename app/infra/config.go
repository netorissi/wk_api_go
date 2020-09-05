package infra

import (
	"fmt"
	"os"
	"sync"

	"github.com/netorissi/wk_api_go/entities"
)

var configMutex = &sync.Mutex{}
var Configurations = &entities.Config{}

func LoadGlobalConfig() error {
	configMutex.Lock()
	defer configMutex.Unlock()

	// var config *entities.Config
	envWK := os.Getenv("WK")

	environment := "LOCAL"
	Configurations.Environment = "local"
	Configurations.Urls.Memcache = ""
	Configurations.Urls.Nats = ""
	Configurations.Urls.MySQL = "root:root@tcp(127.0.0.1:3308)/workon?charset=utf8mb4&parseTime=True&loc=Local"
	Configurations.Urls.NoSQL = ""

	switch envWK {
	case "prod":
		environment = "PRODUÇÃO"
		Configurations.Environment = envWK
		Configurations.Urls.Memcache = ""
		Configurations.Urls.Nats = ""
		Configurations.Urls.MySQL = ""
		Configurations.Urls.NoSQL = ""
	case "preprod":
		environment = "PRÉ-PRODUÇÃO"
		Configurations.Environment = envWK
		Configurations.Urls.Memcache = ""
		Configurations.Urls.Nats = ""
		Configurations.Urls.MySQL = ""
		Configurations.Urls.NoSQL = ""
	case "devtest":
		environment = "TESTE"
		Configurations.Environment = envWK
		Configurations.Urls.Memcache = ""
		Configurations.Urls.Nats = ""
		Configurations.Urls.MySQL = ""
		Configurations.Urls.NoSQL = ""
	}

	fmt.Println("Iniciando API no ambiente de " + environment)

	// config, err = getConfig(envWK)

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// 	return err
	// }

	return nil
}

// func getConfigConsul(env string) (*entities.Config, error) {

// 	url := os.Getenv("WK_URL")

// 	if len(url) == 0 {
// 		log.Fatal("Necessário informar a URL da Workon, exemplo: \"export WK_URL=\"http://127.1.1.1\"\"")
// 	}

// 	configClient := &api.Config{Address: url, Token: "f6fe14bf-2b2b-b58e-60b0-4d4cc205c105"}

// 	if env == "preprod" {
// 		configClient = &api.Config{Address: url, Token: "fd2df2fc-1097-0cd8-686f-211fad134352"}
// 	}

// 	if env == "prod" {
// 		configClient = &api.Config{Address: url, Token: "239e311c-bc07-da0e-425b-415450f7cdb1"}
// 	}

// 	client, err := api.NewClient(configClient)
// 	if err != nil {
// 		log.Fatal("Error to connect consul", err)
// 	}

// 	kv := client.KV()

// 	key := "microservices/TCPayment"

// 	pair, _, err := kv.Get(key, nil)
// 	if err != nil {
// 		log.Fatal("Error to get config "+key+" - ", err)
// 	}

// 	var config *model.Config

// 	err = json.Unmarshal(pair.Value, &config)
// 	if err != nil {
// 		log.Fatal("Error to parse "+key+" - ", err)
// 	}

// 	// Lookup the pair
// 	mysql_url, _, err := kv.Get("server/mysql", nil)
// 	if err != nil {
// 		log.Fatal("Error to get config server/mysql - ", err)
// 	}

// 	// Lookup the pair
// 	nats_url, _, err := kv.Get("server/nats", nil)
// 	if err != nil {
// 		log.Fatal("Error to get config server/nats - ", err)
// 	}

// 	// Lookup the pair
// 	memcached_url, _, err := kv.Get("server/memcached", nil)
// 	if err != nil {
// 		log.Fatal("Error to get config server/memcached - ", err)
// 	}

// 	// Lookup the pair
// 	tc_server, _, err := kv.Get("server/tcserver", nil)
// 	if err != nil {
// 		log.Fatal("Error to get config server/tcserver - ", err)
// 	}

// 	newds := strings.Replace(*config.SqlSettings.DataSource, "ip_mysql", string(mysql_url.Value), 1)
// 	config.SqlSettings.DataSource = &newds

// 	config.Urls.Nats = string(nats_url.Value)
// 	config.Urls.Memcache = string(memcached_url.Value)
// 	config.Urls.TCServer = string(tc_server.Value)

// 	if env == "hml" || env == "preprod" {
// 		config.Environment = "hml"
// 	}

// 	if env == "prod" {
// 		config.Environment = "prod"
// 	}

// 	return config, nil

// }
