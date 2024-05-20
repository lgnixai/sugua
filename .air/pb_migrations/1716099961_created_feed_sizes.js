/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "4p8k3r65917sia0",
    "created": "2024-05-19 06:26:01.058Z",
    "updated": "2024-05-19 06:26:01.058Z",
    "name": "feed_sizes",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "zawj2tkx",
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
        "id": "sh7ruhxb",
        "name": "size",
        "type": "number",
        "required": false,
        "presentable": false,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "noDecimal": false
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
  const collection = dao.findCollectionByNameOrId("4p8k3r65917sia0");

  return dao.deleteCollection(collection);
})
