package handlers

func Update(w http,ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r,"id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v",err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 
	}

	var todo models.Todo

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v",err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 
	}

	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Printf("Erro ao atualizar registo: %v",err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 
	}

	if rows > 1 {
		log.Printf("Error: foram atualizados %d registros",rows)
	}

	resp := map[string]any {
		"Error":false,
		"Message": fmt.Sprintf("todo inserido com sucesso! ID: %d", id),
	}
	
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}