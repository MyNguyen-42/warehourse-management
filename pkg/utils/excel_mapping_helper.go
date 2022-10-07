package utils

import (
	"strings"
)

type ImportOptions struct {
	DefaultValues map[string]interface{} `json:"defaultValues"`
	Layout        string                 `json:"layout"`
	HubAreaId     string                 `json:"hub_area_id"`
	Fields        []string               `json:"fields"`
	StartSheet    int                    `json:"startSheet"`
	StartRow      int                    `json:"startRow"`
	HeaderRow     int                    `json:"headerRow"`
}

func NewExcelMappingHelper() *excelMappingHelper {
	return &excelMappingHelper{}
}

type excelMappingHelper struct {
	RowsData []map[string]string
	Options  *ImportOptions
	errs     []error
}

//func (helper *excelMappingHelper) SetupWithFileData(ctx echo.Context, fileData []byte, defaultOpts *ImportOptions) error {
//	// set default opts
//	helper.Options = new(ImportOptions)
//	if defaultOpts != nil {
//		helper.Options = defaultOpts
//	}
//
//	// open file data as xlsx
//	xf, err := xlsx.OpenBinary(fileData)
//	if err != nil {
//		return errcode.Error(errcode.InvalidImportFile, "File type is not XLSX")
//	}
//
//	// get import options in data
//	if data := ctx.FormValue("data"); data != "" {
//		if err := json.Unmarshal([]byte(data), helper.Options); err != nil {
//			return errcode.Error(errcode.InvalidImportFile, err.Error())
//		}
//	}
//
//	// convert to data matrix
//	xfdata, err := xf.ToSliceUnmerged()
//	if err != nil {
//		return errcode.Error(errcode.InvalidImportFile, err.Error())
//	}
//
//	// không có dữ liệu
//	if len(xfdata) < helper.Options.StartSheet || len(xfdata[helper.Options.StartSheet]) < helper.Options.StartRow {
//		return errcode.Error(errcode.InvalidImportFile, "Data file is empty")
//	}
//
//	// nếu ko cấu hình import field nào, thì fields import sẽ lấy theo dữ liệu hàng header
//	if len(helper.Options.Fields) == 0 {
//		helper.Options.Fields = xfdata[helper.Options.StartSheet][helper.Options.HeaderRow]
//	}
//	log.Logger(ctx.Request().Context()).Error("ExcelMappingHelper SetUp fields: "+strings.Join(helper.Options.Fields, "-"))
//
//	// lấy dữ liệu
//	helper.RowsData = DataInRows(helper.Options.Fields, xfdata[helper.Options.StartSheet][helper.Options.StartRow:])
//	return nil
//}

func (helper *excelMappingHelper) AddField(field string) *excelMappingHelper {
	if helper.Options != nil {
		helper.Options.Fields = append(helper.Options.Fields, field)
	}
	return helper
}

/*func (helper *excelMappingHelper) SetupWithXLSFileData(ctx echo.Context, fileName string, fileData []byte, defaultOpts *request.ImportOptions) error {

	fileBytes := bytes.NewReader(fileData)

	xf, err := xls.OpenReader(fileBytes, "utf-8")
	if err != nil {
		return errors.Make(constant.InvalidImportFormat, err.Error())
	}

	sheet1 := xf.GetSheet(0)
	if sheet1 == nil || sheet1.MaxRow == 0 {
		return errors.Make(constant.InvalidImportFormat, err.Error())
	}

	helper.RowsData = DataInRows(defaultOpts.Fields, ReadAllCells(sheet1, defaultOpts))
	return nil
}

func ReadAllCells(sheet *xls.WorkSheet, defaultOpts *request.ImportOptions) (res [][]string) {
	for i := defaultOpts.StartRow; i <= int(sheet.MaxRow); i++ {
		if sheet.Row(i).LastCol() == -1 {
			continue
		}
		row := sheet.Row(i)
		var data []string
		for index := row.FirstCol(); index < len(defaultOpts.Fields); index++ {
			data = append(data, row.Col(index))
		}

		res = append(res, data)
	}
	return
}*/

func DataInRows(fields []string, rows [][]string) []map[string]string {
	var rowsData = make([]map[string]string, 0)
	for _, cells := range rows {
		data := DataInCells(fields, cells)
		if data != nil {
			rowsData = append(rowsData, data)
		}
	}

	return rowsData
}

func DataInCells(fields []string, cells []string) map[string]string {
	var data = make(map[string]string)
	empty := true
	for i, cell := range cells {
		if len(fields) > i {
			data[fields[i]] = GetValidString(cell)
			if data[fields[i]] != "" {
				empty = false
			}
		}
	}
	if empty {
		return nil
	}
	return data
}

func GetValidString(s string) string {
	s = strings.ToValidUTF8(s, "")
	s = strings.TrimSpace(s)
	return s
}
