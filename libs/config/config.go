package config

import (
	"encoding/json"
	"gostart/libs/system"
	"io/ioutil"
	"sync"
)

var (
	env         string
	redisConfig map[string]interface{}
	logConfig   map[string]interface{}
	mysqlConfig map[string]interface{}
	mongoConfig map[string]interface{}
	lock        sync.Mutex
)

func Init(_env string) {
	env = _env
	load()
}

func load() {
	var redisConfigPath = getConfigPath("redis.json")
	var mysqlConfigPath = getConfigPath("mysql.json")
	var mongoConfigPath = getConfigPath("mongo.json")
	var logConfigPath = getConfigPath("log.json")

	lock.Lock()

	loadConfig(&redisConfig, redisConfigPath)
	loadConfig(&mysqlConfig, mysqlConfigPath)
	loadConfig(&mongoConfig, mongoConfigPath)
	loadConfig(&logConfig, logConfigPath)
	lock.Unlock()
}

func getConfigPath(configFile string) string {
	return system.Root + "/config/" + env + "/" + configFile
}

func loadConfig(data *map[string]interface{}, configPath string) {
	fileData, _ := ioutil.ReadFile(configPath)
	json.Unmarshal(fileData, data)
}

func GetRedisConfig() map[string]interface{} {
	return redisConfig
}

func GetMysqlConfig() map[string]interface{} {
	return mysqlConfig
}

func GetLogConfig() map[string]interface{} {
	return logConfig
}

func GetMongoConfig() map[string]interface{} {
	return mongoConfig
}
