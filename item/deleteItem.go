package item

import (
	"context"
	"log"

	database "grpc-todo/database"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) DeleteItem() error {

	_, err := database.TodoDB.Exec(
		context.Background(),
		`DELETE FROM `+database.ItemsTable+`
		 WHERE id=$1`,
		item.ID,
	)

	if err != nil {
		log.Println("Error deleting item from DB: ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}
