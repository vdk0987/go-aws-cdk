package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.UserStore
}

func NewApiHandler(dbStore database.UserStore) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api *ApiHandler) RegisterUser(event types.RegisterUser) error {
	if event.Password == "" || event.Username == "" {
		return fmt.Errorf("Username and Password fields cannot be empty")
	}
	doesUserExist, err := api.dbStore.UserValidation(event.Username)

	if err != nil {
		return fmt.Errorf("Error while registering user %w", err)
	}
	if doesUserExist {
		return fmt.Errorf("Username is already in use")
	}

	err = api.dbStore.InsertUser(event)
	if err != nil {
		return fmt.Errorf("Error inserting the user into db %w", err)
	}
	return nil
}
