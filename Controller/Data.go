package Controller

import (
	"context"
	"database/sql"
	"log"

	"github.com/kriangkrai/SQL/Models"
)

func GetData() []Models.Data {
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, "SELECT Device,Date,Value FROM Data")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var datas []Models.Data

	for rows.Next() {
		var device string
		var date string
		var value string
		err := rows.Scan(&device, &date, &value)
		if err != nil {
			log.Fatal(err)
		}
		data := Models.Data{
			Device: device,
			Date:   date,
			Value:  value,
		}
		datas = append(datas, data)
	}
	return datas
}

// func Update(id int, value float64) string {
// 	ctx := context.Background()
// 	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	_, execErr := tx.ExecContext(ctx, "UPDATE Data SET Value = ? WHERE ID = ?", value, id)
// 	if execErr != nil {
// 		if rollbackErr := tx.Rollback(); rollbackErr != nil {
// 			log.Fatalf("update failed: %v, unable to rollback: %v\n", execErr, rollbackErr)
// 		}
// 		log.Fatalf("update failed: %v", execErr)
// 	}
// 	if err := tx.Commit(); err != nil {
// 		log.Fatal(err)
// 	}

// 	return "Update Success"
// }

// func Insert(device string, date string, value float64) string {
// 	ctx := context.Background()
// 	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	_, execErr := tx.ExecContext(ctx, "Insert Into Data(Device,Date,Value) VALUES (?,?,?)", device, date, value)
// 	if execErr != nil {
// 		if rollbackErr := tx.Rollback(); rollbackErr != nil {
// 			log.Fatalf("update failed: %v, unable to rollback: %v\n", execErr, rollbackErr)
// 		}
// 		log.Fatalf("update failed: %v", execErr)
// 	}
// 	if err := tx.Commit(); err != nil {
// 		log.Fatal(err)
// 	}

// 	return "Insert Success"
// }
func Insert(datas Models.Data) string {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}
	_, execErr := tx.ExecContext(ctx, "Insert Into Data(Device,Date,Value) VALUES (?,?,?)", datas.Device, datas.Date, datas.Value)
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
func Update(data Models.Data) string {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}
	_, execErr := tx.ExecContext(ctx, "UPDATE Data SET Value = ? WHERE Device = ?", data.Value, data.Device)
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

func Delete(id string) string {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}
	_, execErr := tx.ExecContext(ctx, "Delete From Data WHERE ID = ? ", id)
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
