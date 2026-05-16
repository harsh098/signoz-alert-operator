package controller

import (
	"testing"

	api "github.com/harsh098/signoz-alert-operator/api/v1alpha1"
)

var testcases = []struct {
	name        string
	alert       *api.Alert
	foundByID   bool
	adoptableID string
	wantOp      alertOp
}{
	{
		name:        "fresh-nothing-in-signoz",
		alert:       newAlert(withName("test-alert"), withNamespace("alerting"), withEndpoint("production")),
		foundByID:   false,
		adoptableID: "",
		wantOp:      CREAT,
	},
	{
		name:        "fresh-adoptable-found",
		alert:       newAlert(withName("test-alert"), withNamespace("alerting"), withEndpoint("production")),
		foundByID:   false,
		adoptableID: "weoirunfs-fsdfjlsdf-fsdhdjfh",
		wantOp:      ADOPT,
	},
	{
		name:        "existing-found-by-id",
		alert:       newAlert(withName("test-alert"), withNamespace("alerting"), withEndpoint("production"), withRuleID("weoirunfs-fsdfjlsdf-fsdhdjfh")),
		foundByID:   true,
		adoptableID: "",
		wantOp:      UPDAT,
	},
	{
		name:        "existing-but-deleted-from-ui",
		alert:       newAlert(withName("test-alert"), withNamespace("alerting"), withEndpoint("production"), withRuleID("weoirunfs-fsdfjlsdf-fsdhdjfh")),
		foundByID:   false,
		adoptableID: "",
		wantOp:      CREAT,
	},
	{
		name:        "existing-deleted-from-ui-but-adoptable",
		alert:       newAlert(withName("test-alert"), withNamespace("alerting"), withEndpoint("production"), withRuleID("weoirunfs-fsdfjlsdf-fsdhdjfh")),
		foundByID:   false,
		adoptableID: "weoirunfs-fsdfjlsdf-fsdhdjfh",
		wantOp:      ADOPT,
	},
	{
		name:        "deleting-never-reconciled",
		alert:       newAlert(withName("test-alert"), withNamespace("alerting"), withEndpoint("production"), withDeletionTimestamp()),
		foundByID:   false,
		adoptableID: "",
		wantOp:      NOP,
	},
	{
		name:        "deleting-with-rule-id",
		alert:       newAlert(withName("test-alert"), withNamespace("alerting"), withEndpoint("production"), withRuleID("weoirunfs-fsdfjlsdf-fsdhdjfh"), withDeletionTimestamp()),
		foundByID:   false,
		adoptableID: "",
		wantOp:      DELET,
	},
}

func TestPlan(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := plan(tc.alert, tc.foundByID, tc.adoptableID)
			if got != tc.wantOp {
				t.Errorf("plan() = %v, want %v", got.String(), tc.wantOp.String())
			}
		})
	}
}
