package collection

import (
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func BookmarksCollection() *models.Collection {
	collection := &models.Collection{
		Name:       "bookmarks",
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
				Required: true,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "creator",
				Type:     schema.FieldTypeText,
				Required: true,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "url",
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
				Name:     "tags",
				Type:     schema.FieldTypeRelation,
				Required: true,
				Options: &schema.RelationOptions{
					MaxSelect: types.Pointer(5),

					CascadeDelete: false,
				},
			},
			&schema.SchemaField{
				Name:     "type",
				Type:     schema.FieldTypeSelect,
				Required: true,
				Unique:   false,
				Options: &schema.SelectOptions{
					MaxSelect: 1,
					Values:    []string{"articles", "comics", "podcasts", "videos"},
				},
			},
			&schema.SchemaField{
				Name:     "dead",
				Type:     schema.FieldTypeBool,
				Required: false,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "shared",
				Type:     schema.FieldTypeBool,
				Required: false,
				Unique:   false,
			},
			&schema.SchemaField{
				Name:     "comments",
				Type:     schema.FieldTypeText,
				Required: false,
				Unique:   false,
			},
		),
	}

	return collection
}
