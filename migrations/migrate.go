package migrations

import (
	"authGin/initializers"
	"authGin/models"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
