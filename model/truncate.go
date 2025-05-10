package model

import (
	"fmt"
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty/rqProperty"

	"github.com/fatih/color"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
)

func TruncateDatabase(tConn *Tt.Adapter) {
	fmt.Println(color.BlueString("# Truncating Table Locations #"))
	if !rqProperty.NewLocations(tConn).Truncate() {
		L.LOG.Error(`failed to truncate table locations`)
	}

	fmt.Println(color.BlueString("# Truncating Table Facilities #"))
	if !rqProperty.NewFacilities(tConn).Truncate() {
		L.LOG.Error(`failed to truncate table facilities`)
	}

	fmt.Println(color.BlueString("# Truncating Table Buildings #"))
	if !rqProperty.NewBuildings(tConn).Truncate() {
		L.LOG.Error(`failed to truncate table buildings`)
	}

	fmt.Println(color.BlueString("# Truncating Table Rooms #"))
	if !rqProperty.NewRooms(tConn).Truncate() {
		L.LOG.Error(`failed to truncate table rooms`)
	}

	fmt.Println(color.BlueString("# Truncating Table Bookings #"))
	if !rqProperty.NewBookings(tConn).Truncate() {
		L.LOG.Error(`failed to truncate table bookings`)
	}

	fmt.Println(color.BlueString("# Truncating Table Payments #"))
	if !rqProperty.NewPayments(tConn).Truncate() {
		L.LOG.Error(`failed to truncate table payments`)
	}

	fmt.Println(color.BlueString("# Truncating Table Stocks #"))
	if !rqProperty.NewStocks(tConn).Truncate() {
		L.LOG.Error(`failed to truncate table stocks`)
	}

	fmt.Println(color.BlueString("# Truncating Table Sessions #"))
	if !rqAuth.NewSessions(tConn).Truncate() {
		L.LOG.Error(`failed to truncate table sessions`)
	}

	fmt.Println(color.BlueString("# Truncating Table Users #"))
	if !rqAuth.NewUsers(tConn).Truncate() {
		L.LOG.Error(`failed to truncate table users`)
	}

	fmt.Println(color.BlueString("# Truncating Table Tenants #"))
	if !rqAuth.NewTenants(tConn).Truncate() {
		L.LOG.Error(`failed to truncate table tenants`)
	}
}
