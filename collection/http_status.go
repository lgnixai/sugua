package collection

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func HttpStatusCollection() *models.Collection {
	collection := &models.Collection{
		Name:       "http_status",
		Type:       models.CollectionTypeBase,
		ListRule:   nil,
		ViewRule:   types.Pointer("@request.auth.id != ''"),
		CreateRule: types.Pointer(""),
		UpdateRule: types.Pointer("@request.auth.id != ''"),
		DeleteRule: nil,
		Schema: schema.NewSchema(
			&schema.SchemaField{
				Name:     "feed_id",
				Type:     schema.FieldTypeText,
				Required: false,
				Unique:   false,
			},

			&schema.SchemaField{
				Name:     "last_refreshed",
				Type:     schema.FieldTypeDate,
				Required: false,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "last_modified",
				Type:     schema.FieldTypeText,
				Required: false,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "etag",
				Type:     schema.FieldTypeText,
				Required: false,
				Unique:   false,
			},


		),
	}

	return collection
}
