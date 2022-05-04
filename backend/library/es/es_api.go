package es

import (
	"context"
	"crypto/tls"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"nft_object/library/helper"
	"nft_object/library/logge"

	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
	"github.com/olivere/elastic"
)

// EsApi struct
type EsApi struct {
	Client *elastic.Client
	Host   []string
}

// 初始化客户端
// NewEsApi 参数hosts格式： http://127.0.0.1:5000;http://127.0.0.2:5000
// params[0] username
// params[1] password
// params[2] 或者：params[0] *http.Client
func NewEsApi(hosts string, params ...interface{}) (*EsApi, error) {
	c := new(EsApi)

	httpClient := &http.Client{
		Transport: &http.Transport{
			ExpectContinueTimeout: 1 * time.Second,
			IdleConnTimeout:       5 * time.Second,
			TLSHandshakeTimeout:   5 * time.Second,
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * 10,
			MaxIdleConns:          100,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS11,
				InsecureSkipVerify: true,
			},
			DialContext: (&net.Dialer{
				DualStack: true,
				Timeout:   5 * time.Second,
				KeepAlive: 5 * time.Second,
			}).DialContext,
		},
		Timeout: time.Duration(30) * time.Second,
	}
	username := ""
	password := ""
	if len(params) > 0 {
		for _, row := range params {
			if _, ok := row.(*http.Client); ok {
				httpClient = row.(*http.Client)
			} else if reflect.TypeOf(row).Kind() == reflect.String {
				if username == "" {
					username = row.(string)
				} else if password == "" {
					password = row.(string)
				}
			}
		}
	}

	if hosts == "" {
		return nil, errors.New("create client failed, err: config error")
	} else {
		c.Host = strings.Split(hosts, ";")
		var err error
		c.Client, err = elastic.NewClient(
			elastic.SetHttpClient(httpClient), // 设置http请求参数，如超时时间
			elastic.SetURL(c.Host...),
			elastic.SetSniff(false), // 如果为true高并发存在性能问题出现time wait
			//elastic.SetHealthcheckTimeoutStartup(2), // SetHealthcheck=true 时此参数有效。如果为true高并发下会产生很多连接数
			elastic.SetHealthcheck(false), // 如果为true高并发下会产生很多连接数
			elastic.SetBasicAuth(username, password),
		)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

// PingNode1 ping 连接测试
func PingNode1(host string) {
	start := time.Now()
	var client *elastic.Client
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		fmt.Printf("ping es failed, err: %v", err)
	}
	duration := time.Since(start)
	fmt.Printf("cost time: %v\n", duration)
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}

// IndexExists 校验 index 是否存在
func (c *EsApi) IndexExists(index ...string) bool {
	exists, err := c.Client.IndexExists(index...).Do(context.Background())
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	return exists
}

// CreateIndex 创建 index
func (c *EsApi) CreateIndex(index, mapping string) bool {
	result, err := c.Client.CreateIndex(index).BodyString(mapping).Do(context.Background())
	if err != nil {
		fmt.Printf("create index failed, err: %v\n", err)
	}
	return result.Acknowledged
}

// DeleteIndex 删除 index
func (c *EsApi) DeleteIndex(index ...string) bool {
	response, err := c.Client.DeleteIndex(index...).Do(context.Background())
	if err != nil {
		fmt.Printf("delete index failed, err: %v\n", err)
	}
	return response.Acknowledged
}

// BatchStore 批量插入数据
func (c *EsApi) BatchStore(index, type_ string, data []map[string]interface{}) (bool, error) {

	bulkRequest := c.Client.Bulk()
	for _, row := range data {
		id := ""
		if row["id"] == nil {
			id = helper.CreateId(grand.N(0, 10))
		} else {
			id = row["id"].(string)
		}
		req := elastic.NewBulkIndexRequest().
			Index(index).
			Type(type_).
			Id(id).
			Doc(row)
		bulkRequest = bulkRequest.Add(req)
	}

	bulkResponse, err := bulkRequest.Do(context.Background())
	if err != nil {
		return false, err
	}
	if bulkResponse != nil {

	}
	return true, nil
}

