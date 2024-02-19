package handlers

import (
	"fmt"
	"gdrcdn/utils"
	"net/http"
	"os"
)


func HandlerFiles(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[7:]
	fmt.Println("id: " + id)
	// get the file from the file system
	file, err := os.Open("files/" + id)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	defer file.Close()
	// get the file info
	stat, err := file.Stat()
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	contentType, _ := utils.GetFileContentType(file)
	w.Header().Set("Content-Disposition", "attachment; filename="+id)
	if contentType != "" {
		w.Header().Set("Content-Type", contentType)
	}
	w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))
	http.ServeContent(w, r, id, stat.ModTime(), file)
}