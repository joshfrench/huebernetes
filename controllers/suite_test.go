/*
Copyright 2021 Josh French

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

package controllers

import (
	"path/filepath"
	"testing"

	. "gopkg.in/check.v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	v1alpha1 "huebernetes.dev/huebernetes/api/v1alpha1"
)

func Test(t *testing.T) { TestingT(t) }

type ControllerSuite struct{}

var _ = Suite(&ControllerSuite{})

var cfg *rest.Config
var k8sClient client.Client
var testEnv *envtest.Environment

func (s *ControllerSuite) SetUpSuite(c *C) {
	testEnv = &envtest.Environment{
		CRDDirectoryPaths:     []string{filepath.Join("..", "config", "crd", "bases")},
		ErrorIfCRDPathMissing: false,
	}
	cfg, err := testEnv.Start()
	c.Assert(err, IsNil)
	c.Assert(cfg, NotNil)

	err = v1alpha1.AddToScheme(scheme.Scheme)
	c.Assert(err, IsNil)

	//+kubebuilder:scaffold:scheme

	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	c.Assert(err, IsNil)
	c.Assert(k8sClient, NotNil)
}

func (s *ControllerSuite) TearDownSuite(c *C) {
	err := testEnv.Stop()
	c.Assert(err, IsNil)
}
