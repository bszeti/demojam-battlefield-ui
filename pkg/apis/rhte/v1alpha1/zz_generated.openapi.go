package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/rhte/v1alpha1.Battlefield":       schema_pkg_apis_rhte_v1alpha1_Battlefield(ref),
		"./pkg/apis/rhte/v1alpha1.BattlefieldSpec":   schema_pkg_apis_rhte_v1alpha1_BattlefieldSpec(ref),
		"./pkg/apis/rhte/v1alpha1.BattlefieldStatus": schema_pkg_apis_rhte_v1alpha1_BattlefieldStatus(ref),
	}
}

func schema_pkg_apis_rhte_v1alpha1_Battlefield(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Battlefield is the Schema for the battlefields API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/rhte/v1alpha1.BattlefieldSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/rhte/v1alpha1.BattlefieldStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/rhte/v1alpha1.BattlefieldSpec", "./pkg/apis/rhte/v1alpha1.BattlefieldStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_rhte_v1alpha1_BattlefieldSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BattlefieldSpec defines the desired state of Battlefield",
				Properties: map[string]spec.Schema{
					"duration": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"players": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/rhte/v1alpha1.Player"),
									},
								},
							},
						},
					},
					"hitFrequency": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
				},
				Required: []string{"duration", "players", "hitFrequency"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/rhte/v1alpha1.Player"},
	}
}

func schema_pkg_apis_rhte_v1alpha1_BattlefieldStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BattlefieldStatus defines the observed state of Battlefield",
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"startTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"stopTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"scores": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/rhte/v1alpha1.PlayerStatus"),
									},
								},
							},
						},
					},
				},
				Required: []string{"phase", "scores"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/rhte/v1alpha1.PlayerStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}
