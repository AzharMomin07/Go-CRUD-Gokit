package service

import (
	"AdminPortal/internal/dao"
	"AdminPortal/internal/dbmodels"
	"AdminPortal/model"
	"context"
)

type OrgService interface {
	CreateOrg(ctx context.Context, org model.Org) (*dbmodels.ORG, error)
}

type OrgServiceImpl struct {
	dao dao.Org
}

func (o OrgServiceImpl) CreateOrg(ctx context.Context, org model.Org) (*dbmodels.ORG, error) {

	panic("implement me")
}

func NewOrgService(orgDao dao.Org) OrgService {
	return &OrgServiceImpl{
		dao: orgDao,
	}
}
