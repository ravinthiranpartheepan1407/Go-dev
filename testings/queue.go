package queue

type Queue struct {
	products []int
	id       int
}

func New(id int) Queue {
	return Queue{make([]int, 0, id), id}
}

func (q *Queue) Added(item int) bool {
	if len(q.products) == q.id {
		return false
	} else {
		q.products = append(q.products, item)
		return true
	}
}

func (q *Queue) Next() (int, bool) {
	if len(q.products) > 0 {
		next := q.products[0]
		q.products = q.products[1:]
		return next, false
	} else {
		return 0, false
	}
}
