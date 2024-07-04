package faa

// EngineType is a code representing the type of engine used by an aircraft.
type EngineType int

const (
	EngineTypeNA EngineType = iota
	EngineTypeReciprocating
	EngineTypeTurboProp
	EngineTypeTurboShaft
	EngineTypeTurboJet
	EngineTypeTurboFan
	EngineTypeRamjet
	EngineTypeTwoCycle
	EngineTypeFourCycle
	EngineTypeUnknown
	EngineTypeElectric
	EngineTypeRotary
)
