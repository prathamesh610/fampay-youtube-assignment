package main

import (
	"fmt"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/database"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/server"
	"github.com/prathameshj610/fampay-youtube-assignment/internal/service/thirdparty"
)

func main() {

	db, err := database.NewDatabaseClient()

	if err != nil {
		fmt.Printf("failed to initialize db client: %s", err)
	}

	srv := server.NewEchoServer(db)

	thirdparty.InitializeAndAddKeys("AIzaSyCB5H390Q04K9SawzKYBTHVPc9mE4tU200")
	thirdparty.InitializeAndAddKeys("AIzaSyAkcHIVzjLyJ8APtgVuTc0rOdT8Z7bGX28")
	thirdparty.InitializeAndAddKeys("AIzaSyBdvBaLyW79MEzr7AMCCVRiHMv2-r1qXOo")

	if err := srv.Start(); err != nil {
		fmt.Printf(err.Error())
	}

}
