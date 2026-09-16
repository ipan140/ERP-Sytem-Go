package base

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateCurrency godoc
// @Summary Create a new Currency
// @Description Create a new Currency in the system
// @Tags base
// @Accept json
// @Produce json
// @Success 201 {object} Currency
// @Param request body Currency true "Payload"
// @Router /api/base [post]
// @Security BearerAuth
func CreateCurrencyHandler(c echo.Context) error {
	var data Currency
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateCurrencyService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllCurrency godoc
// @Summary Get all Currency
// @Description Retrieve a list of all Currency
// @Tags base
// @Produce json
// @Success 200 {object} []Currency
// @Router /api/base [get]
// @Security BearerAuth
func GetAllCurrencyHandler(c echo.Context) error {
	data, err := GetAllCurrencyService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetCurrencyByID godoc
// @Summary Get a Currency by ID
// @Description Retrieve a specific Currency by its ID
// @Tags base
// @Produce json
// @Param id path int true "Currency ID"
// @Success 200 {object} Currency
// @Router /api/base/{id} [get]
// @Security BearerAuth
func GetCurrencyByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetCurrencyByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateCurrency godoc
// @Summary Update a Currency
// @Description Update an existing Currency
// @Tags base
// @Accept json
// @Produce json
// @Param id path int true "Currency ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/base/{id} [put]
// @Security BearerAuth
func UpdateCurrencyHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetCurrencyByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateCurrencyService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteCurrency godoc
// @Summary Delete a Currency
// @Description Delete a Currency by ID
// @Tags base
// @Produce json
// @Param id path int true "Currency ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/base/{id} [delete]
// @Security BearerAuth
func DeleteCurrencyHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteCurrencyService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// @Summary Create Country
// @Description Create a new Country
// @Tags core-base
// @Accept json
// @Produce json
// @Success 201 {object} Country
// @Param request body Country true "Payload"
// @Router /api/core/base/country [post]
// @Security BearerAuth
func CreateCountryHandler(c echo.Context) error {
	var data Country
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateCountryService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all Country
// @Description Retrieve a list of all Country
// @Tags core-base
// @Produce json
// @Success 200 {object} Country
// @Router /api/core/base/country [get]
// @Security BearerAuth
func GetAllCountryHandler(c echo.Context) error {
	data, err := GetAllCountryService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetCountryByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetCountryByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update Country
// @Description Update an existing Country
// @Tags core-base
// @Accept json
// @Produce json
// @Param id path int true "Country ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/core/base/country/{id} [put]
// @Security BearerAuth
func UpdateCountryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetCountryByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateCountryService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete Country
// @Description Delete Country by ID
// @Tags core-base
// @Produce json
// @Param id path int true "Country ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/core/base/country/{id} [delete]
// @Security BearerAuth
func DeleteCountryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteCountryService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create CountryState
// @Description Create a new CountryState
// @Tags core-base
// @Accept json
// @Produce json
// @Success 201 {object} CountryState
// @Param request body CountryState true "Payload"
// @Router /api/core/base/countrystate [post]
// @Security BearerAuth
func CreateCountryStateHandler(c echo.Context) error {
	var data CountryState
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateCountryStateService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all CountryState
// @Description Retrieve a list of all CountryState
// @Tags core-base
// @Produce json
// @Success 200 {object} CountryState
// @Router /api/core/base/countrystate [get]
// @Security BearerAuth
func GetAllCountryStateHandler(c echo.Context) error {
	data, err := GetAllCountryStateService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetCountryStateByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetCountryStateByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update CountryState
// @Description Update an existing CountryState
// @Tags core-base
// @Accept json
// @Produce json
// @Param id path int true "CountryState ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/core/base/countrystate/{id} [put]
// @Security BearerAuth
func UpdateCountryStateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetCountryStateByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateCountryStateService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete CountryState
// @Description Delete CountryState by ID
// @Tags core-base
// @Produce json
// @Param id path int true "CountryState ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/core/base/countrystate/{id} [delete]
// @Security BearerAuth
func DeleteCountryStateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteCountryStateService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create Partner
// @Description Create a new Partner
// @Tags core-base
// @Accept json
// @Produce json
// @Success 201 {object} Partner
// @Param request body Partner true "Payload"
// @Router /api/core/base/partner [post]
// @Security BearerAuth
func CreatePartnerHandler(c echo.Context) error {
	var data Partner
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreatePartnerService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all Partner
// @Description Retrieve a list of all Partner
// @Tags core-base
// @Produce json
// @Success 200 {object} Partner
// @Router /api/core/base/partner [get]
// @Security BearerAuth
func GetAllPartnerHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllPartnerService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	data, total, err := GetPaginatedPartnerService(offset, limit, search)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve paginated data", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Retrieved successfully", data, meta)
}
func GetPartnerByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPartnerByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update Partner
// @Description Update an existing Partner
// @Tags core-base
// @Accept json
// @Produce json
// @Param id path int true "Partner ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/core/base/partner/{id} [put]
// @Security BearerAuth
func UpdatePartnerHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPartnerByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdatePartnerService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete Partner
// @Description Delete Partner by ID
// @Tags core-base
// @Produce json
// @Param id path int true "Partner ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/core/base/partner/{id} [delete]
// @Security BearerAuth
func DeletePartnerHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePartnerService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}


