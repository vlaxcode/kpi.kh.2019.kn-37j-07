package models

import "github.com/lancer-kit/armory/db"

var qi QI

// QI is a top level interface for interaction with database.
type QI interface {
	db.Transactional

	Account() AccountsQI

}

// Q implementation of the `QI` interface.
type Q struct {
	*db.SQLConn
}

func DB() QI {
	if qi == nil {
		qi = NewQ(nil)
	}
	return qi
}

// NewQ returns initialized instance of the `QI`.
func NewQ(dbConn *db.SQLConn) *Q {
	if dbConn == nil {
		dbConn = db.GetConnector()
	}
	return &Q{
		SQLConn: dbConn,
	}
}

func (q *Q) Account() AccountsQI {
	return NewAccountsQ(q.SQLConn)
}
