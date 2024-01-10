package handle

import (
	"context"
	protoData "grpc-client/proto"
	"grpc-client/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
)

type UserController struct {
	userClient protoData.UsersClient
}

func NewUserController(r *gin.Engine, userClient protoData.UsersClient) {
	userHandler := &UserController{userClient}

	r.POST("/user", userHandler.addUser)
	r.GET("/user", userHandler.viewUser)
	r.GET("/user/:id", userHandler.viewUserById)
	r.PUT("/user/:id", userHandler.UpdateUser)
	r.DELETE("/user/:id", userHandler.deleteUser)
}

func (e *UserController) addUser(c *gin.Context) {
	var user protoData.User
	err := c.Bind(&user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if user.Name == "" || user.Email == "" || user.Address == "" {
		utils.HandleError(c, http.StatusBadRequest, "fields are required")
		return
	}

	_, err = e.userClient.InsertUser(context.Background(), &user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSuccess(c, "Sucess Add Data")
}

func (e *UserController) viewUser(c *gin.Context) {
	userList, err := e.userClient.GetUserList(context.Background(), new(empty.Empty))
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSuccess(c, userList.List)
}

func (e *UserController) viewUserById(c *gin.Context) {
	id := c.Param("id")
	userid := protoData.UserId{
		Id: id,
	}
	user, err := e.userClient.GetUserById(context.Background(), &userid)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}

	utils.HandleSuccess(c, user)
}

func (e *UserController) UpdateUser(c *gin.Context) {
	var user protoData.User
	id := c.Param("id")
	err := c.Bind(&user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	var up = protoData.UserUpdate{
		Id:   id,
		User: &user,
	}

	_, err = e.userClient.UpdateUser(context.Background(), &up)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSuccess(c, "Update User Success")
}

func (e *UserController) deleteUser(c *gin.Context) {
	id := c.Param("id")
	var up = protoData.UserId{
		Id: id,
	}

	_, err := e.userClient.DeleteUser(context.Background(), &up)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err.Error())
		return
	}

	utils.HandleSuccess(c, "Success delete data")
}
