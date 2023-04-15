package quadtree

type Point interface {
	X() float32
	Y() float32
}

type Boundary interface {
	X() float32
	Y() float32
	Width() float32
	Height() float32
}

type Quadtree struct {
	objects   []Point
	x         float32
	y         float32
	width     float32
	height    float32
	limit     int
	isDivided bool
	nw        *Quadtree
	ne        *Quadtree
	se        *Quadtree
	sw        *Quadtree
}

func NewQuadtree(x, y, width, height float32, limit int) *Quadtree {
	return &Quadtree{
		x:      x,
		y:      y,
		width:  width,
		height: height,
		limit:  limit,
	}
}

// Insert inserts a Point into the Quadtree. If the point is outside the bounds of the
// Quadtree, this method will return false, indicating that the Point was not inserted
func (q *Quadtree) Insert(p Point) bool {
	if len(q.objects) < q.limit {
		if isWithinBounds(p, q) {
			q.objects = append(q.objects, p)
			return true
		}
		return false
	}
	q.divide()
	return q.ne.Insert(p) || q.nw.Insert(p) || q.se.Insert(p) || q.sw.Insert(p)
}

func (q Quadtree) Query(b Boundary) []Point {
	var points []Point
	for _, p := range q.objects {
		if isWithinBounds(p, b) {
			points = append(points, p)
		}
	}
	if q.isDivided {
		points = append(points, q.ne.Query(b)...)
		points = append(points, q.nw.Query(b)...)
		points = append(points, q.sw.Query(b)...)
		points = append(points, q.se.Query(b)...)
	}

	return points
}

func isWithinBounds(p Point, b Boundary) bool {
	return p.X() >= b.X() && p.X() < b.X()+b.Width() &&
		p.Y() >= b.Y() && p.Y() < b.Y()+b.Height()
}

func (q *Quadtree) divide() {
	if q.isDivided {
		return
	}
	q.nw = &Quadtree{
		x:      q.x,
		y:      q.y,
		width:  q.width / 2,
		height: q.height / 2,
		limit:  q.limit,
	}
	q.ne = &Quadtree{
		x:      q.x + q.width/2,
		y:      q.y,
		width:  q.width / 2,
		height: q.height / 2,
		limit:  q.limit,
	}
	q.se = &Quadtree{
		x:      q.x + q.width/2,
		y:      q.y + q.height/2,
		width:  q.width / 2,
		height: q.height / 2,
		limit:  q.limit,
	}
	q.sw = &Quadtree{
		x:      q.x,
		y:      q.y + q.height/2,
		width:  q.width / 2,
		height: q.height / 2,
		limit:  q.limit,
	}
	q.isDivided = true
}

func (q Quadtree) X() float32 {
	return q.x
}

func (q Quadtree) Y() float32 {
	return q.y
}

func (q Quadtree) Width() float32 {
	return q.width
}

func (q Quadtree) Height() float32 {
	return q.height
}
