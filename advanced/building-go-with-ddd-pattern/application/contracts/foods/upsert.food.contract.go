package contracts

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

// UpsertFoodContract information
type UpsertFoodContract struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
}

// Validate validate contract from user input
func (foodContract *UpsertFoodContract) Validate(action string) map[string]string {
	var errMessages = make(map[string]string)

	switch strings.ToLower(action) {
	case "update":
		if foodContract.Title == "" {
			errMessages["title_required"] = "title required"
		}

		if uuid.FromStringOrNil(foodContract.ID.String()) != uuid.Nil {
			errMessages["id_required"] = "id required"
		}
	case "insert":
		if foodContract.Title == "" {
			errMessages["title_required"] = "title required"
		}
	default:
		return nil
	}

	return errMessages
}
