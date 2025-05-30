/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	policyv1 "k8s.io/api/policy/v1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// EvictionsGetter has a method to return a EvictionInterface.
// A group's client should implement this interface.
type EvictionsGetter interface {
	Evictions(namespace string) EvictionInterface
}

// EvictionInterface has methods to work with Eviction resources.
type EvictionInterface interface {
	EvictionExpansion
}

// evictions implements EvictionInterface
type evictions struct {
	*gentype.Client[*policyv1.Eviction]
}

// newEvictions returns a Evictions
func newEvictions(c *PolicyV1Client, namespace string) *evictions {
	return &evictions{
		gentype.NewClient[*policyv1.Eviction](
			"evictions",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *policyv1.Eviction { return &policyv1.Eviction{} },
			gentype.PrefersProtobuf[*policyv1.Eviction](),
		),
	}
}
