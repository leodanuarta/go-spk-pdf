package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")

	// ==== HEADER (centered) ====
	pdf.SetHeaderFunc(func() {
		pdf.SetFont("Times", "B", 12)
		// Move cursor to top margin
		pdf.SetY(5)
		// Full width cell with centered text
		pdf.CellFormat(0, 5, "PERJANJIAN KERJASAMA", "", 2, "C", false, 0, "")
		pdf.CellFormat(0, 5, "APOTEK DARING", "", 2, "C", false, 0, "")
		pdf.CellFormat(0, 5, "No. SP/XXX/LGL-PPP/XXX/2025", "", 2, "C", false, 0, "")
		// Line under header
		pdf.Line(20, 20, 190, 20)
		pdf.Ln(10)
	})

	// ==== FOOTER (centered page number) ====
	pdf.SetFooterFunc(func() {
		pdf.SetY(-15) // 15 mm from bottom
		pdf.SetFont("Times", "I", 10)
		pdf.CellFormat(0, 10,
			fmt.Sprintf("Halaman %d", pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})
	pdf.SetMargins(10, 20, 10)
	pdf.AddPage()

	pageW, pageH := pdf.GetPageSize()
	leftM, topM, rightM, bottomM := pdf.GetMargins()

	colGap := 8.0
	colW := (pageW - leftM - rightM - colGap) / 2
	_ = pageH - topM - bottomM // colH not used here because we place manually

	pdf.SetFont("Times", "", 10)

	// ==== Manual lines in LEFT column ====
	xLeft := leftM
	y := topM

	pdf.SetXY(xLeft, y)
	// pdf.MultiCell(colW, 6, "(1) This is the first line in column 1", "", "L", false)
	pdf.MultiCell(colW, 6, "Perjanjian Kerja Sama Apotek Daring ini (selanjutnya disebut", "", "J", false)

	pdf.SetXY(xLeft, pdf.GetY())
	pdf.Write(6, `sebagai "`)

	pdf.SetFont("Times", "B", 10)
	pdf.Write(6, "Perjanjian")

	pdf.SetFont("Times", "", 10)
	pdf.Write(6, `" ) dibuat dan ditandatangani pada hari ini`)

	pdf.SetXY(xLeft, pdf.GetY()+6)
	pdf.MultiCell(colW, 6, "Jumat, tanggal satu bulan Agustus tahun dua ribu dua puluh lima (01-08-2025), oleh dan antara :", "", "J", false) // perhatikan bagian ini.

	// No 1
	pdf.SetXY(xLeft+5, pdf.GetY()+6)
	pdf.Write(6, "1. ")

	pdf.SetXY(xLeft+13, pdf.GetY())
	pdf.SetFont("TIMES", "B", 10)
	pdf.Write(6, "PT Perintis Pelayanan Paripurna")

	pdf.SetFont("TIMES", "", 10)
	pdf.Write(6, ", suatu perseroan")

	pdf.SetXY(xLeft+13, pdf.GetY()+6)
	pdf.MultiCell(colW-15, 6, "terbatas yang didirikan berdasarkan dan tunduk pada hukum Negara Republik Indonesia, beralamat di Grand ITC Permata Hijau Kantor Emerald, Jl.Letjen Soepeno Arteri Permata Hijau E No.26, Grogol Utara, Kebayoran Lama, Kota Jakarta hal ", "", "J", false)

	pdf.SetXY(xLeft+13, pdf.GetY())
	pdf.Write(6, "ini diwakili oleh ")
	pdf.SetFont("TIMES", "B", 10)
	pdf.Write(6, "Yasinta Yulian Hendrata ")
	pdf.SetFont("TIMES", "", 10)
	pdf.Write(6, "selaku")

	pdf.SetXY(xLeft+13, pdf.GetY()+6)
	pdf.SetFont("TIMES", "B", 10)
	pdf.Write(6, "Presiden Direktur ")
	pdf.SetFont("TIMES", "", 10)
	pdf.Write(6, ", oleh karenanya sah bertindak ")

	pdf.SetXY(xLeft+13, pdf.GetY()+6)
	pdf.SetFont("TIMES", "", 10)
	pdf.Write(6, "untuk dan atas nama ")
	pdf.SetFont("TIMES", "B", 10)
	pdf.Write(6, "PT Perintis Pelayanan ")

	pdf.SetXY(xLeft+13, pdf.GetY()+6)
	pdf.SetFont("TIMES", "B", 10)
	pdf.Write(6, "Paripurna ")
	pdf.SetFont("TIMES", "", 10)
	pdf.Write(6, ` (selanjutnya disebut sebagai "`)
	pdf.SetFont("TIMES", "B", 10)
	pdf.Write(6, "PIHAK ")

	pdf.SetXY(xLeft+13, pdf.GetY()+6)
	pdf.SetFont("TIMES", "B", 10)
	pdf.Write(6, "PERTAMA ")
	pdf.SetFont("TIMES", "", 10)
	pdf.Write(6, `").Dan `)

	// No 2
	pdf.SetXY(xLeft+5, pdf.GetY()+6)
	pdf.Write(6, "2. ")

	pdf.SetXY(xLeft+13, pdf.GetY())
	pdf.MultiCell(colW-15, 6, `_____________________, suatu [*] yang didirikan berdasarkan dan tunduk pada hukum Negara Republik Indonesia beralamat di _________________________.Dalam hal ini diwakili oleh _____________________selaku_____, 
	oleh karenanya sah bertindak untuk dan atas nama _________ sebagai pengelola Apotek ______ dengan Nomor Surat Ijin Apotek (SIA) :______________ (selanjutnya disebut sebagai `, "", "J", false)
	// pdf.Write(6, "PT Perintis Pelayanan Paripurna")

	pdf.SetXY(xLeft+13, pdf.GetY())
	pdf.SetFont("TIMES", "", 10)
	pdf.Write(6, `" `)
	pdf.SetFont("TIMES", "B", 10)
	pdf.Write(6, "PIHAK KEDUA ")
	pdf.SetFont("TIMES", "", 10)
	pdf.Write(6, `"). `)

	pdf.SetXY(xLeft, pdf.GetY()+12)
	pdf.SetFont("TIMES", "B", 10)
	pdf.Write(6, "PIHAK PERTAMA ")
	pdf.SetFont("TIMES", "", 10)
	pdf.Write(6, `dan `)
	pdf.SetFont("TIMES", "B", 10)
	pdf.Write(6, "PIHAK KEDUA ")
	pdf.SetFont("TIMES", "", 10)
	pdf.Write(6, `secara bersama-sama`)

	pdf.SetXY(xLeft, pdf.GetY()+6)
	pdf.SetFont("TIMES", "", 10)
	pdf.Write(6, `disebut sebagai "`)
	pdf.SetFont("TIMES", "B", 10)
	pdf.Write(6, "PARA PIHAK ")
	pdf.SetFont("TIMES", "", 10)
	pdf.Write(6, `" dan secara masing-masing`)

	pdf.SetXY(xLeft, pdf.GetY()+6)
	pdf.SetFont("TIMES", "", 10)
	pdf.Write(6, `disebut sebagai "Pihak". `)
	pdf.SetFont("TIMES", "B", 10)
	pdf.Write(6, "PARA PIHAK ")
	pdf.SetFont("TIMES", "", 10)
	pdf.Write(6, `sepakat untuk `)

	pdf.SetXY(xLeft, pdf.GetY()+6)
	pdf.SetFont("TIMES", "", 10)
	pdf.Write(6, `Perjanjian ini dengan ketentuan dan syarat-syarat sebagai berikut :`)

	pdf.SetXY(xLeft, pdf.GetY()+12)
	pdf.SetFont("TIMES", "B", 10)
	pdf.MultiCell(colW, 6, "PASAL", "", "C", false)
	pdf.MultiCell(colW, 6, "PENDAHULUAN", "", "C", false)





	// ==== Manual lines in RIGHT column ====
	xRight := leftM + colW + colGap
	y = topM

	pdf.SetXY(xRight, y)
	pdf.MultiCell(colW, 6, "(4) This is the first line in column 2", "", "L", false)

	pdf.SetXY(xRight, pdf.GetY())
	pdf.MultiCell(colW, 6, "(5) Another line in column 2", "", "L", false)

	pdf.SetXY(xRight, pdf.GetY())
	pdf.MultiCell(colW, 6, "(6) And another line in column 2", "", "L", false)

	// Save PDF
	err := pdf.OutputFileAndClose("pdf_with_centered_header.pdf")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
