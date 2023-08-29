package main

import (
	"fmt"

	"github.com/eppppi/k8s-object-carrier/carrier"
)

func main() {
	carrier := carrier.K8sObjAntCarrier{}
	carrier.Set("test-key", "test-value")
	fmt.Println(carrier.Get("test-key"))
}
