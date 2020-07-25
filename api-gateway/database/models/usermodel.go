package models

import (
	"database/sql"
	"khanhnguyen234/api-gateway/common"
	"khanhnguyen234/api-gateway/database/items"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserPostForm struct {
	Username string `form:"Username" json:"username" xml:"Username" binding:"required"`
	Password string `form:"Password" json:"password" xml:"Password" binding:"required"`
}

const (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbname   = "go_gateway"
)

func connectDB() (db *sql.DB, err error) {

	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	db, err = sql.Open("postgres", conn)
	if err != nil {
		fmt.Printf("Fail to openDB: %v \n", err)
		return nil, err

	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Fail to conenct: %v \n", err)
		return nil, err
	}

	fmt.Println("Ping OK")
	return db, nil
}
// GetInfoUser Get info user with Username
func GetInfoUser(c *gin.Context) {
	var input UserPostForm
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Connect DB
	db, err := connectDB()
	if err != nil {
		return
	}

	//Get info
	sql := `SELECT username, password FROM account.user WHERE username = ` + fmt.Sprintf("'%v'", input.Username) + ` AND password = ` + fmt.Sprintf("'%v'", input.Password)
	fmt.Println(sql)

	row, err := db.Query(sql)

	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"messages": "Fail",
		})
		return
	}

	_infoUser := items.User{}
	var Username string
	var Password string
	for row.Next() {
		row.Scan(&Username, &Password)
	}
	fmt.Println(row)

	if Username == "" {
		c.JSON(501, gin.H{
			"messages": common.MsgLoginError,
		})
		return
	}
	_infoUser.TypeUsername = Username
	_infoUser.TypePassword = Password

	defer db.Close()

	c.JSON(200, _infoUser)
}

// InsertUser Insert info new user
func InsertUser(c *gin.Context) {
	var input UserPostForm
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := connectDB()
	if err != nil {
		return
	}

	sql := `INSERT INTO account.user(username, password)
		VALUES (` + fmt.Sprintf("'%v'", input.Username) + `,` + fmt.Sprintf("'%v'", input.Password) + `);`
		fmt.Println(sql)

	if _, err = db.Exec(sql); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"messages": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": "Insert Success",
	})

	defer db.Close()
}