package excel

import (
	"fmt"
	"github.com/Fordisk123/ginframe/conf"
	"github.com/Fordisk123/ginframe/log"
	"testing"
)

func init() {
	conf.InitConf("../../conf")
	log.NewDefaultLogger("example", "v1.0.0")
}

func TestExcel(t *testing.T) {
	//db.InitDb()
	//file, _ := os.Create("../../testdata/test1.xlsx")
	//defer file.Close()
	//f, _ := os.Open("../../testdata/test1.xlsx")
	//err := RenderExcelStream(f, map[string]string{
	//	"user.name":     "1",
	//	"user.password": "1",
	//}, file)
	//if err != nil {
	//	panic(err)
	//}

}

func TestReadExcel(t *testing.T) {
	//f, err := os.Open("/Users/onlypiglet/FreeLife/jiangsuquanyi/ginframe/testdata/打印报表模板.xlsx")
	//if err != nil {
	//	panic(err)
	//}
	//reader, err := excelize.OpenReader(f)
	//if err != nil {
	//	panic(err)
	//}
	//list := reader.GetSheetList()
	//for _, s := range list {
	//	rows, _ := reader.GetRows(s)
	//	for i, row := range rows {
	//		for j, s2 := range row {
	//			if s2 == "#NAME?" {
	//				cellName, err := excelize.CoordinatesToCellName(j+1, i+1)
	//				formula, err := reader.GetCellFormula(s, cellName)
	//				if err != nil {
	//					panic(err)
	//				}
	//				println(formula)
	//			}
	//			println(s2)
	//		}
	//	}
	//
	//}
	//err := RenderExcelStream(f, map[string]string{
	//	"user.name":     "1",
	//	"user.password": "1",
	//}, file)
	//if err != nil {
	//	panic(err)
	//}

}

func TestGetExprValue(t *testing.T) {
	//s := GetExpr("BIND(\"BillParam.BillCode\")")
	//println(s)
}

//func TestLookUp(t *testing.T) {
//	s := GetExpr("BIND(\"deptName\")")
//	jsonData := "{\n  \"id\" : 24,\n  \"createdAt\" : \"2024-12-01T13:40:00.959+08:00\",\n  \"updatedAt\" : \"2024-12-01T13:40:00.959+08:00\",\n  \"createTimestamp\" : 1733031600,\n  \"updateTimestamp\" : 1733031600,\n  \"deleteAt\" : null,\n  \"createBy\" : \"\",\n  \"updateBy\" : \"\",\n  \"equipCode\" : \"111234\",\n  \"equipmentLedger\" : {\n    \"id\" : 1,\n    \"createdAt\" : \"2024-11-03T17:21:58.632+08:00\",\n    \"updatedAt\" : \"2024-12-02T08:31:05.432+08:00\",\n    \"createTimestamp\" : 1730625718,\n    \"updateTimestamp\" : 1733099465,\n    \"deleteAt\" : null,\n    \"createBy\" : \"admin\",\n    \"updateBy\" : \"admin\",\n    \"deptId\" : 3,\n    \"Department\" : null,\n    \"equipClassId\" : 3,\n    \"equipModeId\" : 10,\n    \"equipCode\" : \"111234\",\n    \"equipName\" : \"测试设备\",\n    \"registerStatus\" : \"1\",\n    \"validTime\" : 1732094606,\n    \"valid\" : \"2\",\n    \"lastLatitude\" : \"190.2\",\n    \"lastLongitude\" : \"189.2\",\n    \"lastRadius\" : 12,\n    \"lastAddress\" : \"123\",\n    \"closed\" : \"0\",\n    \"lastTime\" : 0,\n    \"fourgIccId\" : \"xxx123132\",\n    \"fourgIccIdValidTime\" : 1739779388,\n    \"soonFourgIccIdValid\" : \"0\"\n  },\n  \"CustomerID\" : 0,\n  \"deptId\" : 0,\n  \"deptName\" : \"测试单位\",\n  \"leaderName\" : \"测试人员\",\n  \"leaderPhone\" : \"12365689635\",\n  \"companyType\" : \"-1\",\n  \"companyTypeDescription\" : \"\",\n  \"deptUscc\" : \"56797543256\",\n  \"address\" : \"花园街\",\n  \"sampleCode\" : \"001\",\n  \"sampleName\" : \"测试样品\",\n  \"sampleType\" : \"777\",\n  \"sampleTesterNumber\" : \"555\",\n  \"sampleTesterRank\" : \"0\",\n  \"sampleQualification\" : \"777\",\n  \"billCode\" : \"TJ0014\",\n  \"certificateNumber\" : \"Cert001\",\n  \"billDate\" : 0,\n  \"billStatus\" : \"1\",\n  \"billStatusDescription\" : \"已试验\",\n  \"standardName\" : \"gv1234\",\n  \"standardNumber\" : \"123456\",\n  \"standardLimit\" : \"gv1234\",\n  \"standardMaxError\" : \"0.01\",\n  \"standardCertNumber\" : \"222\",\n  \"standardCertDate\" : 1709856000,\n  \"standardBasis\" : \"12233\",\n  \"testRange\" : 1000,\n  \"testCount\" : 0,\n  \"collectPoints\" : 2,\n  \"calibratePlace\" : \"花园街\",\n  \"calibratePrice\" : 0,\n  \"calibrateResult\" : \"\",\n  \"environmentTemp\" : 0,\n  \"environmentHumi\" : 0,\n  \"calibrateId\" : 0,\n  \"calibrateName\" : \"校准员1\",\n  \"certificateTime\" : 0,\n  \"approveId\" : 0,\n  \"approveName\" : \"\",\n  \"approveTime\" : 0,\n  \"remark\" : \"\",\n  \"billExternalInfo\" : \"\",\n  \"customer_json_field\" : { },\n  \"experimentData\" : null,\n  \"experimentalDataHeader\" : null\n}"
//	println(JsonLookUp(jsonData, s))
//}

func TestGetExprValue2(t *testing.T) {
	s := "BINDCollectValuesExpr(\"experimentData.#(TorqueType==低扭矩)#|#(LoadDirection==顺时针)#.CollectValues|0\",\"5\",\"index\")"
	expr := GetExpr(s)
	println(fmt.Sprintf("%+v", expr))

}
