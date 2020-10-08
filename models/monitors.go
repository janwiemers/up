package models

import (
	"time"
)

// Application represents the config file up will monitor
type Application struct {
	ID          int           `gorm:"primarykey"  json:"id"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"-"`
	Name        string        `yaml:"name"        json:"name"`
	Protocol    string        `yaml:"protocol"    json:"protocol"`
	Expectation string        `yaml:"expectation" json:"expectation"`
	Target      string        `yaml:"target"      json:"target"`
	Interval    time.Duration `yaml:"interval"    json:"interval"`
	Checks      []Check       `json:"checks"`
	Label       string        `yaml:"label"       json:"label"`
	Degraded    bool          `json:"degraded"`
	Alerted     bool
}

// Check is the datastructure that holds the checks and their results
type Check struct {
	ID            int         `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time   `json:"createdAt"`
	UpdatedAt     time.Time   `json:"-"`
	UP            bool        `json:"up"`
	ApplicationID int         `json:"applicationId"`
	Application   Application `json:"-"`
}
