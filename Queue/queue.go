package queue

import "fmt"

type queue struct {
	list []string
	size int
}

type queueError string

const (
	queueEmptyError  queueError = "queue is empty."
	queueIsFullError queueError = "queue is full."
)

type IQueue interface {
	EnQueue(val string) error
	DeQueue() (string, error) //val-> of dequeue
	Len() int
	IsEmpty() bool
	IsFull() bool
	Front() (string, error)
	Rear() (string, error) //end of the queue
	Print()
}

func InitQueue(size int) IQueue {
	return &queue{
		list: []string{},
		size: size,
	}
}

func (q *queue) Len() int      { return len((*q).list) }
func (q *queue) IsEmpty() bool { return len((*q).list) == 0 }
func (q *queue) IsFull() bool  { return len((*q).list) == (*q).size }
func (q *queue) Front() (string, error) {
	if (*q).IsEmpty() {
		return "", fmt.Errorf(string(queueEmptyError))
	}
	return (*q).list[0], nil
}

func (q *queue) Rear() (string, error) {
	if (*q).IsEmpty() {
		return "", fmt.Errorf(string(queueEmptyError))
	}
	lastvalIndex := len((*q).list)
	return (*q).list[lastvalIndex], nil
}

func (q *queue) EnQueue(val string) error {
	if (*q).IsFull() {
		return fmt.Errorf(string(queueIsFullError))
	}
	(*q).list = append((*q).list, val)
	return nil
}

func (q *queue) DeQueue() (string, error) {
	if (*q).IsEmpty() {
		return "", fmt.Errorf(string(queueEmptyError))
	}
	dequeueItem := (*q).list[0]
	(*q).list = (*q).list[1:]
	return dequeueItem, nil
}

func (q *queue) Print() {
	for _, val := range (*q).list {
		fmt.Printf("%v\n", val)
	}
}
