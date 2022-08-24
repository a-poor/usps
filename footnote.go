package usps

import (
	"fmt"
	"strings"
)

type Footnote uint

const (
	// Zip Code Corrected
	// The address was found to have a different 5-digit Zip Code than given in the submitted list. The correct Zip Code is shown in the output address.
	FootZipCodeCorrected Footnote = 1 << iota

	// City / State Spelling Corrected
	// The spelling of the city name and/or state abbreviation in the submitted address was found to be different than the standard spelling. The standard spelling of the city name and state abbreviation are shown in the output address.
	FootCityStateSpellingCorrected

	// Invalid City / State / Zip
	// The Zip Code in the submitted address could not be found because neither a valid city, state, nor valid 5-digit Zip Code was present. It is also recommended that the requestor check the submitted address for accuracy.
	FootInvalidCityStateZip

	// No Zip+4 Assigned
	// This is a record listed by the United State Postal Service on the national Zip+4 file as a non-deliverable location. It is recommended that the requestor verify the accuracy of the submitted address.
	FootNoZip4Assigned

	// Zip Code Assigned for Multiple Response
	// Multiple records were returned, but each shares the same 5-digit Zip Code.
	FootZipCodeAssignedForMultipleResponse

	// Address Could Not Be Found in The National Directory File Database
	// The address, exactly as submitted, could not be found in the city, state, or Zip Code provided.
	FootAddressNotFoundInNDFD

	// Information In Firm Line Used for Matching
	// Information in the firm line was determined to be a part of the address. It was moved out of the firm line and incorporated into the address line.
	FootInfoInFirmUsedForMatching

	// Missing Secondary Number
	// Zip+4 information indicated this address is a building. The address as submitted does not contain an apartment/suite number.
	FootMissingSecondaryNumber

	// Insufficient / Incorrect Address Data
	// More than one Zip+4 was found to satisfy the address as submitted. The submitted address did not contain sufficiently complete or correct data to determine a single Zip+4 Code.
	FootInsufficientIncorrectAddressData

	// Dual Address
	// The input contained two addresses.
	FootDualAddress

	// Multiple Response Due to Cardinal Rule
	// CASS rule does not allow a match when the cardinal point of a directional changes more than 90%.
	FootMultipleResponseDueToCardinalRule

	// Address Component Changed
	// An address component was added, changed, or deleted in order to achieve a match.
	FootAddressComponentChanged

	// Street Name Changed
	// The spelling of the street name was changed in order to achieve a match.
	FootStreetNameChanged

	// Address Standardized
	// The delivery address was standardized.
	FootAddressStandardized

	// Lowest +4 Tie-Breaker
	// More than one Zip+4 Code was found to satisfy the address as submitted. The lowest Zip+4 addon may be used to break the tie between the records.
	FootLowest4TieBreaker

	// Better Address Exists
	// The delivery address is matchable, but is known by another (preferred) name.
	FootBetterAddressExists

	// Unique Zip Code Match
	// Match to an address with a unique Zip Code.
	FootUniqueZipCodeMatch

	// No Match Due To EWS
	// The delivery address is matchable, but the EWS file indicates that an exact match will be available soon.
	FootNoMatchDueToEWS

	// Incorrect Secondary Address
	// The secondary information does not match that on the national Zip+4 file. This secondary information, although present on the input address, was not valid in the range found on the national Zip+4 file.
	FootIncorrectSecondaryAddress

	// Multiple Response Due to Magnet Street Syndrome
	// The search resulted on a single response; however, the record matched was flagged as having magnet street syndrome.
	FootMultipleResponseMagnetStreet

	// Unofficial Post Office Name
	// The city or post office name in the submitted address is not recognized by the United States Postal Service as an official last line name (preferred city name) and is not acceptable as an alternate name.
	FootUnofficialPostOfficeName

	// Unverifiable City / State
	// The city and state in the submitted address could not be verified as corresponding to the given 5-digit Zip Code.
	FootUnverifiableCityState

	// Invalid Delivery Address
	// The input address record contains a delivery address other than a PO BOX, General Delivery, or Postmaster with a 5-digit Zip Code that is identified as a “small town default.” The United States Postal Service does not provide street delivery for this Zip Code. The United States Postal Service requires use of a PO BOX, General Delivery, or Postmaster for delivery within this Zip Code.
	FootInvalidDeliveryAddress

	// Unique Zip Code Generated
	// Default match inside a unique Zip Code.
	FootUniqueZipCodeGenerated

	// Military Match
	// Match made to a record with a military Zip Code.
	FootMilitaryMatch

	// Match Mode Using the ZIPMOVE Product Data
	// The ZIPMOVE product shows which Zip+4 records have moved from one Zip Code to another.
	FootMatchModeUsingZIPMOVE
)

func (foot *Footnote) UnmarshalText(t []byte) error {
	// Convert the text to a string
	s := strings.ToUpper(string(t))

	// Create a Footnote to hold the results
	var fs Footnote

	// Loop through the string, switching on the characters
	for i, r := range s {
		switch r {
		case 'A':
			fs |= FootZipCodeCorrected
		case 'B':
			fs |= FootCityStateSpellingCorrected
		case 'C':
			fs |= FootInvalidCityStateZip
		case 'D':
			fs |= FootNoZip4Assigned
		case 'E':
			fs |= FootZipCodeAssignedForMultipleResponse
		case 'F':
			fs |= FootAddressNotFoundInNDFD
		case 'G':
			fs |= FootInfoInFirmUsedForMatching
		case 'H':
			fs |= FootMissingSecondaryNumber
		case 'I':
			fs |= FootInsufficientIncorrectAddressData
		case 'J':
			fs |= FootDualAddress
		case 'K':
			fs |= FootMultipleResponseDueToCardinalRule
		case 'L':
			fs |= FootAddressComponentChanged
		case 'M':
			fs |= FootStreetNameChanged
		case 'N':
			fs |= FootAddressStandardized
		case 'O':
			fs |= FootLowest4TieBreaker
		case 'P':
			fs |= FootBetterAddressExists
		case 'Q':
			fs |= FootUniqueZipCodeMatch
		case 'R':
			fs |= FootNoMatchDueToEWS
		case 'S':
			fs |= FootIncorrectSecondaryAddress
		case 'T':
			fs |= FootMultipleResponseMagnetStreet
		case 'U':
			fs |= FootUnofficialPostOfficeName
		case 'V':
			fs |= FootUnverifiableCityState
		case 'W':
			fs |= FootInvalidDeliveryAddress
		case 'X':
			fs |= FootUniqueZipCodeGenerated
		case 'Y':
			fs |= FootMilitaryMatch
		case 'Z':
			fs |= FootMatchModeUsingZIPMOVE
		default:
			return fmt.Errorf("unknown footnote rune %c in position %d", r, i)
		}
	}

	// Update the Footnote
	*foot = fs
	return nil
}
