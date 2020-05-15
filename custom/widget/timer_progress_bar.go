package widget

import (
	"bitbucket.org/avanz/anotherPomodoro/common"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"image/color"
	"time"
)

type CustomProgressBar struct {
	*widget.ProgressBar
	Min, Max, Value time.Duration
	pause, alert    chan bool
	bgColor         color.Color
}

func (bar *CustomProgressBar) Start() {
	ticker := time.NewTicker(1 * time.Second)
	value := bar.Max
	func() {
		for {
			select {
			case p := <-bar.pause:
				if p {
					<-bar.pause
				}
			case <-ticker.C:
				value -= 1 * time.Second
				bar.SetValue(value)
				percentage := (value.Seconds() / bar.Max.Seconds()) * 100
				if percentage == 10 {
					bar.alert <- true
				}
				if value <= bar.Min {
					ticker.Stop()
					return
				}
			}
		}
	}()
}

func NewTimerProgressBar(maxDuration time.Duration, pause, alert chan bool, bgColor color.Color) *CustomProgressBar {
	p := &CustomProgressBar{
		ProgressBar: widget.NewProgressBar(),
		Max:         maxDuration,
		pause:       pause,
		alert:       alert,
		Value:       maxDuration,
		bgColor:     bgColor,
	}
	//p.SetValue(maxDuration)
	widget.Renderer(p).Layout(p.MinSize())
	return p
}

// SetValue changes the current value of this progress bar (from p.Min to p.Max).
// The widget will be refreshed to indicate the change.
func (p *CustomProgressBar) SetValue(v time.Duration) {
	p.Value = v
	widget.Renderer(p).Refresh()
}

// CreateRenderer is a private method to Fyne which links this widget to it's renderer
func (p *CustomProgressBar) CreateRenderer() fyne.WidgetRenderer {
	if p.Min == 0 && p.Max == 0 {
		p.Max = time.Duration(25) * time.Minute
	}

	bar := canvas.NewRectangle(color.Black)
	label := canvas.NewText(common.DurationToString(p.Max), theme.TextColor())
	label.Alignment = fyne.TextAlignCenter
	return &customProgressBarRenderer{[]fyne.CanvasObject{bar, label}, bar, label, p, p.bgColor}
}

type customProgressBarRenderer struct {
	objects []fyne.CanvasObject

	bar   *canvas.Rectangle
	label *canvas.Text

	progress *CustomProgressBar
	bgColor  color.Color
}

// MinSize calculates the minimum size of a progress bar.
// This is simply the "100%" label size plus padding.
func (p *customProgressBarRenderer) MinSize() fyne.Size {
	text := textMinSize(common.DurationToString(25*time.Minute), p.label.TextSize, p.label.TextStyle)

	return fyne.NewSize(text.Width+theme.Padding()*4, text.Height+theme.Padding()*2)
}

func textMinSize(text string, size int, style fyne.TextStyle) fyne.Size {
	t := canvas.NewText(text, color.Black)
	t.TextSize = size
	t.TextStyle = style
	return t.MinSize()
}

func (p *customProgressBarRenderer) updateBar() {
	if p.progress.Value < p.progress.Min {
		p.progress.Value = p.progress.Min
	}
	if p.progress.Value > p.progress.Max {
		p.progress.Value = p.progress.Max
	}

	//ratio := p.progress.Value
	delta := float32(p.progress.Max - p.progress.Min)
	ratio := float32(p.progress.Value-p.progress.Min) / delta

	p.label.Text = common.DurationToString(p.progress.Value)

	size := p.progress.Size()
	//width := int(p.progress.Value.Seconds() / p.progress.Max.Seconds())
	p.bar.Resize(fyne.NewSize(int(float32(size.Width)-float32(size.Width)*ratio), size.Height))
}

// Layout the components of the check widget
func (p *customProgressBarRenderer) Layout(size fyne.Size) {
	p.label.Resize(size)
	p.updateBar()
}

// ApplyTheme is called when the progress bar may need to update it's look
func (p *customProgressBarRenderer) ApplyTheme() {
	p.bar.FillColor = theme.PrimaryColor()
	p.label.Color = theme.TextColor()

	p.Refresh()
}

func (p *customProgressBarRenderer) BackgroundColor() color.Color {
	return p.bgColor
	//return theme.ButtonColor()
}

func (p *customProgressBarRenderer) Refresh() {
	p.updateBar()
	canvas.Refresh(p.progress)
}

func (p *customProgressBarRenderer) Objects() []fyne.CanvasObject {
	return p.objects
}

func (p *customProgressBarRenderer) Destroy() {
}
