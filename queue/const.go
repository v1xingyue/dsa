package queue

import(
    "errors"
)

var (
    ErrorQueueEmpty error = errors.New("Queue is Empty")
)