// PackageBoolQuery 封装query,参数有两种格式（查询操作类型有三种：term,terms,match）
// 1、简单数组格式默认为must: params = [][3]interface{}{{"id","term","1"}}
// 2、复杂map格式:
// params["must"] = [][3]interface{}{{"id","term","1"}}
// params["filter"] = [][3]interface{}{{"id","terms",[]string{"1"}}}
// params["mustNot"] = [][3]interface{}{{"id","match","1"}}
// params["should"] = [][3]interface{}{{"id","term","1"}}
func (c *EsApi) PackageBoolQuery(params interface{}) string {
	query := elastic.NewBoolQuery()
	data := map[string]interface{}{}
	switch reflect.TypeOf(params).Kind() {
	case reflect.Map:
		data = gconv.Map(params)
	case reflect.Slice:
		data["must"] = params
	}

	for k, rows := range data {
		row := rows.([][3]interface{})
		switch k {
		case "must":
			for _, v := range row {
				query = query.Must(c.PackageOperator(v))
			}
		case "filter":
			for _, v := range row {
				query = query.Filter(c.PackageOperator(v))
			}
		case "mustNot":
			for _, v := range row {
				query = query.MustNot(c.PackageOperator(v))
			}
		case "should":
			for _, v := range row {
				query = query.Should(c.PackageOperator(v))
			}
		}
	}

	src, err := query.Source()
	if err != nil {
		fmt.Println("query error")
		return ""
	}
	b, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		fmt.Println("json error")
		return ""
	}
	queryStr := string(b)
	return queryStr
}

// PackageOperator 根据查询操作来封装query，查询操作：term,terms,match
func (c *EsApi) PackageOperator(data [3]interface{}) elastic.Query {
	op := gconv.String(data[1])
	field := gconv.String(data[0])
	val := data[2]
	switch op {
	case "terms":
		vals := gconv.SliceAny(val)
		return elastic.NewTermsQuery(field, vals...)
	case "match":
		return elastic.NewMatchQuery(field, val)
	default:
		return elastic.NewTermQuery(field, val)
	}
}

// TermQuery term 查询
func (c *EsApi) TermQuery(index, type_, fieldName, fieldValue string) *elastic.SearchResult {
	query := elastic.NewTermQuery(fieldName, fieldValue)

	fmt.Println(query, "============")
	//_ = elastic.NewQueryStringQuery(fieldValue) //关键字查询

	searchResult, err := c.Client.Search().
		Index(index).Type(type_).
		Query(query).
		From(0).Size(10).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		panic(err)
	}
	fmt.Printf("query cost %d millisecond.\n", searchResult.TookInMillis)

	return searchResult
}

// QueryString 查询条件为字符串
// queryJson = `{"bool": {"must": [{"range": {"operator_time": {"gte": "2019-11-27","lte": "2019-12-02"}}}]}}`
func (c *EsApi) QueryString(index, type_ string, queryJson string, page int,
	limit int) (*elastic.SearchHits, error) {

	query := elastic.RawStringQuery(queryJson)
	// page 从1开始（分页时offset会减1）
	if page == 0 {
		page = 1
	}

	// 不允许limit为0，以及大于5000
	if limit == 0 || limit > 5000 {
		limit = 5000
	}
	offset := (page - 1) * limit
	searchResult, err := c.Client.Search().
		Index(index).
		Type(type_).
		Query(query).
		From(offset).
		Size(limit).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		return nil, err
	}
	return searchResult.Hits, nil

}

// AggsQuery 分组统计
func (c *EsApi) AggsQuery(index, type_ string, where map[string]interface{},
	dateParams map[string]string, groupFields []string) (*elastic.SearchResult, error) {

	query := c.SetQuery(where)
	searchService := c.Client.Search().
		Index(index).
		Type(type_).
		Query(query).
		Size(0).
		Pretty(true)

	limit := 10000

	if dateParams["format"] == "" {
		dateParams["format"] = "yyyy-MM-dd"
	}

	th := elastic.NewDateHistogramAggregation().
		Interval(dateParams["interval"]).
		Field(dateParams["field"]).
		Format(dateParams["format"]).
		TimeZone("Asia/Shanghai")

	for _, v := range groupFields {
		sec_name := elastic.NewTermsAggregation().Field(v).
			SubAggregation(v, th).Size(limit)
		searchService.Aggregation("@"+v, sec_name)
	}

	searchResult, err := searchService.Do(context.Background())

	if err != nil {
		return nil, err
	}
	return searchResult, nil
}

