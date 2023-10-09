package endpoints

import (
	"AdminPortal/internal/service"
	"AdminPortal/model"
	"context"
	"github.com/friendsofgo/errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/volatiletech/null/v8"
	"log"
)

type OrgEndpoint struct {
	CreateOrgEndpoint endpoint.Endpoint
}

func MakeOrgEndpoints(s service.OrgService) OrgEndpoint {
	return OrgEndpoint{
		CreateOrgEndpoint: MakeCreateOrgEndpoint(s),
	}
}

type CreateOrgRequest struct {
	Student CreateOrgRequestBody `json:"org"`
}

type CreateOrgRequestBody struct {
	Id       int64  `json:"id"`
	Name     string `json:"fullName"`
	Location string `json:"location"`
}

type createResponseBody struct {
	Org int64 `json:"org"`
}

type createOrgResponse struct {
	// in:body
	Body createResponseBody `json:",inline"`
}

func MakeCreateOrgEndpoint(s service.OrgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		log.Print("call create org endpoint ", request)

		req, ok := request.(CreateOrgRequestBody)
		if !ok {
			log.Printf("Endpoint () - invalid request %v:", request)
			return nil, errInvalidRequest
		}

		ORG := model.Org{
			Id:       req.Id,
			Name:     null.StringFrom(req.Name),
			Location: null.StringFrom(req.Location),
		}

		serviceReq, errService := s.CreateOrg(ctx, ORG)

		if errService != nil {
			log.Printf("Endpoint () - error requesting the service with %v", errService)
			return nil, errService
		}

		res := createOrgResponse{
			Body: createResponseBody{
				Org: serviceReq.ID, // Assuming that `serviceReq` has an `Id` field of type int64
			},
		}
		log.Printf("Endpoint () - student response body %v", res.Body)
		return res.Body, nil
	}
}

var errInvalidRequest = errors.New("invalid request")
