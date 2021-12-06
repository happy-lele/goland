package main

import "fmt"

// 类型重命名
type ServiceType string

const (
	ServiceTypeClusterIP    ServiceType = "ClusterIP"
	ServiceTypeNodePort     ServiceType = "NodePort"
	ServiceTypeLoadBalancer ServiceType = "LoadBalancer"
	ServiceTypeExternalName ServiceType = "ExternalName"
)

func main()  {
	fmt.Println(ServiceTypeClusterIP)
	fmt.Println(ServiceTypeNodePort)
	fmt.Println(ServiceTypeLoadBalancer)
	fmt.Println(ServiceTypeExternalName)
}

