package calendar

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

// GetEventsHandler godoc
// @Summary      Ambil Daftar Jadwal Kalender
// @Description  Mengambil jadwal rapat, cuti, atau acara lainnya yang sesuai dengan filter privasi user
// @Tags         Calendar
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        model query string false "Filter berdasarkan Modul (contoh: hr.leaves)"
// @Param        start query string false "Format: YYYY-MM-DD"
// @Param        end   query string false "Format: YYYY-MM-DD"
// @Success      200  {array}   CalendarEvent
// @Failure      401  {object}  map[string]string "Unauthorized"
// @Failure      500  {object}  map[string]string "Gagal mengambil data"
// @Router       /api/calendar/events [get]
func GetEventsHandler(c echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

	modelFilter := c.QueryParam("model")
	if modelFilter == "" {
		modelFilter = c.QueryParam("res_model")
	}
	start := c.QueryParam("start")
	end := c.QueryParam("end")

	events, err := FetchEventsService(userID, modelFilter, start, end)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal mengambil data"})
	}

	return c.JSON(http.StatusOK, events)
}

// CreateEventHandler godoc
// @Summary      Buat Jadwal Kalender Baru
// @Description  Menyimpan acara kalender baru ke database (Otomatis mengirimkan ke peserta yang di-tag)
// @Tags         Calendar
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body CalendarEvent true "Isian Jadwal (Kirim array attendee_ids untuk peserta)"
// @Success      201  {object}  CalendarEvent
// @Failure      400  {object}  map[string]string "Data tidak valid"
// @Failure      500  {object}  map[string]string "Gagal menyimpan jadwal"
// @Router       /api/calendar/events [post]
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

// UpdateEventHandler godoc
// @Summary      Update Jadwal Kalender
// @Description  Memperbarui judul, tanggal, url, atau peserta dari sebuah acara kalender
// @Tags         Calendar
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id      path string        true "ID Jadwal (Contoh: 1)"
// @Param        request body CalendarEvent true "Data baru untuk jadwal ini"
// @Success      200  {object}  CalendarEvent
// @Failure      400  {object}  map[string]string "Data tidak valid"
// @Failure      404  {object}  map[string]string "Event tidak ditemukan"
// @Failure      500  {object}  map[string]string "Gagal update jadwal"
// @Router       /api/calendar/events/{id} [put]
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

// DeleteEventHandler godoc
// @Summary      Hapus Jadwal Kalender
// @Description  Menghapus acara dari kalender secara permanen
// @Tags         Calendar
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path string true "ID Jadwal (Contoh: 1)"
// @Success      200  {object}  map[string]string "Jadwal dihapus"
// @Failure      500  {object}  map[string]string "Gagal menghapus jadwal"
// @Router       /api/calendar/events/{id} [delete]
func DeleteEventHandler(c echo.Context) error {
	id := c.Param("id")
	if err := RemoveEventService(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal menghapus jadwal"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Jadwal dihapus"})
}

// GetCategoriesHandler godoc
// @Summary      Ambil Daftar Kategori Kalender
// @Description  Mengambil daftar kategori kalender berdasarkan modul (HR, Finance, SupplyChain, Sales, dll)
// @Tags         Calendar
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        module query string false "Nama Modul (contoh: HR, Finance, SupplyChain)"
// @Success      200  {array}   CalendarCategory
// @Failure      500  {object}  map[string]string "Gagal mengambil kategori"
// @Router       /api/calendar/categories [get]
func GetCategoriesHandler(c echo.Context) error {
	module := c.QueryParam("module")
	if module == "" {
		module = c.QueryParam("res_model")
	}

	categories, err := FetchCategoriesService(module)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal mengambil kategori"})
	}

	return c.JSON(http.StatusOK, categories)
}

// CreateCategoryHandler godoc
// @Summary      Tambah Kategori Kalender Baru
// @Description  Membuat kategori baru beserta warna dan ikon untuk modul kalender tertentu
// @Tags         Calendar
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body CalendarCategory true "Data Kategori"
// @Success      201  {object}  CalendarCategory
// @Failure      400  {object}  map[string]string "Data tidak valid"
// @Failure      500  {object}  map[string]string "Gagal menyimpan kategori"
// @Router       /api/calendar/categories [post]
func CreateCategoryHandler(c echo.Context) error {
	var req CalendarCategory
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Data tidak valid"})
	}

	if req.Name == "" || req.Module == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nama dan Modul wajib diisi"})
	}

	if req.Color == "" {
		req.Color = "primary"
	}

	if err := SaveCategoryService(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal menyimpan kategori"})
	}

	return c.JSON(http.StatusCreated, req)
}

// UpdateCategoryHandler godoc
// @Summary      Update Kategori Kalender
// @Description  Memperbarui nama, warna, atau ikon kategori kalender
// @Tags         Calendar
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id      path string           true "ID Kategori"
// @Param        request body CalendarCategory true "Data baru kategori"
// @Success      200  {object}  CalendarCategory
// @Failure      400  {object}  map[string]string "Data tidak valid"
// @Failure      500  {object}  map[string]string "Gagal memperbarui kategori"
// @Router       /api/calendar/categories/{id} [put]
func UpdateCategoryHandler(c echo.Context) error {
	id := c.Param("id")
	var req CalendarCategory
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Data tidak valid"})
	}

	updated, err := ModifyCategoryService(id, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal memperbarui kategori"})
	}

	return c.JSON(http.StatusOK, updated)
}

// DeleteCategoryHandler godoc
// @Summary      Hapus Kategori Kalender
// @Description  Menghapus kategori kalender
// @Tags         Calendar
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path string true "ID Kategori"
// @Success      200  {object}  map[string]string "Kategori dihapus"
// @Failure      500  {object}  map[string]string "Gagal menghapus kategori"
// @Router       /api/calendar/categories/{id} [delete]
func DeleteCategoryHandler(c echo.Context) error {
	id := c.Param("id")
	if err := RemoveCategoryService(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal menghapus kategori"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Kategori berhasil dihapus"})
}

