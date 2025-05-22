package model

import (
	"errors"
	"fmt"
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty/rqProperty"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/pierrec/lz4/v4"
)

const backupDir = `./backup`
const backupTimeFormat = `20060102150405`

func getBackupTableFileOutput(tableName string, offset, limit int) string {
	datetime := time.Now().Format(backupTimeFormat)
	host, err := os.Hostname()
	if err != nil {
		host = `unknown`
	}
	return fmt.Sprintf(
		"%s/%s_%s-%s_%d_%d.jsonline.lz4",
		backupDir, tableName, datetime, host, offset, limit,
	)
}

func BackupDatabase(tConn *Tt.Adapter) {
	err := os.MkdirAll(backupDir, os.ModePerm)
	if err != nil {
		L.LOG.Error(`failed to create backup directory :`, err)
		return
	}

	fmt.Println(color.BlueString("# Backup Table Users #"))
	if err := backupTable(tConn, rqAuth.NewUsers, `users`); err != nil {
		L.LOG.Fatal(err)
	}

	fmt.Println(color.BlueString("# Backup Table Tenants #"))
	if err := backupTable(tConn, rqAuth.NewTenants, `tenants`); err != nil {
		L.LOG.Fatal(err)
	}

	fmt.Println(color.BlueString("# Backup Table Locations #"))
	if err := backupTable(tConn, rqProperty.NewLocations, `locations`); err != nil {
		L.LOG.Fatal(err)
	}

	fmt.Println(color.BlueString("# Backup Table Facilities #"))
	if err := backupTable(tConn, rqProperty.NewFacilities, `facilities`); err != nil {
		L.LOG.Fatal(err)
	}

	fmt.Println(color.BlueString("# Backup Table Buildings #"))
	if err := backupTable(tConn, rqProperty.NewBuildings, `buildings`); err != nil {
		L.LOG.Fatal(err)
	}

	fmt.Println(color.BlueString("# Backup Table Rooms #"))
	if err := backupTable(tConn, rqProperty.NewRooms, `rooms`); err != nil {
		L.LOG.Fatal(err)
	}

	fmt.Println(color.BlueString("# Backup Table Bookings #"))
	if err := backupTable(tConn, rqProperty.NewBookings, `bookings`); err != nil {
		L.LOG.Fatal(err)
	}

	fmt.Println(color.BlueString("# Backup Table Payments #"))
	if err := backupTable(tConn, rqProperty.NewPayments, `payments`); err != nil {
		L.LOG.Fatal(err)
	}

	fmt.Println(color.BlueString("# Backup Table Stocks #"))
	if err := backupTable(tConn, rqProperty.NewStocks, `stocks`); err != nil {
		L.LOG.Fatal(err)
	}
}

type newTableFunc interface {
	CountTotalAllRows() (total uint64)
	GetRows(offset, limit uint32) [][]any
	Truncate() bool
}

func backupTable[T newTableFunc](conn *Tt.Adapter, newFunc func(tt *Tt.Adapter) T, tableName string) error {
	table := newFunc(conn)

	totalAllRows := table.CountTotalAllRows()

	batchSize := 10_000

	for i := 0; i < int(totalAllRows); i += batchSize {
		err := func() error {
			data := table.GetRows(uint32(i), uint32(batchSize))
			if len(data) == 0 {
				return errors.New("no data")
			}

			outputFile := getBackupTableFileOutput(tableName, i, batchSize)
			file, err := os.Create(outputFile)
			if err != nil {
				L.LOG.Error(`Err os.Create(outputFile): `, err)
				return err
			}
			defer file.Close()

			lz4Writer := lz4.NewWriter(file)
			defer lz4Writer.Close()

			for _, row := range data {
				jsonRow, err := json.Marshal(row)
				if err != nil {
					L.LOG.Error(`Err json.Marshal(row): `, err)
					return err
				}

				_, err = lz4Writer.Write(append(jsonRow, '\n'))
				if err != nil {
					L.LOG.Error(`Err lz4Writer.Write(append(jsonRow, '\n')): `, err)
					return err
				}
			}

			fmt.Println(color.GreenString("[OK] Backed up to file " + outputFile))

			return nil
		}()
		if err != nil {
			return err
		}
	}

	return nil
}
