package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/gorilla/websocket"
	"github.com/janwiemers/up/helper"
	"github.com/janwiemers/up/models"
	"github.com/janwiemers/up/websockets"
)

var (
	monitors                 []models.Application
	selectedListItem         int      = 0
	previousSelectedListItem int      = 0
	grid                     *ui.Grid = ui.NewGrid()
)

const (
	iconWarning = "⚠️"
	iconOK      = "🙂"
)

func main() {
	helper.InitViperConfig()
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// monitors, _ = connection.GetMonitors()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	connect()

	t := widgets.NewTable()
	t.Title = "Monitor Overview"
	t.RowSeparator = true
	t.ColumnWidths = []int{22, 10, -1}
	t.Rows = [][]string{
		{"Loading", ""},
	}

	p := widgets.NewParagraph()
	p.Title = "Check Overview"
	p.Text = "0"

	d := widgets.NewTable()
	d.Title = "Monitor Details"
	d.Rows = [][]string{
		{"Loading", ""},
	}

	l := widgets.NewList()
	l.Title = "Monitors"
	l.Rows = createList()
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false

	grid.Set(
		ui.NewRow(0.8, t),
		ui.NewRow(0.2,
			ui.NewCol(0.2, l),
			ui.NewCol(0.8, d),
		),
	)

	ui.Render(grid)

	draw := func() {
		if monitors == nil {
			return
		}
		monitor := monitors[selectedListItem]
		_ = monitor

		t.Rows = buildTable(monitors)
		l.Rows = createList()
		d.Rows = buildMonitorDetails()

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
		}
	}
}

func buildChecks(checks []models.Check) string {
	var b bytes.Buffer
	b.WriteString(" ")
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

func buildTable(monitors []models.Application) [][]string {
	tableRows := [][]string{}

	for i := range monitors {
		monitor := monitors[i]
		checks := buildChecks(monitor.Checks)
		tableRows = append(tableRows, []string{getMonitorName(monitor), fmt.Sprintf(" [%v](fg:blue)", getLastCheckDate(monitor.Checks)), checks})
	}

	return tableRows
}

func connect() {
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/ws",
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println(err)
	}
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			var msg websockets.BroadcastData
			err = json.Unmarshal(message, &msg)
			if err != nil {
				log.Println(err)
			}

			handleMessage(msg)

		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
}

func handleMessage(m websockets.BroadcastData) {
	if m.Type == "monitors" {
		monitors = m.Monitors
	}

	if m.Type == "addCheck" {
		addCheckToMonitor(m.Check)
	}

	if m.Type == "addChecks" {
		for i := range m.Checks {
			addCheckToMonitor(m.Checks[i])
		}
	}
}

func addCheckToMonitor(c models.Check) {
	for i := range monitors {
		monitor := monitors[i]
		if monitor.ID == c.ApplicationID {
			monitors[i].Checks = append(monitors[i].Checks, c)
			if len(monitors[i].Checks) > 24 {
				_, monitors[i].Checks = monitors[i].Checks[0], monitors[i].Checks[1:]
			}
		}
	}
}

func buildMonitorDetails() [][]string {
	monitor := monitors[selectedListItem]
	r := [][]string{
		{"Target", monitor.Target},
		{"Interval", fmt.Sprintf("%.0fs", monitor.Interval.Seconds())},
		{"CreatedAt", monitor.CreatedAt.Format("Mon Jan _2 15:04:05 2006")},
		{"Protocol", monitor.Protocol},
	}
	return r
}

func createList() []string {
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

func getFirstCheckDate(checks []models.Check) string {
	return checks[0].CreatedAt.Format("Mon Jan _2 15:04:05 2006")
}

func getLastCheckDate(checks []models.Check) string {
	max := len(checks) - 1
	return checks[max].CreatedAt.Format("15:04:05")
}

func getMonitor(idx int) models.Application {
	return monitors[idx]
}

func getMonitorName(m models.Application) string {
	color := "green"
	icon := ""
	if m.Degraded == true {
		color = "red"
		icon = iconWarning
	}
	return fmt.Sprintf("[%v %v](fg:%v)", m.Name, icon, color)
}
