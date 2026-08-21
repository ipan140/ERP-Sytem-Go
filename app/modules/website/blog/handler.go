package blog

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateBlogPost godoc
// @Summary Create a new BlogPost
// @Description Create a new BlogPost in the system
// @Tags website-blog
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/website/blog [post]
// @Security BearerAuth
func CreateBlogPostHandler(c echo.Context) error {
	var data BlogPost
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateBlogPostService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllBlogPost godoc
// @Summary Get all BlogPost
// @Description Retrieve a list of all BlogPost
// @Tags website-blog
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/website/blog [get]
// @Security BearerAuth
func GetAllBlogPostHandler(c echo.Context) error {
	data, err := GetAllBlogPostService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetBlogPostByID godoc
// @Summary Get a BlogPost by ID
// @Description Retrieve a specific BlogPost by its ID
// @Tags website-blog
// @Produce json
// @Param id path int true "BlogPost ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/blog/{id} [get]
// @Security BearerAuth
func GetBlogPostByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetBlogPostByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateBlogPost godoc
// @Summary Update a BlogPost
// @Description Update an existing BlogPost
// @Tags website-blog
// @Accept json
// @Produce json
// @Param id path int true "BlogPost ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/blog/{id} [put]
// @Security BearerAuth
func UpdateBlogPostHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetBlogPostByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateBlogPostService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteBlogPost godoc
// @Summary Delete a BlogPost
// @Description Delete a BlogPost by ID
// @Tags website-blog
// @Produce json
// @Param id path int true "BlogPost ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/blog/{id} [delete]
// @Security BearerAuth
func DeleteBlogPostHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteBlogPostService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
