package views

import (
	"fmt"

	"github.com/HubertBel/lazyorg/internal/utils"
	"github.com/jroimartin/gocui"
)

type TimeView struct {
	*BaseView
	Body string
}

func NewTimeView() *TimeView {
	tv := &TimeView{
        BaseView: NewBaseView("time"),
	}

	return tv
}

func (tv *TimeView) Update(g *gocui.Gui) error {
	v, err := g.SetView(
		tv.Name,
		tv.X,
		tv.Y,
		tv.X+tv.W,
		tv.Y+tv.H,
	)

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
        v.FgColor = gocui.ColorGreen
	}

	tv.updateBody(v)

	return nil
}

func (tv *TimeView) updateBody(v *gocui.View) {
	initialTime := 12 - tv.H/4
	halfTime := 0
	tv.Body = ""

	for range tv.H {
        s := utils.FormatHour(initialTime, halfTime)
		if halfTime == 0 {
			tv.Body += fmt.Sprintf("%s - \n", s)
			halfTime = 30
		} else {
			tv.Body += fmt.Sprintf("%s \n", s)
			initialTime++
			halfTime = 0
		}
	}

	v.Clear()
	fmt.Fprintln(v, tv.Body)
}
