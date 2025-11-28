package employee_controller

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"TP_Andreev/internal/service"
	"TP_Andreev/internal/transport/http/router"
)

type EmployeeController struct {
	service service.Service
}

type jsData struct {
	Table template.JS
	Chart template.JS
}

type tmplData struct {
	Title string
	JS    jsData
}

var tmpl = template.Must(
	template.Must(
		template.New("jsData").Parse(src),
	).ParseFiles("web/templates/employee.html"),
)

func New(service service.Service) *EmployeeController {
	return &EmployeeController{service: service}
}

func (c *EmployeeController) GetEmployee(w http.ResponseWriter, r *http.Request, params router.Params) {
	idParam := params["id"]
	id, _ := strconv.Atoi(idParam)

	employeeData := c.service.GetEmployeeStat(id)
	employeeDataJ, _ := json.Marshal(employeeData)

	employeeTripData := c.service.GetEmployeeTripCountByAllYears(id)
	employeeTripDataJ, _ := json.Marshal(employeeTripData)

	jsData := jsData{
		Table: template.JS(employeeDataJ),
		Chart: template.JS(employeeTripDataJ),
	}
	data := tmplData{
		Title: employeeData.Name,
		JS:    jsData,
	}

	tmpl.ExecuteTemplate(w, "employee.html", data)
}

const src = `
	<script>
        const employeeData = {{.Table}};
        const chartData = {{.Chart}};
    </script>`
