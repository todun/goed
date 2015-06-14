package ui

import (
	"fmt"

	"github.com/tcolar/goed/core"
	"github.com/tcolar/termbox-go"
)

type Col struct {
	WidthRatio float64
	Views      []*View
}

func (e *Editor) NewCol(width float64, views []*View) *Col {
	if len(views) < 1 {
		panic("Column must have at least one view !")
	}
	return &Col{
		WidthRatio: width,
		Views:      views,
	}
}

func (e *Editor) WidgetAt(x, y int) Renderer {
	_, h := e.term.Size()
	if y == 0 {
		return e.Cmdbar
	}
	if y == h-1 {
		return e.Statusbar
	}
	for _, c := range e.Cols {
		for _, v := range c.Views {
			if x >= v.x1 && x <= v.x2 && y >= v.y1 && y <= v.y2 {
				return v
			}
		}
	}
	return nil
}

func (e Editor) Render() {
	e.TermFB(e.theme.Fg, e.theme.Bg)
	e.term.Clear(e.Bg.Uint16(), e.Bg.Uint16())

	for _, c := range e.Cols {
		for _, v := range c.Views {
			v.Render()
		}
	}

	// cursor
	v := e.curView
	cc, cl := v.CurCol(), v.CurLine()
	c, _, _ := v.CurChar()
	// With some terminals & color schemes the cursor might be "invisible" if we are at a
	// location with no text (ie: end of line)
	// so in that case put as space there to cause the cursor to appear.
	var car = ' '
	if c != nil {
		car = *c
	}
	// Note theterminal inverse the colors where the cursor is
	// this is why this statement might appear "backward"
	e.TermFB(e.theme.BgCursor, e.theme.FgCursor)
	e.TermChar(cc+v.x1-v.offx+2, cl+v.y1-v.offy+2, car)
	e.TermFB(e.theme.Fg, e.theme.Bg)

	e.Cmdbar.Render()
	e.Statusbar.Render()

	e.term.Flush()
}

type Renderer interface {
	Bounds() (int, int, int, int)
	Render()
	SetBounds(x1, y1, x2, y2 int)
	Event(e *Editor, ev *termbox.Event)
}

// TODO: optimize, for example might only need to resize a single column
func (e *Editor) Resize(width, height int) {
	e.Cmdbar.SetBounds(0, 0, width, 0)
	e.Statusbar.SetBounds(0, height-1, width, height-1)
	wc := 0
	wr := 0.0
	for i, c := range e.Cols {
		hc := 1
		hr := 0.0
		w := int(float64(width) * c.WidthRatio)
		if i == len(e.Cols)-1 {
			w = width - wc // las column gets rest of width
			c.WidthRatio = 1.0 - wr
		}
		for j, v := range c.Views {
			h := int(float64(height-2) * v.HeightRatio)
			if h < 1 {
				h = 1
			}
			if j == len(c.Views)-1 {
				h = height - hc - 1 // last view gets rest of height
				v.HeightRatio = 1.0 - hr
			}
			v.SetBounds(wc, hc, wc+w-1, hc+h-1)
			hc += h
			hr += v.HeightRatio
		}
		wc += w
		wr += c.WidthRatio
	}
}

