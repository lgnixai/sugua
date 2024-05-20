/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "nfpzlmvpp531zmg",
    "created": "2024-05-18 15:39:11.800Z",
    "updated": "2024-05-18 15:39:11.800Z",
    "name": "feeds",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "ltabg8io",
        "name": "title",
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
  const collection = dao.findCollectionByNameOrId("nfpzlmvpp531zmg");

  return dao.deleteCollection(collection);
})
