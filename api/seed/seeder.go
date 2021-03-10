package seed

import (
	"backend/api/models"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "Admin Celerates",
		Email:    "celerate.indonesia@gmail.com",
		Password: "celerate123",
	},
}

var roles = []models.Roles{
	models.Roles{
		UserEmail: "celerate.indonesia@gmail.com",
		Access:    "private",
	},
	models.Roles{
		UserEmail: "luther@gmail.com",
		Access:    "public",
	},
	models.Roles{
		UserEmail: "steven@gmail.com",
		Access:    "public",
	},
}

//Load function
func Load(db *gorm.DB) {

	// err := db.Debug().DropTableIfExists(&models.User{}, &models.Roles{}).Error
	// if err != nil {
	// 	log.Fatalf("cannot drop table: %v", err)
	// }
	// err = db.Debug().AutoMigrate(&models.User{}, &models.Roles{}).Error
	// if err != nil {
	// 	log.Fatalf("cannot migrate table: %v", err)
	// }

	// err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	// if err != nil {
	// 	log.Fatalf("attaching foreign key error: %v", err)
	// }

	//users for
	// for i, _ := range users {
	// 	err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
	// 	if err != nil {
	// 		log.Fatalf("cannot seed users table: %v", err)
	// 	}
	// 	// posts[i].AuthorID = users[i].ID

	// 	// err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
	// 	// if err != nil {
	// 	// 	log.Fatalf("cannot seed posts table: %v", err)
	// 	// }
	// }

	// for j, _ := range roles {
	// 	err = db.Debug().Model(&models.Roles{}).Create(&roles[j]).Error
	// 	if err != nil {
	// 		log.Fatalf("cannot seed roles table: %v", err)
	// 	}
	// }
}