// ViewMove handles moving & resizing views/columns, typically using the mouse
func (e *Editor) ViewMove(x1, y1, x2, y2 int) {
	w, h := e.term.Size()
	v1 := e.WidgetAt(x1, y1).(*View)
	v2 := e.WidgetAt(x2, y2).(*View)
	c1 := e.ViewColumn(v1)
	c2 := e.ViewColumn(v2)
	c1i := e.ColIndex(c1)
	v1i := e.ViewIndex(c1, v1)
	v2i := e.ViewIndex(c2, v2)
	c2i := e.ColIndex(c2)
	onSep := x2 == v2.x1 // dropped on a column "scrollbar"
	if x1 == x2 && y1 == y2 {
		// noop
		e.SetStatus("")
		return
	}
	if onSep {
		// We are moving a single view
		if v1 == v2 {
			if v1i > 0 {
				// Reducing a view
				ratio := float64(y2-y1) / float64(h-2)
				v1.HeightRatio -= ratio
				c1.Views[v1i-1].HeightRatio += ratio // giving space to prev view
			}
		} else if c1i == c2i && v2i == v1i-1 && y2 != v2.y1 {
			// Expanding the view
			ratio := float64(y1-y2) / float64(h-2)
			v1.HeightRatio += ratio
			c1.Views[v1i-1].HeightRatio -= ratio // taking space from prev view
		} else if c1i == c2i {
			// moved within same column
			i1, i2 := v1i, v2i
			if i1 < i2 { // down
				copy(c1.Views[i1:i2], c2.Views[i1+1:i2+1])
				c1.Views[v2i] = v1
			} else { // up
				copy(c1.Views[i2+1:i1+1], c1.Views[i2:i1])
				c1.Views[v2i+1] = v1
			}
		} else {
			// moved to a different column
			// taking space out of target view
			ratio := float64(y2-v2.y1) / float64(v2.y2-v2.y1)
			e.DelView(v1, false)
			v1.HeightRatio = v2.HeightRatio * (1.0 - ratio)
			v2.HeightRatio *= ratio
			c2.Views = append(c2.Views, nil)
			copy(c2.Views[v2i+2:], c2.Views[v2i+1:])
			c2.Views[v2i+1] = v1
		}
	} else {
		// Moving a whole column or a view to it's own column
		if y2 == 1 && v1i > 0 && x2 > v2.x1 && x2 < v2.x2 {
			// Moving a view to it's own column
			ratio := float64(x2-v2.x1) / float64(v2.x2-v2.x1)
			nc := e.AddCol(c2, 1.0-ratio)
			e.DelView(v1, false)
			nc.Views[0] = v1
		} else if c1i == c2i {
			// Reducing a column
			if c1i > 0 {
				ratio := float64(x2-x1) / float64(w)
				c1.WidthRatio -= ratio
				e.Cols[c1i-1].WidthRatio += ratio // giving space to prev col
			}
		} else if c2i == c1i-1 {
			// Expanding the column
			ratio := float64(x1-x2) / float64(w)
			c1.WidthRatio += ratio
			e.Cols[c1i-1].WidthRatio -= ratio // taking space from prev col
		} else {
			//reorder
			i1, i2 := c1i, c2i
			if i1 < i2 { // right
				copy(e.Cols[i1:i2], e.Cols[i1+1:i2+1])
				e.Cols[c2i] = c1
			} else { // left
				copy(e.Cols[i2+1:i1+1], e.Cols[i2:i1])
				e.Cols[c2i+1] = c1
			}

		}
	}
	e.SetStatus("")
	e.Resize(e.term.Size())
}

// ViewColumn returns the column that is holding a given view
func (e *Editor) ViewColumn(v *View) *Col {
	for _, c := range e.Cols {
		for _, view := range c.Views {
			if v == view {
				return c
			}
		}
	}
	return nil
}

// ViewIndex returns the index of a view in the column
func (e *Editor) ViewIndex(col *Col, view *View) int {
	for i, v := range col.Views {
		if v == view {
			return i
		}
	}
	return -1
}

// ColIndex returns the index of a column in the editor
func (e *Editor) ColIndex(col *Col) int {
	for i, c := range e.Cols {
		if c == col {
			return i
		}
	}
	return -1
}

// AddCol adds a new column, space is "taken" from toCol
func (e *Editor) AddCol(toCol *Col, ratio float64) *Col {
	r := toCol.WidthRatio
	nv := e.NewView()
	nv.HeightRatio = 1.0
	c := e.NewCol(r*ratio, []*View{nv})
	toCol.WidthRatio = r - (r * ratio)
	// Insert it after curcol
	i := e.ColIndex(toCol) + 1
	e.Cols = append(e.Cols, nil)
	copy(e.Cols[i+1:], e.Cols[i:])
	e.Cols[i] = c

	e.ActivateView(nv, 0, 0)
	e.Resize(e.term.Size())
	return c
}

// TODO: Not all that smart yet
func (e *Editor) AddViewSmart() *View {
	var nv *View
	var emptiestCol *Col
	for _, c := range e.Cols {
		// if we have room for a new column,favor that
		if c.Views[0].x2-c.Views[0].x1 > 120 {
			nc := e.AddCol(c, 0.5)
			nv = nc.Views[0]
			break
		}
		// otherwise favor emptiest, most to the right col
		if emptiestCol == nil || len(c.Views) <= len(emptiestCol.Views) {
			emptiestCol = c
		}
	}
	if nv == nil {
		// will take some of emptiest col tallest view
		var emptiestView *View
		for _, v := range emptiestCol.Views {
			if emptiestView == nil || emptiestView.HeightRatio <= v.HeightRatio {
				emptiestView = v
			}
		}
		// TODO : consider buffer text length ?
		nv = e.AddView(emptiestView, 0.5)
	}
	e.ActivateView(nv, 0, 0)
	e.Resize(e.term.Size())
	return nv
}

func (e *Editor) InsertViewSmart(view *View) {
	nv := e.AddViewSmart()
	e.ReplaceView(nv, view)
}

