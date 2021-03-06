// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/postgres/v1alpha1.DeployPostgres":       schema_pkg_apis_postgres_v1alpha1_DeployPostgres(ref),
		"./pkg/apis/postgres/v1alpha1.DeployPostgresSpec":   schema_pkg_apis_postgres_v1alpha1_DeployPostgresSpec(ref),
		"./pkg/apis/postgres/v1alpha1.DeployPostgresStatus": schema_pkg_apis_postgres_v1alpha1_DeployPostgresStatus(ref),
	}
}

func schema_pkg_apis_postgres_v1alpha1_DeployPostgres(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DeployPostgres is the Schema for the deploypostgres API",
				Type:        []string{"object"},
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
							Ref: ref("./pkg/apis/postgres/v1alpha1.DeployPostgresSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/postgres/v1alpha1.DeployPostgresStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/postgres/v1alpha1.DeployPostgresSpec", "./pkg/apis/postgres/v1alpha1.DeployPostgresStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_postgres_v1alpha1_DeployPostgresSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DeployPostgresSpec defines the desired state of DeployPostgres",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_postgres_v1alpha1_DeployPostgresStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DeployPostgresStatus defines the observed state of DeployPostgres",
				Type:        []string{"object"},
			},
		},
	}
}
