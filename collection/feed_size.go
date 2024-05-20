package collection

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func FeedSizeCollection() *models.Collection {
	collection := &models.Collection{
		Name:       "feed_size",
		Type:       models.CollectionTypeBase,
		ListRule:   nil,
		ViewRule:   types.Pointer("@request.auth.id != ''"),
		CreateRule: types.Pointer(""),
		UpdateRule: types.Pointer("@request.auth.id != ''"),
		DeleteRule: nil,
		Schema: schema.NewSchema(


			&schema.SchemaField{
				Name:     "size",
				Type:     schema.FieldTypeNumber,
				Required: false,
				Unique:   false,
			},


		),
	}

	return collection
}
