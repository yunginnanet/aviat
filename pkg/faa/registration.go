package faa

import "strings"

// RT (Registration Type) is a code indicating the type of registrant.
type RT int

const (
	RTUnknown RT = iota
	RTIndividual
	RTPartnership
	RTCorporation
	RTCoOwned
	RTGovernment
	RTLLC
	RTNonCitizenCorporation
	RTNonCitizenCoOwned
)

var registrantTypeToString = map[RT]string{
	RTIndividual: "Individual", RTPartnership: "Partnership",
	RTCorporation: "Corporation", RTCoOwned: "Co-owned",
	RTGovernment: "Government", RTLLC: "LLC",
	RTNonCitizenCorporation: "Non-citizen corporation",
	RTNonCitizenCoOwned:     "Non-citizen co-owned",
	RTUnknown:               "N/A",
}

var stringToRegistrantType = map[string]RT{
	"individual": RTIndividual, "partnership": RTPartnership,
	"corporation": RTCorporation, "co-owned": RTCoOwned,
	"government": RTGovernment, "llc": RTLLC,
	"non-citizen corporation": RTNonCitizenCorporation,
	"non-citizen co-owned":    RTNonCitizenCoOwned,
}

func (rt RT) String() string {
	return registrantTypeToString[rt]
}

func parseRegistrantType(s string) RT {
	s = strings.ToLower(s)
	if rt, ok := stringToRegistrantType[s]; ok {
		return rt
	}
	return RTUnknown
}
