package model

import (
	"bufio"
	"errors"
	"fmt"
	"kostjc/model/mAuth/wcAuth"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/pierrec/lz4/v4"
)

func RestoreDatabase(tConn *Tt.Adapter) {
	fmt.Println(color.GreenString("# Restore Table Users #"))
	if err := restoreTableUsers(tConn); err != nil {
		L.LOG.Fatal(err)
	}
}

func restoreTableUsers(tConn *Tt.Adapter) error {
	backupFiles, err := getBackupFilesByTableName(`users`)
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

func getBackupFilesByTableName(tableName string) (files []string, err error) {
	err = filepath.Walk(backupDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasPrefix(info.Name(), tableName+"_") {
			files = append(files, info.Name())
		}

		return nil
	})
	if err != nil {
		L.LOG.Error("error walking through directory:", err)
		return
	}

	if len(files) == 0 {
		err = errors.New("no backup files available")
		return
	}

	return
}
