package manufacturing

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateMrpProduction godoc
// @Summary Create a new MrpProduction
// @Description Create a new MrpProduction in the system
// @Tags supply_chain-manufacturing
// @Accept json
// @Produce json
// @Success 201 {object} MrpProduction
// @Param request body MrpProduction true "Payload"
// @Router /api/supply_chain/manufacturing [post]
// @Security BearerAuth
func CreateMrpProductionHandler(c echo.Context) error {
	var data MrpProduction
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateMrpProductionService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllMrpProduction godoc
// @Summary Get all MrpProduction
// @Description Retrieve a list of all MrpProduction
// @Tags supply_chain-manufacturing
// @Produce json
// @Success 200 {object} []MrpProduction
// @Router /api/supply_chain/manufacturing [get]
// @Security BearerAuth
func GetAllMrpProductionHandler(c echo.Context) error {
	data, err := GetAllMrpProductionService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetMrpProductionByID godoc
// @Summary Get a MrpProduction by ID
// @Description Retrieve a specific MrpProduction by its ID
// @Tags supply_chain-manufacturing
// @Produce json
// @Param id path int true "MrpProduction ID"
// @Success 200 {object} MrpProduction
// @Router /api/supply_chain/manufacturing/{id} [get]
// @Security BearerAuth
func GetMrpProductionByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMrpProductionByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateMrpProduction godoc
// @Summary Update a MrpProduction
// @Description Update an existing MrpProduction
// @Tags supply_chain-manufacturing
// @Accept json
// @Produce json
// @Param id path int true "MrpProduction ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/manufacturing/{id} [put]
// @Security BearerAuth
func UpdateMrpProductionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMrpProductionByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateMrpProductionService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteMrpProduction godoc
// @Summary Delete a MrpProduction
// @Description Delete a MrpProduction by ID
// @Tags supply_chain-manufacturing
// @Produce json
// @Param id path int true "MrpProduction ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/manufacturing/{id} [delete]
// @Security BearerAuth
func DeleteMrpProductionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteMrpProductionService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// @Summary Create MrpWorkcenter
// @Description Create a new MrpWorkcenter
// @Tags supply_chain-manufacturing
// @Accept json
// @Produce json
// @Success 201 {object} MrpWorkcenter
// @Param request body MrpWorkcenter true "Payload"
// @Router /api/supply_chain/manufacturing/mrpworkcenter [post]
// @Security BearerAuth
func CreateMrpWorkcenterHandler(c echo.Context) error {
	var data MrpWorkcenter
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateMrpWorkcenterService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all MrpWorkcenter
// @Description Retrieve a list of all MrpWorkcenter
// @Tags supply_chain-manufacturing
// @Produce json
// @Success 200 {object} MrpWorkcenter
// @Router /api/supply_chain/manufacturing/mrpworkcenter [get]
// @Security BearerAuth
func GetAllMrpWorkcenterHandler(c echo.Context) error {
	data, err := GetAllMrpWorkcenterService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetMrpWorkcenterByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMrpWorkcenterByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update MrpWorkcenter
// @Description Update an existing MrpWorkcenter
// @Tags supply_chain-manufacturing
// @Accept json
// @Produce json
// @Param id path int true "MrpWorkcenter ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/manufacturing/mrpworkcenter/{id} [put]
// @Security BearerAuth
func UpdateMrpWorkcenterHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMrpWorkcenterByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateMrpWorkcenterService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete MrpWorkcenter
// @Description Delete MrpWorkcenter by ID
// @Tags supply_chain-manufacturing
// @Produce json
// @Param id path int true "MrpWorkcenter ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/manufacturing/mrpworkcenter/{id} [delete]
// @Security BearerAuth
func DeleteMrpWorkcenterHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteMrpWorkcenterService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create MrpBom
// @Description Create a new MrpBom
// @Tags supply_chain-manufacturing
// @Accept json
// @Produce json
// @Success 201 {object} MrpBom
// @Param request body MrpBom true "Payload"
// @Router /api/supply_chain/manufacturing/mrpbom [post]
// @Security BearerAuth
func CreateMrpBomHandler(c echo.Context) error {
	var data MrpBom
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateMrpBomService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all MrpBom
// @Description Retrieve a list of all MrpBom
// @Tags supply_chain-manufacturing
// @Produce json
// @Success 200 {object} MrpBom
// @Router /api/supply_chain/manufacturing/mrpbom [get]
// @Security BearerAuth
func GetAllMrpBomHandler(c echo.Context) error {
	data, err := GetAllMrpBomService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetMrpBomByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMrpBomByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update MrpBom
// @Description Update an existing MrpBom
// @Tags supply_chain-manufacturing
// @Accept json
// @Produce json
// @Param id path int true "MrpBom ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/manufacturing/mrpbom/{id} [put]
// @Security BearerAuth
func UpdateMrpBomHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMrpBomByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateMrpBomService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete MrpBom
// @Description Delete MrpBom by ID
// @Tags supply_chain-manufacturing
// @Produce json
// @Param id path int true "MrpBom ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/manufacturing/mrpbom/{id} [delete]
// @Security BearerAuth
func DeleteMrpBomHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteMrpBomService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create MrpBomLine
// @Description Create a new MrpBomLine
// @Tags supply_chain-manufacturing
// @Accept json
// @Produce json
// @Success 201 {object} MrpBomLine
// @Param request body MrpBomLine true "Payload"
// @Router /api/supply_chain/manufacturing/mrpbomline [post]
// @Security BearerAuth
func CreateMrpBomLineHandler(c echo.Context) error {
	var data MrpBomLine
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateMrpBomLineService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all MrpBomLine
// @Description Retrieve a list of all MrpBomLine
// @Tags supply_chain-manufacturing
// @Produce json
// @Success 200 {object} MrpBomLine
// @Router /api/supply_chain/manufacturing/mrpbomline [get]
// @Security BearerAuth
func GetAllMrpBomLineHandler(c echo.Context) error {
	data, err := GetAllMrpBomLineService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetMrpBomLineByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMrpBomLineByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update MrpBomLine
// @Description Update an existing MrpBomLine
// @Tags supply_chain-manufacturing
// @Accept json
// @Produce json
// @Param id path int true "MrpBomLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/manufacturing/mrpbomline/{id} [put]
// @Security BearerAuth
func UpdateMrpBomLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMrpBomLineByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateMrpBomLineService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete MrpBomLine
// @Description Delete MrpBomLine by ID
// @Tags supply_chain-manufacturing
// @Produce json
// @Param id path int true "MrpBomLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/manufacturing/mrpbomline/{id} [delete]
// @Security BearerAuth
func DeleteMrpBomLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteMrpBomLineService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create MrpBomByproduct
// @Description Create a new MrpBomByproduct
// @Tags supply_chain-manufacturing
// @Accept json
// @Produce json
// @Success 201 {object} MrpBomByproduct
// @Param request body MrpBomByproduct true "Payload"
// @Router /api/supply_chain/manufacturing/mrpbombyproduct [post]
// @Security BearerAuth
func CreateMrpBomByproductHandler(c echo.Context) error {
	var data MrpBomByproduct
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateMrpBomByproductService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all MrpBomByproduct
// @Description Retrieve a list of all MrpBomByproduct
// @Tags supply_chain-manufacturing
// @Produce json
// @Success 200 {object} MrpBomByproduct
// @Router /api/supply_chain/manufacturing/mrpbombyproduct [get]
// @Security BearerAuth
func GetAllMrpBomByproductHandler(c echo.Context) error {
	data, err := GetAllMrpBomByproductService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetMrpBomByproductByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMrpBomByproductByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update MrpBomByproduct
// @Description Update an existing MrpBomByproduct
// @Tags supply_chain-manufacturing
// @Accept json
// @Produce json
// @Param id path int true "MrpBomByproduct ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/manufacturing/mrpbombyproduct/{id} [put]
// @Security BearerAuth
func UpdateMrpBomByproductHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMrpBomByproductByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateMrpBomByproductService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete MrpBomByproduct
// @Description Delete MrpBomByproduct by ID
// @Tags supply_chain-manufacturing
// @Produce json
// @Param id path int true "MrpBomByproduct ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/manufacturing/mrpbombyproduct/{id} [delete]
// @Security BearerAuth
func DeleteMrpBomByproductHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteMrpBomByproductService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create MrpWorkorder
// @Description Create a new MrpWorkorder
// @Tags supply_chain-manufacturing
// @Accept json
// @Produce json
// @Success 201 {object} MrpWorkorder
// @Param request body MrpWorkorder true "Payload"
// @Router /api/supply_chain/manufacturing/mrpworkorder [post]
// @Security BearerAuth
func CreateMrpWorkorderHandler(c echo.Context) error {
	var data MrpWorkorder
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateMrpWorkorderService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all MrpWorkorder
// @Description Retrieve a list of all MrpWorkorder
// @Tags supply_chain-manufacturing
// @Produce json
// @Success 200 {object} MrpWorkorder
// @Router /api/supply_chain/manufacturing/mrpworkorder [get]
// @Security BearerAuth
func GetAllMrpWorkorderHandler(c echo.Context) error {
	data, err := GetAllMrpWorkorderService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetMrpWorkorderByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMrpWorkorderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update MrpWorkorder
// @Description Update an existing MrpWorkorder
// @Tags supply_chain-manufacturing
// @Accept json
// @Produce json
// @Param id path int true "MrpWorkorder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/manufacturing/mrpworkorder/{id} [put]
// @Security BearerAuth
func UpdateMrpWorkorderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMrpWorkorderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateMrpWorkorderService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete MrpWorkorder
// @Description Delete MrpWorkorder by ID
// @Tags supply_chain-manufacturing
// @Produce json
// @Param id path int true "MrpWorkorder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/manufacturing/mrpworkorder/{id} [delete]
// @Security BearerAuth
func DeleteMrpWorkorderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteMrpWorkorderService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}


