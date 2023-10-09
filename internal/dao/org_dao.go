package dao

import (
	"AdminPortal/internal/database"
	"AdminPortal/internal/dbmodels"
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
)

type Org interface {
	CreatOrg(ctx context.Context, org *dbmodels.ORG) (dbmodels.ORG, error)
}
type orgDaoImpl struct {
	con database.DbConnection
}

func (o orgDaoImpl) CreatOrg(ctx context.Context, org *dbmodels.ORG) (dbmodels.ORG, error) {
	log.Print("creating new org")
	_ = org.Insert(ctx, o.con.Conn, boil.Infer())
	return *org, nil
}

func OrgDao(conn database.DbConnection) Org {
	return &orgDaoImpl{
		con: conn,
	}

}
