package main

func main() {
	loader := new(commonLoader)

	fr := new(fileReader)
	ar := new(apiReader)

	loader.load(fr)
	loader.load(ar)

	// e := newExcel()
	// e.addSheetFromFile("../localization/BI MSW.BANK.de-de.csv")
	// //e.addSheetsFromDir("../localization", true)
	// e.load(true)
	// e.export("./export.xlsx")
}
