package yahoo

import (
	"fmt"
	"strconv"
	"strings"

	"net/url"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/nestgo/utils"
)

const yqlUrl = `https://query.yahooapis.com/v1/public/yql?q=%s&format=json&diagnostics=true&env=store://datatables.org/alltableswithkeys&callback=`

//Exchange 通过yahoo提供的接口转换
func Exchange(rateFrom, rateTo string) (ret float64, err error) {
	pair := fmt.Sprintf(`select * from yahoo.finance.xchange  where pair in ("%s%s")`, strings.ToUpper(rateFrom), strings.ToUpper(rateTo))
	pair = url.QueryEscape(pair)
	requestUri := fmt.Sprintf(yqlUrl, pair)
	var dataByte []byte
	dataByte, err = utils.HTTPGet(requestUri)
	if err != nil {
		err = fmt.Errorf("yahoo接口转换失败,err=%v", err)
		return
	}
	var json *simplejson.Json
	json, err = simplejson.NewJson(dataByte)
	if err != nil {
		return
	}
	//解析返回结果 正常: query->results->rate->Rate
	var rateString string
	rateString, err = json.GetPath("query", "results", "rate", "Rate").String()
	if err != nil {
		err = fmt.Errorf("yahoo接口转换失败,err=%v", err)
		return
	}
	ret, err = strconv.ParseFloat(rateString, 64)
	return
}
