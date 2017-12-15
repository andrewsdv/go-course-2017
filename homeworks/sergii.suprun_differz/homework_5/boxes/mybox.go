package boxes

import (
	"math/rand"
	"sort"
	"strconv"
	"time"

	"github.com/differz/go-course-2017/homeworks/sergii.suprun_differz/homework_5/shapes"
)

type MyBox struct {
	n     int
	x     int
	state int
	box   []shapes.Shaper
}

func NewMyBox(n int, x int) (BlackBoxer, error) {
	return &MyBox{
			n:     n,
			x:     x,
			state: 0,
			box:   []shapes.Shaper{}},
		nil
}

func (b *MyBox) IsEmpty() bool {
	return len(b.box) == 0
}

func (b *MyBox) GetState() int {
	return b.state
}

func (b *MyBox) PrintWeight() string {
	k := b.n / b.x
	s := ""
	for key, val := range b.box {
		s += strconv.Itoa(val.Weight()) + " "
		if (key+1)%k == 0 {
			s += "\n"
		}
	}
	return s
}

func (b *MyBox) String() string {
	return "(box x=" + strconv.Itoa(b.x) + ")"
}

func (b *MyBox) Generate() {
	usedNames := shapes.GetAvailable()
	x := strconv.Itoa(b.x)
	k := b.n / b.x

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < k*k; i++ {
		j := rand.Intn(len(usedNames))
		shape, _ := shapes.Create(map[string]string{"SHAPE": usedNames[j], "X": x})
		b.box = append(b.box, shape)
	}
}

func (b *MyBox) Shake(corner int) {
	index := 0
	k := b.n / b.x
	t := make([]shapes.Shaper, k*k)

	sort.Sort(shapes.ByWeight(b.box))

	for i := 0; i < k; i++ {
		for j, ii := 0, i; j <= i; j, ii = j+1, ii-1 {
			direct := b.box[index]
			reverse := b.box[k*k-index-1]
			switch corner {
			case 1:
				t[ii*k+j] = direct
				t[(k-ii-1)*k+k-j-1] = reverse
			case 2:
				t[ii*k+k-j-1] = direct
				t[(k-ii-1)*k+j] = reverse
			case 3:
				t[(k-ii-1)*k+j] = direct
				t[ii*k+k-j-1] = reverse
			case 4:
				t[(k-ii-1)*k+k-j-1] = direct
				t[ii*k+j] = reverse
			default:
				return
			}
			index++
		}
	}
	b.box = t
	b.state = corner
}
