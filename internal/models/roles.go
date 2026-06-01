package models

import (
	"database/sql"
	"sync"
)

type RoleManager struct {
	AdminID  int
	SellerID int
	ClientID int
}

var (
	Roles RoleManager
	once  sync.Once
)

func InitRoles(db *sql.DB) error {
	var err error
	once.Do(func() {
		err = db.QueryRow("SELECT id FROM roledictionary WHERE role_name = 'admin'").Scan(&Roles.AdminID)
		err = db.QueryRow("SELECT id FROM roledictionary WHERE role_name = 'seller'").Scan(&Roles.SellerID)
		err = db.QueryRow("SELECT id FROM roledictionary WHERE role_name = 'client'").Scan(&Roles.ClientID)
	})
	return err
}
