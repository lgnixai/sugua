package collection

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func FeedErrorCollection() *models.Collection {
	collection := &models.Collection{
		Name:       "feed_error",
		Type:       models.CollectionTypeBase,
		ListRule:   nil,
		ViewRule:   types.Pointer("@request.auth.id != ''"),
		CreateRule: types.Pointer(""),
		UpdateRule: types.Pointer("@request.auth.id != ''"),
		DeleteRule: nil,
		Schema: schema.NewSchema(


			&schema.SchemaField{
				Name:     "error",
				Type:     schema.FieldTypeText,
				Required: false,
				Unique:   false,
			},


		),
	}

	return collection
}
