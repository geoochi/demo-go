package handlers

import (
	"net/http"
	"strconv"

	"demo/internal/models"
	"github.com/gin-gonic/gin"
)

// 模拟数据库
var users = make(map[uint]*models.User)
var nextID uint = 1

// CreateUser handles POST /users
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = nextID
	users[user.ID] = &user
	nextID++

	c.JSON(http.StatusCreated, user)
}

// GetUser handles GET /users/:id
func GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if user, exists := users[uint(id)]; exists {
		c.JSON(http.StatusOK, user)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// ListUsers handles GET /users
func ListUsers(c *gin.Context) {
	userList := make([]*models.User, 0, len(users))
	for _, user := range users {
		userList = append(userList, user)
	}
	c.JSON(http.StatusOK, userList)
}

// UpdateUser handles PUT /users/:id
func UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if _, exists := users[uint(id)]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = uint(id)
	users[uint(id)] = &user
	c.JSON(http.StatusOK, user)
}

// DeleteUser handles DELETE /users/:id
func DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if _, exists := users[uint(id)]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	delete(users, uint(id))
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
} 