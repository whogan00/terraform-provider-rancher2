package rancher2

import (
	"reflect"
	"testing"

	managementClient "github.com/rancher/rancher/pkg/client/generated/management/v3"
)

var (
	testClusterRKEConfigDNSNodelocalConf                   *managementClient.Nodelocal
	testClusterRKEConfigDNSLinearAutoscalerParamsConf      *managementClient.LinearAutoscalerParams
	testClusterRKEConfigDNSNodelocalInterface              []interface{}
	testClusterRKEConfigDNSLinearAutoscalerParamsInterface []interface{}
	testClusterRKEConfigDNSConf                            *managementClient.DNSConfig
	testClusterRKEConfigDNSInterface                       []interface{}
)

func init() {
	testClusterRKEConfigDNSNodelocalConf = &managementClient.Nodelocal{
		NodeSelector: map[string]string{
			"sel1": "value1",
			"sel2": "value2",
		},
		IPAddress: "ip_address",
	}
	testClusterRKEConfigDNSLinearAutoscalerParamsConf = &managementClient.LinearAutoscalerParams{
		CoresPerReplica:           128,
		Max:                       0,
		Min:                       1,
		NodesPerReplica:           4,
		PreventSinglePointFailure: true,
	}
	testClusterRKEConfigDNSNodelocalInterface = []interface{}{
		map[string]interface{}{
			"node_selector": map[string]interface{}{
				"sel1": "value1",
				"sel2": "value2",
			},
			"ip_address": "ip_address",
		},
	}
	testClusterRKEConfigDNSLinearAutoscalerParamsInterface = []interface{}{
		map[string]interface{}{
			"cores_per_replica":            128,
			"max":                          0,
			"min":                          1,
			"nodes_per_replica":            4,
			"prevent_single_point_failure": true,
		},
	}
	testClusterRKEConfigDNSConf = &managementClient.DNSConfig{
		NodeSelector: map[string]string{
			"sel1": "value1",
			"sel2": "value2",
		},
		Nodelocal:              testClusterRKEConfigDNSNodelocalConf,
		LinearAutoscalerParams: testClusterRKEConfigDNSLinearAutoscalerParamsConf,
		Provider:               "kube-dns",
		ReverseCIDRs:           []string{"rev1", "rev2"},
		UpstreamNameservers:    []string{"up1", "up2"},
	}
	testClusterRKEConfigDNSInterface = []interface{}{
		map[string]interface{}{
			"node_selector": map[string]interface{}{
				"sel1": "value1",
				"sel2": "value2",
			},
			"nodelocal":                testClusterRKEConfigDNSNodelocalInterface,
			"linear_autoscaler_params": testClusterRKEConfigDNSLinearAutoscalerParamsInterface,
			"provider":                 "kube-dns",
			"reverse_cidrs":            []interface{}{"rev1", "rev2"},
			"upstream_nameservers":     []interface{}{"up1", "up2"},
		},
	}
}

func TestFlattenClusterRKEConfigDNSNodelocal(t *testing.T) {

	cases := []struct {
		Input          *managementClient.Nodelocal
		ExpectedOutput []interface{}
	}{
		{
			testClusterRKEConfigDNSNodelocalConf,
			testClusterRKEConfigDNSNodelocalInterface,
		},
	}

	for _, tc := range cases {
		output := flattenClusterRKEConfigDNSNodelocal(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenClusterRKEConfigDNS(t *testing.T) {

	cases := []struct {
		Input          *managementClient.DNSConfig
		ExpectedOutput []interface{}
	}{
		{
			testClusterRKEConfigDNSConf,
			testClusterRKEConfigDNSInterface,
		},
	}

	for _, tc := range cases {
		output, err := flattenClusterRKEConfigDNS(tc.Input)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandClusterRKEConfigDNSNodelocal(t *testing.T) {

	cases := []struct {
		Input          []interface{}
		ExpectedOutput *managementClient.Nodelocal
	}{
		{
			testClusterRKEConfigDNSNodelocalInterface,
			testClusterRKEConfigDNSNodelocalConf,
		},
	}

	for _, tc := range cases {
		output := expandClusterRKEConfigDNSNodelocal(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandClusterRKEConfigDNS(t *testing.T) {

	cases := []struct {
		Input          []interface{}
		ExpectedOutput *managementClient.DNSConfig
	}{
		{
			testClusterRKEConfigDNSInterface,
			testClusterRKEConfigDNSConf,
		},
	}

	for _, tc := range cases {
		output, err := expandClusterRKEConfigDNS(tc.Input)
		if err != nil {
			t.Fatalf("[ERROR] on expander: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
