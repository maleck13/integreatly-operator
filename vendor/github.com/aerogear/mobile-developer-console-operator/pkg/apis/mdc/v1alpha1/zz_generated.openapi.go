// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/aerogear/mobile-developer-console-operator/pkg/apis/mdc/v1alpha1.MobileDeveloperConsole":       schema_pkg_apis_mdc_v1alpha1_MobileDeveloperConsole(ref),
		"github.com/aerogear/mobile-developer-console-operator/pkg/apis/mdc/v1alpha1.MobileDeveloperConsoleSpec":   schema_pkg_apis_mdc_v1alpha1_MobileDeveloperConsoleSpec(ref),
		"github.com/aerogear/mobile-developer-console-operator/pkg/apis/mdc/v1alpha1.MobileDeveloperConsoleStatus": schema_pkg_apis_mdc_v1alpha1_MobileDeveloperConsoleStatus(ref),
	}
}

func schema_pkg_apis_mdc_v1alpha1_MobileDeveloperConsole(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MobileDeveloperConsole is the Schema for the mobiledeveloperconsoles API",
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
							Ref: ref("github.com/aerogear/mobile-developer-console-operator/pkg/apis/mdc/v1alpha1.MobileDeveloperConsoleSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/aerogear/mobile-developer-console-operator/pkg/apis/mdc/v1alpha1.MobileDeveloperConsoleStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/aerogear/mobile-developer-console-operator/pkg/apis/mdc/v1alpha1.MobileDeveloperConsoleSpec", "github.com/aerogear/mobile-developer-console-operator/pkg/apis/mdc/v1alpha1.MobileDeveloperConsoleStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_mdc_v1alpha1_MobileDeveloperConsoleSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MobileDeveloperConsoleSpec defines the desired state of MobileDeveloperConsole",
				Properties: map[string]spec.Schema{
					"oAuthClientId": {
						SchemaProps: spec.SchemaProps{
							Description: "OAuthClientId is the id of the OAuthClient to use when protecting the Mobile Developer Console instance with OpenShift OAuth Proxy.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"oAuthClientSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "OAuthClientSecret is the secret of the OAuthClient to use when protecting the Mobile Developer Console instance with OpenShift OAuth Proxy.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"oAuthClientId", "oAuthClientSecret"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_mdc_v1alpha1_MobileDeveloperConsoleStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MobileDeveloperConsoleStatus defines the observed state of MobileDeveloperConsole",
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"phase"},
			},
		},
		Dependencies: []string{},
	}
}