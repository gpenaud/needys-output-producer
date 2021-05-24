package producer

import (
  "github.com/jung-kurt/gofpdf"
)

func ProducePdf(filename string) error {
  pdf := gofpdf.New("P", "mm", "A4", "")
  pdf.AddPage()
  pdf.SetFont("Arial", "B", 16)

  // CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
  pdf.CellFormat(190, 7, "This is needys event dumper PDF page", "0", 0, "CM", false, 0, "")
  return pdf.OutputFileAndClose(filename)
}
