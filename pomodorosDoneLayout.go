package main

import (
	"fyne.io/fyne"
)

type pomodorosDoneLayout struct {
	Cols int
}

func NewPomodorosDoneLayout() *pomodorosDoneLayout {
	return &pomodorosDoneLayout{}
}

func (d *pomodorosDoneLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := 0, 0
	for _, o := range objects {
		childSize := o.Size()

		w += childSize.Width
		if childSize.Height > h {
			h = childSize.Height
		}
	}
	return fyne.NewSize(w, h)
}

func (d *pomodorosDoneLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	//row := 0
	//padWidth := (d.Cols - 1) * theme.Padding()
	//
	//pos := fyne.NewPos(10, containerSize.Height-d.MinSize(objects).Height)
	//for i, o := range objects {
	//	size := o.Size()
	//	o.Resize(size)
	//	o.Move(pos)
	//
	//	if (i+1)%d.Cols == 0 {
	//		row++
	//		col = 0
	//
	//	} else {
	//		col++
	//
	//	}
	//	pos = pos.Add(fyne.NewPos(size.Width+theme.Padding(), 0))
	//}
}
