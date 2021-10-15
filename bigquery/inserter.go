package bigquery

import "cloud.google.com/go/bigquery"

type ValueSaver interface {
	Save() (map[string]bigquery.Value, string, error)
}

func (c Client) InsertRows(datasetID string, tableID string, items []ValueSaver) error {
	// Create inserter
	inserter := c.BQClient.Dataset(datasetID).Table(tableID).Inserter()

	// Insert list of items
	if err := inserter.Put(c.ctx, items); err != nil {
		log.Fatalf("Could not save to BigQuery %v", err)
		return err
	}

	return nil
}

func (c Client) InsertRow(datasetID string, tableID string, item ValueSaver) error {
	// Create inserter
	inserter := c.BQClient.Dataset(datasetID).Table(tableID).Inserter()

	// Insert list of items
	if err := inserter.Put(c.ctx, item); err != nil {
		log.Fatalf("Could not save to BigQuery %v", err)
		return err
	}

	return nil
}
