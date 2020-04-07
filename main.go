package main

import (
	"com.mark/ginFirst/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_=router.Run()
}
