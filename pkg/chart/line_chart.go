package chart

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/vicanso/go-charts/v2"
	"io"
	"sort"
)

type LineChart struct {
	BasePicChart
	// ChannelData map[string] 为不同隧道 map[int64]float64 int64 为时间点，float64 为对应的y轴
	ChannelData map[int64]float64
}

func (lc *LineChart) GenChart(title string, XLabel, YLabel string, data interface{}, format string) (io.Reader, error) {
	err := lc.LoadFont()
	if err != nil {
		return nil, err
	}
	md, ok := data.(map[int64]float64)
	if !ok {
		return nil, errors.New("line chart data is not map[string]map[int64]float64")
	}

	tss := make([]int64, 0, len(md))
	for k, _ := range md {
		tss = append(tss, k)
	}
	sort.Slice(tss, func(i, j int) bool {
		return tss[i] < tss[j]
	})

	values := make([]float64, 0)
	for _, ts := range tss {
		values = append(values, md[ts])
	}

	xs := make([]string, 0)

	for _, ts := range tss {
		xs = append(xs, fmt.Sprintf("%v%s", ts, XLabel))
	}

	p, err := charts.LineRender(
		[][]float64{
			values,
		},
		charts.PaddingOptionFunc(charts.Box{
			Top:    30,
			Left:   30,
			Right:  30,
			Bottom: 30,
			IsSet:  true,
		}),
		func() charts.OptionFunc {
			if format == Svg {
				return charts.SVGTypeOption()
			} else {
				return charts.PNGTypeOption()
			}
		}(),
		charts.TitleTextOptionFunc(title),
		charts.XAxisDataOptionFunc(xs, charts.FalseFlag()),
		charts.YAxisOptionFunc(charts.YAxisOption{
			Formatter: "{value} " + YLabel,
		}),
		func(opt *charts.ChartOption) {
			//opt.Title.Top = "-20"
			opt.XAxis.SplitNumber = 5
			opt.XAxis.TextRotation = 0
			opt.SymbolShow = charts.FalseFlag()
			opt.LineStrokeWidth = 1
			opt.ValueFormatter = func(f float64) string {
				return fmt.Sprintf("%.0f", f)
			}
		},
	)
	buf, err := p.Bytes()
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(buf[:]), nil
}
