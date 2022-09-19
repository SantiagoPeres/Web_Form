package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/aprendagolang/webform/controllers"
)

func SubscriptionHandler() func(http.ResponseWriter, *http.Request) {
	tmpl, err := template.ParseFiles("handlers/templates/subscription_form.html")
	if err != nil {
		log.Fatal(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		form := map[string]string{
			"Name":  "",
			"email": "",
			"Error": "",
		}

		if r.Method == http.MethodPost {
			err := parseForm(form, r)
			if err != nil {
				fmt.Fprintf(w, "erro ao fazer parse do form : %v", err)
				return
			}

			validateform(form)
			if form["Error"] == "" {

				err = controllers.Create(form["name"], form["email"])
				if err != nil {
					fmt.Fprintf(w, "erro ao salvar os dados : %v", err)
					return
				}
			}
		}

		tmpl.Execute(w, nil)
	}
}

func parseForm(form map[string]string, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	form["Name"] = r.PostForm.Get("name")
	form["Email"] = r.PostForm.Get("email")

	return nil
}

func validateform(form map[string]string) {
	if form["Name"] == "" {
		form["Error"] = "O campo nome não pode ser vazio."
		return
	}
	if form["Email"] == "" {
		form["Error"] = "O campo email não pode ser vazio."
		return
	}
}
