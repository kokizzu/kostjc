package model

import (
	"bufio"
	"errors"
	"fmt"
	"kostjc/model/mAuth/wcAuth"
	"kostjc/model/mProperty/wcProperty"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/pierrec/lz4/v4"
)

func RestoreDatabase(tConn *Tt.Adapter, dateTime string) {
	fmt.Println(color.GreenString("# Restore Table Users #"))
	if err := restoreTableUsers(tConn, dateTime); err != nil {
		L.LOG.Error(err)
	}

	fmt.Println(color.GreenString("# Restore Table Tenants #"))
	if err := restoreTableTenants(tConn, dateTime); err != nil {
		L.LOG.Error(err)
	}

	fmt.Println(color.GreenString("# Restore Table Locations #"))
	if err := restoreTableLocations(tConn, dateTime); err != nil {
		L.LOG.Error(err)
	}

	fmt.Println(color.GreenString("# Restore Table Facilities #"))
	if err := restoreTableFacilities(tConn, dateTime); err != nil {
		L.LOG.Error(err)
	}

	fmt.Println(color.GreenString("# Restore Table Buildings #"))
	if err := restoreTableBuildings(tConn, dateTime); err != nil {
		L.LOG.Error(err)
	}

	fmt.Println(color.GreenString("# Restore Table Rooms #"))
	if err := restoreTableRooms(tConn, dateTime); err != nil {
		L.LOG.Error(err)
	}

	fmt.Println(color.GreenString("# Restore Table Bookings #"))
	if err := restoreTableBookings(tConn, dateTime); err != nil {
		L.LOG.Error(err)
	}

	fmt.Println(color.GreenString("# Restore Table Payments #"))
	if err := restoreTablePayments(tConn, dateTime); err != nil {
		L.LOG.Error(err)
	}

	fmt.Println(color.GreenString("# Restore Table Stocks #"))
	if err := restoreTableStocks(tConn, dateTime); err != nil {
		L.LOG.Error(err)
	}
}

func restoreTableUsers(tConn *Tt.Adapter, dateTime string) error {
	backupFiles, err := getBackupFilesByTableNameByDatetime(`users`, dateTime)
	if err != nil {
		return err
	}

	for idx, f := range backupFiles {
		filePath := (backupDir + `/` + f)
		if err := func() error {
			defer subTaskPrint(color.BlueString("[%d] Restoring file %s", (idx + 1), filePath))()

			file, err := os.Open(filePath)
			if err != nil {
				L.LOG.Error("failed to open file: ", err)
				return err
			}
			defer file.Close()

			reader := lz4.NewReader(file)
			scanner := bufio.NewScanner(reader)

			stat := &ImporterStat{}
			defer stat.Print(`last`)

			var idxLine int = 0
			for scanner.Scan() {
				stat.Total++
				stat.Print()
				idxLine++

				line := scanner.Text()

				err := func() error {
					var row []any
					err := json.Unmarshal([]byte(line), &row)
					if err != nil {
						stat.Fail(`Err json.Unmarshal([]byte(line), &row): ` + err.Error())
						return err
					}

					usr := wcAuth.NewUsersMutator(tConn)
					usr.FromUncensoredArray(row)

					if usr.FindById() {
						if !usr.DoOverwriteById() {
							errMsg := `Err usr.DoOverwriteById(): ` + usr.SpaceName()
							stat.Fail(errMsg)
							return errors.New(errMsg)
						}
					}

					if !usr.DoUpsert() {
						errMsg := `Err usr.DoUpsert(): ` + usr.SpaceName()
						stat.Fail(errMsg)
						return errors.New(errMsg)
					}

					stat.Ok(true)

					return nil
				}()
				if err != nil {
					return err
				}
			}

			return nil
		}(); err != nil {
			return err
		}
	}

	return nil
}

func restoreTableTenants(tConn *Tt.Adapter, dateTime string) error {
	backupFiles, err := getBackupFilesByTableNameByDatetime(`tenants`, dateTime)
	if err != nil {
		return err
	}

	for idx, f := range backupFiles {
		filePath := (backupDir + `/` + f)
		if err := func() error {
			defer subTaskPrint(color.BlueString("[%d] Restoring file %s", (idx + 1), filePath))()

			file, err := os.Open(filePath)
			if err != nil {
				L.LOG.Error("failed to open file: ", err)
				return err
			}
			defer file.Close()

			reader := lz4.NewReader(file)
			scanner := bufio.NewScanner(reader)

			stat := &ImporterStat{}
			defer stat.Print(`last`)

			var idxLine int = 0
			for scanner.Scan() {
				stat.Total++
				stat.Print()
				idxLine++

				line := scanner.Text()

				err := func() error {
					var row []any
					err := json.Unmarshal([]byte(line), &row)
					if err != nil {
						stat.Fail(`Err json.Unmarshal([]byte(line), &row): ` + err.Error())
						return err
					}

					tnt := wcAuth.NewTenantsMutator(tConn)
					tnt.FromUncensoredArray(row)

					if tnt.FindById() {
						if !tnt.DoOverwriteById() {
							errMsg := `Err tnt.DoOverwriteById(): ` + tnt.SpaceName()
							stat.Fail(errMsg)
							return errors.New(errMsg)
						}
					}

					if !tnt.DoUpsert() {
						errMsg := `Err tnt.DoUpsert(): ` + tnt.SpaceName()
						stat.Fail(errMsg)
						return errors.New(errMsg)
					}

					stat.Ok(true)

					return nil
				}()
				if err != nil {
					return err
				}
			}

			return nil
		}(); err != nil {
			return err
		}
	}

	return nil
}

