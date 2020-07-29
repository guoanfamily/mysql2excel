package service

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

func toaxis(row int, col string) string {
	return col + strconv.Itoa(row)
}
func SaveExcel(filename string, tables []*TableModel) {
	f := excelize.NewFile()
	// Create a new sheet.
	sheetName := ""
	rowNum := 0

	contentRow := 2

	style, err := f.NewStyle(`{"font":{"bold":true,"underline":"single","family":"Berlin Sans FB Demi","size":11,"color":"#0000ee"}}`)
	if err != nil {
		fmt.Println(err)
	}

	headStyle, err := f.NewStyle(`{"font":{"bold":true,"size":11,"color":"#000000"}}`)
	if err != nil {
		fmt.Println(err)
	}

	borderStyle, err := f.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}]}`)
	if err != nil {
		fmt.Println(err)
	}
	f.SetSheetName("Sheet1", "content")
	f.SetColWidth("content", "A", "B", 20)
	f.SetCellValue("content", "A1", "表名")
	f.SetCellValue("content", "B1", "表注释")

	for _, table := range tables {
		if sheetName != table.TableName {

			//设置边框样式
			if sheetName != "" {
				f.SetCellStyle(sheetName, "A5", toaxis(rowNum-1, "E"), borderStyle)
			}
			sheetName = table.TableName

			//目录页面增加表
			f.SetCellValue("content", toaxis(contentRow, "A"), sheetName)
			f.SetCellValue("content", toaxis(contentRow, "B"), table.TableComment)
			f.SetCellHyperLink("content", toaxis(contentRow, "A"), sheetName+"!A1", "Location")
			f.SetCellStyle("content", toaxis(contentRow, "A"), toaxis(contentRow, "A"), style)

			contentRow++

			//创建新sheet，添加页头
			f.NewSheet(sheetName)
			f.SetColWidth(sheetName, "A", "B", 20)
			f.SetCellValue(sheetName, "A1", "表名")
			f.SetCellValue(sheetName, "B1", sheetName)
			f.SetCellValue(sheetName, "C1", "返回")
			f.SetCellStyle(sheetName, "C1", "C1", style)
			f.SetCellHyperLink(sheetName, "C1", "content!"+toaxis(contentRow-1, "A"), "Location")

			f.SetCellValue(sheetName, "A2", "表注释")
			f.SetCellValue(sheetName, "B2", table.TableComment)

			// 拼接表头信息
			f.SetCellValue(sheetName, "A4", "字段名称")
			f.SetCellValue(sheetName, "B4", "字段注释")
			f.SetCellValue(sheetName, "C4", "数据类型")
			f.SetCellValue(sheetName, "D4", "主键")
			f.SetCellValue(sheetName, "E4", "可空")
			f.SetCellStyle(sheetName, "A4", "E4", headStyle)
			rowNum = 5
		}
		f.SetCellValue(sheetName, toaxis(rowNum, "A"), table.ColumnName)
		f.SetCellValue(sheetName, toaxis(rowNum, "B"), table.ColumnComment)
		f.SetCellValue(sheetName, toaxis(rowNum, "C"), table.ColumnType)
		f.SetCellValue(sheetName, toaxis(rowNum, "D"), table.IsPrimary)
		f.SetCellValue(sheetName, toaxis(rowNum, "E"), table.CanNull)
		rowNum++
	}

	f.SetActiveSheet(1)
	// Save xlsx file by the given path.
	if err := f.SaveAs(filename); err != nil {
		fmt.Println(err)
	}
}
