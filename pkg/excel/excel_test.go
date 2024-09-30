package excel

import (
	"github.com/Fordisk123/ginframe/conf"
	"github.com/Fordisk123/ginframe/db"
	"github.com/Fordisk123/ginframe/log"
	"os"
	"testing"
)

func init() {
	conf.InitConf("../../conf")
	log.NewDefaultLogger("example", "v1.0.0")
}

func TestExcel(t *testing.T) {
	db.InitDb()
	file, _ := os.Create("./test1.xlsx")
	defer file.Close()
	f, _ := os.Open("./test.xlsx")
	err := RenderExcelStream(f, map[string]string{
		"user.name":     "1",
		"user.password": "1",
	}, file)
	if err != nil {
		panic(err)
	}

}
