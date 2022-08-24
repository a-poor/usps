package usps

import (
	"fmt"
	"strings"
)

// DPVConfirmation represents the status DPV Confirmation Indicator.
//
// The DPV Confirmation Indicator is the primary method used by the
// USPS to determine whether an address was considered deliverable
// or undeliverable.
type DPVConfirmation rune

const (
	// Address was DPV confirmed for both primary and (if present) secondary numbers.
	DVP1Confirm2Confirm DPVConfirmation = 'Y'

	// Address was DPV confirmed for the primary number only, and the secondary number information was missing.
	DVP1Confirm2Missing DPVConfirmation = 'D'

	// Address was DPV confirmed for the primary number only, and the secondary number information was present by not confirmed.
	DVP1Confirm2NotConfirm DPVConfirmation = 'S'

	// Both primary and (if present) secondary number information failed to DPV confirm.
	DVP1NotConfirm2NotConfirm DPVConfirmation = 'N'
)

func (c DPVConfirmation) String() string {
	return string(c)
}

func (c *DPVConfirmation) UnmarshalText(text []byte) error {
	// Normalize the text to upper case.
	rs := []rune(strings.ToUpper(string(text)))

	// Can be empty?
	if len(rs) == 0 {
		return nil
	}

	// Check that there's no more than one character.
	if len(rs) > 1 {
		return fmt.Errorf("failed to unmarshal DPVConfirmation %q, expected 1 char", text)
	}

	// Check that the character is valid.
	switch r := DPVConfirmation(rs[0]); r {
	case DVP1Confirm2Confirm:
		*c = DVP1Confirm2Confirm
	case DVP1Confirm2Missing:
		*c = DVP1Confirm2Missing
	case DVP1Confirm2NotConfirm:
		*c = DVP1Confirm2NotConfirm
	case DVP1NotConfirm2NotConfirm:
		*c = DVP1NotConfirm2NotConfirm

	default:
		return fmt.Errorf("failed to unmarshal DPVConfirmation %q, invalid char", text)
	}

	// Return success!
	return nil
}
