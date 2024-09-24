package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/shafiqsaaidin/go-crud-api-tutorial/240920-gin/pkg/recipes"
)

func main() {
	// Create Gin router
	router := gin.Default()

	// Instantiate recipe Handler and provide a data store implementation
	store := recipes.NewMemStore()
	recipesHandlers := NewRecipesHandler(store)

	// Register routes
	router.GET("/", homePage)
	router.GET("/recipes", recipesHandlers.ListRecipes)
	router.POST("/recipes", recipesHandlers.CreateRecipe)
	router.GET("/recipes/:id", recipesHandlers.GetRecipe)
	router.PUT("/recipes/:id", recipesHandlers.UpdateRecipe)
	router.DELETE("/recipes/:id", recipesHandlers.DeleteRecipe)

	// Start the server
	router.Run()
}

func homePage(c *gin.Context) {
	c.String(http.StatusOK, "This is my home page")
}

type RecipesHandler struct {
	store recipeStore
}

// recipeStore is an interface for the data store
type recipeStore interface {
	Add(name string, recipe recipes.Recipe) error
	Get(name string) (recipes.Recipe, error)
	List() (map[string]recipes.Recipe, error)
	Update(name string, recipe recipes.Recipe) error
	Remove(name string) error
}

// Define handler function signature
func (h RecipesHandler) CreateRecipe(c *gin.Context) {
	// Get request body and convert it to recipes.Recipe
	var recipe recipes.Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a URL-friendly name
	id := slug.Make(recipe.Name)

	// Add to the store
	h.store.Add(id, recipe)

	// Return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h RecipesHandler) ListRecipes(c *gin.Context) {
	// Call the store to get the list of recipes
	r, err := h.store.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Return the list, JSON encoding is implicit
	c.JSON(200, r)
}

func (h RecipesHandler) GetRecipe(c *gin.Context) {
	// Retrive the url parameter
	id := c.Param("id")

	// Get the recipe by ID from the store
	recipe, err := h.store.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	// Return the recipe, JSON encoding is implicit
	c.JSON(200, recipe)
}

func (h RecipesHandler) UpdateRecipe(c *gin.Context) {
	// Get request body and update it to tje recipes.Recipe
	var recipe recipes.Recipe
	if err := c.ShouldBindBodyWithJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve URL parameter
	id := c.Param("id")

	// Call the store to update the recipe
	err := h.store.Update(id, recipe)
	if err != nil {
		if err == recipes.NotFoundErr {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h RecipesHandler) DeleteRecipe(c *gin.Context) {
	// Retrieve URL parameter
	id := c.Param("id")

	// call the store to delete the recipe
	err := h.store.Remove(id)
	if err != nil {
		if err == recipes.NotFoundErr {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return success payload
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// NewRecipesHandler is a constructor for RecipesHandler
func NewRecipesHandler(s recipeStore) *RecipesHandler {
	return &RecipesHandler{
		store: s,
	}
}
