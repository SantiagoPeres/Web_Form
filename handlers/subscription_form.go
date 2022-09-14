package handlers

import (
	"fmt"
	"net/http"
)

func SubscriptionHandler(w http.ResponseWriter, r* http.Request){
	if r.Method == http.MethodPost{
		if err := r.ParseForm(); err != nil{
			fmt.Printf(w, "erro ao fazer parse do form : %v", err)
			return
		}
		
		err := controllers.Create(r.PostForm.Get("name"), r.PostForm.Get("email"))
		if err != nil {
			fmt.Printf(w, "erro ao salvar os dados : %v", err)
			return
		}
	}

	http.ServeFile(w, r "handlers/templates/subscription_form.html")
}