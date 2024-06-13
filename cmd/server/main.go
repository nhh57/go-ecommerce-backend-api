package main

import (
	"github.com/nhh57/go-ecommerce-backend-api/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run(":8083") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
