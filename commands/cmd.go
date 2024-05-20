package commands

import (
	`log`
	`os`

	`github.com/pocketbase/pocketbase`
	`github.com/spf13/cobra`

	`github.com/lgnixai/sugua/export`
)

var (
	configDir    = "."
	buildersPath = "./bin"
)

func AddCommands(app *pocketbase.PocketBase) {
	var filename string
	//var runFlags taskfile.RunFlags
	exportCmd := &cobra.Command{
		Use:       "export",
		Short:     "Exports the collection to JSON on STDOUT",
		ValidArgs: []string{"filename"},
		Run: func(cmd *cobra.Command, args []string) {
			collectionFile, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
			if err != nil {
				log.Fatal(err)
			}
			defer collectionFile.Close()
			err = export.ExportCollections(app, collectionFile)
			if err != nil {
				panic(err)
			}
		},
	}

	exportCmd.Flags().StringVarP(&filename, "output", "o", "collections.json", "Export to file")

	importCmd := &cobra.Command{
		Use:       "import",
		Short:     "Imports the collection from JSON",
		ValidArgs: []string{"filename"},
		Run: func(cmd *cobra.Command, args []string) {
			err := export.ImportCollections(app, filename)
			if err != nil {
				panic(err)
			}
		},
	}
	importCmd.Flags().StringVarP(&filename, "from", "f", "collections.json", "Import from file")

	runCmd := &cobra.Command{
		Use:   "run [environment] [task...]",
		Short: "Run taskfile.dev tasks",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			//doubleDashPos := cmd.ArgsLenAtDash()
			//if err := taskfile.Run(configDir, buildersPath, server, runFlags, args[0], doubleDashPos, args[1:]); err != nil {
			//	return err
			//}
			return nil
		},
	}
	//runCmd.Flags()..(&runFlags, "from", "f", "collections.json", "Import from file")

	app.RootCmd.AddCommand(exportCmd)
	app.RootCmd.AddCommand(importCmd)
	app.RootCmd.AddCommand(runCmd)
}
