package handlers

import (
	"ARCTIC-WOLF/code/models"
	"ARCTIC-WOLF/code/risks"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRisks(ctx context.Context, w http.ResponseWriter, request *http.Request) {
	resp, err := risks.GetRisks()
	if err != nil {
		SendCommonError(err, w, 0)
		return
	}
	SendResponse(w, resp)

}

func CreateRisk(ctx context.Context, w http.ResponseWriter, request *http.Request) {
	var newRisk models.Risk
	if err := json.NewDecoder(request.Body).Decode(&newRisk); err != nil {
		SendCommonError(errors.New("invalid request payload"), w, http.StatusBadRequest)
	}
	id, status, err := risks.CreateRisk(newRisk)
	if err != nil {
		SendCommonError(err, w, status)
		return
	}
	resp := map[string]interface{}{
		"id":  id,
		"msg": "Risk Created Successfully",
	}
	SendResponse(w, resp)

}

func GetRiskById(ctx context.Context, w http.ResponseWriter, request *http.Request, c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		SendCommonError(nil, w, http.StatusBadRequest)
	}
	resp, status, err := risks.GetRiskById(id)
	if err != nil {
		SendCommonError(err, w, status)
		return
	}
	SendResponse(w, resp)
}

func SendResponse(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func SendCommonError(err error, w http.ResponseWriter, statusCode int) {
	resp := map[string]interface{}{}
	if err != nil {
		resp["msg"] = err.Error()
	} else {
		resp["msg"] = "Sorry,Not able to find the issue"
	}
	if statusCode != 0 {
		w.WriteHeader(statusCode)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
