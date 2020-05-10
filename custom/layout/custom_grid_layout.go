package layout

import (
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"math"
)

type customGridLayout struct {
	Cols int
}

func (g *customGridLayout) countRows(objects []fyne.CanvasObject) int {
	return int(math.Ceil(float64(len(objects)) / float64(g.Cols)))
}

// Get the leading (top or left) edge of a grid cell.
// size is the ideal cell size and the offset is which col or row it's on.
func getLeading(position float64) int {
	ret := position + float64(theme.Padding())

	return int(math.Round(ret))
}

// Layout is called to pack all child objects into a specified size.
// For a GridLayout this will pack objects into a table format with the number
// of columns specified in our constructor.
func (g *customGridLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	//rows := g.countRows(objects)

	//padWidth := (g.Cols - 1) * theme.Padding()
	//padHeight := (rows - 1) * theme.Padding()

	row, col := 0, 0
	var prevX1, prevY1 float64
	for i, child := range objects {
		//cellWidth := float64(child.Size().Width-padWidth)
		//cellHeight := float64(child.Size().Height-padHeight)
		x1 := getLeading(prevX1)
		y1 := (getLeading(prevY1)+child.Size().Height)*row + 1
		prevX1 += float64(child.Size().Width + theme.Padding())
		//prevY1 = float64(child.Size().Height+padHeight)
		//x2 := getTrailing(cellWidth, col)
		//y2 := getTrailing(cellHeight, row)

		child.Move(fyne.NewPos(x1, y1))
		//child.Resize(fyne.NewSize(x2-x1, y2-y1))

		if (i+1)%g.Cols == 0 {
			row++
			col = 0
			prevX1 = 0
		} else {
			col++

		}
	}
}

// MinSize finds the smallest size that satisfies all the child objects.
// For a GridLayout this is the size of the largest child object multiplied by
// the required number of columns and rows, with appropriate padding between
// children.
func (g *customGridLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	rows := g.countRows(objects)
	minSize := fyne.NewSize(0, 0)
	for _, child := range objects {
		minSize = minSize.Union(child.MinSize())
	}

	minContentSize := fyne.NewSize(minSize.Width*g.Cols, minSize.Height*rows)
	return minContentSize.Add(fyne.NewSize(theme.Padding()*(g.Cols-1), theme.Padding()*(rows-1)))
}

// NewGridLayout returns a new GridLayout instance
func NewResizableGridLayout(cols int) fyne.Layout {
	return &customGridLayout{cols}
}
