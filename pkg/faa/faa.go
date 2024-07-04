package faa

import (
	"encoding/hex"
	"strconv"
	"strings"
)

/*
CREATE TABLE "MASTER" (
	"N-NUMBER" VARCHAR NOT NULL,
	"SERIAL NUMBER" VARCHAR NOT NULL,
	"MFR MDL CODE" VARCHAR NOT NULL,
	"ENG MFR MDL" DECIMAL,
	"YEAR MFR" DECIMAL,
	"TYPE REGISTRANT" DECIMAL,
	"NAME" VARCHAR,
	"STREET" VARCHAR,
	"STREET2" VARCHAR,
	"CITY" VARCHAR,
	"STATE" VARCHAR,
	"ZIP CODE" VARCHAR,
	"REGION" VARCHAR,
	"COUNTY" DECIMAL,
	"COUNTRY" VARCHAR,
	"LAST ACTION DATE" DECIMAL NOT NULL,
	"CERT ISSUE DATE" DECIMAL,
	"CERTIFICATION" VARCHAR,
	"TYPE AIRCRAFT" VARCHAR NOT NULL,
	"TYPE ENGINE" DECIMAL NOT NULL,
	"STATUS CODE" VARCHAR NOT NULL,
	"MODE S CODE" DECIMAL NOT NULL,
	"FRACT OWNER" BOOLEAN,
	"AIR WORTH DATE" DECIMAL,
	"OTHER NAMES(1)" VARCHAR,
	"OTHER NAMES(2)" VARCHAR,
	"OTHER NAMES(3)" VARCHAR,
	"OTHER NAMES(4)" VARCHAR,
	"OTHER NAMES(5)" VARCHAR,
	"EXPIRATION DATE" DECIMAL,
	"UNIQUE ID" DECIMAL NOT NULL,
	"KIT MFR" VARCHAR,
	" KIT MODEL" VARCHAR,
	"MODE S CODE HEX" VARCHAR NOT NULL,
	ii BOOLEAN
);
*/

/*

 */

type (
	// NNumber is an aircraft registration number. It is the unique identifier for each aircraft in the FAA database.
	NNumber string
	// SerialNumber is the serial number assigned to an aircraft by its manufacturer.
	SerialNumber string
	// AircraftManufacturerModelCode is a code assigned to the manufacturer and model of an aircraft.
	AircraftManufacturerModelCode struct {
		ManufacturerCode string // Positions (38-40)
		ModelCode        string // Positions (41-42)
		SeriesCode       string // Positions (43-44)
	}
	// EngineManufacturerModelCode is a code assigned to the manufacturer and model of an engine.
	EngineManufacturerModelCode struct {
		ManufacturerCode string // Positions (200-202)
		ModelCode        string // Positions (203-204)
	}
	// YearManufactured is the year of manufacture.
	YearManufactured int
)

type Registrant struct {
	Name    string
	Address string
	City    string
	State   string
	Zip     string
	Region  int
}

// Region is a geographic region of a registrant of an aircraft.
type Region int

const (
	RegionUnknown Region = iota
	RegionEastern
	RegionSouthwestern
	RegionCentral
	RegionWesternPacific
	RegionAlaskan
	RegionSouthern                 = 7
	RegionEuropean                 = 8
	RegionGreatLakes               = 43 // C
	RegionNewEngland        Region = 45
	RegionNorthwestMountain Region = 53
)

func regionIsValid(i Region) bool {
	switch {
	case i < RegionEastern, i > RegionNorthwestMountain,
		i == 6, i < RegionGreatLakes && i > RegionEuropean, i == 0, i < 0:
		return false
	default:
		return true
	}
}

func strToRegion(s string) Region {
	s = strings.TrimSpace(s)
	numeric, err := strconv.Atoi(s)
	if err != nil {
		numeric, err = strconv.Atoi(hex.EncodeToString([]byte(s)))
	}
	switch {
	case err != nil:
		return RegionUnknown
	case regionIsValid(Region(numeric)):
		return Region(numeric)
	default:
		return RegionUnknown
	}
}

// AirWorthiness is a code representing the airworthiness classification of an aircraft.
type AirWorthiness int

const (
	AirWorthinessUnknown AirWorthiness = iota
	AirworthinessStandard
	AirWorthinessLimited
	AirWorthinessRestricted
	AirWorthinessExperimental
	AirWorthinessProvisional
	AirWorthinessMultiple
	AirWorthinessPrimary
	AirWorthinessSpecialFlightPermit
	AirWorthinessLightSport
)

// Type of aircraft
type TypeAircraft int

const (
	Glider                TypeAircraft = 1
	Balloon               TypeAircraft = 2
	BlimpDirigible        TypeAircraft = 3
	FixedWingSingleEngine TypeAircraft = 4
	FixedWingMultiEngine  TypeAircraft = 5
	Rotorcraft            TypeAircraft = 6
	WeightShiftControl    TypeAircraft = 7
	PoweredParachute      TypeAircraft = 8
	Gyroplane             TypeAircraft = 9
	HybridLift            TypeAircraft = 10
	OtherTypeAircraft     TypeAircraft = 11
)

// StatusCode is a code representing the registration status of an aircraft.
type StatusCode string

