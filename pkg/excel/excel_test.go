package excel

import (
	"github.com/Fordisk123/ginframe/conf"
	"github.com/Fordisk123/ginframe/db"
	"github.com/Fordisk123/ginframe/log"
	"github.com/xuri/excelize/v2"
	"os"
	"testing"
)

func init() {
	conf.InitConf("../../conf")
	log.NewDefaultLogger("example", "v1.0.0")
}

func TestExcel(t *testing.T) {
	db.InitDb()
	file, _ := os.Create("../../testdata/test1.xlsx")
	defer file.Close()
	f, _ := os.Open("../../testdata/test1.xlsx")
	err := RenderExcelStream(f, map[string]string{
		"user.name":     "1",
		"user.password": "1",
	}, file)
	if err != nil {
		panic(err)
	}

}

func TestReadExcel(t *testing.T) {
	f, err := os.Open("/Users/onlypiglet/FreeLife/jiangsuquanyi/ginframe/testdata/打印报表模板.xlsx")
	if err != nil {
		panic(err)
	}
	reader, err := excelize.OpenReader(f)
	if err != nil {
		panic(err)
	}
	list := reader.GetSheetList()
	for _, s := range list {
		rows, _ := reader.GetRows(s)
		for i, row := range rows {
			for j, s2 := range row {
				if s2 == "#NAME?" {
					cellName, err := excelize.CoordinatesToCellName(j+1, i+1)
					formula, err := reader.GetCellFormula(s, cellName)
					if err != nil {
						panic(err)
					}
					println(formula)
				}
				println(s2)
			}
		}

	}
	//err := RenderExcelStream(f, map[string]string{
	//	"user.name":     "1",
	//	"user.password": "1",
	//}, file)
	//if err != nil {
	//	panic(err)
	//}

}

func TestGetExprValue(t *testing.T) {
	s := GetExpr("BIND(\"BillParam.BillCode\")")
	println(s)
}

func TestLookUp(t *testing.T) {
	jsonData := "{\"BillParam\":{\"BillCode\":123},\"a\":\"v\"}"
	s := GetExpr("BIND(\"BillParam.BillCode\")")
	println(JsonLookUp(jsonData, s))
}
