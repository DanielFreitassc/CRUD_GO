package handlers

func Delete(w http,ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r,"id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v",err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Erro ao remover registo: %v",err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 
	}

	if rows > 1 {
		log.Printf("Error: foram removidos %d registros",rows)
	}

	resp := map[string]any {
		"Error":false,
		"Message": fmt.Sprintf("registro removido com sucesso! ID: %d", id),
	}
	
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}