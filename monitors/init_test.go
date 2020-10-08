package monitors

import (
	"reflect"
	"testing"
	"time"

	"github.com/janwiemers/up/models"
)

func TestPopulateDefaults(t *testing.T) {
	tests := []struct {
		name string
		args models.Application
		want models.Application
	}{
		{
			name: "Populate all defaults",
			want: models.Application{
				Name:        "test",
				Protocol:    "http",
				Expectation: "200",
				Interval:    5 * time.Minute,
			},
			args: models.Application{
				Name: "test",
			},
		},
		{
			name: "Populate some defaults",
			want: models.Application{
				Name:        "test",
				Protocol:    "dns",
				Expectation: "200",
				Interval:    5 * time.Minute,
			},
			args: models.Application{
				Name:     "test",
				Protocol: "dns",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := populateApplicationDefaults(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("populateApplicationDefaults() = %v, want %v", got, tt.want)
			}
		})
	}
}
