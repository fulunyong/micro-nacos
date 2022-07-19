package nacos

import (
	"errors"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"net"
	"strconv"
	"strings"
)

//RegisterServe 注册服务
//port 当前运行服务端口
func RegisterServe(port int) {

	clientConfig := constant.ClientConfig{
		NamespaceId:         "", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "log",
		CacheDir:            "cache",
		LogLevel:            "debug",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "10.10.27.176",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}

	// 将服务注册到nacos
	client, _ := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	_, err := client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          MustGetIntranetIp(), //默认获取本机ip4地址
		Port:        uint64(port),        //默认获取项目运行地址
		ServiceName: "goMicroServer",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: "DEFAULT",       // 默认值DEFAULT
		GroupName:   "DEFAULT_GROUP", // 默认值DEFAULT_GROUP
	})
	if err != nil {
		fmt.Printf("注册服务异常%s\n", err)
		return
	}
}

// MustGetIntranetIp performs as GetIntranetIp, but it panics if any error occurs.
func MustGetIntranetIp() string {
	ip, err := GetIntranetIp()
	if err != nil {
		panic(err)
	}
	return ip
}

// GetIntranetIp retrieves and returns the first intranet ip of current machine.
func GetIntranetIp() (ip string, err error) {
	ips, err := GetIntranetIpArray()
	if err != nil {
		return "", err
	}
	if len(ips) == 0 {
		return "", nil
	}
	return ips[0], nil
}

// GetIntranetIpArray retrieves and returns the intranet ip list of current machine.
func GetIntranetIpArray() (ips []string, err error) {
	var (
		addresses  []net.Addr
		interFaces []net.Interface
	)
	interFaces, err = net.Interfaces()
	if err != nil {
		return ips, err
	}
	for _, interFace := range interFaces {
		if interFace.Flags&net.FlagUp == 0 {
			// interface down
			continue
		}
		if interFace.Flags&net.FlagLoopback != 0 {
			// loop back interface
			continue
		}
		// ignore warden bridge
		if strings.HasPrefix(interFace.Name, "w-") {
			continue
		}
		addresses, err = interFace.Addrs()
		if err != nil {
			err = errors.New("interFace.Addrs failed")
			return ips, err
		}
		for _, addr := range addresses {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				// not an ipv4 address
				continue
			}
			ipStr := ip.String()
			if IsIntranet(ipStr) {
				ips = append(ips, ipStr)
			}
		}
	}
	return ips, nil
}

// IsIntranet checks and returns whether given ip an intranet ip.
//
// Local: 127.0.0.1
// A: 10.0.0.0--10.255.255.255
// B: 172.16.0.0--172.31.255.255
// C: 192.168.0.0--192.168.255.255
func IsIntranet(ip string) bool {
	if ip == "127.0.0.1" {
		return true
	}
	array := strings.Split(ip, ".")
	if len(array) != 4 {
		return false
	}
	// A
	if array[0] == "10" || (array[0] == "192" && array[1] == "168") {
		return true
	}
	// C
	if array[0] == "192" && array[1] == "168" {
		return true
	}
	// B
	if array[0] == "172" {
		second, err := strconv.ParseInt(array[1], 10, 64)
		if err != nil {
			return false
		}
		if second >= 16 && second <= 31 {
			return true
		}
	}
	return false
}
