package collection

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func FeedCollection() *models.Collection {
	collection := &models.Collection{
		Name:       "feeds",
		Type:       models.CollectionTypeBase,
		ListRule:   nil,
		ViewRule:   types.Pointer("@request.auth.id != ''"),
		CreateRule: types.Pointer(""),
		UpdateRule: types.Pointer("@request.auth.id != ''"),
		DeleteRule: nil,
		Schema: schema.NewSchema(
			&schema.SchemaField{
				Name:     "title",
				Type:     schema.FieldTypeText,
				Required: false,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "folder_id",
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
				Name:     "feed_link",
				Type:     schema.FieldTypeUrl,
				Required: true,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "archive",
				Type:     schema.FieldTypeUrl,
				Required: false,
				Unique:   false,
			},


			&schema.SchemaField{
				Name:     "has_icon",
				Type:     schema.FieldTypeBool,
				Required: false,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "description",
				Type:     schema.FieldTypeText,
				Required: false,
				Unique:   false,
			},

		),
	}

	return collection
}
