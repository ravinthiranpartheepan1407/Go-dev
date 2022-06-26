package queue

import "testing"

func TestIsValidQueue(t *testing.T) {
	q := New(5)
	for i := 0; i < 5; i++ {
		if len(q.products) != i {
			t.Errorf("The products are not listed(%v), (%v)=false", len(q.products), i)
		}

		if !q.Added(i) {
			t.Errorf("The Products cannot be added after its limit(%v)=false", i)
		}

		if q.Added(6) {
			t.Fail()
		}
	}
}

func TestNext(t *testing.T) {
	q := New(3)

	for i := 0; i < 3; i++ {
		q.Added(i)
	}

	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("Not able to get products")
		}
		if item != i {
			t.Errorf("Product listed in wrong order(%v),(%v)=false", item, i)
		}
	}
	item, ok := q.Next()
	if ok {
		t.Errorf("Cannot list anymore items(%v)=false", item)
	}
}
