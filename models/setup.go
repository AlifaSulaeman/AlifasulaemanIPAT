package models
import(
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB
func ConnectDatabase ()  {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_retapi_gin"), &gorm.Config{})
	if err != nil{
		panic(err)
	}

	database.AutoMigrate(&Users{})

	DB = database
}