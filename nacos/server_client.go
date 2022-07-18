package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func RegisterServiceInstance(client naming_client.INamingClient, param vo.RegisterInstanceParam) {
	success, _ := client.RegisterInstance(param)
	fmt.Printf("RegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
}

func DeRegisterServiceInstance(client naming_client.INamingClient, param vo.DeregisterInstanceParam) {
	success, _ := client.DeregisterInstance(param)
	fmt.Printf("DeRegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
}

func GetService(client naming_client.INamingClient, param vo.GetServiceParam) model.Service {
	service, _ := client.GetService(param)
	fmt.Printf("GetService,param:%+v, result:%+v \n\n", param, service)
	return service
}

func SelectAllInstances(client naming_client.INamingClient, param vo.SelectAllInstancesParam) []model.Instance {
	instances, _ := client.SelectAllInstances(param)
	fmt.Printf("SelectAllInstance,param:%+v, result:%+v \n\n", param, instances)
	return instances
}

func SelectInstances(client naming_client.INamingClient, param vo.SelectInstancesParam) []model.Instance {
	instances, _ := client.SelectInstances(param)
	fmt.Printf("SelectInstances,param:%+v, result:%+v \n\n", param, instances)
	return instances
}

func SelectOneHealthyInstance(client naming_client.INamingClient, param vo.SelectOneHealthInstanceParam) *model.Instance {
	instances, err := client.SelectOneHealthyInstance(param)
	if err != nil {
		fmt.Printf("SelectOneHealthyInstance error:%s\n", err)
	}
	fmt.Printf("SelectInstances,param:%+v, result:%+v \n\n", param, instances)
	return instances
}

func Subscribe(client naming_client.INamingClient, param *vo.SubscribeParam) {
	err := client.Subscribe(param)
	if err != nil {
		fmt.Printf("Subscribe error:%s\n", err)
	}
}

func UnSubscribe(client naming_client.INamingClient, param *vo.SubscribeParam) {
	err := client.Unsubscribe(param)
	if err != nil {
		fmt.Printf("UnSubscribe error:%s\n", err)
	}
}

func GetAllService(client naming_client.INamingClient, param vo.GetAllServiceInfoParam) model.ServiceList {
	serviceList, err := client.GetAllServicesInfo(param)
	if err != nil {
		fmt.Printf("GetAllService error:%s\n", err)
		return model.ServiceList{}
	}
	fmt.Printf("GetAllService,param:%+v, result:%+v \n\n", param, serviceList)
	return serviceList
}
