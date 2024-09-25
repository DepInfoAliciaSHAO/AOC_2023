package utils

// This code is adapted from swift, licensed under the MIT License.
// Copyright (c) 2016 Matthijs Hollemans and contributors
type FIFO[T any] struct {
	Queue []T
}

func (f FIFO[T]) Enqueue(element T) FIFO[T] {
	f.Queue = append(f.Queue, element) // Simply append to enqueue.
	return FIFO[T]{f.Queue}
}

func (f FIFO[T]) Dequeue() (FIFO[T], T) {
	element := f.Queue[0]                // The first element is the one to be dequeued.
	return FIFO[T]{f.Queue[1:]}, element // Slice off the element once it is dequeued.
}

func (f FIFO[T]) Len() int {
	return len(f.Queue)
}
