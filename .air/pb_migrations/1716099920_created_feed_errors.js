/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "9vhtpll87z2ui9w",
    "created": "2024-05-19 06:25:20.584Z",
    "updated": "2024-05-19 06:25:20.584Z",
    "name": "feed_errors",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "r4mphxru",
        "name": "feed_id",
        "type": "text",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      },
      {
        "system": false,
        "id": "c2dpuoxw",
        "name": "error",
        "type": "text",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      }
    ],
    "indexes": [],
    "listRule": null,
    "viewRule": null,
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {}
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("9vhtpll87z2ui9w");

  return dao.deleteCollection(collection);
})
