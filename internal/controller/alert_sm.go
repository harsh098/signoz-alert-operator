package controller

import (
	api "github.com/harsh098/signoz-alert-operator/api/v1alpha1"
)

type alertOp int

// Enums for Operations
const (
	NOP alertOp = iota
	CREAT
	UPDAT
	ADOPT
	DELET
)

func (op alertOp) String() string {
	value := map[alertOp]string{
		NOP:   "no-op",
		CREAT: "create",
		UPDAT: "update",
		ADOPT: "adopt",
		DELET: "delete",
	}

	return value[op]
}

// State Machine:
// This state machine is used by the Reconiler to
// decide the next course of action for a given alert.
func plan(alert *api.Alert, foundByID bool, adoptableID string) alertOp {
	delTS := alert.DeletionTimestamp
	ruleID := alert.Status.RuleID

	switch {
	case delTS != nil && ruleID != "":
		return DELET
	case delTS != nil && ruleID == "":
		return NOP
	case foundByID:
		return UPDAT
	case adoptableID != "":
		return ADOPT
	}

	return CREAT
}
