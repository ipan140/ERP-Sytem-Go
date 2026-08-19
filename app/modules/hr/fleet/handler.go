package fleet

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateVehicle godoc
// @Summary Create a new Vehicle
// @Description Create a new Vehicle in the system
// @Tags hr-fleet
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/fleet [post]
// @Security BearerAuth
func CreateVehicleHandler(c echo.Context) error {
	var data Vehicle
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateVehicleService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllVehicle godoc
// @Summary Get all Vehicle
// @Description Retrieve a list of all Vehicle
// @Tags hr-fleet
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet [get]
// @Security BearerAuth
func GetAllVehicleHandler(c echo.Context) error {
	data, err := GetAllVehicleService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetVehicleByID godoc
// @Summary Get a Vehicle by ID
// @Description Retrieve a specific Vehicle by its ID
// @Tags hr-fleet
// @Produce json
// @Param id path int true "Vehicle ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet/{id} [get]
// @Security BearerAuth
func GetVehicleByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetVehicleByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateVehicle godoc
// @Summary Update a Vehicle
// @Description Update an existing Vehicle
// @Tags hr-fleet
// @Accept json
// @Produce json
// @Param id path int true "Vehicle ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet/{id} [put]
// @Security BearerAuth
func UpdateVehicleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetVehicleByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateVehicleService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteVehicle godoc
// @Summary Delete a Vehicle
// @Description Delete a Vehicle by ID
// @Tags hr-fleet
// @Produce json
// @Param id path int true "Vehicle ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet/{id} [delete]
// @Security BearerAuth
func DeleteVehicleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteVehicleService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// @Summary Create VehicleLogContract
// @Description Create a new VehicleLogContract
// @Tags hr-fleet
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/fleet/vehiclelogcontract [post]
// @Security BearerAuth
func CreateVehicleLogContractHandler(c echo.Context) error {
	var data VehicleLogContract
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateVehicleLogContractService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all VehicleLogContract
// @Description Retrieve a list of all VehicleLogContract
// @Tags hr-fleet
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet/vehiclelogcontract [get]
// @Security BearerAuth
func GetAllVehicleLogContractHandler(c echo.Context) error {
	data, err := GetAllVehicleLogContractService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetVehicleLogContractByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetVehicleLogContractByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// @Summary Update VehicleLogContract
// @Description Update an existing VehicleLogContract
// @Tags hr-fleet
// @Accept json
// @Produce json
// @Param id path int true "VehicleLogContract ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet/vehiclelogcontract/{id} [put]
// @Security BearerAuth
func UpdateVehicleLogContractHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetVehicleLogContractByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateVehicleLogContractService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// @Summary Delete VehicleLogContract
// @Description Delete VehicleLogContract by ID
// @Tags hr-fleet
// @Produce json
// @Param id path int true "VehicleLogContract ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet/vehiclelogcontract/{id} [delete]
// @Security BearerAuth
func DeleteVehicleLogContractHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteVehicleLogContractService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create VehicleLogFuel
// @Description Create a new VehicleLogFuel
// @Tags hr-fleet
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/fleet/vehiclelogfuel [post]
// @Security BearerAuth
func CreateVehicleLogFuelHandler(c echo.Context) error {
	var data VehicleLogFuel
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateVehicleLogFuelService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all VehicleLogFuel
// @Description Retrieve a list of all VehicleLogFuel
// @Tags hr-fleet
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet/vehiclelogfuel [get]
// @Security BearerAuth
func GetAllVehicleLogFuelHandler(c echo.Context) error {
	data, err := GetAllVehicleLogFuelService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetVehicleLogFuelByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetVehicleLogFuelByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// @Summary Update VehicleLogFuel
// @Description Update an existing VehicleLogFuel
// @Tags hr-fleet
// @Accept json
// @Produce json
// @Param id path int true "VehicleLogFuel ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet/vehiclelogfuel/{id} [put]
// @Security BearerAuth
func UpdateVehicleLogFuelHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetVehicleLogFuelByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateVehicleLogFuelService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// @Summary Delete VehicleLogFuel
// @Description Delete VehicleLogFuel by ID
// @Tags hr-fleet
// @Produce json
// @Param id path int true "VehicleLogFuel ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet/vehiclelogfuel/{id} [delete]
// @Security BearerAuth
func DeleteVehicleLogFuelHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteVehicleLogFuelService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create VehicleLogServices
// @Description Create a new VehicleLogServices
// @Tags hr-fleet
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/fleet/vehiclelogservices [post]
// @Security BearerAuth
func CreateVehicleLogServicesHandler(c echo.Context) error {
	var data VehicleLogServices
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateVehicleLogServicesService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all VehicleLogServices
// @Description Retrieve a list of all VehicleLogServices
// @Tags hr-fleet
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet/vehiclelogservices [get]
// @Security BearerAuth
func GetAllVehicleLogServicesHandler(c echo.Context) error {
	data, err := GetAllVehicleLogServicesService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetVehicleLogServicesByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetVehicleLogServicesByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// @Summary Update VehicleLogServices
// @Description Update an existing VehicleLogServices
// @Tags hr-fleet
// @Accept json
// @Produce json
// @Param id path int true "VehicleLogServices ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet/vehiclelogservices/{id} [put]
// @Security BearerAuth
func UpdateVehicleLogServicesHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetVehicleLogServicesByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateVehicleLogServicesService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// @Summary Delete VehicleLogServices
// @Description Delete VehicleLogServices by ID
// @Tags hr-fleet
// @Produce json
// @Param id path int true "VehicleLogServices ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet/vehiclelogservices/{id} [delete]
// @Security BearerAuth
func DeleteVehicleLogServicesHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteVehicleLogServicesService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}
