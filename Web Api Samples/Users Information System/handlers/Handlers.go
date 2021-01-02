package handlers

import (
	"encoding/json"
	"net/http"

	. "../dataloaders"
	. "../models"
)

func Run() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	page := Page{ID: 1, Name: "Users", Description: "User List", URI: "/users"}

	users := LoadUsers()
	interests := LoadInterests()
	interestMappings := LoadInterestMappings()

	var newUsers []User

	for _, user := range users {

		for _, interestMapping := range interestMappings {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interestMapping.InterestID == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
	}

	viewModel := UserViewModel{Page: page, Users: newUsers}

	data, _ := json.Marshal(viewModel)

	w.Write([]byte(data))

}
