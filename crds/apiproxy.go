// Copyright (c) 2017 Northwestern Mutual.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package crds

import (
	"fmt"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var apiProxyCRD = &apiextensionsv1beta1.CustomResourceDefinition{
	ObjectMeta: metav1.ObjectMeta{
		Name: fmt.Sprintf("apiproxies.%s", KanaliGroupName),
	},
	Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
		Group:   KanaliGroupName,
		Version: Version,
		Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
			Plural:   "apiproxies",
			Singular: "apiproxy",
			ShortNames: []string{
				"ap",
			},
			Kind:     "ApiProxy",
			ListKind: "ApiProxyList",
		},
		Scope: apiextensionsv1beta1.NamespaceScoped,
		Validation: &apiextensionsv1beta1.CustomResourceValidation{
			OpenAPIV3Schema: &apiextensionsv1beta1.JSONSchemaProps{
				Required: []string{
					"spec",
				},
				AdditionalProperties: &apiextensionsv1beta1.JSONSchemaPropsOrBool{
					Allows: false,
				},
				Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
					"spec": {
						Required: []string{
							"source",
							"target",
						},
						AdditionalProperties: &apiextensionsv1beta1.JSONSchemaPropsOrBool{
							Allows: false,
						},
						Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
							"source": {
								Required: []string{
									"path",
								},
								AdditionalProperties: &apiextensionsv1beta1.JSONSchemaPropsOrBool{
									Allows: false,
								},
								Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
									"path": {
										Type:      "string",
										Pattern:   `^\/.*`,
										MinLength: int64Ptr(1),
									},
									"virtualHost": {
										Type: "string",
									},
								},
							},
							"target": {
								Required: []string{
									"backend",
								},
								AdditionalProperties: &apiextensionsv1beta1.JSONSchemaPropsOrBool{
									Allows: false,
								},
								Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
									"path": {
										Type:      "string",
										Pattern:   `^\/.*`,
										MinLength: int64Ptr(1),
									},
									"mock": {
										Type: "object",
										Required: []string{
											"configMapName",
										},
										AdditionalProperties: &apiextensionsv1beta1.JSONSchemaPropsOrBool{
											Allows: false,
										},
										Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
											"configMapName": {
												Ref: stringPtr("#/definitions/name"),
											},
										},
									},
									"backend": {
										OneOf: []apiextensionsv1beta1.JSONSchemaProps{
											{
												Type: "object",
												Required: []string{
													"endpoint",
												},
												AdditionalProperties: &apiextensionsv1beta1.JSONSchemaPropsOrBool{
													Allows: false,
												},
												Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
													"endpoint": {
														Type: "string",
													},
												},
											},
											{
												Type: "object",
												Required: []string{
													"service",
												},
												AdditionalProperties: &apiextensionsv1beta1.JSONSchemaPropsOrBool{
													Allows: false,
												},
												Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
													"service": {
														Type: "object",
														OneOf: []apiextensionsv1beta1.JSONSchemaProps{
															{
																Type: "object",
																Required: []string{
																	"name",
																	"port",
																},
																AdditionalProperties: &apiextensionsv1beta1.JSONSchemaPropsOrBool{
																	Allows: false,
																},
																Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
																	"name": {
																		Ref: stringPtr("#/definitions/name"),
																	},
																	"port": {
																		Ref: stringPtr("#/definitions/port"),
																	},
																},
															},
															{
																Type: "object",
																Required: []string{
																	"labels",
																	"port",
																},
																AdditionalProperties: &apiextensionsv1beta1.JSONSchemaPropsOrBool{
																	Allows: false,
																},
																Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
																	"port": {
																		Ref: stringPtr("#/definitions/port"),
																	},
																	"labels": {
																		Type:        "array",
																		UniqueItems: true,
																		Items: &apiextensionsv1beta1.JSONSchemaPropsOrArray{
																			Schema: &apiextensionsv1beta1.JSONSchemaProps{
																				OneOf: []apiextensionsv1beta1.JSONSchemaProps{
																					{
																						Type: "object",
																						Required: []string{
																							"name",
																							"value",
																						},
																						AdditionalProperties: &apiextensionsv1beta1.JSONSchemaPropsOrBool{
																							Allows: false,
																						},
																						Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
																							"name": {
																								Ref: stringPtr("#/definitions/name"),
																							},
																							"value": {
																								Type: "string",
																							},
																						},
																					},
																					{
																						Type: "object",
																						Required: []string{
																							"name",
																							"header",
																						},
																						AdditionalProperties: &apiextensionsv1beta1.JSONSchemaPropsOrBool{
																							Allows: false,
																						},
																						Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
																							"name": {
																								Ref: stringPtr("#/definitions/name"),
																							},
																							"header": {
																								Type: "string",
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"ssl": {
										Type: "object",
										Required: []string{
											"secretName",
										},
										AdditionalProperties: &apiextensionsv1beta1.JSONSchemaPropsOrBool{
											Allows: false,
										},
										Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
											"secretName": {
												Ref: stringPtr("#/definitions/name"),
											},
										},
									},
								},
							},
							"plugins": {
								Type:        "array",
								UniqueItems: true,
								Items: &apiextensionsv1beta1.JSONSchemaPropsOrArray{
									Schema: &apiextensionsv1beta1.JSONSchemaProps{
										Required: []string{
											"name",
										},
										AdditionalProperties: &apiextensionsv1beta1.JSONSchemaPropsOrBool{
											Allows: false,
										},
										Type: "object",
										Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
											"name": {
												Type: "string",
											},
											"version": {
												Type:    "string",
												Pattern: "^v?(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$",
											},
											"config": {
												Type: "object",
											},
										},
									},
								},
							},
						},
					},
				},
				Definitions: map[string]apiextensionsv1beta1.JSONSchemaProps{
					"name": {
						Type:      "string",
						MinLength: int64Ptr(1),
						MaxLength: int64Ptr(63),
						Pattern:   "[a-z0-9]([-a-z0-9]*[a-z0-9])?",
					},
					"port": {
						Type:    "integer",
						Minimum: float64Ptr(0),
						Maximum: float64Ptr(65535),
					},
				},
			},
		},
	},
}
