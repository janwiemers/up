package connection

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/janwiemers/up/models"
	"github.com/spf13/viper"
)

// GetMonitors returns the monitors
func GetMonitors() ([]models.Application, error) {
	resp, err := http.Get(fmt.Sprintf("%v/applications", viper.GetString("UP_BASE_URL")))

	if err != nil {
		panic(err.Error())
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apps []models.Application
	err = json.Unmarshal(data, &apps)

	if err != nil {
		return nil, err
	}
	return apps, nil
}

// GetChecks returns the monitors
func GetChecks(id int) ([]models.Check, error) {
	resp, err := http.Get(fmt.Sprintf("%v/application/%v/checks", viper.GetString("UP_BASE_URL"), id))

	if err != nil {
		panic(err.Error())
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var checks []models.Check
	err = json.Unmarshal(data, &checks)

	if err != nil {
		return nil, err
	}
	return checks, nil
}
