package Controller

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/kriangkrai/SQL/Models"
)

func GetLocation() []Models.Location {
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "SELECT ID,Name,Location FROM Location")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var locations []Models.Location

	for rows.Next() {
		var id int
		var name string
		var location string
		err := rows.Scan(&id, &name, &location)
		if err != nil {
			log.Fatal(err)
		}
		data := Models.Location{
			ID:       id,
			Name:     name,
			Location: location,
		}
		locations = append(locations, data)
	}
	return locations
}

func InsertLocation(locations Models.Location) string {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}
	_, execErr := tx.ExecContext(ctx, "Insert Into Location(Name,Location) VALUES (?,?)", locations.Name, locations.Location)
	if execErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("update failed: %v, unable to rollback: %v\n", execErr, rollbackErr)
		}
		log.Fatalf("update failed: %v", execErr)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	return "Insert Success"
}

func UpdateLocation(locations Models.Location) string {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}
	_, execErr := tx.ExecContext(ctx, "UPDATE Location SET Name = ? , Location = ? WHERE ID = ?", locations.Name, locations.Location, locations.ID)
	if execErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("update failed: %v, unable to rollback: %v\n", execErr, rollbackErr)
		}
		log.Fatalf("update failed: %v", execErr)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	return "Update Success"
}

func DeleteLocation(id string) string {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}
	_, execErr := tx.ExecContext(ctx, "Delete From Location WHERE ID = ? ", id)
	if execErr != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("update failed: %v, unable to rollback: %v\n", execErr, rollbackErr)
		}
		log.Fatalf("update failed: %v", execErr)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	return "Delete Success"
}
