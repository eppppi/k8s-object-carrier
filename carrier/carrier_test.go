package carrier

import (
	"reflect"
	"testing"
)

func Test_isTraceKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty key",
			args: args{key: ""},
			want: false,
		},
		{
			name: "sample key",
			args: args{key: "test1"},
			want: false,
		},
		{
			name: "sample key with prefix",
			args: args{key: KOC_PREFIX + "test-key"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTraceKey(tt.args.key); got != tt.want {
				t.Errorf("isTraceKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestK8sObjAntCarrier_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		objMeta *K8sObjAntCarrier
		args    args
		want    string
	}{
		{
			name: "empty annotation",
			objMeta: &K8sObjAntCarrier{
				Annotations: map[string]string{},
			},
			args: args{key: ""},
			want: "",
		},
		{
			name: "get non-exist annotation",
			objMeta: &K8sObjAntCarrier{
				Annotations: map[string]string{},
			},
			args: args{key: "no-such-key"},
			want: "",
		},
		{
			name: "get proper annotation",
			objMeta: &K8sObjAntCarrier{
				Annotations: map[string]string{KOC_PREFIX + "test-key": "test-value"},
			},
			args: args{key: KOC_PREFIX + "test-key"},
			want: "test-value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.objMeta.Get(tt.args.key); got != tt.want {
				t.Errorf("K8sObjAntCarrier.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestK8sObjAntCarrier_Set(t *testing.T) {
// 	type args struct {
// 		key   string
// 		value string
// 	}
// 	tests := []struct {
// 		name    string
// 		objMeta *K8sObjAntCarrier
// 		args    args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.objMeta.Set(tt.args.key, tt.args.value)
// 		})
// 	}
// }

func TestK8sObjAntCarrier_Keys(t *testing.T) {
	tests := []struct {
		name    string
		objMeta *K8sObjAntCarrier
		want    []string
	}{
		// TODO: Add test cases.
		{
			name: "empty annotation",
			objMeta: &K8sObjAntCarrier{
				Annotations: map[string]string{},
			},
			want: []string{},
		},
		{
			name: "get one annotation",
			objMeta: &K8sObjAntCarrier{
				Annotations: map[string]string{KOC_PREFIX + "test-key": "test-value"},
			},
			want: []string{KOC_PREFIX + "test-key"},
		},
		{
			name: "get two annotation",
			objMeta: &K8sObjAntCarrier{
				Annotations: map[string]string{KOC_PREFIX + "test-key": "test-value", KOC_PREFIX + "test-key2": "test-value2"},
			},
			want: []string{KOC_PREFIX + "test-key", KOC_PREFIX + "test-key2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.objMeta.Keys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("K8sObjAntCarrier.Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}
