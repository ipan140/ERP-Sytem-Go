package social_marketing

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateSocialPost godoc
// @Summary Create a new SocialPost
// @Description Create a new SocialPost in the system
// @Tags marketing-social_marketing
// @Accept json
// @Produce json
// @Success 201 {object} SocialPost
// @Param request body SocialPost true "Payload"
// @Router /api/marketing/social_marketing [post]
// @Security BearerAuth
func CreateSocialPostHandler(c echo.Context) error {
	var data SocialPost
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateSocialPostService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllSocialPost godoc
// @Summary Get all SocialPost
// @Description Retrieve a list of all SocialPost
// @Tags marketing-social_marketing
// @Produce json
// @Success 200 {object} []SocialPost
// @Router /api/marketing/social_marketing [get]
// @Security BearerAuth
func GetAllSocialPostHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllSocialPostService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	data, total, err := GetPaginatedSocialPostService(offset, limit, search)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data retrieved successfully", data, meta)
}

// GetSocialPostByID godoc
// @Summary Get a SocialPost by ID
// @Description Retrieve a specific SocialPost by its ID
// @Tags marketing-social_marketing
// @Produce json
// @Param id path int true "SocialPost ID"
// @Success 200 {object} SocialPost
// @Router /api/marketing/social_marketing/{id} [get]
// @Security BearerAuth
func GetSocialPostByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSocialPostByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateSocialPost godoc
// @Summary Update a SocialPost
// @Description Update an existing SocialPost
// @Tags marketing-social_marketing
// @Accept json
// @Produce json
// @Param id path int true "SocialPost ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/social_marketing/{id} [put]
// @Security BearerAuth
func UpdateSocialPostHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSocialPostByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateSocialPostService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteSocialPost godoc
// @Summary Delete a SocialPost
// @Description Delete a SocialPost by ID
// @Tags marketing-social_marketing
// @Produce json
// @Param id path int true "SocialPost ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/social_marketing/{id} [delete]
// @Security BearerAuth
func DeleteSocialPostHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSocialPostService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}


