/*
Copyright 2018 The Kubernetes Authors.

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

package test

import (
	apiv1 "k8s.io/api/core/v1"
	vpa_types "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/poc.autoscaling.k8s.io/v1alpha1"
)

// RecommendationBuilder helps building test instances of RecommendedPodResources.
type RecommendationBuilder interface {
	WithContainer(containerName string) RecommendationBuilder
	WithTarget(cpu, memory string) RecommendationBuilder
	WithMinRecommended(cpu, memory string) RecommendationBuilder
	WithMaxRecommended(cpu, memory string) RecommendationBuilder
	Get() *vpa_types.RecommendedPodResources
}

// Recommendation returns a new RecommendationBuilder.
func Recommendation() RecommendationBuilder {
	return &recommendationBuilder{}
}

type recommendationBuilder struct {
	containerName  string
	target         apiv1.ResourceList
	minRecommended apiv1.ResourceList
	maxRecommended apiv1.ResourceList
}

func (b *recommendationBuilder) WithContainer(containerName string) RecommendationBuilder {
	c := *b
	c.containerName = containerName
	return &c
}

func (b *recommendationBuilder) WithTarget(cpu, memory string) RecommendationBuilder {
	c := *b
	c.target = Resources(cpu, memory)
	return &c
}

func (b *recommendationBuilder) WithMinRecommended(cpu, memory string) RecommendationBuilder {
	c := *b
	c.minRecommended = Resources(cpu, memory)
	return &c
}

func (b *recommendationBuilder) WithMaxRecommended(cpu, memory string) RecommendationBuilder {
	c := *b
	c.maxRecommended = Resources(cpu, memory)
	return &c
}

func (b *recommendationBuilder) Get() *vpa_types.RecommendedPodResources {
	if b.containerName == "" {
		panic("Must call WithContainer() before Get()")
	}
	return &vpa_types.RecommendedPodResources{
		ContainerRecommendations: []vpa_types.RecommendedContainerResources{
			{
				Name:           b.containerName,
				Target:         b.target,
				MinRecommended: b.minRecommended,
				MaxRecommended: b.maxRecommended,
			},
		}}
}