const (
	TriennialAircraftRegistration           StatusCode = "A"
	ExpiredDealer                           StatusCode = "D"
	CertificateRevoked                      StatusCode = "E"
	DealerCertificate                       StatusCode = "M"
	NonCitizenCorporations                  StatusCode = "N"
	RegistrationPending                     StatusCode = "R"
	SecondTriennialAircraftRegistration     StatusCode = "S"
	ValidRegistration                       StatusCode = "T"
	ValidRegistrationTrainee                StatusCode = "V"
	CertificateInvalid                      StatusCode = "W"
	EnforcementLetter                       StatusCode = "X"
	PermanentReserved                       StatusCode = "Z"
	Undeliverable                           StatusCode = "1"
	NNumberAssignedNotRegistered            StatusCode = "2"
	NNumberAssignedAmateurBuilt             StatusCode = "3"
	NNumberAssignedImport                   StatusCode = "4"
	ReservedNNumber                         StatusCode = "5"
	AdministrativelyCanceled                StatusCode = "6"
	SaleReported                            StatusCode = "7"
	SecondAttempt                           StatusCode = "8"
	CertificateRevokedEnforcement           StatusCode = "9"
	PendingCancellation                     StatusCode = "10"
	NonTypeCertificatedAmateur              StatusCode = "11"
	ImportPendingCancellation               StatusCode = "12"
	RegistrationExpired                     StatusCode = "13"
	FirstNoticeReRegistrationRenewal        StatusCode = "14"
	SecondNoticeReRegistrationRenewal       StatusCode = "15"
	RegistrationExpiredPendingCancellation  StatusCode = "16"
	SaleReportedPendingCancellation         StatusCode = "17"
	SaleReportedCanceled                    StatusCode = "18"
	RegistrationPendingPendingCancellation  StatusCode = "19"
	RegistrationPendingCanceled             StatusCode = "20"
	RevokedPendingCancellation              StatusCode = "21"
	RevokedCanceled                         StatusCode = "22"
	ExpiredDealerPendingCancellation        StatusCode = "23"
	ThirdNoticeReRegistrationRenewal        StatusCode = "24"
	FirstNoticeRegistrationRenewal          StatusCode = "25"
	SecondNoticeRegistrationRenewal         StatusCode = "26"
	RegistrationExpiredPendingCancellation2 StatusCode = "27"
	ThirdNoticeRegistrationRenewal          StatusCode = "28"
	RegistrationExpiredPendingCancellation3 StatusCode = "29"
)

// ModeSCode is the mode S code of an aircraft.
type ModeSCode string

type AircraftRegistration struct {
	NNumber             string       `master:"1-5"`     // Identification number assigned to aircraft.
	SerialNumber        string       `master:"7-36"`    // The complete aircraft serial number assigned to the aircraft by the manufacturer.
	AircraftMfrModel    string       `master:"38-44"`   // A code assigned to the aircraft manufacturer, model and series.
	EngineMfrModel      string       `master:"46-50"`   // A code assigned to the engine manufacturer and model.
	YearMfr             int          `master:"52-55"`   // Year manufactured.
	TypeRegistrant      RT           `master:"57"`      // Type of registrant (1- Individual, 2- Partnership, etc.)
	RegistrantName      string       `master:"59-108"`  // The first registrant’s name which appears on the Application for Registration.
	Street1             string       `master:"110-142"` // The street address which appears on the Application for Registration.
	Street2             string       `master:"144-176"` // The 2nd street address which appears on the Application for Registration.
	City                string       `master:"178-195"` // The city name which appears on the Application for Registration.
	State               string       `master:"197-198"` // The state name which appears on the Application for Registration.
	ZipCode             string       `master:"200-209"` // The postal Zip Code which appears on the Application for Registration.
	Region              string       `master:"211"`     // Region code (1 - Eastern, 2 - Southwestern, etc.)
	CountyMail          string       `master:"213-215"` // A code representing the county which appears on the Application for Registration.
	CountryMail         string       `master:"217-218"` // A code representing the country which appears on the Application for Registration.
	LastActivityDate    string       `master:"220-227"` // Last activity date in the format YYYY/MM/DD.
	CertificateDate     string       `master:"229-236"` // Certificate issue date in the format YYYY/MM/DD.
	AirworthinessCode   string       `master:"238"`     // AirWorthiness certificate class.
	TypeAircraft        TypeAircraft `master:"249"`     // Type of aircraft (1 - Glider, 2 - Balloon, etc.)
	TypeEngine          EngineType   `master:"251-252"` // Type of engine (0 - None, 1 - Reciprocating, etc.)
	StatusCode          string       `master:"254-255"` // Status code of the aircraft.
	ModeSCode           string       `master:"257-264"` // Aircraft Transponder Code.
	FractionalOwnership string       `master:"266"`     // Indicates if registration has fractional ownership.
	AirworthinessDate   string       `master:"268-275"` // Date of AirWorthiness.
	OtherName1          string       `master:"277-326"` // 1st co-owner or partnership name.
	OtherName2          string       `master:"328-377"` // 2nd co-owner or partnership name.
	OtherName3          string       `master:"379-428"` // 3rd co-owner or partnership name.
	OtherName4          string       `master:"430-479"` // 4th co-owner or partnership name.
	OtherName5          string       `master:"481-530"` // 5th co-owner or partnership name.
	ExpirationDate      string       `master:"532-539"` // Expiration date in the format YYYY/MM/DD.
	UniqueID            string       `master:"541-548"` // Unique Identification Number.
	KitMfr              string       `master:"550-579"` // Kit Manufacturer Name.
	KitModel            string       `master:"581-600"` // Kit Model Name.
	ModeSCodeHex        string       `master:"602-611"` // Mode S Code in Hexadecimal Format.
}
