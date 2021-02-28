package main

func main() {
	loader := new(CommonLoader)
	ew := new(CSVWriter)

	fn := "../localization/BI MSW.BANK.de-de.csv"
	fr := newFileReader(fn)

	//ar := newAPIReader("ACHLA URL")

	// bf := loader.load(fr)
	buff := loader.load(fr)
	//loader.load(ar)

	loader.write(ew, buff)

	// e := newExcel()
	// e.addSheetFromFile("../localization/BI MSW.BANK.de-de.csv")
	// //e.addSheetsFromDir("../localization", true)
	// e.load(true)
	// e.export("./export.xlsx")
}
