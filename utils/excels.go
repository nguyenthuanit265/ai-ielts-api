package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"gopkg.in/guregu/null.v4"
	"os"
	"reflect"
	"time"
)

func SliceOfStructsToSlice(s interface{}) [][]interface{} {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Slice {
		ShowInfoLogs("func SliceOfStructsToSlice only accepts slice input")
		return nil
	}

	var result [][]interface{}
	for i := 0; i < v.Len(); i++ {
		structVal := v.Index(i)
		if structVal.Kind() != reflect.Struct {
			panic("SliceOfStructsToSlice only accepts slice of struct input")
		}

		var row []interface{}
		for j := 0; j < structVal.NumField(); j++ {
			fieldVal := structVal.Field(j)
			switch fieldVal.Kind() {
			case reflect.Slice, reflect.Array:
				slice := make([]interface{}, fieldVal.Len())
				for k := 0; k < fieldVal.Len(); k++ {
					slice[k] = fieldVal.Index(k).Interface()
				}
				row = append(row, slice)
			default:
				switch t := fieldVal.Interface().(type) {
				case time.Time:
					if t.IsZero() && IsZeroTime(t) {
						row = append(row, "")
					} else {
						row = append(row, fieldVal.Interface())
					}
				case null.Time:
					if t.IsZero() && IsZeroTime(t.ValueOrZero()) {
						row = append(row, "")
					} else {
						row = append(row, fieldVal.Interface().(null.Time).ValueOrZero())
					}
				case null.Int:
					if !t.Valid {
						row = append(row, "")
					} else {
						row = append(row, fieldVal.Interface().(null.Int).ValueOrZero())
					}
				case null.Float:
					if !t.Valid {
						row = append(row, "")
					} else {
						row = append(row, fieldVal.Interface().(null.Float).ValueOrZero())
					}
				case null.String:
					if !t.Valid {
						row = append(row, "")
					} else {
						row = append(row, fieldVal.Interface().(null.String).ValueOrZero())
					}
				case null.Bool:
					if !t.Valid {
						row = append(row, "")
					} else {
						row = append(row, fieldVal.Interface().(null.Bool).ValueOrZero())
					}

				default:
					row = append(row, fieldVal.Interface())
				}
			}
		}
		result = append(result, row)
	}
	return result
}

func BuildHeaderResponseExcel(ctx *gin.Context, fileName string) *gin.Context {
	// Set header
	ctx.Writer.Header().Set("Content-Disposition", "attachment;filename="+fileName)
	ctx.Writer.Header().Set("Content-Transfer-Encoding", "binary")
	ctx.Writer.Header().Set("Expires", "0")
	ctx.Writer.Header().Set("Cache-Control", "must-revalidate")
	ctx.Header("Content-Type", "application/octet-stream")

	return ctx
}

func IsExistFile(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		log.Printf("%v file does not exist\n", filePath)
		return false
	} else {
		log.Printf("%v file exist\n", filePath)
		return true
	}
}

func DeleteFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		log.Errorf("Error remove file %v, error %v", filePath, err)
		return err
	}
	return nil
}

func CreateDirectory(filePath string) {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		errCreate := os.Mkdir(filePath, os.ModePerm)
		if errCreate != nil {
			log.Errorf("Error create directory %v, error %v", filePath, errCreate)
		} else {
			log.Printf("Created directory %v successfully", filePath)

		}
	}
}

func WriteFile[T any](sw *excelize.StreamWriter, listData []T, rowNumber *int, colStart int) {
	for _, data := range listData {
		dataExport := SliceOfStructsToSlice([]T{data})
		if dataExport == nil || len(dataExport) == 0 {
			continue
		}
		cell, errGetCell := excelize.CoordinatesToCellName(colStart, *rowNumber)
		if errGetCell != nil {
			log.Errorf("Error convert col %v, row %v to cell name", colStart, *rowNumber)
			break
		}
		if errSetRow := sw.SetRow(cell, dataExport[0]); errSetRow != nil {
			log.Errorf("Error set row, error %v", errSetRow)
			break
		}
		*rowNumber++
	}
}

func BuildHeaderTitleFrom(sw *excelize.StreamWriter, modelHeader any) {
	var headers []interface{}
	errSetRow := sw.SetRow("A1", headers)
	if errSetRow != nil {
		log.Errorf("Error write header title excel, error %v", errSetRow)
		return
	}
}

func ReadTitleFrom(srcFilePath string) []interface{} {
	fileTemplate, errOpenFile := excelize.OpenFile(srcFilePath)
	var title []interface{}
	if errOpenFile != nil {
		log.Errorf("Error open file template %v, error %v", srcFilePath, errOpenFile)
		return title
	}
	defer func() {
		err := fileTemplate.Close()
		if err != nil {
			log.Errorf("Error close file template %v", err)
		}
	}()

	for _, sheetName := range fileTemplate.GetSheetMap() {
		srcSheet, err := fileTemplate.GetRows(sheetName)
		log.Printf("srcSheet: %v", srcSheet)
		if err != nil {
			log.Errorf("Error get rows of source file, error %v", err)
		}
		for _, row := range srcSheet {
			for _, col := range row {
				title = append(title, col)
			}
		}
	}

	return title
}

func CopyFile(sourcePathFile string, destPathFile string) bool {
	// Open template
	fileTemplate, errOpenFile := excelize.OpenFile(sourcePathFile)
	if errOpenFile != nil {
		log.Errorf("Error open file template %v, error %v", sourcePathFile, errOpenFile)
		return false
	}
	defer func() {
		err := fileTemplate.Close()
		if err != nil {
			log.Errorf("Error close file template %v", err)
		}
	}()

	// Create a new file
	dstFile := excelize.NewFile()
	defer func() {
		if err := dstFile.Close(); err != nil {
			log.Errorf("Error close file %v", err)
		}
	}()

	// Copy all the sheets from the source file to the new file
	for _, sheetName := range fileTemplate.GetSheetMap() {
		srcSheet, err := fileTemplate.GetRows(sheetName)
		log.Printf("srcSheet: %v", srcSheet)
		if err != nil {
			log.Errorf("Error get rows of source file, error %v", err)
			return false
		}
		for _, row := range srcSheet {
			dstFile.SetSheetRow(sheetName, "A1", &row)
		}
	}

	// Save the new file with a new name
	errSaveFile := dstFile.SaveAs(destPathFile)
	if errSaveFile != nil {
		fmt.Println(errSaveFile)
		return false
	}
	return true
}
