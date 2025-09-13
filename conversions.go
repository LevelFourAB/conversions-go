package conversions

import (
	"container/heap"
	"context"
	"fmt"
	"sync"
)

// Conversions contains converters between different types.
type Conversions[T comparable, D any] struct {
	converters sync.Map

	typeExtractor func(D) (T, error)
}

// New creates a new Conversions instance.
func New[T comparable, D any](opts ...Option[T, D]) (*Conversions[T, D], error) {
	o := &options[T, D]{}

	for _, opt := range opts {
		if err := opt(o); err != nil {
			return nil, err
		}
	}

	if o.typeExtractor == nil {
		return nil, fmt.Errorf("type extractor is required")
	}

	c := &Conversions[T, D]{
		typeExtractor: o.typeExtractor,
	}

	// Add the conversions specified via options
	for _, conv := range o.converters {
		c.AddConversion(conv.From, conv.To, conv.Converter)
	}
	return c, nil
}

// AddConversion adds a new converter to the conversions.
func (c *Conversions[T, D]) AddConversion(from T, to T, conv Converter[D]) {
	m, _ := c.converters.LoadOrStore(from, &sync.Map{})
	m.(*sync.Map).Store(to, conv)
}

// Convert converts a value from one type to another.
func (c *Conversions[T, D]) Convert(ctx context.Context, value D, to T) (any, error) {
	from, err := c.typeExtractor(value)
	if err != nil {
		return nil, err
	}

	if from == to {
		// No conversion needed
		return value, nil
	}

	conv, err := c.findConverter(from, to)
	if err != nil {
		return nil, err
	}

	return conv(ctx, value)
}

// findConverter will lookup the converter to use between two types. If no
// direct converter is found it will try to find a converter that can convert
// from the source type to an intermediate type and then from that intermediate
// type to the destination type. It does this using a priority queue search,
// where the we always try to find the shortest path first.
func (c *Conversions[T, D]) findConverter(from T, to T) (Converter[D], error) {
	// Get if there is a direct converter
	if value, ok := c.converters.Load(from); ok {
		m := value.(*sync.Map)
		conv, ok := m.Load(to)
		if ok {
			return conv.(Converter[D]), nil
		}
	}

	// Get all converters that can convert from the source type
	converters, ok := c.converters.Load(from)
	if !ok {
		return nil, fmt.Errorf("no converter found from %v to %v", from, to)
	}

	// Create a priority queue to find the shortest path
	queue := &priorityQueue[T, D]{}
	heap.Init(queue)

	// Add all converters that can convert from the source type
	converters.(*sync.Map).Range(func(key, value any) bool {
		convertsTo := key.(T)
		conv := value.(Converter[D])
		heap.Push(queue, &queueItem[T, D]{
			conv:  conv,
			to:    convertsTo,
			depth: 1,
		})
		return true
	})

	// Keep track of previously visited types to avoid loops
	visited := make(map[T]bool)
	visited[from] = true

	// Search for a converter that can convert to the destination type
	for queue.Len() > 0 {
		item := heap.Pop(queue).(*queueItem[T, D])
		intermediateConverter := item.conv

		// Check if the converter can convert to the destination type
		if item.to == to {
			// We can convert to the destination type, so cache the converter
			// for future use
			c.AddConversion(from, to, item.conv)
			return item.conv, nil
		}

		// No direct converter found, so check if there is a converter that
		// can convert from the intermediate type to the destination type
		if converters, ok := c.converters.Load(item.to); ok {
			// Check if we have already visited this type
			if visited[item.to] {
				continue
			}

			// Mark the type as visited
			visited[item.to] = true

			// Add all converters that can convert from the intermediate type
			converters.(*sync.Map).Range(func(key, value any) bool {
				convertsTo := key.(T)
				nextConverter := value.(Converter[D])

				// Create a new converter that will first convert to the
				// intermediate type and then to the destination type
				conv := func(previousConverter, nextConverter Converter[D]) Converter[D] {
					return func(ctx context.Context, value D) (D, error) {
						value, err := previousConverter(ctx, value)
						if err != nil {
							return *new(D), err
						}
						return nextConverter(ctx, value)
					}
				}(intermediateConverter, nextConverter)

				heap.Push(queue, &queueItem[T, D]{
					conv:  conv,
					to:    convertsTo,
					depth: item.depth + 1,
				})
				return true
			})
		}
	}

	return nil, fmt.Errorf("no converter found from %v to %v", from, to)
}

type priorityQueue[T any, D any] []*queueItem[T, D]

type queueItem[T any, D any] struct {
	// The converter to use
	conv Converter[D]

	// The type that the converter converts to
	to T

	// The number of converters that have been used to get to this point
	depth int
}

func (q priorityQueue[T, D]) Len() int {
	return len(q)
}

func (q priorityQueue[T, D]) Less(i, j int) bool {
	return q[i].depth < q[j].depth
}

func (q priorityQueue[T, D]) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *priorityQueue[T, D]) Push(x any) {
	*q = append(*q, x.(*queueItem[T, D]))
}

func (q *priorityQueue[T, D]) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	*q = old[0 : n-1]
	return item
}
