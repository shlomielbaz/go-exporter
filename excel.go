package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// Excel represents excel file
type Excel struct {
	sheets *[]Sheet
}

func newExcel() Excel {
	e := Excel{
		sheets: &[]Sheet{},
	}
	return e
}

// add sheet to excel sheets collection
func (e Excel) addSheet(s Sheet) {
	*e.sheets = append(*e.sheets, s)
}

// lazy loading sheets
func (e Excel) load(printFilename bool) {
	for _, s := range *e.sheets {
		if printFilename {
			fmt.Println("filename: ", s.filename)
		}
		s.load()
	}
}

// addSheetFromFile
func (e Excel) addSheetFromFile(filename string) {
	s := newSheet(filename)
	// s.load()
	e.addSheet(s)
}

// AddSheetsFromDir gets a folder path
func (e Excel) addSheetsFromDir(root string, printFilename bool) error {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".csv" {
			return nil
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		s := newSheet(file)
		// s.load()
		e.addSheet(s)

		if printFilename == true {
			fmt.Println(file)
		}
	}

	return nil
}

func (e Excel) export(filename string) error {
	f := excelize.NewFile()
	f.Path = filename
	f.Save()

	// table, err := queries.NewLocalizationTableMetadata(e.Connection).Get()

	//     if err != nil {
	//        e.Logger.Error.Println("DONE READING LOCALIZATION TABLE METADATA WITH ERROR: ", err)
	//        return err
	//     }

	//     qo := queries.NewExportKeysQuery(e.Logger)
	//     query, values := qo.Get(req, table.Columns)
	//     record := map[string]interface{}{}
	//     iter := e.Session.Query(query, values...).Iter()

	//    defer func() {
	//        if err := iter.Close(); err != nil {
	//           e.Logger.Warning.Print("Export ExportToFile", err)
	//        }
	//    }()

	//     f := excelize.NewFile()

	//     ignore := queries.Ignore()
	//     for cn, _ := range table.Columns {
	//        if ignore[cn] == false {
	//           f.NewSheet(cn)
	//           f.SetColWidth(cn, "A", "C", 70)
	//        }
	//     }

	//     f.DeleteSheet("Sheet1")
	//     var rcount = 1
	//     for iter.MapScan(record) {
	//        for rk, _ := range record {
	//           if ignore[rk] == false {
	//              var lc = strings.Replace(rk, "-", "_", -1)
	//              var ct = strconv.Itoa(rcount)
	//              var A = "A" + ct
	//              var B = "B" + ct
	//              var C = "C" + ct
	//              var D = "D" + ct

	//              var key = record["key"].(string)
	//              var module = record["module"].(string)
	//              var source = record["en_us"].(string)
	//              var val = record[rk].(string)

	//              f.SetCellValue(lc, A, key)
	//              f.SetCellValue(lc, B, module)
	//              f.SetCellValue(lc, C, source)
	//              f.SetCellValue(lc, D, val)
	//           }
	//        }
	//        record = map[string]interface{}{}
	//        rcount++
	//     }

	//     f.Path = "/Users/assaf/Downloads/localization/export.xlsx"
	//     f.Save()

	return nil
}
