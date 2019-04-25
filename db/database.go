package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type PlayerInfo struct {
	ID   string
	PosX float64
	PosY float64
}

var Database *sql.DB

type DbConfig struct {
	DbHost string
	DbPort string
	DbName string
	DbUser string
	DbPwd  string
}

var DbCfg DbConfig = DbConfig{
	DbHost: "0.0.0.0",
	DbPort: "5432",
	DbName: "postgres",
	DbUser: "postgres",
	DbPwd:  "1Qh0RjfU7T!",
}

func InitDB() (err error) {

	connInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DbCfg.DbHost, DbCfg.DbPort, DbCfg.DbUser, DbCfg.DbPwd, DbCfg.DbName)

	Database, err = sql.Open("postgres", connInfo)
	if err != nil {
		return err
	}
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(i) * time.Second)
		if err = Database.Ping(); err == nil {
			break
		}
		return err
	}

	fmt.Println("Successfully connected to database")

	return nil
}

func CreateDbTable(DB *sql.DB) error {

	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS "public"."users"(
		"username" character varying(50) NOT NULL UNIQUE,
		"password" TEXT NOT NULL,
		"id" character varying(50) UNIQUE,
		"playername" character varying(50),
		"positionX" float,
		"positionY" float,
		"world" character varying(50)
	)`)
	if err != nil {

		return err

	}
	fmt.Println("Players Table Created")

	return nil
}

func DeleteDbTable(DB *sql.DB) error {
	_, err := DB.Exec(`DROP TABLE users`)
	if err != nil {
		fmt.Println(err)
		return (err)
	}
	fmt.Println("Deleted Users")

	_, err = DB.Exec(`DROP TABLE inventory`)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Deleted inventories")
	return nil
}

func DeleteInventory(DB *sql.DB, ID string) error {

	query := `DELETE FROM "public"."inventory" WHERE playerid = $1`
	_, err := DB.Exec(query, ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Inventory of playerid " + ID + " deleted")
	return nil
}

func RegisterPlayer(DB *sql.DB, username string, password string, ID string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		fmt.Println("Problem with hashing password")
		return err
	}

	query := `INSERT INTO users(username,password,id) VALUES ($1,$2,$3)`

	_, err = DB.Exec(query, username, hashedPassword, ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Player registered")
	return nil
}

func LoginPlayer(DB *sql.DB, username string, password string) (authorized bool, player PlayerInfo) {

	var hashed_pass string
	fmt.Println(player)
	err := DB.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&hashed_pass)
	if err != nil {
		fmt.Println("Could not parse password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashed_pass), []byte(password))
	if err != nil {
		fmt.Println("Unauthorized")
		return false, PlayerInfo{}
	}
	fmt.Println("Logged in")

	err = DB.QueryRow("SELECT id FROM users WHERE username = $1", username).Scan(&player.ID)
	if err != nil {
		fmt.Println("Could not parse id")
		return false, PlayerInfo{}
	}
	err = DB.QueryRow("SELECT positionX FROM users WHERE username = $1", username).Scan(&player.PosX)
	if err != nil {
		fmt.Println("Could not parse positionX")
		return false, PlayerInfo{}
	}
	err = DB.QueryRow("SELECT id FROM users WHERE username = $1", username).Scan(&player.PosY)
	if err != nil {
		fmt.Println("Could not parse positionY")
		return false, PlayerInfo{}
	}

	return true, player

}

func CheckIfValid(DB *sql.DB) {

}

func CreateInventoryTable(DB *sql.DB) error {

	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS "public"."inventory"(
		"playerid" character varying(50) NOT NULL,
		"slot" character varying(20)
		
	)`)
	if err != nil {
		return err
	}
	fmt.Println("Inventory Created")

	return nil
}

func AddNewInventorySlot(DB *sql.DB, ID string) error {
	query := `INSERT INTO inventory(playerid,slot) VALUES ($1, $2)`

	_, err := DB.Exec(query, ID, "")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}