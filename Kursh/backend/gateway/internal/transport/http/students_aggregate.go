package http

import (
	"context"
	"net/http"
	"sync"

	"golang.org/x/sync/errgroup"

	"gateway/internal/clients/courses"
	"gateway/internal/clients/students"
)

type StudentView struct {
	ID          string              `json:"id"`
	FullName    string              `json:"fullName"`
	Email       string              `json:"email"`
	Age         int                 `json:"age"`
	Performance string              `json:"performance"`
	PhotoURL    string              `json:"photoUrl"`
	City        string              `json:"city"`
	Phone       string              `json:"phone"`
	Format      string              `json:"format"`
	Progress    int                 `json:"progress"`
	Course      *courses.CourseInfo `json:"course,omitempty"`
}

func (h *Handler) handleListStudentsAggregated(w http.ResponseWriter, r *http.Request) {
	base, err := h.students.List(r.Context())
	if err != nil {
		writeError(w, http.StatusBadGateway, "failed to fetch students")
		return
	}

	views, err := h.buildStudentViews(r.Context(), base)
	if err != nil {
		writeError(w, http.StatusBadGateway, "failed to aggregate students")
		return
	}

	writeJSON(w, http.StatusOK, views)
}

func (h *Handler) handleGetStudentAggregated(w http.ResponseWriter, r *http.Request) {
	id, ok := extractID(r.URL.Path, "/api/students/")
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid student id")
		return
	}

	s, err := h.students.GetByID(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusNotFound, "student not found")
		return
	}

	views, err := h.buildStudentViews(r.Context(), []students.StudentBase{*s})
	if err != nil || len(views) == 0 {
		writeError(w, http.StatusBadGateway, "failed to aggregate student")
		return
	}

	writeJSON(w, http.StatusOK, views[0])
}

func (h *Handler) buildStudentViews(ctx context.Context, base []students.StudentBase) ([]StudentView, error) {
	courseIDs := make([]string, 0, len(base))
	seen := map[string]struct{}{}
	for _, s := range base {
		if s.CourseID != nil && *s.CourseID != "" {
			if _, ok := seen[*s.CourseID]; !ok {
				seen[*s.CourseID] = struct{}{}
				courseIDs = append(courseIDs, *s.CourseID)
			}
		}
	}

	courseMap := map[string]*courses.CourseInfo{}
	var mu sync.Mutex
	eg, egCtx := errgroup.WithContext(ctx)
	for _, cid := range courseIDs {
		cid := cid
		eg.Go(func() error {
			c, err := h.courses.GetByID(egCtx, cid)
			if err == nil && c != nil {
				mu.Lock()
				courseMap[cid] = c
				mu.Unlock()
			}
			return nil
		})
	}
	_ = eg.Wait()

	views := make([]StudentView, 0, len(base))
	for _, s := range base {
		var course *courses.CourseInfo
		if s.CourseID != nil {
			course = courseMap[*s.CourseID]
		}

		views = append(views, StudentView{
			ID: s.ID, FullName: s.FullName, Email: s.Email,
			Age: s.Age, Performance: s.Performance,
			PhotoURL: s.PhotoURL, City: s.City, Phone: s.Phone,
			Format:   s.Format,
			Progress: s.Progress,
			Course:   course,
		})
	}

	return views, nil
}
