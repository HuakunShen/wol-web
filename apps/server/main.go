package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "server/migrations"

	wol "github.com/HuakunShen/wol/wol-go"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()
	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Dashboard
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), true))
		// se.Router.GET("/auth", apis.Static(os.DirFS("./pb_public/auth.html"), false))
		se.Router.GET("/api/hosts", func(e *core.RequestEvent) error {
			return e.JSON(200, map[string]string{"message": "Host waked"})
		})
		se.Router.POST("/api/wake", func(e *core.RequestEvent) error {
			data := struct {
				// unexported to prevent binding
				Id string `json:"id" form:"id"`
			}{}
			if err := e.BindBody(&data); err != nil {
				return e.BadRequestError("Failed to read request data", err)
			}
			info, err := e.RequestInfo()
			if err != nil {
				return e.BadRequestError("Failed to read request info", err)
			}
			id, ok := info.Body["id"].(string)
			if !ok {
				return e.BadRequestError("Failed to read id from body", err)
			}

			isAuthenticated := e.Auth != nil
			if !isAuthenticated {
				return e.JSON(401, map[string]string{"message": "Unauthorized"})
			}
			userId := e.Auth.Id
			requestedHost, err := app.FindRecordById("hosts", id)
			if err != nil {
				return e.BadRequestError("Failed to find host", err)
			}
			// Get the user field - it will return an interface{} that you can type assert
			recordOwnerId := requestedHost.GetString("user")
			if recordOwnerId != userId {
				return e.JSON(401, map[string]string{"message": "Unauthorized"})
			}
			macValue := requestedHost.GetString("mac")
			portValue := requestedHost.GetInt("port")
			fmt.Println("portValue", portValue)
			if !ok {
				return e.JSON(400, map[string]string{"message": "Failed to parse port"})
			}
			err = wol.WakeOnLan(macValue, "255.255.255.255", strconv.Itoa(portValue))
			if err != nil {
				return e.JSON(400, map[string]string{"message": "Failed to wake host", "error": err.Error()})
			}
			return e.JSON(200, map[string]string{"message": "WakeOnLan Magic Packet Sent"})
		})
		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
