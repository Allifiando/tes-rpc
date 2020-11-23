package models

import (
	"santara.co.id/set/market-common-service-grpc/pkg/proto/common"
)

// FindByID ...
func FindByID(id int64) (data *common.DataGetUser, err error) {
	d := &common.DataGetUser{}
	query := `SELECT id,email,pin FROM users where id = ?`
	err = Init().QueryRow(query, id).Scan(&d.Id, &d.Email, &d.Pin)
	data = d
	if err != nil {
		return data, err
	}

	return data, err
}
