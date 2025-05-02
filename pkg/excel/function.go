package excel

import (
	"strings"
	"unicode"
)

type FunctionInfo struct {
	Name   string
	Params []string
}

func ExtractFunctionInfo(s string) *FunctionInfo {
	// 阶段1：提取函数名和参数部分
	name, paramsPart, ok := splitFunctionParts(s)
	if !ok {
		return nil
	}

	// 阶段2：解析参数列表
	params := parseParams(paramsPart)

	return &FunctionInfo{
		Name:   name,
		Params: params,
	}
}

// 分离函数名和参数部分（支持嵌套括号）
func splitFunctionParts(s string) (name, params string, ok bool) {
	// 寻找第一个 '('
	start := strings.IndexByte(s, '(')
	if start == -1 || start == 0 {
		return "", "", false
	}

	// 验证函数名
	for _, r := range s[:start] {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			return "", "", false
		}
	}

	// 匹配括号层级
	level := 1
	end := -1
	for i := start + 1; i < len(s); i++ {
		switch s[i] {
		case '(':
			level++
		case ')':
			level--
			if level == 0 {
				end = i
				goto found
			}
		}
	}
found:
	if end == -1 {
		return "", "", false
	}

	return s[:start], s[start+1 : end], true
}

// 解析带嵌套结构的参数列表
func parseParams(paramsStr string) []string {
	var params []string
	var buffer strings.Builder
	quoteOpen := false
	parenLevel := 0

	for _, r := range paramsStr {
		switch {
		case r == '"':
			quoteOpen = !quoteOpen
			buffer.WriteRune(r)
		case !quoteOpen && r == '(':
			parenLevel++
			buffer.WriteRune(r)
		case !quoteOpen && r == ')':
			parenLevel--
			buffer.WriteRune(r)
		case !quoteOpen && parenLevel == 0 && r == ',':
			params = append(params, processParam(buffer.String()))
			buffer.Reset()
		default:
			buffer.WriteRune(r)
		}
	}

	// 处理最后一个参数
	if buffer.Len() > 0 {
		params = append(params, processParam(buffer.String()))
	}

	return params
}

// 清理参数格式
func processParam(p string) string {
	p = strings.TrimSpace(p)
	if len(p) >= 2 && p[0] == '"' && p[len(p)-1] == '"' {
		return p[1 : len(p)-1]
	}
	return p
}

//func main() {
//	testCase := `BINDCollectValuesExpr("experimentData.#(TorqueType==低扭矩)#|#(LoadDirection==顺时针)#.CollectValues","5")`
//
//	fmt.Printf("测试字符串: %s\n", testCase)
//	if result := ExtractFunctionInfo(testCase); result != nil {
//		fmt.Println("解析结果:")
//		fmt.Printf("函数名: %s\n", result.Name)
//		fmt.Println("参数列表:")
//		for i, p := range result.Params {
//			fmt.Printf("[%d] %s\n", i, p)
//		}
//	} else {
//		fmt.Println("解析失败")
//	}
//}
