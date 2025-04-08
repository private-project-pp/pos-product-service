package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type ConfigApp struct {
	Service  ServiceConfig  `mapstructure:"service" yaml:"service"`
	Db       DbConfig       `mapstructure:"db" yaml:"db"`
	Auth     AuthConfig     `mapstructure:"auth" yaml:"auth"`
	Internal InternalConfig `mapstructure:"internal" yaml:"internal"`
}

type ServiceConfig struct {
	Name    string `mapstructure:"name" yaml:"name"`
	Port    string `mapstructure:"port" yaml:"port"`
	LogFile string `mapstructure:"log_file" yaml:"log_file"`
}

type DbConfig struct {
	Host string `mapstructure:"host" yaml:"host"`
}

type AuthConfig struct {
	PassHashCode string `mapstructure:"pass_hash_code" yaml:"pass_hash_code"`
	JwtSecureKey string `mapstructure:"jwt_secure_key" yaml:"jwt_secure_key"`
}

type InternalConfig struct {
	MiddlewareRpcService RpcServiceConfig `mapstructure:"middleware_rpc_service" yaml:"middleware_rpc_service"`
}

type RpcServiceConfig struct {
	Address string `mapstructure:"address" yaml:"address"`
}

var (
	Service ServiceConfig
	DB      DbConfig
	Auth    AuthConfig
)

func SetupConfig() (out ConfigApp) {
	var err error
	viper.AddRemoteProvider("vault", "https://127.0.0.1:8443", "/secret/whatever/config.json")
	viper.SetConfigType("json")
	err = viper.ReadRemoteConfig()
	if err != nil {
		err = readLocalConfig()
	}
	if err != nil {
		panic(err)
	}

	viper.Unmarshal(&out)

	Service = out.Service
	DB = out.Db
	Auth = out.Auth
	return out
}

func readLocalConfig() (err error) {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(filepath.Join(path, "shared/config"))
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
