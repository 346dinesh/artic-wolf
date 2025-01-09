package risks

import (
	"ARCTIC-WOLF/code/database"
	"ARCTIC-WOLF/code/models"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

func GetRisks() ([]models.Risk, error) {
	risks := database.MemoryOfRisks
	var riskList []models.Risk
	for _, v := range risks {
		riskList = append(riskList, v)
	}
	if len(riskList) == 0 {
		return nil, errors.New("no risks found in database")
	}
	return riskList, nil
}

func CreateRisk(newRisk models.Risk) (*string, int, error) {

	// Validate the RiskState
	if !newRisk.State.IsValid() {
		return nil, http.StatusBadRequest, errors.New("invalid state provided")
	}

	// Generate a new UUID for this risk.
	newRisk.ID = uuid.New().String()
	database.MemoryOfRisks[newRisk.ID] = newRisk //This is temporary as no db connection is happening

	return &newRisk.ID, http.StatusCreated, nil
}

func GetRiskById(id string) (*models.Risk, int, error) {
	risk, ok := database.MemoryOfRisks[id]
	if !ok {
		return nil, http.StatusNotFound, errors.New("risk not found")
	}

	return &risk, http.StatusOK, nil
}
