package test

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type Groups struct {
	User  string `mapstructure:"user"`
	Role  string `mapstructure:"Role"`
	Maner string `mapstructure:"maner"`
}

type Config struct {
	TestEnv string `mapstructure:"test"`
}

func Test_load(t *testing.T) {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("packer-server")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	assert.NoError(t, err)

	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		//if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
		//	viper.Set(k, (strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		//}

		t.Logf("key:%s value: %s", k, value)
	}

}
func Test_env(t *testing.T) {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("packer-server")
	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	assert.NoError(t, err)
	t.Logf("%s", viper.Get("webhook.url"))
}
func Test_set_env(t *testing.T) {
	vp := viper.New()
	vp.SetConfigName("test")
	vp.SetConfigType("yaml")
	vp.AddConfigPath(".")
	// 设置为true 可自动获取相应的环境变量
	vp.AutomaticEnv()
	// 设置与本项目相关的环境变量的前缀
	vp.SetEnvPrefix("app")
	vp.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := vp.ReadInConfig()
	assert.NoError(t, err)

	config := &Config{}
	err = vp.Unmarshal(&config)
	assert.NoError(t, err)

	t.Logf("config: %+v", config)
	groups := &Groups{}
	err = vp.UnmarshalKey("groups", &groups)
	t.Logf("groups: %+v", groups)
	assert.NoError(t, err)
}
