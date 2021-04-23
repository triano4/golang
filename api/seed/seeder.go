package seed

import (
	"backend/api/models"
	"log"

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

var navs = []models.Nav{
	models.Nav{
		ID:       "1",
		Name:     "Home",
		ParentId: "0",
		SortId:   1,
		NavIcon:  "",
		NavPath:  "/",
		NavAcc:   "public",
	},
	models.Nav{
		ID:       "2",
		Name:     "User",
		ParentId: "0",
		SortId:   2,
		NavIcon:  "",
		NavPath:  "/Report",
		NavAcc:   "public",
	},
	models.Nav{
		ID:       "3",
		Name:     "Setting",
		ParentId: "0",
		SortId:   3,
		NavIcon:  "",
		NavPath:  "/",
		NavAcc:   "public",
	},
}

//Load function
func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}, &models.Roles{}, &models.Nav{}, &models.Activity{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Roles{}, &models.Nav{}, &models.Activity{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	// err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	// if err != nil {
	// 	log.Fatalf("attaching foreign key error: %v", err)
	// }

	//users for
	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		// posts[i].AuthorID = users[i].ID

		// err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		// if err != nil {
		// 	log.Fatalf("cannot seed posts table: %v", err)
		// }
	}

	for j, _ := range roles {
		err = db.Debug().Model(&models.Roles{}).Create(&roles[j]).Error
		if err != nil {
			log.Fatalf("cannot seed roles table: %v", err)
		}
	}

	for n, _ := range navs {
		err = db.Debug().Model(&models.Nav{}).Create(&navs[n]).Error
		if err != nil {
			log.Fatalf("cannot seed nav table: %v", err)
		}
	}
}
