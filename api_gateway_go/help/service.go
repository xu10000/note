package help

import (
	"api_gateway/syslog"
	"errors"
	"math/rand"
	"strconv"

	"github.com/micro/go-micro/registry"
)

var (
	services_map = map[string][]*registry.Node{}
)

func Get_service_url(name string) (string, error) {
	services, exists := services_map[name]

	if exists != true {
		account_service, err := Get_registry().GetService(name)

		if err != nil {
			syslog.Debug.Printf("获取登录服务出错:", err)
			return "", err
		}

		length := len(account_service)

		if length == 0 {
			return "", errors.New("未获取到登录服务")
		}
		services_map[name] = account_service[0].Nodes
		services = services_map[name]
	}

	index := rand.Intn(len(services))
	service_node := services[index]
	url := service_node.Address + ":" + strconv.Itoa(service_node.Port)

	return url, nil

}
