package courses

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/knr1997/quiz-tracker-backend/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

// ListCourses - GET /courses
func (h *handler) ListCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := h.service.ListCourses(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, courses)
}

// GetCourseByID - GET /courses/{id}
func (h *handler) GetCourseByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		json.Write(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid course ID",
		})
		return
	}

	course, err := h.service.GetCourseByID(r.Context(), id)
	if err != nil {
		log.Println(err)
		// You can add more specific error handling here
		status := http.StatusInternalServerError
		errorMsg := err.Error()

		// Example: Check for "not found" errors
		// if errors.Is(err, sql.ErrNoRows) {
		//     status = http.StatusNotFound
		//     errorMsg = "Course not found"
		// }

		json.Write(w, status, map[string]string{
			"error": errorMsg,
		})
		return
	}

	json.Write(w, http.StatusOK, course)
}

// CreateCoruse - POST /courses
func (h *handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var tempCourse createCourseParams
	if err := json.Read(r, &tempCourse); err != nil {
		log.Println(err)
		json.Write(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	// Validate input
	if tempCourse.Name == "" || tempCourse.Code == "" {
		json.Write(w, http.StatusBadRequest, map[string]string{
			"error": "Name and Code are required",
		})
		return
	}

	createdCourse, err := h.service.CreateCourse(r.Context(), tempCourse)
	if err != nil {
		log.Println(err)
		json.Write(w, http.StatusInternalServerError, map[string]string{
			"error": "Failed to create course",
		})
		return
	}

	json.Write(w, http.StatusCreated, createdCourse)
}

// UpdateCourse - PUT /courses/{id}
func (h *handler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		json.Write(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid course ID",
		})
		return
	}

	var tempCourse updateCourseParams
	if err := json.Read(r, &tempCourse); err != nil {
		log.Println(err)
		json.Write(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	if tempCourse.Name == "" && tempCourse.Code == "" {
		json.Write(w, http.StatusBadRequest, map[string]string{
			"error": "At least one field (Name or Code) must be provided",
		})
		return
	}

	updatedCourse, err := h.service.UpdateCourse(r.Context(), id, tempCourse)
	if err != nil {
		log.Println(err)
		status := http.StatusInternalServerError
		errorMsg := "Failed to update course"

		// You can add specific error handling here
		// if errors.Is(err, sql.ErrNoRows) {
		//     status = http.StatusNotFound
		//     errorMsg = "Course not found"
		// }

		json.Write(w, status, map[string]string{
			"error": errorMsg,
		})
		return
	}

	json.Write(w, http.StatusOK, updatedCourse)
}

// DeleteCourse - DELETE /courses/{id}
func (h *handler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		json.Write(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid course ID",
		})
		return
	}

	err = h.service.DeleteCourse(r.Context(), id)
	if err != nil {
		log.Println(err)
		status := http.StatusInternalServerError
		errorMsg := "Failed to delete course"

		// You can add specific error handling here
		// if errors.Is(err, sql.ErrNoRows) {
		//     status = http.StatusNotFound
		//     errorMsg = "Course not found"
		// }

		json.Write(w, status, map[string]string{
			"error": errorMsg,
		})
		return
	}

	// Return 204 No Content for successful deletion
	w.WriteHeader(http.StatusNoContent)
}

// Helper function to extract ID from request
func getIDFromRequest(r *http.Request) (int64, error) {
	idStr := chi.URLParam(r, "id")
	return strconv.ParseInt(idStr, 10, 64)
}