func restoreTableLocations(tConn *Tt.Adapter, dateTime string) error {
	backupFiles, err := getBackupFilesByTableNameByDatetime(`locations`, dateTime)
	if err != nil {
		return err
	}

	for idx, f := range backupFiles {
		filePath := (backupDir + `/` + f)
		if err := func() error {
			defer subTaskPrint(color.BlueString("[%d] Restoring file %s", (idx + 1), filePath))()

			file, err := os.Open(filePath)
			if err != nil {
				L.LOG.Error("failed to open file: ", err)
				return err
			}
			defer file.Close()

			reader := lz4.NewReader(file)
			scanner := bufio.NewScanner(reader)

			stat := &ImporterStat{}
			defer stat.Print(`last`)

			var idxLine int = 0
			for scanner.Scan() {
				stat.Total++
				stat.Print()
				idxLine++

				line := scanner.Text()

				err := func() error {
					var row []any
					err := json.Unmarshal([]byte(line), &row)
					if err != nil {
						stat.Fail(`Err json.Unmarshal([]byte(line), &row): ` + err.Error())
						return err
					}

					loc := wcProperty.NewLocationsMutator(tConn)
					loc.FromUncensoredArray(row)

					if loc.FindById() {
						if !loc.DoOverwriteById() {
							errMsg := `Err loc.DoOverwriteById(): ` + loc.SpaceName()
							stat.Fail(errMsg)
							return errors.New(errMsg)
						}
					}

					if !loc.DoUpsert() {
						errMsg := `Err loc.DoUpsert(): ` + loc.SpaceName()
						stat.Fail(errMsg)
						return errors.New(errMsg)
					}

					stat.Ok(true)

					return nil
				}()
				if err != nil {
					return err
				}
			}

			return nil
		}(); err != nil {
			return err
		}
	}

	return nil
}

func restoreTableFacilities(tConn *Tt.Adapter, dateTime string) error {
	backupFiles, err := getBackupFilesByTableNameByDatetime(`facilities`, dateTime)
	if err != nil {
		return err
	}

	for idx, f := range backupFiles {
		filePath := (backupDir + `/` + f)
		if err := func() error {
			defer subTaskPrint(color.BlueString("[%d] Restoring file %s", (idx + 1), filePath))()

			file, err := os.Open(filePath)
			if err != nil {
				L.LOG.Error("failed to open file: ", err)
				return err
			}
			defer file.Close()

			reader := lz4.NewReader(file)
			scanner := bufio.NewScanner(reader)

			stat := &ImporterStat{}
			defer stat.Print(`last`)

			var idxLine int = 0
			for scanner.Scan() {
				stat.Total++
				stat.Print()
				idxLine++

				line := scanner.Text()

				err := func() error {
					var row []any
					err := json.Unmarshal([]byte(line), &row)
					if err != nil {
						stat.Fail(`Err json.Unmarshal([]byte(line), &row): ` + err.Error())
						return err
					}

					fac := wcProperty.NewFacilitiesMutator(tConn)
					fac.FromUncensoredArray(row)

					if fac.FindById() {
						if !fac.DoOverwriteById() {
							errMsg := `Err fac.DoOverwriteById(): ` + fac.SpaceName()
							stat.Fail(errMsg)
							return errors.New(errMsg)
						}
					}

					if !fac.DoUpsert() {
						errMsg := `Err fac.DoUpsert(): ` + fac.SpaceName()
						stat.Fail(errMsg)
						return errors.New(errMsg)
					}

					stat.Ok(true)

					return nil
				}()
				if err != nil {
					return err
				}
			}

			return nil
		}(); err != nil {
			return err
		}
	}

	return nil
}

