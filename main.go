package main

import (
	`fmt`
	`log`
	`os`
	`path/filepath`
	`strings`

	`github.com/pocketbase/pocketbase`
	`github.com/pocketbase/pocketbase/core`
	`github.com/pocketbase/pocketbase/models`
	`github.com/pocketbase/pocketbase/plugins/jsvm`
	`github.com/pocketbase/pocketbase/plugins/migratecmd`

	`github.com/lgnixai/sugua/collection`
	`github.com/lgnixai/sugua/commands`
	`github.com/lgnixai/sugua/crontab`
	_ "github.com/lgnixai/sugua/data/migrations"
	`github.com/lgnixai/sugua/handlers`
	`github.com/lgnixai/sugua/hooks`
	`github.com/lgnixai/sugua/pkg/env`
	`github.com/lgnixai/sugua/pkg/watcher`
	`github.com/lgnixai/sugua/utils`
	worker `github.com/lgnixai/sugua/work`
)

func defaultPublicDir() string {
	if strings.HasPrefix(os.Args[0], os.TempDir()) {
		// most likely ran with go run
		return "./public"
	}

	return filepath.Join(os.Args[0], "../public")
}

func init() {
	env.Init()

	go func() {
		surrealErr := utils.CMD("task", "surrealdb")
		if surrealErr != nil {
			fmt.Println("Error downloading keck task: ", surrealErr)
		}
	}()
	go func() {
		kechErr := utils.CMD("task", "keck")
		if kechErr != nil {
			fmt.Println("Error downloading keck task: ", kechErr)
		}
	}()

	fmt.Println("running keck task")
}

