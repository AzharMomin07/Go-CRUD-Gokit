package transport

import (
	"AdminPortal/internal/endpoints"
	"context"
	"encoding/json"
	httpTrans "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
)

func CreatesOrgHttpHandler(endpoint endpoints.OrgEndpoint, router *mux.Router) {

	router.Handle("/orgs",
		httpTrans.NewServer(
			endpoint.CreateOrgEndpoint,
			decodeCreateOrg,
			encodeOrg,
		)).Methods(http.MethodPost)

	return
}

func encodeOrg(ctx context.Context, writer http.ResponseWriter, i interface{}) error {
	if err, ok := i.(error); ok && err != nil {
		logrus.Warnf("Encode ()  - error %v", ok)
	}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(writer).Encode(i)
}

func decodeCreateOrg(ctx context.Context, request2 *http.Request) (request interface{}, err error) {

	var org endpoints.CreateOrgRequestBody

	body, err := io.ReadAll(request2.Body)

	if err != nil {
		log.Printf("Decode () - nill error in body %v", err)
	}
	log.Print("Decode () - transport add organization - request body :", string(body))

	err = json.Unmarshal(body, &org)

	return org, nil
}
