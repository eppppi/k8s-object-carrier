package main

import (
	"fmt"

	"github.com/eppppi/k8s-object-carrier/carrier"
)

func main() {
	koc := carrier.K8sObjAntCarrier{}
	koc.Set(carrier.KOC_PREFIX+"test-key", "test-value")
	fmt.Println(koc.Get(carrier.KOC_PREFIX + "test-key"))
}