// Pagination 分页查询数据
// params map[string]interface{} where参数格式：where["id"] = {"value":"","operator":"term","type":"must"}
// params map[string]interface{} params["highlight"] []string 要高亮字段
// params map[string]interface{} params["query"] string query json：{bool:{}}
// params string query json：{bool:{}}
func (c *EsApi) Pagination(index, type_ string, params interface{}, column []string,
	page int, limit int) map[string]interface{} {
	// page 从1开始（分页时offset会减1）
	if page == 0 {
		page = 1
	}
	lists := map[string]interface{}{
		"items":        []interface{}{},
		"total":        0,
		"current_page": page,
		"per_page":     limit,
		"pages":        0,
		"error":        "",
	}

	// 不允许limit为0，以及大于10000
	if limit == 0 || limit > 10000 {
		limit = 10000
	}
	if len(column) == 0 {
		res, _ := c.Client.GetMapping().Index(index).Do(context.Background())
		list := gconv.Map(gconv.Map(gconv.Map(gconv.Map(res[index])["mappings"])[type_])["properties"])
		for k := range list {
			column = append(column, k)
		}
	}
	offset := (page - 1) * limit

	s := c.Client.Search().Index(index).Type(type_).From(offset).
		DocvalueFields(column...).Size(limit).Pretty(true)

	t := reflect.TypeOf(params).Kind()
	switch t {
	case reflect.String:
		s = s.Query(elastic.RawStringQuery(params.(string)))
	case reflect.Map:
		if _, ok := params.(map[string]interface{})["query"]; ok {
			s = s.Query(elastic.RawStringQuery(gconv.String(params.(map[string]interface{})["query"])))
		} else {
			s = s.Query(c.SetQuery(params.(map[string]interface{})))
		}
		if _, ok := params.(map[string]interface{})["highlight"]; ok {
			highlight := gconv.SliceStr(params.(map[string]interface{})["highlight"])
			h := elastic.NewHighlight()
			var hf []*elastic.HighlighterField
			for _, c := range highlight {
				hf = append(hf, elastic.NewHighlighterField(c))
			}
			h = h.Fields(hf...)
			s = s.Highlight(h)
		}
	}
	searchResult, err := s.Do(context.Background())
	if err != nil {
		lists["error"] = err.Error()
		return lists
	}
	return c.PaginationData(lists, searchResult, page, limit)
}

// PaginationData 处理查询分页结果
func (c *EsApi) PaginationData(lists map[string]interface{},
	searchResult *elastic.SearchResult, page, limit int) map[string]interface{} {
	count := searchResult.Hits.TotalHits
	list := searchResult.Hits.Hits
	items := []interface{}{}
	if len(list) > 0 {
		for _, v := range list {
			js, err := v.Source.MarshalJSON()
			if err != nil {
				continue
			}
			row := gconv.Map(js)
			if len(v.Highlight) > 0 {
				row["highlight"] = v.Highlight
			}
			items = append(items, row)
		}
	}
	lists["items"] = items
	lists["total"] = count
	lists["current_page"] = page
	lists["per_page"] = limit
	pages := 0
	if limit > 0 {
		pages = int(count) / limit
	}
	lists["pages"] = math.Ceil(float64(pages))
	return lists
}

