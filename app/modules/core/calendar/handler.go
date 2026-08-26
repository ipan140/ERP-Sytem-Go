package calendar

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func GetEventsHandler(c echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

	modelFilter := c.QueryParam("model")
	start := c.QueryParam("start")
	end := c.QueryParam("end")

	events, err := FetchEventsService(userID, modelFilter, start, end)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal mengambil data"})
	}

	return c.JSON(http.StatusOK, events)
}

func CreateEventHandler(c echo.Context) error {
	var req CalendarEvent
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Data tidak valid"})
	}

	if userID, ok := c.Get("user_id").(uint); ok {
		req.UserID = userID
	}

	if err := SaveEventService(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal menyimpan jadwal"})
	}

	return c.JSON(http.StatusCreated, req)
}

func UpdateEventHandler(c echo.Context) error {
	id := c.Param("id")
	var req CalendarEvent
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Data tidak valid"})
	}

	updatedEvent, err := ModifyEventService(id, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal update jadwal"})
	}

	return c.JSON(http.StatusOK, updatedEvent)
}

func DeleteEventHandler(c echo.Context) error {
	id := c.Param("id")
	if err := RemoveEventService(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal menghapus jadwal"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Jadwal dihapus"})
}
