package grid

type Cell struct {
	Value int
	Row   int
	Col   int
}

type Grid struct {
	Area
	Cells []Cell
}

type Area struct {
	Width  int
	Height int
}

func NewGrid() *Grid {
	return &Grid{
		Area: Area {
			Width:  0,
			Height: 0,
		},
		Cells:  make([]Cell, 0, 1),
	}
}

func (g *Grid) AddValues(iRow int, values []int) {
	numValues := len(values)
	if numValues == 0 {
		return
	}

	if len(g.Cells)%numValues != 0 {
		return
	}

	for i, value := range values {
		g.Cells = append(g.Cells, Cell{value, iRow, i})
	}

	g.Width = numValues
	g.Height = iRow + 1
}
