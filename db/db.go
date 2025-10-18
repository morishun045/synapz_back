// db.go

package db

import (
	"fmt"
	"log"
	"os"
	"regexp"
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

/* データベースの初期化 */
func Init() {
	var err error
	// 専用ユーザーでの接続をデフォルトに
	dsn := "synapz_user:synapz_password@tcp(localhost:3306)/synapz_db?charset=utf8mb4&parseTime=True&loc=Local"
	
	// 環境変数が設定されている場合はそちらを優先
	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		dbPass := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")
		if dbName == "" {
			dbName = "synapz_db"
		}
		dsn = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
			dbUser, dbPass, dbName)
	}
	
	log.Printf("Attempting to connect to database with DSN: %s", 
		// パスワード部分を隠して表示
		regexp.MustCompile(`:[^:@]*@`).ReplaceAllString(dsn, ":****@"))
	
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Warning: database cannot connect: %v", err)
		log.Println("Continuing without database connection...")
		log.Println("To connect to MySQL, please ensure:")
		log.Println("  1. MySQL server is running")
		log.Println("  2. Database 'synapz_db' exists")
		log.Println("  3. User 'synapz_user' exists with password 'synapz_password'")
		log.Println("  4. Run setup_db.sql to create the database and user")
		DB = nil
		return
	}
	log.Println("Database connected successfully")
}
