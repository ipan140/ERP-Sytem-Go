package knowledge

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateArticle godoc
// @Summary Create a new Article
// @Description Create a new Article in the system
// @Tags knowledge
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/knowledge [post]
// @Security BearerAuth
func CreateArticleHandler(c echo.Context) error {
	var data Article
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateArticleService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllArticle godoc
// @Summary Get all Article
// @Description Retrieve a list of all Article
// @Tags knowledge
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/knowledge [get]
// @Security BearerAuth
func GetAllArticleHandler(c echo.Context) error {
	data, err := GetAllArticleService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetArticleByID godoc
// @Summary Get a Article by ID
// @Description Retrieve a specific Article by its ID
// @Tags knowledge
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/knowledge/{id} [get]
// @Security BearerAuth
func GetArticleByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetArticleByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateArticle godoc
// @Summary Update a Article
// @Description Update an existing Article
// @Tags knowledge
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/knowledge/{id} [put]
// @Security BearerAuth
func UpdateArticleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetArticleByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateArticleService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteArticle godoc
// @Summary Delete a Article
// @Description Delete a Article by ID
// @Tags knowledge
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/knowledge/{id} [delete]
// @Security BearerAuth
func DeleteArticleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteArticleService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
