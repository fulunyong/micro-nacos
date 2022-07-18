package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func Config() {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("console.nacos.io", 80),
	}
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("e525eafa-f7d7-4029-83d9-008937f9d468"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("log"),
		constant.WithCacheDir("cache"),
		constant.WithLogLevel("debug"),
	)

	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}
	//注册服务到nacos
	RegisterServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          "10.10.27.176",
		Port:        8848,
		ServiceName: "GoMicroServer",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	})

}
