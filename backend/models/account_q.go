package models

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/lancer-kit/armory/db"
	"github.com/pkg/errors"
)

const tableAccounts = "accounts"

type AccountsQI interface {
	Insert(acc *Account) (int64, error)
	GetLast() (*Account, error)
}

type AccountsQ struct {
	*db.SQLConn
	db.Table
}

func NewAccountsQ(conn *db.SQLConn) *AccountsQ {
	return &AccountsQ{
		SQLConn: conn.Clone(),
		Table: db.Table{
			Name:     tableAccounts,
			QBuilder: sq.Select("*").From(tableAccounts),
		},
	}
}

func (q *AccountsQ) Insert(acc *Account) (int64, error) {
	if acc == nil {
		return 0, errors.New("no data passed")
	}

	fld := map[string]interface{}{
		"name":     acc.Name,
		"surname":  acc.Surname,
		"login":    acc.Login,
		"password": acc.Password,
		"sex":      acc.Sex,
		"email":    acc.Email,
		"date":     acc.Date,
	}

	query := sq.Insert(q.Name).SetMap(fld)

	id, err := q.SQLConn.Insert(query)
	intID, ok := id.(int64)
	if !ok {
		return 0, errors.New("unable to parse id to int64")
	}
	return intID, err
}

func (q *AccountsQ) GetLast() (*Account, error) {
	res := make([]Account, 0)
	q.ApplyPage("id desc")

	err := q.SQLConn.Select(q.QBuilder, &res)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &res[0], err
}
