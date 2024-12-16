package utils

import (
	models "backend/internal/models/common"
	"bytes"
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func GenerateShippingLabelPDF(fromAddr, toAddr *models.Address, festivalName, attendeeFullName, attendeePhoneNumber string) ([]byte, error) {

	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		OrientationStr: gofpdf.OrientationPortrait,
		UnitStr:        "mm",
		Size: gofpdf.SizeType{
			Wd: 100,
			Ht: 150,
		},
	})
	pdf.AddUTF8Font("DejaVu", "", "DejaVuSansCondensed.ttf")
	pdf.AddPage()

	pdf.Ln(2)

	pdf.SetFont("DejaVu", "", 12)
	pdf.MultiCell(0, 8, fmt.Sprintf(
		"FROM:\n%s\n%s %s\n%s, %s\n%s, %s",
		festivalName,
		fromAddr.Street,
		fromAddr.Number,
		fromAddr.City.Name,
		fromAddr.City.PostalCode,
		fromAddr.City.Country.Name,
		fromAddr.City.Country.ISO,
	), "1", "L", false)
	pdf.Ln(2)

	pdf.MultiCell(0, 8, fmt.Sprintf(
		"SHIP TO:\n%s\n%s\n%s %s\n%s, %s\n%s, %s\n",
		attendeeFullName,
		attendeePhoneNumber,
		toAddr.Street,
		toAddr.Number,
		toAddr.City.Name,
		toAddr.City.PostalCode,
		toAddr.City.Country.Name,
		toAddr.City.Country.ISO,
	), "1", "L", false)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
