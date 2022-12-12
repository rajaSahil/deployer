package deployer

import (
	"fmt"
	"github.com/docker/go-connections/nat"
)

// const PULSAR

var (
	CONSTANTS = map[string]map[string]interface{}{
		"pulsar": {},
		"mongo":  {},
		"mysql":  {},
	}
	PULSAR = initializeConstant("pulsar", []nat.Port{nat.Port(6650), nat.Port(8080)})
	MONGO  = map[string]string{}
	MYSQL  = map[string]string{}
)

func GetConstantsFromJson() {

}

func initializeConstant(imageName string, ports []nat.Port) map[string]interface{} {
	for k, v := range CONSTANTS {
		fmt.Println(k, v)
	}
	return map[string]interface{}{
		"imageName": imageName,
		"port":      ports,
	}
}

func GetConstants(name string) interface{} {
	return CONSTANTS[name]
}