func restoreTableBuildings(tConn *Tt.Adapter, dateTime string) error {
	backupFiles, err := getBackupFilesByTableNameByDatetime(`buildings`, dateTime)
	if err != nil {
		return err
	}

	for idx, f := range backupFiles {
		filePath := (backupDir + `/` + f)
		if err := func() error {
			defer subTaskPrint(color.BlueString("[%d] Restoring file %s", (idx + 1), filePath))()

			file, err := os.Open(filePath)
			if err != nil {
				L.LOG.Error("failed to open file: ", err)
				return err
			}
			defer file.Close()

			reader := lz4.NewReader(file)
			scanner := bufio.NewScanner(reader)

			stat := &ImporterStat{}
			defer stat.Print(`last`)

			var idxLine int = 0
			for scanner.Scan() {
				stat.Total++
				stat.Print()
				idxLine++

				line := scanner.Text()

				err := func() error {
					var row []any
					err := json.Unmarshal([]byte(line), &row)
					if err != nil {
						stat.Fail(`Err json.Unmarshal([]byte(line), &row): ` + err.Error())
						return err
					}

					bld := wcProperty.NewBuildingsMutator(tConn)
					bld.FromUncensoredArray(row)

					if bld.FindById() {
						if !bld.DoOverwriteById() {
							errMsg := `Err bld.DoOverwriteById(): ` + bld.SpaceName()
							stat.Fail(errMsg)
							return errors.New(errMsg)
						}
					}

					if !bld.DoUpsert() {
						errMsg := `Err bld.DoUpsert(): ` + bld.SpaceName()
						stat.Fail(errMsg)
						return errors.New(errMsg)
					}

					stat.Ok(true)

					return nil
				}()
				if err != nil {
					return err
				}
			}

			return nil
		}(); err != nil {
			return err
		}
	}

	return nil
}

func restoreTableRooms(tConn *Tt.Adapter, dateTime string) error {
	backupFiles, err := getBackupFilesByTableNameByDatetime(`rooms`, dateTime)
	if err != nil {
		return err
	}

	for idx, f := range backupFiles {
		filePath := (backupDir + `/` + f)
		if err := func() error {
			defer subTaskPrint(color.BlueString("[%d] Restoring file %s", (idx + 1), filePath))()

			file, err := os.Open(filePath)
			if err != nil {
				L.LOG.Error("failed to open file: ", err)
				return err
			}
			defer file.Close()

			reader := lz4.NewReader(file)
			scanner := bufio.NewScanner(reader)

			stat := &ImporterStat{}
			defer stat.Print(`last`)

			var idxLine int = 0
			for scanner.Scan() {
				stat.Total++
				stat.Print()
				idxLine++

				line := scanner.Text()

				err := func() error {
					var row []any
					err := json.Unmarshal([]byte(line), &row)
					if err != nil {
						stat.Fail(`Err json.Unmarshal([]byte(line), &row): ` + err.Error())
						return err
					}

					room := wcProperty.NewRoomsMutator(tConn)
					room.FromUncensoredArray(row)

					if room.FindById() {
						if !room.DoOverwriteById() {
							errMsg := `Err room.DoOverwriteById(): ` + room.SpaceName()
							stat.Fail(errMsg)
							return errors.New(errMsg)
						}
					}

					if !room.DoUpsert() {
						errMsg := `Err room.DoUpsert(): ` + room.SpaceName()
						stat.Fail(errMsg)
						return errors.New(errMsg)
					}

					stat.Ok(true)

					return nil
				}()
				if err != nil {
					return err
				}
			}

			return nil
		}(); err != nil {
			return err
		}
	}

	return nil
}

func restoreTableBookings(tConn *Tt.Adapter, dateTime string) error {
	backupFiles, err := getBackupFilesByTableNameByDatetime(`bookings`, dateTime)
	if err != nil {
		return err
	}

	for idx, f := range backupFiles {
		filePath := (backupDir + `/` + f)
		if err := func() error {
			defer subTaskPrint(color.BlueString("[%d] Restoring file %s", (idx + 1), filePath))()

			file, err := os.Open(filePath)
			if err != nil {
				L.LOG.Error("failed to open file: ", err)
				return err
			}
			defer file.Close()

			reader := lz4.NewReader(file)
			scanner := bufio.NewScanner(reader)

			stat := &ImporterStat{}
			defer stat.Print(`last`)

			var idxLine int = 0
			for scanner.Scan() {
				stat.Total++
				stat.Print()
				idxLine++

				line := scanner.Text()

				err := func() error {
					var row []any
					err := json.Unmarshal([]byte(line), &row)
					if err != nil {
						stat.Fail(`Err json.Unmarshal([]byte(line), &row): ` + err.Error())
						return err
					}

					bkn := wcProperty.NewBookingsMutator(tConn)
					bkn.FromUncensoredArray(row)

					if bkn.FindById() {
						if !bkn.DoOverwriteById() {
							errMsg := `Err bkn.DoOverwriteById(): ` + bkn.SpaceName()
							stat.Fail(errMsg)
							return errors.New(errMsg)
						}
					}

					if !bkn.DoUpsert() {
						errMsg := `Err bkn.DoUpsert(): ` + bkn.SpaceName()
						stat.Fail(errMsg)
						return errors.New(errMsg)
					}

					stat.Ok(true)

					return nil
				}()
				if err != nil {
					return err
				}
			}

			return nil
		}(); err != nil {
			return err
		}
	}

	return nil
}

