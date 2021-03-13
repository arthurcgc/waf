/*
Copyright 2016 The Kubernetes Authors.

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

// Note: the example only works with the code within the same release/branch.
package main

import (
	"os"
	"strconv"

	"github.com/arthurcgc/tcc/internal"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

const envRunOutsideCluster = "WAF_OUTSIDE_CLUSTER"

func isOutsideCluster() bool {
	outside, err := strconv.ParseBool(os.Getenv(envRunOutsideCluster))
	if err != nil {
		panic(err)
	}

	return outside
}

func main() {
	var mgr internal.WafManager
	var err error
	if isOutsideCluster() {
		mgr, err = internal.NewOutsideCluster()
		if err != nil {
			panic(err)
		}
	} else {
		mgr, err = internal.NewInCluster()
		if err != nil {
			panic(err)
		}
	}
	mgr.Start()
}