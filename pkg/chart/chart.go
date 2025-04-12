package chart

import (
	_ "embed"
	"fmt"
	"github.com/vicanso/go-charts/v2"
	"io"
)

const (
	Png = "PNG"
	Svg = "SVG"
)

//go:embed wqy-microhei.ttc
var fontData []byte

type PicChart interface {
	GenChart(title string, XLabel, YLabel string, data interface{}, format string) (io.Reader, error)
}

type BasePicChart struct {
	FontPath string
}

func (bpc BasePicChart) LoadFont() error {
	//// Path to the Chinese font file
	//fontPath := "./resource/wqy-microhei.ttc"
	//
	//// Read the font file
	//fontData, err := os.ReadFile(fontPath)
	//if err != nil {
	//	return fmt.Errorf("failed to read font file: %v", err)
	//}

	// Install the font with a name
	err := charts.InstallFont("chinese", fontData)
	if err != nil {
		return fmt.Errorf("failed to install font: %v", err)
	}

	// Get the installed font
	font, err := charts.GetFont("chinese")
	if err != nil {
		return fmt.Errorf("failed to get installed font: %v", err)
	}

	// Set it as the default font
	charts.SetDefaultFont(font)

	return nil

}
