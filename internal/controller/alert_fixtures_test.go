/*
Copyright 2026.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0
*/

package controller

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	api "github.com/harsh098/signoz-alert-operator/api/v1alpha1"
)

// alertOpt mutates an Alert in-place. Used to override defaults in newAlert.
type alertOpt func(*api.Alert)

// newAlert returns an Alert with sensible defaults: namespace=test-ns,
// name=test-alert, endpointRef=signoz, and a minimal-but-valid Spec.Rule
// (matches the required fields of RuletypesPostableRule). Apply opts to
// override defaults.
func newAlert(opts ...alertOpt) *api.Alert {
	a := &api.Alert{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-alert",
			Namespace: "test-ns",
		},
		Spec: api.AlertSpec{
			EndpointRef: api.EndpointRefSpec{Name: "signoz"},
			Rule: runtime.RawExtension{
				Raw: []byte(`{
					"alert": "test-alert",
					"alertType": "METRIC_BASED_ALERT",
					"ruleType": "threshold_rule",
					"condition": {"compositeQuery": {}}
				}`),
			},
		},
	}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

// withName overrides metadata.name.
func withName(name string) alertOpt {
	return func(a *api.Alert) { a.Name = name }
}

// withNamespace overrides metadata.namespace.
func withNamespace(ns string) alertOpt {
	return func(a *api.Alert) { a.Namespace = ns }
}

// withEndpoint overrides spec.endpointRef.name.
func withEndpoint(name string) alertOpt {
	return func(a *api.Alert) { a.Spec.EndpointRef.Name = name }
}

// withRuleJSON overrides spec.rule with the given raw JSON.
func withRuleJSON(raw string) alertOpt {
	return func(a *api.Alert) { a.Spec.Rule = runtime.RawExtension{Raw: []byte(raw)} }
}

// withRuleID pre-populates status.ruleID. Use to simulate an Alert that has
// already been reconciled successfully at least once.
func withRuleID(id string) alertOpt {
	return func(a *api.Alert) { a.Status.RuleID = id }
}

// withDeletionTimestamp marks the Alert as being deleted. For envtest, prefer
// k8s Delete() — but this is fine for plan() unit tests where we just need the
// in-memory bit set.
func withDeletionTimestamp() alertOpt {
	return func(a *api.Alert) {
		now := metav1.Now()
		a.DeletionTimestamp = &now
		// k8s API server only allows DeletionTimestamp to be set when a finalizer
		// is present; mirror that here so envtest paths see a consistent object.
		a.Finalizers = append(a.Finalizers, "monitoring.hmx86.cloud/finalizer")
	}
}

// withFinalizer adds the operator's finalizer. Use when testing reconcile
// paths that expect the finalizer to already be installed.
func withFinalizer() alertOpt {
	return func(a *api.Alert) {
		a.Finalizers = append(a.Finalizers, "monitoring.hmx86.cloud/finalizer")
	}
}
