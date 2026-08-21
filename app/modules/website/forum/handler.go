package forum

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateForumPost godoc
// @Summary Create a new ForumPost
// @Description Create a new ForumPost in the system
// @Tags website-forum
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/website/forum [post]
// @Security BearerAuth
func CreateForumPostHandler(c echo.Context) error {
	var data ForumPost
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateForumPostService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllForumPost godoc
// @Summary Get all ForumPost
// @Description Retrieve a list of all ForumPost
// @Tags website-forum
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/website/forum [get]
// @Security BearerAuth
func GetAllForumPostHandler(c echo.Context) error {
	data, err := GetAllForumPostService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetForumPostByID godoc
// @Summary Get a ForumPost by ID
// @Description Retrieve a specific ForumPost by its ID
// @Tags website-forum
// @Produce json
// @Param id path int true "ForumPost ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/forum/{id} [get]
// @Security BearerAuth
func GetForumPostByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetForumPostByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateForumPost godoc
// @Summary Update a ForumPost
// @Description Update an existing ForumPost
// @Tags website-forum
// @Accept json
// @Produce json
// @Param id path int true "ForumPost ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/forum/{id} [put]
// @Security BearerAuth
func UpdateForumPostHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetForumPostByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateForumPostService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteForumPost godoc
// @Summary Delete a ForumPost
// @Description Delete a ForumPost by ID
// @Tags website-forum
// @Produce json
// @Param id path int true "ForumPost ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/forum/{id} [delete]
// @Security BearerAuth
func DeleteForumPostHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteForumPostService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
