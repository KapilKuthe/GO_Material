package data

import (
	"context"
	"fmt"
	"nfscGofiber/database"
	"nfscGofiber/msg"
)

// DataRetriever retrieves data from the database
type DataRetriever struct {
	db database.Database
}

// NewDataRetriever creates a new DataRetriever
func NewDataRetriever(db database.Database) *DataRetriever {
	return &DataRetriever{db: db}
}

// RetrieveData fetches data from the database
func (dr *DataRetriever) RetrieveData(ctx context.Context) ([]msg.MsgRequest, error) {
	rows, err := dr.db.Query(ctx, "SELECT id, timestamp, message, is_active FROM msg_requests where is_active=true ")
	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}
	defer rows.Close()

	var data []msg.MsgRequest
	for rows.Next() {
		var d msg.MsgRequest
		err := rows.Scan(&d.Id, &d.Timestamp, &d.Message, &d.IsActive)
		if err != nil {
			return nil, fmt.Errorf("error scanning rows: %w", err)
		}
		data = append(data, d)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("no data found")
	}
	return data, nil
}