func main() {
	app := pocketbase.NewWithConfig(pocketbase.Config{
		DefaultDataDir: "./data/data",
	})

	worker.InitWorker(app)
	commands.AddCommands(app)
	//var publicDirFlag string
	crontab.RegisterCronJobs(app)
	handlers.RegisterHandlers(app)
	hooks.RegisterHooks(app)
	// add "--publicDir" option flag
	//app.RootCmd.PersistentFlags().StringVar(
	//	&publicDirFlag,
	//	"publicDir",
	//	defaultPublicDir(),
	//	"the directory to serve static files",
	//)
	migrationsDir := "./data/migrations" // default to "pb_migrations" (for js) and "migrations" (for go)

	// load js files to allow loading external JavaScript migrations
	jsvm.MustRegister(app, jsvm.Config{
		// Dir: migrationsDir,
		MigrationsDir: migrationsDir,
	})
	collections := []*models.Collection{
		collection.BookmarksCollection(),
		collection.FeedCollection(),
		collection.ItemCollection(),
		collection.FeedSizeCollection(),
		collection.FeedErrorCollection(),
		collection.HttpStatusCollection(),
	}

	// manually declare schemas
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		for _, collection := range collections {
			existing, _ := app.Dao().FindCollectionByNameOrId(collection.Name)

			if existing == nil {
				if err := app.Dao().SaveCollection(collection); err != nil {
					log.Fatal("[OnBeforeServe]: %w", err)
				}
			}
		}

		return nil
	})

	// register the `migrate` command
	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		//TemplateLang: migratecmd.TemplateLangJS, // or migratecmd.TemplateLangGo (default)
		Dir:         migrationsDir,
		Automigrate: true,
	})

	//app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
	//	// serves static files from the provided public dir (if exists)
	//	e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS(publicDirFlag), true))
	//
	//	return nil
	//})
	//

	//
	//app.OnRecordBeforeUpdateRequest().Add(func(e *core.RecordUpdateEvent) error {
	//	switch e.Collection.Name {
	//	case "rollouts":
	//		return controller.HandleRolloutUpdate(e, app)
	//	}
	//	return nil
	//})
	//
	//app.OnRecordBeforeDeleteRequest().Add(func(e *core.RecordDeleteEvent) error {
	//	switch e.Collection.Name {
	//	case "rollouts":
	//		return controller.HandleRolloutDelete(e, app)
	//	case "deployments":
	//		return controller.HandleDeploymentDelete(e, app)
	//	case "projects":
	//		return controller.HandleProjectDelete(e, app)
	//	}
	//	return nil
	//})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {

		// get status of a specific deployment

		//e.Router.GET("/pb/:projectId/:deploymentId/metrics", func(c echo.Context) error {
		//	projectId := c.PathParam("projectId")
		//	deploymentId := c.PathParam("deploymentId")
		//
		//	return controller.HandleDeploymentMetrics(c, app, projectId, deploymentId)
		//}, apis.RequireRecordAuth("users"))
		//
		//e.Router.GET("/pb/:projectId/:deploymentId/events", func(c echo.Context) error {
		//	projectId := c.PathParam("projectId")
		//	deploymentId := c.PathParam("deploymentId")
		//
		//	return controller.HandleDeploymentEvents(c, app, projectId, deploymentId)
		//}, apis.RequireRecordAuth("users"))
		//
		//e.Router.GET("/pb/:projectId/:podName/logs", func(c echo.Context) error {
		//	projectId := c.PathParam("projectId")
		//	podName := c.PathParam("podName")
		//
		//	return k8s.GetRolloutLogs(c.Response().Writer, projectId, podName)
		//}, apis.RequireRecordAuth("users"))
		//
		//e.Router.GET("/pb/blueprints/:blueprintId", func(c echo.Context) error {
		//	blueprintId := c.PathParam("blueprintId")
		//
		//	return controller.HandleBlueprint(c, app, blueprintId)
		//}, apis.RequireRecordAuth("users"))
		//
		//e.Router.POST("/pb/blueprints/shared/:blueprintId", func(c echo.Context) error {
		//	blueprintId := c.PathParam("blueprintId")
		//
		//	return controller.HandleBlueprintAdd(c, app, blueprintId)
		//}, apis.RequireRecordAuth("users"))
		//
		//e.Router.POST("/pb/auto-update/:autoUpdateId", func(c echo.Context) error {
		//	// TODO: change the auth to be a token generated by the user
		//	autoUpdateId := c.PathParam("autoUpdateId")
		//
		//	return controller.HandleAutoUpdate(c, app, autoUpdateId)
		//})
		//
		//e.Router.GET("/pb/cluster-info", func(c echo.Context) error {
		//	return controller.HandleClusterInfo(c, app)
		//	// }, apis.RequireRecordAuth("users"))
		//})
		//
		//// delete a pod of a rollout by pod name
		//e.Router.DELETE("/pb/:projectId/:podName", func(c echo.Context) error {
		//	projectId := c.PathParam("projectId")
		//	podName := c.PathParam("podName")
		//
		//	return controller.HandlePodDelete(c, app, projectId, podName)
		//}, apis.RequireRecordAuth("users"))
		//
		//// websocket for deployments status
		//e.Router.GET("/ws/k8s/deployments", watcher.WsK8sDeploymentsHandler)
		//
		//// websocket for pod logs
		//e.Router.GET("/ws/k8s/logs", watcher.WsK8sRolloutLogsHandler)

		// websocket for rollout events
		e.Router.GET("/ws/k8s/events", watcher.WsK8sRolloutEventsHandler)

		return nil
	})

	//app.OnAfterBootstrap().Add(func(e *core.BootstrapEvent) error {
	//	scheduler := cron.New()
	//
	//	// Run image update every minute
	//	scheduler.MustAdd("autoUpdate", env.Config.CronTick, func() {
	//		fmt.Println("Running auto update")
	//		//err := controller.AutoUpdateController(app)
	//		//if err != nil {
	//		//	log.Printf("Error updating image: %v\n", err)
	//		//}
	//	})
	//
	//	scheduler.Start()
	//	return nil
	//})

	//app.RootCmd.AddCommand(&cobra.Command{
	//	Use:   "generate-music-urls",
	//	Short: "Generate music urls",
	//	Run: func(cmd *cobra.Command, args []string) {
	//		utils.ParseMusicListToUrls("./assets/reference/musics.json")
	//	},
	//})

	app.Bootstrap()
	//serveCmd := cmd.NewServeCommand(app, false)
	//serveCmd.Execute()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
