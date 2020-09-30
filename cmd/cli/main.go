package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/janwiemers/up/connection"
	"github.com/janwiemers/up/helper"
	"github.com/janwiemers/up/models"
)

var (
	monitors                 []models.Application
	checks                   []models.Check
	selectedListItem         int      = 0
	previousSelectedListItem int      = 0
	tickerCount              int      = 1
	grid                     *ui.Grid = ui.NewGrid()
	d                        *widgets.Table
	l                        *widgets.List

	statusOK     string = "OK"
	statusFailed string = "FAILED"
)

func initialize() {
	monitors, _ = connection.GetMonitors()
	time.Sleep(30 * time.Second)
	go initialize()
}

func main() {
	helper.InitViperConfig()
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	go initialize()

	createList := func() []string {
		b := []string{}
		for i := range monitors {
			monitor := monitors[i]
			color := "green"
			if monitor.Degraded == true {
				color = "red"
			}
			b = append(b, fmt.Sprintf("[%v](fg:%v)", monitor.Name, color))
		}

		return b
	}

	// monitors, _ = connection.GetMonitors()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	l := widgets.NewList()
	l.Title = "Monitors"
	l.Rows = createList()
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false

	t := widgets.NewTable()
	t.Title = "Monitor Details"
	t.TextStyle = ui.NewStyle(ui.ColorWhite)
	t.RowSeparator = false
	t.Rows = [][]string{
		{"Loading", ""},
	}

	p := widgets.NewParagraph()
	p.Title = "Check Overview"
	p.Text = "0"

	s := widgets.NewParagraph()
	s.Title = "Status"
	s.Text = ""

	s.BorderStyle = ui.NewStyle(ui.ColorGreen)
	s.TextStyle = ui.NewStyle(ui.ColorGreen)

	grid.Set(
		ui.NewRow(1.0,
			ui.NewCol(0.2, l),
			ui.NewCol(0.8,
				ui.NewRow(1.0,
					ui.NewRow(0.4,
						ui.NewCol(0.2, s),
						ui.NewCol(0.8, t),
					),
					ui.NewRow(0.6, p),
				),
			),
		),
	)

	ui.Render(grid)

	draw := func() {
		if monitors == nil {
			return
		}
		monitor := monitors[selectedListItem]

		if selectedListItem != previousSelectedListItem {
			checks, _ = connection.GetChecks(monitor.ID)
		}

		color := ui.NewStyle(ui.ColorGreen)
		word := calcCenter(s, statusOK)
		if monitor.Degraded == true {
			color = ui.NewStyle(ui.ColorRed)
			word = calcCenter(s, statusFailed)
		}

		s.Text = word
		s.BorderStyle = color
		s.TextStyle = color

		t.Rows = [][]string{
			{"Target", monitor.Target},
			{"Interval", fmt.Sprintf("%.0fs", monitor.Interval.Seconds())},
			{"CreatedAt", monitor.CreatedAt.Format("Mon Jan _2 15:04:05 2006")},
			{"Protocol", monitor.Protocol},
		}

		l.Rows = createList()

		p.Text = buildChecks(checks)
		ui.Render(grid)
		previousSelectedListItem = selectedListItem
	}

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				grid.SetRect(0, 0, payload.Width, payload.Height)
				ui.Clear()
				draw()
			case "j", "<Down>":
				l.ScrollDown()
				selectedListItem = l.SelectedRow
				draw()
			case "k", "<Up>":
				l.ScrollUp()
				selectedListItem = l.SelectedRow
				draw()
			}
		case <-ticker:
			draw()
			tickerCount++
		}
	}
}

func buildChecks(checks []models.Check) string {
	var b bytes.Buffer
	b.WriteString("\n ")
	for c := range checks {
		check := checks[c]

		color := "green"
		if check.UP == false {
			color = "red"
		}
		b.WriteString(fmt.Sprintf("[%v](fg:%v) ", "██", color))
	}

	return b.String()
}

func calcCenter(block *widgets.Paragraph, subject string) string {
	var b bytes.Buffer
	lines := block.Dy() / 2
	length := len(subject)
	for i := 1; i < lines; i++ {
		b.WriteString("\n")
	}

	indent := (block.Dx() - length) / 2
	for i := 1; i < indent; i++ {
		b.WriteString(" ")
	}
	b.WriteString(subject)
	return b.String()
}
