package http

import "net/http"

type DashboardView struct {
	CoursesCount  int `json:"coursesCount"`
	StudentsCount int `json:"studentsCount"`
}

func (h *Handler) handleDashboard(w http.ResponseWriter, r *http.Request) {
	courses, err1 := h.courses.List(r.Context())
	students, err2 := h.students.List(r.Context())

	if err1 != nil || err2 != nil {
		writeError(w, http.StatusBadGateway, "failed to load dashboard data")
		return
	}

	writeJSON(w, http.StatusOK, DashboardView{
		CoursesCount:  len(courses),
		StudentsCount: len(students),
	})
}
