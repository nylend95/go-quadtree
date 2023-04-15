package quadtree

import (
	"reflect"
	"testing"
)

func Test_Insert(t *testing.T) {
	qt := NewQuadtree(0, 0, 100, 100, 4)
	type args struct {
		p Point
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{Position{0, 0}}, true},
		{args{Position{100, 100}}, false},
		{args{Position{50, 0}}, true},
		{args{Position{0, 50}}, true},
		{args{Position{20, 10}}, true},
		{args{Position{99, 99}}, true},
		{args{Position{120, 110}}, false},
		{args{Position{-1, -2}}, false},
	}
	for _, tt := range tests {
		t.Run("Testing", func(t *testing.T) {
			if got := qt.Insert(tt.args.p); got != tt.want {
				t.Errorf("qt.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_InsertDividedAtLimit(t *testing.T) {
	qt := NewQuadtree(0, 0, 100, 100, 4)
	type args struct {
		p Point
	}
	tests := []struct {
		args      args
		isDivided bool
	}{
		{args{Position{0, 0}}, false},
		{args{Position{0, 0}}, false},
		{args{Position{0, 0}}, false},
		{args{Position{0, 0}}, false},
		{args{Position{0, 0}}, true},
	}
	for _, tt := range tests {
		t.Run("Testing", func(t *testing.T) {
			qt.Insert(tt.args.p)
			if qt.isDivided != tt.isDivided {
				t.Errorf("qt.Insert(), quadtree should be divided %v, got: %v", qt.isDivided, tt.isDivided)
			}
		})
	}
}

func Test_isWithinBounds(t *testing.T) {
	type args struct {
		p Position
		b rect
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{Position{0, 0}, defaultRect()}, true},
		{args{Position{100, 100}, defaultRect()}, false},
		{args{Position{50, 0}, defaultRect()}, true},
		{args{Position{0, 50}, defaultRect()}, true},
		{args{Position{20, 10}, defaultRect()}, true},
		{args{Position{99, 99}, defaultRect()}, true},
		{args{Position{120, 110}, defaultRect()}, false},
		{args{Position{-1, -2}, defaultRect()}, false},
	}
	for _, tt := range tests {
		t.Run("Testing", func(t *testing.T) {
			if got := isWithinBounds(tt.args.p, tt.args.b); got != tt.want {
				t.Errorf("isWithinBounds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuadtree_Query(t *testing.T) {
	tests := []struct {
		name string
		q    Quadtree
		b    Boundary
		want []Point
	}{
		{"Not yet divided", getQuadtree([]Point{Position{0, 0}, Position{0, 55}, Position{55, 0}, Position{55, 55}}), rect{0, 0, 50, 50}, []Point{Position{0, 0}}},
		{"Divided, all in north east", getQuadtree([]Point{Position{0, 0}, Position{1, 1}, Position{0, 55}, Position{55, 0}, Position{55, 55}}), rect{0, 0, 50, 50}, []Point{Position{0, 0}, Position{1, 1}}},
		{"Divided, get in first and second quad", getQuadtree([]Point{Position{0, 0}, Position{0, 55}, Position{55, 0}, Position{55, 55}, Position{1, 1}}), rect{0, 0, 50, 50}, []Point{Position{0, 0}, Position{1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Query(tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quadtree.Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getQuadtree(points []Point) Quadtree {
	qt := NewQuadtree(0, 0, 100, 100, 4)
	for _, p := range points {
		qt.Insert(p)
	}
	return *qt
}

type Position struct {
	x float32
	y float32
}

func defaultRect() rect {
	return rect{0, 0, 100, 100}
}

type rect struct {
	x      float32
	y      float32
	width  float32
	height float32
}

func (q Position) X() float32 {
	return q.x
}

func (q Position) Y() float32 {
	return q.y
}

func (q rect) X() float32 {
	return q.x
}

func (q rect) Y() float32 {
	return q.y
}

func (q rect) Width() float32 {
	return q.width
}

func (q rect) Height() float32 {
	return q.height
}
