package util

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

func ShowPeers(table *tablewriter.Table) {
	table.Append([]string{time.Now().String(), time.Now().String(), time.Now().String()})
	table.Render()

}

func ConfigTable() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"IP", "PORT", "ID"})

	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
	)

	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor})
	return table
}
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
