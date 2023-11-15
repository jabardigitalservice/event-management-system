package handler

import (
	"encoding/json"
	"net/http"

	helperresponse "github.com/jabardigitalservice/super-app-services/event/src/helper/http/response"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/response"
	httpresponse "github.com/jabardigitalservice/super-app-services/event/src/response/http"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (h *Handler) CreateObject(w http.ResponseWriter, r *http.Request) {
	txn := h.app.GetNewRelic().StartTransaction("CreateObject")
	defer txn.End()

	ctx := newrelic.NewContext(r.Context(), txn)
	handlerSegment := txn.StartSegment("CreateObjectHandler")

	var respObj response.Object
	if err := json.NewDecoder(r.Body).Decode(&respObj); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reqObjJSON, err := json.Marshal(respObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var reqObj request.Object
	if err := json.Unmarshal(reqObjJSON, &reqObj); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createdObj, err := h.endpoint.CreateObject(ctx, reqObj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer handlerSegment.End()
	helperresponse.Render(w, r, h.responseMapping, httpresponse.Created, createdObj, nil)
}
