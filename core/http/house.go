package http

import (
	"net/http"
	"path/filepath"
)

func (s *Server) createHousePage(w http.ResponseWriter, r *http.Request) {
	// New house handler will be implemented later
}

func (s *Server) createHouse(w http.ResponseWriter, r *http.Request) {
	// New house form submission handler will be implemented later
}

func (s *Server) modifyHousePage(w http.ResponseWriter, r *http.Request) {
	// Edit house handler will be implemented later
}

func (s *Server) modifyHouse(w http.ResponseWriter, r *http.Request) {
	// Edit house form submission handler will be implemented later
}

func (s *Server) deleteHousePage(w http.ResponseWriter, r *http.Request) {
	// Delete house confirmation handler will be implemented later
}

func (s *Server) deleteHouse(w http.ResponseWriter, r *http.Request) {
	// Delete house submission handler will be implemented later
}

func (s *Server) housePage(w http.ResponseWriter, r *http.Request) {
	// House detail handler will be implemented later
}

func (s *Server) houseFile(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join(s.uploadsDir, r.PathValue("id"), r.PathValue("filename"))
	http.ServeFile(w, r, filePath)
}
