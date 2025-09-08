package main

import (
	"github.com/Tweenitwits/easy_rest_api_service.git/internal/config"
)

func main() {

	cfg := config.MustLoad()

	// TODO: init config: cleanenv

	//TODO: init logger: slog

	//TODO: init storage: sqlite

	//TODO: init router: chi, "chi render"

	//TODO: run server
}
