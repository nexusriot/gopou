package sorting

import (
	"reflect"
	"sort"
)

type lessFunc func(p1, p2 interface{}) bool
type Sortable interface{}

// Sorter implements the Sort interface, sorting the changes within.
type Sorter struct {
	items []Sortable
	less  []lessFunc
}

// Criteria is a criteria tp perform sorting includes field name and sorting order.
type Criteria struct {
	FieldName string
	Ascending bool
}

// Sorted returns the sorted slice of type of Sortable Interface.
func (s *Sorter) Sorted(changes interface{}) []Sortable {

	reflectCopy := reflect.Indirect(reflect.ValueOf(changes))
	for i := 0; i < reflectCopy.Len(); i++ {
		s.items = append(s.items, reflectCopy.Index(i).Interface())
	}
	sort.Sort(s)
	return s.items
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
func OrderedBy(less ...lessFunc) *Sorter {
	return &Sorter{
		less: less,
	}
}

// Len is part of sort.Interface.
func (s *Sorter) Len() int {
	return len(s.items)
}

// Swap is part of sort.Interface.
func (s *Sorter) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}

// Less is part of sort.Interface.
func (s *Sorter) Less(i, j int) bool {
	p, q := s.items[i], s.items[j]
	var k int
	for k = 0; k < len(s.less)-1; k++ {
		less := s.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return s.less[k](p, q)
}

// function for getting "comparator" function
func getComparator(fieldName string, ascending bool) func(c1, c2 interface{}) bool {
	return func(c1, c2 interface{}) bool {
		rv1 := reflect.Indirect(reflect.ValueOf(c1)).FieldByName(fieldName)
		rv2 := reflect.Indirect(reflect.ValueOf(c2)).FieldByName(fieldName)
		// TODO: check rv1, rv2 should have the same type ??
		switch rv1.Kind() {
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if ascending {
				return rv1.Int() < rv2.Int()
			}
			return rv1.Int() > rv2.Int()
		case reflect.Float32, reflect.Float64:
			if ascending {
				return rv1.Float() < rv2.Float()
			}
			return rv1.Float() > rv2.Float()
		default:
			if ascending {
				return rv1.String() < rv2.String()
			}
			return rv1.String() > rv2.String()
		}
	}
}

func Sorted(criteria []*Criteria, items interface{}) ([]Sortable, error) {
	var lessFuncs []lessFunc
	for _, v := range criteria {
		lessFuncs = append(lessFuncs, getComparator(v.FieldName, v.Ascending))
	}
	return OrderedBy(lessFuncs...).Sorted(items), nil
}
