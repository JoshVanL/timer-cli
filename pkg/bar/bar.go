package bar

import ()

type Bar struct {
	total   int
	current int
	width   int
	head    string
	empty   string
	fill    string
	left    string
	right   string
}

func New(total int) *Bar {
	return &Bar{
		total:   total,
		current: 0,
		width:   70,
		head:    ">",
		empty:   "-",
		fill:    "=",
		left:    "[",
		right:   "]",
	}
}

func (b *Bar) Increment() (finished bool) {
	b.current++
	if b.current >= b.total {
		return true
	}

	return false
}

func (b *Bar) GetString() string {
	step := float64(b.width) / float64(b.total)

	out := b.left
	for n := 0; n < int((float64(b.current) * step)); n++ {
		out += b.fill
	}
	out += b.head

	for n := 0; n < int((float64(b.total-b.current) * step)); n++ {
		out += b.empty
	}
	out += b.right

	return out
}
