package curupd

import(
	"net/url"
	"net/http"
	"strings"
	"strconv"
	"io/ioutil"
	"encoding/xml"
	"time"
	"fmt"
)

const (
	CURUPD_URL = "http://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx"
	
	CURUPD_REQ = `<?xml version="1.0" encoding="utf-8"?>
	<soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
	  <soap12:Body>
	    <GetCursOnDateXML xmlns="http://web.cbr.ru/">
	      <On_date>%s</On_date>
	    </GetCursOnDateXML>
	  </soap12:Body>
	</soap12:Envelope>`
	
	TIME_LAYOUT = "20060102"
	TIME_ISO_LAYOUT = "2006-01-02T15:04:05"
)

// CreateSoapEnvelope struct
type CreateSoapEnvelope struct {
	CreateBody createBody `xml:"Body"`
}
type createBody struct {
	GetCursOnDateXMLResponse createCursOnDateResponse `xml:"GetCursOnDateXMLResponse"`
}

type createCursOnDateResponse struct {
	GetCursOnDateXMLResult createCursOnDateResult `xml:"GetCursOnDateXMLResult"`
}

type createCursOnDateResult struct {	
	ValuteData CreateValuteData `xml:"ValuteData"`
}

type CreateValuteData struct {
	OnDate string `xml:"OnDate,attr" json:"-"`
	Date time.Time `json:"date"`
	ValuteCursOnDate []Create `xml:"ValuteCursOnDate" json:"rates"`
}

type Create struct {
	Vname string `xml:"Vname" json:"name_full"`
	Vnom int `xml:"Vnom"`
	Vcurs float32 `xml:"Vcurs" json:"value"`
	Vcode string `xml:"Vcode" json:"code"`
	VchCode string `xml:"VchCode" json:"name"`
}

func GetCurrencyRates() (*CreateValuteData, error){

	u, err := url.ParseRequestURI(CURUPD_URL)
	if err != nil {
		return nil, err
	}

	req := fmt.Sprintf(CURUPD_REQ, time.Now().Format(TIME_ISO_LAYOUT))
	r, _ := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(req))
	r.Header.Add("Content-Type", "application/soap+xml; charset=utf-8")
	r.Header.Add("Content-Length", strconv.Itoa(len(req)))

	client := &http.Client{}
	resp, err := client.Do(r)	
	if err != nil {				
		return nil, err
	}
	
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	var createEnv CreateSoapEnvelope
	err = xml.Unmarshal([]byte(b), &createEnv)	
	if err != nil {
		return nil, err
	}

	t, err := time.Parse(TIME_LAYOUT, createEnv.CreateBody.GetCursOnDateXMLResponse.GetCursOnDateXMLResult.ValuteData.OnDate)
	if err != nil {
		return nil, err
	}
	createEnv.CreateBody.GetCursOnDateXMLResponse.GetCursOnDateXMLResult.ValuteData.Date = t
	
	return &createEnv.CreateBody.GetCursOnDateXMLResponse.GetCursOnDateXMLResult.ValuteData, nil

}
