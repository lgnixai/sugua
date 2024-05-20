/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("dvmhbl8lsnt9p0z")

  collection.name = "items"

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("dvmhbl8lsnt9p0z")

  collection.name = "item"

  return dao.saveCollection(collection)
})
