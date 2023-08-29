package carrier

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OpenTelemetry carrier using annotation of kubernetes object
// type K8sObjAntCarrier metav1.ObjectMeta.Annotations
type K8sObjAntCarrier metav1.ObjectMeta

func (objMeta *K8sObjAntCarrier) Get(key string) string {
	anotations := (*metav1.ObjectMeta)(objMeta).GetAnnotations()
	fmt.Println(anotations)
	return "true" // TODO
}

func (objMeta *K8sObjAntCarrier) Set(key string, value string) {
	(*metav1.ObjectMeta)(objMeta).SetAnnotations(map[string]string{key: value})
}

// func (objMeta K8sObjAntCarrier) Keys() []string {
// 	keys := make([]string, 0, len(hc))
// 	for k := range hc {
// 		if IS_RELATED_TO_TRACE(k) {
// 			keys = append(keys, k)
// 		}
// 	}
// 	return keys
// }
