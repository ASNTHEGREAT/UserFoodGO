package entity

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var Users = []User {
	{ID: 1, Username: "test@test.com", Password: "testPassword", Firstname: "testfn", Lastname: "testln"},
}

func AddUser(ctx *gin.Context){
	var newUser User
	err := ctx.BindJSON(&newUser)

	if err != nil {	
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User has been added to the database."})
	}
	Users = append(Users, newUser)
}