package carrier

import (
	"log"
	"strings"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const KOC_PREFIX = "github.com/eppppi/koc/"

func isTraceKey(key string) bool {
	return strings.HasPrefix(key, KOC_PREFIX)
}

// OpenTelemetry carrier using annotation of kubernetes object
type K8sObjAntCarrier struct {
	metav1.Object
}

func NewK8sAntCarrierFromInterface(objInterface interface{}) (*K8sObjAntCarrier, error) {
	obj, err := meta.Accessor(objInterface)
	if err != nil {
		return nil, err
	}
	return NewK8sAntCarrierFromObj(obj)
}

func NewK8sAntCarrierFromObj(obj metav1.Object) (*K8sObjAntCarrier, error) {
	return &K8sObjAntCarrier{obj}, nil
}

func (objCarrier *K8sObjAntCarrier) Get(key string) string {
	annotations := objCarrier.GetAnnotations()
	if val, ok := annotations[key]; ok {
		return val
	} else {
		log.Printf("warning: key %s is not found", key)
		return ""
	}
}

func (objCarrier *K8sObjAntCarrier) Set(key string, value string) {
	if !isTraceKey(key) {
		log.Printf("warning: key %s is invalid for trace key", key)
	}
	objCarrier.SetAnnotations(map[string]string{key: value})
}

func (objCarrier *K8sObjAntCarrier) Keys() []string {
	annotations := objCarrier.GetAnnotations()
	keys := make([]string, 0, len(annotations))
	for k := range annotations {
		if isTraceKey(k) {
			keys = append(keys, k)
		}
	}
	return keys
}
