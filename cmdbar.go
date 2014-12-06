package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Cmdbar widget
type Cmdbar struct {
	Widget
	Cmd     []rune
	History [][]rune // TODO : cmd history
}

func (c *Cmdbar) Render() {
	Ed.FB(Ed.Theme.Cmdbar.Fg, Ed.Theme.Cmdbar.Bg)
	Ed.Fill(Ed.Theme.Cmdbar.Rune, c.x1, c.y1, c.x2, c.y2)
	if Ed.CmdOn {
		Ed.FB(Ed.Theme.CmdbarTextOn, Ed.Theme.Cmdbar.Bg)
		Ed.Str(c.x1, c.y1, fmt.Sprintf("> %s", string(c.Cmd)))
	} else {
		Ed.FB(Ed.Theme.CmdbarText, Ed.Theme.Cmdbar.Bg)
		Ed.Str(c.x1, c.y1, fmt.Sprintf("> %s", string(c.Cmd)))
	}
	Ed.FB(Ed.Theme.CmdbarText, Ed.Theme.Cmdbar.Bg)
	Ed.Str(c.x2-11, c.y1, fmt.Sprintf("|GoEd %s", Version))
}

func (c *Cmdbar) RunCmd() {
	// TODO: This is temporary until I create real fs based events & actions
	s := string(c.Cmd)
	parts := strings.Fields(s)
	if len(parts) < 1 {
		return
	}
	args := parts[1:]
	var err error
	switch parts[0] {
	//case "d", "del": // as vi del
	//case "dd":
	case "dc", "delcol":
		Ed.DelCol(Ed.CurCol)
	case "dv", "delview":
		Ed.DelView(Ed.CurView)
	case "e", "exec":
		c.exec(args)
	//case "gf", "gofmt":
	//	Ed.SetStatus("TBD gofmt")
	//case "h", "help":
	//	Ed.SetStatus("TBD help")
	case "nc", "newcol":
		c.newCol(args)
	case "nv", "newview":
		c.newView(args)
	case "o", "open":
		err = c.open(args)
	case "p", "paste": // as vi
		c.paste(args)
	case "s", "save":
		c.save(args)
	case "y", "yank": // as vi copy
		err = c.yank(args)
	case "yy":
		err = c.yank([]string{"1"})
	default:
		Ed.SetStatusErr("Unexpected command " + parts[0])
	}
	if err == nil {
		Ed.CmdOn = false
	} else {
		Ed.SetStatus(err.Error())
	}
}

func (c *Cmdbar) paste(args []string) {
	v := Ed.CurView
	v.MoveCursor(-v.CurCol(), 1)
	l := v.CurLine()
	v.Paste()
	v.InsertNewLine()
	v.MoveCursor(-v.CurCol(), l-v.CurLine())
	v.Dirty = true
}

func (c *Cmdbar) yank(args []string) error {
	v := Ed.CurView
	if len(args) == 0 {
		return fmt.Errorf("Expected an argument")
	}
	nb, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("Expected a numericargument")
	}
	nb--
	Ed.CurView.Copy(
		Selection{
			LineFrom: v.CurLine(),
			LineTo:   v.CurLine() + nb,
			ColTo:    v.LineLen(v.CurLine() + nb),
		})
	return nil
}

func (c *Cmdbar) open(args []string) error {
	if len(args) < 1 {
		// try to expand a location from the current view
		return fmt.Errorf("No path provided")
	}
	v := Ed.NewView()
	err := Ed.Open(args[0], v, "")
	if err != nil {
		return err
	}
	if Ed.CurView.Dirty {
		Ed.InsertViewSmart(v)
		Ed.SetStatus("insert")
	} else {
		Ed.ReplaceView(Ed.CurView, v)
		Ed.SetStatus("replace")
	}
	Ed.CurView = v
	return nil
}

// Open what's selected or under the cursor
// if newView is true then open in a new view, otherwise
// replace content of v
func (c *Cmdbar) OpenSelection(v *View, newView bool) {
	newView = newView || v.Dirty
	text := ""
	if len(v.Selections) > 0 {
		text = Ed.RunesToString(v.Selections[0].Text(v))
	} else {
		// TODO: parse line !
		text = string(v.Line(v.CurLine()))
	}
	v2 := Ed.NewView()
	if err := Ed.Open(text, v2, v.Buffer.file); err != nil {
		Ed.SetStatusErr(err.Error())
	}
	if newView {
		if strings.HasSuffix(text, string(os.PathSeparator)) {
			Ed.InsertView(v2, v, 0.5)
		} else {
			Ed.InsertViewSmart(v2)
		}
	} else {
		Ed.ReplaceView(v, v2)
	}
	Ed.CurView = v
}

func (c *Cmdbar) save(args []string) {
	//TODO: might need a new id etc.../
	if Ed.CurView != nil {
		Ed.CurView.Save()
	}
}

func (c *Cmdbar) newCol(args []string) {
	// nc : newblank col
	// nc [file], new col, open file
	// nc 40 -> new col 40% width
	// nc 40 [file] -> new col 40% width, open file
	loc := ""
	pct := 50
	if len(args) > 0 {
		p, err := strconv.Atoi(args[0])
		if err == nil {
			pct = p
			if len(args) > 1 {
				loc = strings.Join(args[1:], " ")
			}
		} else {
			loc = strings.Join(args, " ")
		}
	}
	v := Ed.AddCol(Ed.CurCol, float64(pct)/100.0).Views[0]
	if len(loc) > 0 {
		Ed.Open(loc, v, "")
	}
}

func (c *Cmdbar) newView(args []string) {
	loc := ""
	pct := 50
	if len(args) > 0 {
		p, err := strconv.Atoi(args[0])
		if err == nil {
			pct = p
			if len(args) > 1 {
				loc = strings.Join(args[1:], " ")
			}
		} else {
			loc = strings.Join(args, " ")
		}
	}
	v := Ed.AddView(Ed.CurView, float64(pct)/100.0)
	if len(loc) > 0 {
		Ed.Open(loc, v, "")
	}
}

func (c *Cmdbar) exec(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	v := Ed.AddViewSmart()
	v.Cmd = cmd
	go v.Exec()
}
