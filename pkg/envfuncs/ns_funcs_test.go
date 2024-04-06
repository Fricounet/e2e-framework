/*
Copyright 2024 The Kubernetes Authors.

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

package envfuncs_test

import (
	"context"
	"os"
	"testing"
	"time"

	"sigs.k8s.io/e2e-framework/pkg/env"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/features"
)

var testenv env.Environment

func TestMain(m *testing.M) {
	testenv = env.New().BeforeEachFeature(func(ctx context.Context, cfg *envconf.Config, t *testing.T, feature features.Feature) (context.Context, error) {
		t.Parallel()
		return ctx, nil
	})
	os.Exit(testenv.Run(m))
}

func TestFailNow(t *testing.T) {
	feat1 := features.New("fail now").
		Assess("Assess 1", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			t.Log("Assess 1")
			time.Sleep(5 * time.Second)
			return ctx
		}).
		Feature()

	feat2 := features.New("succeed").
		Assess("Assess 1", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			t.Log("Assess 1 (should be printed)")
			return ctx
		}).
		Feature()

	testenv.Test(t, feat1, feat2)
}
