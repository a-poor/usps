package usps

import (
	"fmt"
	"strings"
)

// DVPFootnote
//
// DPV® Standardized Footnotes - EZ24x7Plus and Mail*STAR are required to express DPV
// results using USPS standard two character footnotes.
//
// Example: AABB
//
// Footnotes Reporting CASS™ ZIP+4™ Certification:
// AA, A1
//
// Footnotes Reporting DPV Validation Observations:
// BB, CC, N1, M1, M3, P1, P3, F1, G1, U1
type DVPFootnote uint

const (
	// (AA) Input address matched to the ZIP+4 file.
	DVPFootInputMatchedZip4File DVPFootnote = 1 << iota

	// (A1) Input address not matched to the ZIP+4 file.
	DVPFootInputDidntMatchZip4File

	// (BB) Matched to DPV (all components).
	DVPFootMatchedAll

	// (CC) Secondary number not matched (present but invalid).
	DVPFootSecondaryInvalid

	// (N1) High-rise address missing secondary number.
	DVPFootHighRiseMissingSecondaryNumber

	// (M1) Primary number missing.
	DVPFootPrimaryMissing

	// (M3) Primary number invalid.
	DVPFootPrimaryInvalid

	// (P1) Input Address RR or HC Box number Missing.
	DVPFootAddrRR_HCNumberMissing

	// (P3) Input Address PO, RR, or HC Box number Invalid.
	DVPFootAddrPO_RR_HCNumberInvalid

	// (F1) Input Address Matched to a Military Address.
	DVPFootInputMatchedMilitaryAddress

	// (G1) Input Address Matched to a General Delivery Address.
	DVPFootInputMatchedGeneralDeliveryAddress

	// (U1) Input Address Matched to a Unique ZIP Code™.
	DVPFootInputMatchedUniqueZipCode
)

func (f DVPFootnote) String() string {
	var s string
	if f&DVPFootInputMatchedZip4File > 0 {
		s += "AA"
	}
	if f&DVPFootInputDidntMatchZip4File > 0 {
		s += "A1"
	}
	if f&DVPFootMatchedAll > 0 {
		s += "BB"
	}
	if f&DVPFootSecondaryInvalid > 0 {
		s += "CC"
	}
	if f&DVPFootHighRiseMissingSecondaryNumber > 0 {
		s += "N1"
	}
	if f&DVPFootPrimaryMissing > 0 {
		s += "M1"
	}
	if f&DVPFootPrimaryInvalid > 0 {
		s += "M3"
	}
	if f&DVPFootAddrRR_HCNumberMissing > 0 {
		s += "P1"
	}
	if f&DVPFootAddrPO_RR_HCNumberInvalid > 0 {
		s += "P3"
	}
	if f&DVPFootInputMatchedMilitaryAddress > 0 {
		s += "F1"
	}
	if f&DVPFootInputMatchedGeneralDeliveryAddress > 0 {
		s += "G1"
	}
	if f&DVPFootInputMatchedUniqueZipCode > 0 {
		s += "U1"
	}
	return s
}

func (foot *DVPFootnote) UnmarshalText(t []byte) error {
	// Clean the text
	rs := []rune(strings.ToUpper(string(t)))

	// Validate the result.
	// If empty, return an empty DVPFootnote.
	if len(rs) == 0 {
		*foot = DVPFootnote(0)
		return nil
	}

	// Assert that the length is divisible by 2.
	if len(rs)%2 != 0 {
		return fmt.Errorf("invalid dvp footnote %q, expected length to be divisible by 2", string(t))
	}

	// Create a variable to hold the result.
	var fs DVPFootnote

	// Iterate over the runes.
	for i := 0; i+1 < len(rs); i += 2 {
		// Get the two characters.
		cs := string(rs[i : i+2])

		switch cs {
		case "AA":
			fs |= DVPFootInputMatchedZip4File
		case "A1":
			fs |= DVPFootInputDidntMatchZip4File
		case "BB":
			fs |= DVPFootMatchedAll
		case "CC":
			fs |= DVPFootSecondaryInvalid
		case "N1":
			fs |= DVPFootHighRiseMissingSecondaryNumber
		case "M1":
			fs |= DVPFootPrimaryMissing
		case "M3":
			fs |= DVPFootPrimaryInvalid
		case "P1":
			fs |= DVPFootAddrRR_HCNumberMissing
		case "P3":
			fs |= DVPFootAddrPO_RR_HCNumberInvalid
		case "F1":
			fs |= DVPFootInputMatchedMilitaryAddress
		case "G1":
			fs |= DVPFootInputMatchedGeneralDeliveryAddress
		case "U1":
			fs |= DVPFootInputMatchedUniqueZipCode
		default:
			return fmt.Errorf("unknown dvp footnote %q in position %d", string(t), i)
		}
	}

	// Set the result.
	*foot = fs
	return nil
}
