package excel

import (
	"context"
	"fmt"
	"github.com/Fordisk123/ginframe/db"
	"github.com/Fordisk123/ginframe/log"
	"github.com/Fordisk123/ginframe/pkg/file"
	gerrors "github.com/pkg/errors"
	"github.com/tidwall/gjson"
	excelize "github.com/xuri/excelize/v2"
	"io"
	"strings"
)

const TemplatePrefixFieldName = "$_$_"
const TableAndFieldSplitChar = "."

func dbExtraceExpr(expr string) (string, string, error) {
	split := strings.Split(expr[4:], TableAndFieldSplitChar)
	if len(split) != 2 {
		return "", "", fmt.Errorf("%s is invalid,should like %stable_name%sfield_name", expr, TemplatePrefixFieldName, TableAndFieldSplitChar)
	}
	return split[0], split[1], nil
}

func DbValue(expr string, indexs map[string]string) (interface{}, error) {
	tn, fn, err := dbExtraceExpr(expr)
	if err != nil {
		return nil, err
	}
	id := indexs[fmt.Sprintf("%s.%s", tn, fn)]
	if id == "" {
		return nil, fmt.Errorf("didn't find %s.%s index value of", tn, fn)
	}
	data := make(map[string]interface{})
	if err := db.GetDb(nil).Exec(fmt.Sprintf("select %s from %s where id = %s", fn, tn, id)).Table(tn).Take(data).Error; err != nil {
		return nil, err
	}

	if data == nil || len(data) == 0 {
		return nil, fmt.Errorf("get data failed,sql is %s", fmt.Sprintf("select %s from %s where id = %s", fn, tn, id))
	}
	return data[fn], err
}

type CellRenderInfo struct {
	Expr      string      `json:"expr"`
	SheetName string      `json:"sheet_name"`
	CellName  string      `json:"cell_name"`
	Value     interface{} `json:"value"`
}

// GetCellRenderInfos 解析出 excel 中的变量内容
func GetCellRenderInfos(fileTmpStream io.ReadCloser) (*excelize.File, []*CellRenderInfo, error) {
	cris := make([]*CellRenderInfo, 0)
	excelizeFile, err := excelize.OpenReader(fileTmpStream)
	if err != nil {
		return nil, nil, err
	}
	for _, sheet := range excelizeFile.GetSheetList() {
		// 遍历工作表中所有行
		rows, err := excelizeFile.GetRows(sheet)
		if err != nil {
			log.GetLogger(context.Background()).Warn(err.Error())
			return excelizeFile, nil, err
		}

		for rowIdx, row := range rows {
			// 遍历每行中的单元格
			for colIdx, cell := range row {
				if len(cell) >= 4 && cell[:4] == TemplatePrefixFieldName {
					cellName, err := excelize.CoordinatesToCellName(colIdx+1, rowIdx+1)
					if err != nil {
						log.GetLogger(context.Background()).Warn(err.Error())
						return excelizeFile, nil, err
					}
					cellType, err := excelizeFile.GetCellType(sheet, cellName)
					if err != nil {
						return excelizeFile, nil, err
					}
					// 只输出文字类型的单元格
					if cellType == excelize.CellTypeSharedString || cellType == excelize.CellTypeInlineString {
						if err != nil {
							return excelizeFile, nil, err
						}
						cris = append(cris, &CellRenderInfo{
							Expr:      cell,
							SheetName: sheet,
							CellName:  cellName,
						})
					}

				}
			}
		}
	}
	return excelizeFile, cris, nil
}

// GetExprValue getExprValue
func GetExprValue(infos []*CellRenderInfo, indexs map[string]string) error {
	var err error
	for _, info := range infos {
		info.Value, err = DbValue(info.Expr, indexs)
		if err != nil {
			return gerrors.WithStack(err)
		}
	}
	return nil
}

// ReplaceExcelExpr 替换excel 中的表达式为对应的真实数据
func ReplaceExcelExpr(ef *excelize.File, infos []*CellRenderInfo) error {
	for _, info := range infos {
		err := ef.SetCellValue(info.SheetName, info.CellName, info.Value)
		if err != nil {
			return err
		}
	}
	return nil
}

