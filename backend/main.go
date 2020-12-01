package main
 
import (
   "context"
   "log"
 
   "github.com/Piichet/app/controllers"
   _ "github.com/Piichet/app/docs"
   "github.com/Piichet/app/ent"
   "github.com/gin-contrib/cors"
   "github.com/gin-gonic/gin"
   _ "github.com/mattn/go-sqlite3"
   swaggerFiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
)

type User struct {
	Name     string
	Email    string
	Password string
}
type Users struct {
	User []User
}
type Titles struct {
	Title []Title
}
type Title struct {
	Title string
}
type Positions struct {
	Position []Position
}
type Position struct {
	Position string
}
type Genders struct {
	Gender []Gender
}
type Gender struct {
	Gender string
}
// @title SUT SA Example API Playlist Vidoe
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:user.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
 
   v1 := router.Group("/api/v1")
   controllers.NewUserController(v1, client)
   controllers.NewTitleController(v1, client)
   controllers.NewGenderController(v1, client)
   controllers.NewPositionController(v1, client)
   
   // Set title
	titles := Titles{
		Title: []Title{
			Title{"เภสัชกรชาย"},
			Title{"เภสัชกรหญิง"},
		},
	}

	for _, t := range titles.Title {
		client.Title.
			Create().
			SetTitle(t.Title).
			Save(context.Background())
	}

	// Set Gender
	genders := Genders{
		Gender: []Gender{
			Gender{"ชาย"},
			Gender{"หญิง"},
		},
	}

	for _, g := range genders.Gender {
		client.Gender.
			Create().
			SetGender(g.Gender).
			Save(context.Background())
	}

	// Set Position
	positions := Positions{
		Position: []Position{
			Position{"เบิกยาจากคลังสำหรับห้องยา"},
			Position{"ลงทะเบียนยาเข้าคลัง"},
			Position{"บันทึกการจ่ายยา"},
			Position{"บันทึกข้อมูลยา"},
			Position{"บันทึกประวัติผู้ป่วย"},
		},
	}

	for _, p := range positions.Position {
		client.Position.
			Create().
			SetPosition(p.Position).
			Save(context.Background())
	}

   router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
   router.Run()
}