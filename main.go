package main

func main() {
	e := newExcel()
	e.addSheetFromFile("../localization/BI MSW.BANK.de-de.csv")
	// e.addSheetsFromDir("../localization")
	e.export("./export.xlsx")
}