func restoreTablePayments(tConn *Tt.Adapter, dateTime string) error {
	backupFiles, err := getBackupFilesByTableNameByDatetime(`payments`, dateTime)
	if err != nil {
		return err
	}

	for idx, f := range backupFiles {
		filePath := (backupDir + `/` + f)
		if err := func() error {
			defer subTaskPrint(color.BlueString("[%d] Restoring file %s", (idx + 1), filePath))()

			file, err := os.Open(filePath)
			if err != nil {
				L.LOG.Error("failed to open file: ", err)
				return err
			}
			defer file.Close()

			reader := lz4.NewReader(file)
			scanner := bufio.NewScanner(reader)

			stat := &ImporterStat{}
			defer stat.Print(`last`)

			var idxLine int = 0
			for scanner.Scan() {
				stat.Total++
				stat.Print()
				idxLine++

				line := scanner.Text()

				err := func() error {
					var row []any
					err := json.Unmarshal([]byte(line), &row)
					if err != nil {
						stat.Fail(`Err json.Unmarshal([]byte(line), &row): ` + err.Error())
						return err
					}

					pym := wcProperty.NewPaymentsMutator(tConn)
					pym.FromUncensoredArray(row)

					if pym.FindById() {
						if !pym.DoOverwriteById() {
							errMsg := `Err pym.DoOverwriteById(): ` + pym.SpaceName()
							stat.Fail(errMsg)
							return errors.New(errMsg)
						}
					}

					if !pym.DoUpsert() {
						errMsg := `Err pym.DoUpsert(): ` + pym.SpaceName()
						stat.Fail(errMsg)
						return errors.New(errMsg)
					}

					stat.Ok(true)

					return nil
				}()
				if err != nil {
					return err
				}
			}

			return nil
		}(); err != nil {
			return err
		}
	}

	return nil
}

func restoreTableStocks(tConn *Tt.Adapter, dateTime string) error {
	backupFiles, err := getBackupFilesByTableNameByDatetime(`stocks`, dateTime)
	if err != nil {
		return err
	}

	for idx, f := range backupFiles {
		filePath := (backupDir + `/` + f)
		if err := func() error {
			defer subTaskPrint(color.BlueString("[%d] Restoring file %s", (idx + 1), filePath))()

			file, err := os.Open(filePath)
			if err != nil {
				L.LOG.Error("failed to open file: ", err)
				return err
			}
			defer file.Close()

			reader := lz4.NewReader(file)
			scanner := bufio.NewScanner(reader)

			stat := &ImporterStat{}
			defer stat.Print(`last`)

			var idxLine int = 0
			for scanner.Scan() {
				stat.Total++
				stat.Print()
				idxLine++

				line := scanner.Text()

				err := func() error {
					var row []any
					err := json.Unmarshal([]byte(line), &row)
					if err != nil {
						stat.Fail(`Err json.Unmarshal([]byte(line), &row): ` + err.Error())
						return err
					}

					stc := wcProperty.NewStocksMutator(tConn)
					stc.FromUncensoredArray(row)

					if stc.FindById() {
						if !stc.DoOverwriteById() {
							errMsg := `Err stc.DoOverwriteById(): ` + stc.SpaceName()
							stat.Fail(errMsg)
							return errors.New(errMsg)
						}
					}

					if !stc.DoUpsert() {
						errMsg := `Err stc.DoUpsert(): ` + stc.SpaceName()
						stat.Fail(errMsg)
						return errors.New(errMsg)
					}

					stat.Ok(true)

					return nil
				}()
				if err != nil {
					return err
				}
			}

			return nil
		}(); err != nil {
			return err
		}
	}

	return nil
}

func getBackupFilesByTableNameByDatetime(tableName string, dateTime string) (files []string, err error) {
	err = filepath.Walk(backupDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasPrefix(info.Name(), tableName+"_"+dateTime) {
			files = append(files, info.Name())
		}

		return nil
	})
	if err != nil {
		L.LOG.Error("error walking through directory:", err)
		return
	}

	if len(files) == 0 {
		dateTimeFormattedStr := ``
		dateTimeFormatted, err := time.Parse(backupTimeFormat, dateTime)
		if err == nil {
			dateTimeFormattedStr = `(` + dateTimeFormatted.Format(time.RFC1123Z) + `)`
		}
		err = errors.New("no backup files available for table " + tableName + " at " + dateTime + ` ` + dateTimeFormattedStr)
		return files, err
	}

	return
}
