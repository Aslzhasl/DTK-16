package excel

import (
	"github.com/xuri/excelize/v2"
	"violation-type-service/internal/model"
	"violation-type-service/internal/repository"
)

func ImportFromExcel(path string, repo repository.ViolationRepository) error {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return err
	}
	defer f.Close()

	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		return err
	}

	var list []model.ViolationType
	for i, row := range rows {
		if i == 0 || len(row) < 1 {
			continue
		}
		vt := model.ViolationType{Name: row[0]}
		if len(row) > 1 {
			vt.OtherInfo = row[1]
		}
		list = append(list, vt)
	}
	return repo.BulkInsert(list)
}