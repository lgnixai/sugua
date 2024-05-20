package worker

import (
	`github.com/pocketbase/pocketbase`
)

var (
	MyWorker *Worker
)

func InitWorker(app *pocketbase.PocketBase) {

	MyWorker = NewWorker(app)
}
