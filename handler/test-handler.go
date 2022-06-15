package handler

import (
	"bc_melomingoo/processor"
	"encoding/json"
	"net/http"
)

type TestHandler BaseHandler

func (h *TestHandler) TestHandlerCheck(w http.ResponseWriter, r *http.Request) {
	db := h.TestDB

	testList, err := processor.GetTestList(db)

	if err != nil {

		return
	}
	if err := json.NewEncoder(w).Encode(testList); err != nil {

	}
}
