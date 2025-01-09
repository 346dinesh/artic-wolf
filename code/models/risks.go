package models

//Currently Storing in Memory

// Risk represents a Risk object in our system.
type Risk struct {
	ID          string    `json:"id"`
	State       RiskState `json:"state"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

// RiskState is a custom type that represents valid risk states.
type RiskState string

// Constants that define valid risk states.
const (
	RiskStateOpen          RiskState = "open"
	RiskStateClosed        RiskState = "closed"
	RiskStateAccepted      RiskState = "accepted"
	RiskStateInvestigating RiskState = "investigating"
)

// IsValid checks if a given RiskState is valid.
func (rs RiskState) IsValid() bool {
	switch rs {
	case
		RiskStateOpen,
		RiskStateClosed,
		RiskStateAccepted,
		RiskStateInvestigating:
		return true
	}
	return false
}
