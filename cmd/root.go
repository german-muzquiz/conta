package cmd

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "conta",
	Short: "Conta es una herramienta para hacer calculos de declaraciones del SAT",
	Long: `Conta es una herramienta para hacer calculos de declaraciones del SAT.
Permite escanear XML de facturas y calcular las retenciones y deducciones personales.`,
	Run: run,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(_ *cobra.Command, _ []string) {
	box := tview.NewBox().
		SetBorder(true).
		SetTitle("Hello world!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