// ScrollQuery scroll第一次查询数据，(通过scrollId查询可直接简单调用ScrollById方法)
// params queryJson，如：{bool:{}}，或者是map,如map[string]interface{}{"query":{"bool":{}},"highlight":[]string{}}
// def[0] = keepAlive 如1m
// def[1] = scrollId
func (c *EsApi) ScrollQuery(index, type_ string, params interface{},
	limit int, column []string, def ...string) (*elastic.SearchResult, error) {

	var s *elastic.ScrollService
	keepAlive := "1m"
	scrollId := ""
	if len(def) >= 1 {
		keepAlive = def[0]
	}
	if len(def) >= 2 {
		scrollId = def[1]
	}
	queryJson := ""
	highlight := []string{}
	t := reflect.TypeOf(params).Kind()
	switch t {
	case reflect.String:
		queryJson = params.(string)
	case reflect.Map:
		if _, ok := params.(map[string]interface{})["query"]; ok {
			queryJson = gconv.String(params.(map[string]interface{})["query"])
		}
		if _, ok := params.(map[string]interface{})["highlight"]; ok {
			highlight = gconv.SliceStr(params.(map[string]interface{})["highlight"])
		}
	}

	if scrollId == "" {
		// 第一次查询
		s = c.Client.Scroll(index).
			Type(type_).
			Scroll(keepAlive).
			Size(limit)
		if len(column) > 0 {
			columnSource := elastic.NewSearchSource().
				FetchSourceContext(elastic.NewFetchSourceContext(true).Include(column...))
			s.SearchSource(columnSource)
		}
		if queryJson != "" {
			// 查询条件
			s = s.Query(elastic.RawStringQuery(queryJson))
		}
		if len(highlight) > 0 {
			h := elastic.NewHighlight()
			var hf []*elastic.HighlighterField
			for _, c := range highlight {
				hf = append(hf, elastic.NewHighlighterField(c))
			}
			h = h.Fields(hf...)
			s = s.Highlight(h)
		}
	} else {
		// 根据游标查询
		s = c.Client.Scroll(keepAlive).ScrollId(scrollId)
	}

	res, err := s.Do(context.Background())

	if scrollId != "" && err != nil && !strings.Contains(err.Error(), "EOF") {
		fmt.Println("游标查询数据失败", c.Host, index, type_, scrollId, elastic.RawStringQuery(queryJson))
		logge.Write("es_api", "error", "游标查询数据失败",
			c.Host, index, elastic.RawStringQuery(queryJson))
		return nil, err
	} else if scrollId != "" && err != nil && strings.Contains(err.Error(), "EOF") {
		fmt.Println("查询结束", index, type_, scrollId, err.Error())
		res.ScrollId = ""
		return res, nil
	} else if err != nil {
		fmt.Println("首次查询游标数据失败", c.Host, index, type_, scrollId, elastic.RawStringQuery(queryJson))
		logge.Write("es_api", "error", "首次查询游标数据失败",
			c.Host, index, scrollId, elastic.RawStringQuery(queryJson))
		return nil, err
	}
	return res, err
}

// ScrollById 根据scrollId滚动查询数据
func (c *EsApi) ScrollById(scrollId string, keepAlives ...string) (*elastic.SearchResult, error) {
	keepAlive := "1m"
	if len(keepAlives) >= 1 {
		keepAlive = keepAlives[0]
	}
	return c.ScrollQuery("", "", "", 0, []string{}, keepAlive, scrollId)
}

// SetQuery 前端请求的参数转换查询query参数
// where参数格式：where["id"] = {"value":"","operator":"term","type":"must"}
// where["operator_time"] = {"value":{"gte":"","lt":""},"operator":"range","type":"must"}
// where["content"] = {"value":"关键词","operator":"match","type":"should"}
// where["content"] = {"value":"关键词","operator":"match","type":"should","field":["content","reply_content"]}
func (c *EsApi) SetQuery(where map[string]interface{}) *elastic.BoolQuery {
	boolQuery := elastic.NewBoolQuery()
	for k, v := range where {
		row := gconv.Map(v)
		fmt.Println(row, k, v)
		var q elastic.Query

		if gconv.String(row["operator"]) == "term" {
			q = elastic.NewTermQuery(k, gconv.String(row["value"]))
		} else if gconv.String(row["operator"]) == "range" {
			time_val := gconv.Map(row["value"])
			rq := elastic.NewRangeQuery(k)
			for k1, v1 := range time_val {
				if k1 == "gte" {
					rq.Gte(v1)
				} else if k1 == "lt" {
					rq.Lte(v1)
				}
			}
			q = rq
		} else if gconv.String(row["operator"]) == "terms" {
			q = elastic.NewTermsQuery(k, gconv.Interfaces(row["value"])...)
		} else if gconv.String(row["operator"]) == "match" {
			fields := gconv.SliceStr(row["field"])
			if len(fields) > 0 {
				q = elastic.NewMultiMatchQuery(row["value"], fields...)
			} else {
				q = elastic.NewMatchQuery(k, row["value"])
			}
		}

		if gconv.String(row["type"]) == "must_not" {
			boolQuery.MustNot(q)
		} else if gconv.String(row["type"]) == "should" {
			boolQuery.Should(q)
		} else {
			boolQuery.Must(q)
		}
	}
	return boolQuery
}

