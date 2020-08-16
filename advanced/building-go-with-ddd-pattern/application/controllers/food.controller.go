package controllers

import (
	contracts "building-go-with-ddd-pattern/application/contracts/foods"
	"building-go-with-ddd-pattern/application/services"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"github.com/jinzhu/gorm"
)

// FoodController
type FoodController struct {
	FoodService *services.FoodService
}

func NewFoodController(db *gorm.DB) *FoodController {
	return &FoodController{
		FoodService: services.NewFoodService(db),
	}
}

func (c *FoodController) Insert(g *gin.Context) {
	var foodError = make(map[string]string)

	title := g.PostForm("title")
	description := g.PostForm("description")

	foodContract := contracts.UpsertFoodContract{}
	foodContract.Title = title
	foodContract.Description = description

	savedFood, foodError := c.FoodService.InsertFood(foodContract)

	if foodError != nil {
		g.JSON(http.StatusBadRequest, foodError)
		return
	}

	g.JSON(http.StatusCreated, savedFood)
}

func (c *FoodController) Update(g *gin.Context) {
	var foodError = make(map[string]string)

	title := g.PostForm("title")
	description := g.PostForm("description")
	id, _ := uuid.FromString(g.Param("id"))

	foodContract := contracts.UpsertFoodContract{}
	foodContract.ID = id
	foodContract.Title = title
	foodContract.Description = description

	savedFood, foodError := c.FoodService.UpdateFood(foodContract)

	if foodError != nil {
		g.JSON(http.StatusBadRequest, foodError)
		return
	}

	g.JSON(http.StatusCreated, savedFood)
}

func (c *FoodController) GetAll(g *gin.Context) {
	var err error

	foods, err := c.FoodService.GetAll()

	if err != nil {
		g.JSON(http.StatusBadRequest, err)
	}

	g.JSON(http.StatusOK, foods)
}

func (c *FoodController) Get(g *gin.Context) {
	var err error

	id, _ := uuid.FromString(g.Param("id"))
	food, err := c.FoodService.GetFood(id)

	if err != nil {
		g.JSON(http.StatusBadRequest, err)
	}

	g.JSON(http.StatusOK, food)
}

func (c *FoodController) Delete(g *gin.Context) {
	id, _ := uuid.FromString(g.Param("id"))
	err := c.FoodService.Delete(id)

	if err != nil {
		g.JSON(http.StatusBadRequest, err)
	}

	g.JSON(http.StatusOK, "delete is successfully")
}

func (c *FoodController) FoodControllerRoute(r *gin.Engine) {
	// food routes
	r.POST("/foods", c.Insert)
	r.GET("/foods/:id", c.Get)
	r.GET("/foods", c.GetAll)
	r.DELETE("/foods/:id", c.Delete)
	r.PUT("/foods/:id", c.Update)
}