// AddCol adds a new view in the current column, space is "taken" from toView
func (e *Editor) AddView(toView *View, ratio float64) *View {
	nv := e.NewView()
	e.InsertView(nv, toView, ratio)
	e.ActivateView(nv, 0, 0)
	return nv
}

func (e *Editor) InsertView(view, toView *View, ratio float64) {
	if ratio > 1.0 {
		ratio = 1.0
	}
	r := toView.HeightRatio
	view.HeightRatio = r * ratio
	toView.HeightRatio = r - (r * ratio)
	col := e.ViewColumn(toView)
	// Insert it at after toView
	i := e.ViewIndex(col, toView) + 1
	col.Views = append(col.Views, nil)
	copy(col.Views[i+1:], col.Views[i:])
	col.Views[i] = view
	e.Resize(e.term.Size())
}

func (e *Editor) ReplaceView(oldView, newView *View) {
	newView.x1 = oldView.x1
	newView.x2 = oldView.x2
	newView.y1 = oldView.y1
	newView.y2 = oldView.y2
	newView.HeightRatio = oldView.HeightRatio
	col := e.ViewColumn(oldView)
	i := e.ViewIndex(col, oldView)
	col.Views[i] = newView
	e.TerminateView(oldView)
}

func (e *Editor) DelCol(col *Col, terminateViews bool) {
	if len(e.Cols) <= 1 {
		e.SetStatusErr("Only one column left !")
		return
	}
	var prev *Col
	for i, c := range e.Cols {
		if c == col {
			if prev != nil {
				prev.WidthRatio += c.WidthRatio
				e.CurCol = prev
			} else {
				e.Cols[i+1].WidthRatio += c.WidthRatio
				e.CurCol = e.Cols[i+1]
			}
			e.ActivateView(e.CurCol.Views[0], 0, 0)
			e.Cols = append(e.Cols[:i], e.Cols[i+1:]...)
			break
		}
		prev = e.Cols[i]
	}
	if terminateViews {
		for _, v := range col.Views {
			e.TerminateView(v)
		}
	}
	col = nil
	e.Resize(e.term.Size())
}

func (e *Editor) DelView(view *View, terminate bool) {
	c := e.ViewColumn(view)
	if len(e.Cols) <= 1 && len(c.Views) <= 1 {
		e.SetStatusErr("Only one view left !")
		return
	}
	// only one view left in col, delcol
	if len(c.Views) <= 1 {
		e.DelCol(c, terminate)
		return
	}
	// otherwise remove curview and reassign space
	var prev *View
	for i, v := range c.Views {
		if v == view {
			if prev != nil {
				prev.HeightRatio += v.HeightRatio
				e.curView = prev
			} else {
				c.Views[i+1].HeightRatio += v.HeightRatio
				e.curView = c.Views[i+1]
			}
			c.Views = append(c.Views[:i], c.Views[i+1:]...)
			e.ActivateView(e.curView, 0, 0)
			if terminate {
				e.TerminateView(v)
			}
			break
		}
		prev = c.Views[i]
	}
	e.Resize(e.term.Size())
}

func (e *Editor) TerminateView(v *View) {
	if v == nil {
		return
	}
	if v.backend != nil {
		v.backend.Close()
	}
	v = nil
}

// Delete (close) a view, but with dirty check
func (e *Editor) DelViewCheck(v *View) {
	if !v.canClose() {
		e.SetStatusErr("Unsaved changes. Save or request close again.")
		return
	}
	e.DelView(v, true)
}

// Delete (close) a col, but with dirty check
func (e *Editor) DelColCheck(c *Col) {
	ok := true
	for _, v := range c.Views {
		ok = ok && v.canClose()
	}
	if !ok {
		e.SetStatusErr("Unsaved changes. Save or request close again.")
		return
	}
	e.DelCol(c, true)
}

func (e *Editor) ActivateView(v *View, cursorx, cursory int) {
	e.curView = v
	e.CurCol = e.ViewColumn(v)
	v.MoveCursor(cursorx-v.CurCol(), cursory-v.CurLine())
}

func (e *Editor) SetCurView(id int64) error {
	v := e.ViewById(id)
	if v == nil {
		return fmt.Errorf("No such view %d", id)
	}
	e.ActivateView(v.(*View), 0, 0)
	return nil
}

func (e Editor) ViewById(id int64) core.Viewable {
	for _, c := range e.Cols {
		for _, v := range c.Views {
			if v.Id() == id {
				return v
			}
		}
	}
	return nil
}

func (e *Editor) ViewByLoc(loc string) *View {
	if len(loc) == 0 {
		return nil
	}
	for _, c := range e.Cols {
		for _, v := range c.Views {
			if v.backend.SrcLoc() == loc {
				return v
			}
		}
	}
	return nil
}