func RenderExcelStream(fileTmpStream io.ReadCloser, indexs map[string]string, renderReceiveWriter io.Writer) error {
	excelFile, infos, err := GetCellRenderInfos(fileTmpStream)
	if err != nil {
		return err
	}
	defer func() {
		if excelFile != nil {
			excelFile.Close()
		}
		if fileTmpStream != nil {
			fileTmpStream.Close()
		}
	}()
	err = GetExprValue(infos, indexs)
	if err != nil {
		return gerrors.WithStack(err)
	}

	err = ReplaceExcelExpr(excelFile, infos)
	if err != nil {
		return gerrors.WithStack(err)
	}

	for _, info := range infos {
		fmt.Println(excelFile.GetCellValue(info.SheetName, info.CellName))
		fmt.Println(info.Value)
	}
	return gerrors.WithStack(excelFile.Write(renderReceiveWriter))
}

type ExprTypeStr string

const (
	Bind                  ExprTypeStr = "BIND"
	BindImage             ExprTypeStr = "BINDIMAGE"
	BindDataExpr          ExprTypeStr = "BINDDataExpr"
	BindExpr              ExprTypeStr = "BINDExpr"
	BINDCollectValuesExpr ExprTypeStr = "BINDCollectValuesExpr"
	BINDRepeat            ExprTypeStr = "BINDRepeat"
	BINDIndex             ExprTypeStr = "BINDIndex"
)

// GetExpr 查找匹配
func GetExpr(expr string) Expr {
	info := ExtractFunctionInfo(expr)
	if info == nil {
		return Expr{
			Type:  Unknown,
			Value: "",
		}
	}
	switch info.Name {
	case string(Bind):
		return Expr{
			Type:  Str,
			Value: info.Params[0],
			Args:  info.Params[1:],
		}
	case string(BINDIndex):
		return Expr{
			Type:  Index,
			Value: info.Params[0],
			Args:  info.Params[1:],
		}
	case string(BindImage):
		return Expr{
			Type:  Img,
			Value: info.Params[0],
			Args:  info.Params[1:],
		}
	case string(BindDataExpr):
		return Expr{
			Type:  DataExpr,
			Value: info.Params[0],
			Args:  info.Params[1:],
		}
	case string(BindExpr):
		return Expr{
			Type:  ExprExpr,
			Value: info.Params[0],
			Args:  info.Params[1:],
		}
	case string(BINDCollectValuesExpr):
		return Expr{
			Type:  CollectValuesDataExpr,
			Value: info.Params[0],
			Args:  info.Params[1:],
		}
	case string(BINDRepeat):
		return Expr{
			Type:  RepeatExpr,
			Value: info.Params[0],
			Args:  info.Params[1:],
		}
	default:
		return Expr{
			Type:  Unknown,
			Value: "",
		}
	}

}

// JsonLookUp 获取json对应字段内容
func JsonLookUp(ctx context.Context, jsonData string, expr Expr, file file.File) (interface{}, error) {
	switch expr.Type {
	case Str:
		raw := gjson.Get(jsonData, expr.Value).Raw
		if strings.HasPrefix(raw, "\"") {
			raw = raw[1:]
		}
		// 去掉结尾的双引号
		if strings.HasSuffix(raw, "\"") {
			raw = raw[:len(raw)-1]
		}
		return raw, nil
	case Img:
		raw := gjson.Get(jsonData, expr.Value).Raw
		if strings.HasPrefix(raw, "\"") {
			raw = raw[1:]
		}
		// 去掉结尾的双引号
		if strings.HasSuffix(raw, "\"") {
			raw = raw[:len(raw)-1]
		}
		download, err := file.Download(ctx, raw)
		if err != nil {
			return "", err
		}
		return download, nil
	default:
		return nil, fmt.Errorf("unsupported type %v", expr.Type)
	}
}

type ExprType int

const (
	Unknown ExprType = -1
	Str     ExprType = iota
	Img
	DataExpr
	ExprExpr
	CollectValuesDataExpr
	RepeatExpr
	Index
)

type Expr struct {
	Type  ExprType
	Value string
	Args  []string
}
