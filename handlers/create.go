package handlers

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.todo
	
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v",err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return 
	}

	id, err := model.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any {
			"Erro": true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v",err )
		} else {
			resp = map[string]any {
				"Error":false,
				"Message": fmt.Sprintf("Todo inserido com sucesso! ID: %v",id),

			}
		}

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}