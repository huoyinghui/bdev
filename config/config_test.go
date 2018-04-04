package config

import (
	"github.com/astaxie/beego/logs"
	"gopkg.in/yaml.v2"
	iot "io/ioutil"
	"testing"
)

func TestNewConf(t *testing.T) {
	conf := AppConfInfo{
		AppName:         "bdev",
		HttpPort:        4008,
		GrpcListen:      ":5008",
		RunMode:         "prod",
		AutoRender:      false,
		CopyRequestBody: true,
		EnableDocs:      true,
		OrmDebug:        true,
		PgDataSource:    "user=postgres password=postgres dbname=test host=127.0.0.1 port=5432 sslmode=disable",
		LogLevel:        logs.LevelDebug,
		JwtSalt:         "testsalt",
		AccessKeyID:		"",
		SecretAccessKey:	"",
		Region:				"",
	}

	data, err := yaml.Marshal(conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("data:%v", string(data))
	if err := iot.WriteFile("../conf/app_template.yml", data, 0644); err != nil {
		t.Fatal(err)
	}

}

func TestLoadConf(t *testing.T) {
	data, err := LoadConf("../conf/app.yml")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("data:%v", data)
}

func TestLoadConfFromData(t *testing.T) {
	data, err := LoadConfFromData(TestData, "yml")
	if err != nil {
		t.Logf("data:%v", data)
		t.Fatal(err)
	}
	t.Logf("data:%v", data)
}
