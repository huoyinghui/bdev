package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	iot "io/ioutil"
	"strings"
)

type AppConfInfo struct {
	AppName         string `yaml:"app_name" json:"app_name"`
	HttpPort        int    `yaml:"http_port" json:"http_port"`
	RunMode         string `yaml:"run_mode" json:"run_mode"`
	AutoRender      bool   `yaml:"auto_render" json:"auto_render"`
	CopyRequestBody bool   `yaml:"copy_request_body" json:"copy_request_body"`
	GrpcListen      string `yaml:"grpc_listen" json:"grpc_listen"`
	PgDataSource    string `yaml:"pg_data_source" json:"pg_data_source"`
	OrmDebug        bool   `yaml:"orm_debug" json:"orm_debug"`
	EnableDocs      bool   `yaml:"enable_docs" json:"enable_docs"`
	LogLevel        int    `yaml:"log_level" json:"log_level"`
	JwtSalt         string `yaml:"jwt_salt" json:"jwt_salt"`
	AccessKeyID     string `yaml:"access_key_id", json:"access_key_id"`
	SecretAccessKey string `yaml:"secret_access_key", json:"secret_access_key"`
	Region          string `yaml:"region" json:"region"`
}

var AppConf *AppConfInfo

func init() {
	Init("debug")
}

var TestData = []byte(`
app_name: bdev
http_port: 4008
run_mode: prod
auto_render: false
copy_request_body: true
grpc_listen: :5008
pg_data_source: user=postgres password=postgres dbname=test host=127.0.0.1 port=5432
  sslmode=disable
orm_debug: true
enable_docs: true
log_level: 7
jwt_salt: testsalt
access_key_id: ""
secret_access_key: ""
region: ""
`)

func Init(t string) {
	var err error

	if t == "debug" {
		AppConf, err = LoadConfFromData(TestData, "yml")
		if err != nil {
			panic(err)
		}
	}

	if AppConf == nil {
		AppConf, err = LoadConf("conf/app.yml")
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("config-err:%v\nAppConf:%v\n", err, AppConf)
}

func LoadConf(filepath string) (item *AppConfInfo, err error) {
	if filepath == "" {
		return nil, errors.New("filepath is empty, must use --config xxx.yml/json")
	}

	data, err := iot.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	item = &AppConfInfo{}
	if strings.HasSuffix(filepath, ".json") {
		err = json.Unmarshal(data, item)
	} else if strings.HasSuffix(filepath, ".yml") || strings.HasSuffix(filepath, ".yaml") {
		err = yaml.Unmarshal(data, item)
	} else {
		return nil, errors.New("you config file must be json/yml")
	}

	if err != nil {
		return nil, err
	}

	return item, nil
}

func LoadConfFromData(data []byte, t string) (item *AppConfInfo, err error) {
	// data 可以是json/yml格式的数据. 调用方需要指定t.
	item = &AppConfInfo{}
	if t == "json" {
		err = json.Unmarshal(data, item)
	} else if t == "yml" {
		err = yaml.Unmarshal(data, item)
	}
	if err != nil {
		return nil, err
	}

	return item, nil
}
