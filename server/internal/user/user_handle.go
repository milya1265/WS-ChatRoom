package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type handler struct {
	Service Service
}

func NewHandler(s *Service) Handler {
	return &handler{Service: *s}
}

func (h *handler) CreateUser(c *gin.Context) {
	var u CreateUserReq

	if err := c.BindJSON(&u); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Service.CreateUser(c.Request.Context(), &u)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user":    user,
		"message": "sign up success",
	})
}

func (h *handler) Login(c *gin.Context) {
	var loginUser LoginUserReq
	err := c.BindJSON(&loginUser)
	if err != nil {
		log.Println("ERROR --> scan row:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.Service.Login(c.Request.Context(), &loginUser)
	if err != nil {
		log.Println("ERROR --> scan row:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("jwt", u.accessToken, 60*60*24, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "sign-in success", "id": u.ID})
}

func (h *handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout success"})
}
