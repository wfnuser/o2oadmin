package controllers

import (
	"github.com/beego/admin/src/rbac"
	"time"
	"github.com/astaxie/beego"
	"o2oadmin/models"
)

type ReportController struct {
	rbac.CommonController
}

type PdbMediaRequest struct {
	Dimension 	[]string `form:"dimension[]"`
	Medias		[]string `form:"media[]"`
	StartDate	time.Time `form:"startDate,2006-1-2"`
	EndDate	time.Time `form:"endDate,2006-1-2"`
	Page		int `form:"page"`
	Rows		int `form:"rows"`
	Sortby		string `form:"sortby"`
	Order		string `form:"order"`
}

type DspRequest struct {
	Dimension 	int	`form:"dimension"`
	Date		time.Time `form:"date"`
	QueryStr    string	`form:"q"`
}


func (this *ReportController) GetPdbMediaReport() {


	request := PdbMediaRequest{}
	this.ParseForm(request)
	this.TplNames = this.GetTemplatetype() + "/report/pmp_media_report.tpl"
	beego.Debug(request)

}

func (this *ReportController) GetPdbMediaReportData() {

	request := PdbMediaRequest{}
	this.ParseForm(&request)

//	excludedFields := []string{}
//	if request.Dimension != nil {
//
//	} else {
//		excludedFields = []string{"0", "1"}
//	}

	report, count, err := models.GetGroupedPmpDailyRequestReport(request.Dimension, request.Medias, request.StartDate, request.EndDate, request.Sortby, request.Order,(request.Page-1)*request.Rows, request.Rows)

	if err != nil {
		beego.Debug("failed to get pmp demand daily report")
	} else {
		// set PdbMediaName and PdbAdspaceName
		for idx, reportItem := range report {

			// because range copy values from the slice, we need to use index to change the original item
			report[idx].ReqAll = reportItem.ReqError + reportItem.ReqNoad + reportItem.ReqSuccess
			if report[idx].ReqAll > 0 {
				report[idx].FillRate = float32(reportItem.ReqSuccess) / float32(report[idx].ReqAll)
			}
		}

		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &report}
	}
	this.ServeJson()

}