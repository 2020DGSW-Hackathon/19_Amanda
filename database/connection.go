package database

import (
	"Amanda_Server/config"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type connectionMethod interface {
	Connect()
}

var DB *gorm.DB

func Connect() {
	dbConf := config.Config.DB

	connectOptions := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbConf.Host,
		dbConf.Port,
		dbConf.Username,
		dbConf.Name,
		dbConf.Password)

	db, err := gorm.Open("postgres", connectOptions)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&User{},
		&UserStar{},
		&UserComment{},
		&UserReplyComment{},
	)

	db.Model(&UserStar{}).AddForeignKey("fk_user_idx", "users(idx)", "RESTRICT", "RESTRICT")
	db.Model(&UserComment{}).AddForeignKey("fk_user_idx", "users(idx)", "RESTRICT", "RESTRICT")
	db.Model(&UserComment{}).AddForeignKey("fk_object_idx", "users(idx)", "RESTRICT", "RESTRICT")
	db.Model(&UserReplyComment{}).AddForeignKey("comment_idx", "user_comments(idx)", "RESTRICT", "RESTRICT")

	DB = db

	log.Print("[DATABASE] 연결 완료")
}
