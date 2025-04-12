package captcha

import (
	"github.com/Fordisk123/ginframe/pkg/chart"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestChart(t *testing.T) {

	data := map[int64]float64{
		0:  12,
		1:  123.2,
		2:  -123.2,
		6:  10000,
		4:  12,
		3:  13.2,
		5:  1000,
		7:  133.2,
		8:  -113.2,
		9:  113.2,
		10: 1223.2,
		11: 143.2,
		12: 153.2,
		13: 113.2,
		15: 163.2,
		14: 173.2,
		16: 193.2,
		17: 133.2,
		18: 113.2,
		19: 183.2,
		20: 143.2,
	}

	chartReader, err := (&chart.LineChart{}).GenChart("动态扭矩", "ms", "Nj", data, chart.Png)
	if err != nil {
		t.Fatal(err)
	}

	all, err := io.ReadAll(chartReader)

	writeFile(all)

}

func writeFile(buf []byte) error {
	tmpPath := "./tmp"
	err := os.MkdirAll(tmpPath, 0700)
	if err != nil {
		return err
	}

	file := filepath.Join(tmpPath, "time-line-chart.png")
	err = os.WriteFile(file, buf, 0600)
	if err != nil {
		return err
	}
	return nil
}
