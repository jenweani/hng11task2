package main

import "hng11task2/api"

func main() {

	srv := api.NewServer(8080, api.BuildRoutesHandler())
	srv.Listen()
}