package collection

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func ItemCollection() *models.Collection {
	collection := &models.Collection{
		Name:       "items",
		Type:       models.CollectionTypeBase,
		ListRule:   nil,
		ViewRule:   types.Pointer("@request.auth.id != ''"),
		CreateRule: types.Pointer(""),
		UpdateRule: types.Pointer("@request.auth.id != ''"),
		DeleteRule: nil,
		Schema: schema.NewSchema(
			&schema.SchemaField{
				Name:     "guid",
				Type:     schema.FieldTypeText,
				Required: false,
				Unique:   false,
			},

			&schema.SchemaField{
				Name:     "feed_id",
				Type:     schema.FieldTypeText,
				Required: false,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "title",
				Type:     schema.FieldTypeText,
				Required: false,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "link",
				Type:     schema.FieldTypeUrl,
				Required: false,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "content",
				Type:     schema.FieldTypeText,
				Required: true,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "date",
				Type:     schema.FieldTypeDate,
				Required: false,
				Unique:   false,
			},


			&schema.SchemaField{
				Name:     "Status",
				Type:     schema.FieldTypeNumber,
				Required: false,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "image_url",
				Type:     schema.FieldTypeText,
				Required: false,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "audio_url",
				Type:     schema.FieldTypeText,
				Required: false,
				Unique:   false,
			},

		),
	}

	return collection
}
