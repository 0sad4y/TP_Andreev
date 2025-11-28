package main_controller

import (
	"encoding/json"
	"html/template"
	"net/http"

	"TP_Andreev/internal/service"
	"TP_Andreev/internal/transport/http/router"
)

type MainController struct {
	service service.Service
}

type tmplData struct {
	Table  template.JS
	Chart1 template.JS
	Chart2 template.JS
}

var tmpl = template.Must(
	template.Must(
		template.New("jsData").Parse(src),
	).ParseFiles("web/templates/main.html"),
)

func New(service service.Service) *MainController {
	return &MainController{service: service}
}

func (c *MainController) GetMainPage(w http.ResponseWriter, r *http.Request, params router.Params) {
	employeeTripsData := c.service.GetAllEmployeeTrips()
	employeeTripsDataJ, _ := json.Marshal(employeeTripsData)

	moneySpentData := c.service.GetMoneySpentByAllYears()
	moneySpentDataJ, _ := json.Marshal(moneySpentData)

	tripCountData := c.service.GetTripCountByAllYears()
	tripCountDataJ, _ := json.Marshal(tripCountData)

	data := tmplData{
		Table:  template.JS(employeeTripsDataJ),
		Chart1: template.JS(moneySpentDataJ),
		Chart2: template.JS(tripCountDataJ),
	}

	tmpl.ExecuteTemplate(w, "main.html", data)
}

const src = `
	<script>
        const tableData = {{.Table}};
        const chartData1 = {{.Chart1}};
        const chartData2 = {{.Chart2}};
    </script>`