// ExportHeader 初始化导出csv文件header
func (c *EsApi) ExportHeader(index, _type string,
	headerParams map[string]string) ([]string, []string, []string) {
	column := []string{}
	hkey := make([]string, 0)
	headerCsv := make([]string, 0)
	if len(headerParams) > 0 {
		for k, v := range headerParams {
			column = append(column, k)
			hkey = append(hkey, k)
			headerCsv = append(headerCsv, v)
		}
	} else {
		res, _ := c.Client.GetMapping().Index(index).Do(context.Background())
		list := gconv.Map(gconv.Map(gconv.Map(gconv.Map(res[index])["mappings"])[_type])["properties"])
		for k := range list {
			column = append(column, k)
			hkey = append(hkey, k)
			headerCsv = append(headerCsv, k)
		}
	}
	return column, hkey, headerCsv
}

// ExportCsv 导出数据到csv文件
func (c *EsApi) ExportCsv(fileName string, index, _type, queryJson string,
	limit int, headerParams map[string]string, def ...string) (int64, error) {
	column, hkey, headerCsv := c.ExportHeader(index, _type, headerParams)
	query, err := c.ScrollQuery(index, _type, queryJson, limit, column, def...)
	if err != nil {
		logge.Write("es_api", "error", "查询失败"+err.Error(), c.Host, index, queryJson)
		return 0, err
	}
	// 导出数据有限制
	total := query.Hits.TotalHits
	searchTotal := len(query.Hits.Hits)
	if total == 0 {
		return 0, errors.New("no data")
	}
	if len(column) == 0 {
		for _, hit := range query.Hits.Hits {
			item := make(map[string]interface{})
			err := json.Unmarshal(*hit.Source, &item)
			if err != nil {
				fmt.Println("json Unmarshal error: ", err.Error())
				continue
			}
			for k := range item {
				column = append(column, k)
				hkey = append(hkey, k)
				headerCsv = append(headerCsv, k)
			}
			break
		}
	}
	newFile, err := os.Create(fileName)
	if err != nil {
		logge.Write("es_api", "error", "创建文件失败"+err.Error(), c.Host, index, queryJson)
		return 0, err
	}
	// 写入utf-8 BOM，防止中文乱码
	newFile.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(newFile)
	w.Write(headerCsv)
	c.WriteCsv(w, query, hkey)
	scorllId := query.ScrollId
	if scorllId != "" && query.Hits.TotalHits >= int64(limit) {
		for {
			query, err = c.ScrollById(scorllId)
			if err != nil {
				return 0, err
			}
			if query.ScrollId == "" {
				err = errors.New("查询数据结束")
				break
			}
			searchTotal += len(query.Hits.Hits)
			scorllId = query.ScrollId
			if query.Hits == nil {
				err = errors.New("查询数据结束")
				break
			}
			c.WriteCsv(w, query, hkey)
			if searchTotal >= 2000000 {
				err = errors.New("数据超过200万，结束任务")
				break
			}
		}
		if err != nil {
			logge.Write("es_api", "info", "ExportCsv", err.Error(), c.Host, index, queryJson)
		}
	}
	w.Flush()
	newFile.Close()
	logge.Write("es_api", "info", "ExportCsv", total, c.Host, index, queryJson)
	return total, nil
}

// WriteCsv 导出数据到csv文件
func (c *EsApi) WriteCsv(w *csv.Writer, data *elastic.SearchResult, hkey []string) {
	for _, hit := range data.Hits.Hits {
		item := make(map[string]interface{})
		err := json.Unmarshal(*hit.Source, &item)
		if err != nil {
			fmt.Println("json Unmarshal error: ", err.Error())
			continue
		}
		row := make([]string, 0)
		for _, k := range hkey {
			val := ""
			if _, ok := item[k]; ok {
				val = gconv.String(item[k]) + "\t"
			}
			row = append(row, val)
		}
		if len(row) == 0 {
			continue
		}
		// 写入数据
		w.Write(row)
		w.Flush()
	}
}
