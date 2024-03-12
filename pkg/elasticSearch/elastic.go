package elasticsearch_pkg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

/**
create a client
create all the operations on the data
step-1 : install the package
step-2 : operations
	a. create an index operation
	b. add document to the index
	c. search for documents in the index
	d. update a document in the index
	e. delete a document from the index

*/

func CreateDefaultClient() (*elasticsearch.Client, error) {
	client, err := elasticsearch.NewDefaultClient() // returns a client using the default options (localhost:9200, no auth)
	if err != nil {
		return nil, err
	}
	return client, nil
}

type IndexRequestData struct {
	Index string
	ID    string
	Body  []byte
	Ctx   context.Context
}

// create index
func CreateIndex(client *elasticsearch.Client, r *IndexRequestData) (*esapi.Response, error) {

	req := esapi.IndexRequest{
		Index:      r.Index,
		DocumentID: r.ID,
		Body:       bytes.NewBuffer(r.Body),
		Refresh:    "true", // if we want changes to be visible immediately or not immediate
	}
	res, err := req.Do(r.Ctx, client)
	if err != nil {
		return res, fmt.Errorf("could not perform index request: %s", err)
	}
	defer func() { _ = res.Body.Close() }()
	if res.IsError() {
		return res, fmt.Errorf("index failed: %s", res.Status())
	}
	return res, nil
}

// search for the document in the index
func SearchDoc(c *elasticsearch.Client, i *esapi.SearchRequest) (*esapi.Response, error) {
	res, err := i.Do(context.Background(), c)
	if err != nil {
		return res, err
	}
	defer func() { _ = res.Body.Close() }()
	if res.IsError() {
		var e map[string]interface{}
		b, _ := io.ReadAll(res.Body)
		_ = json.Unmarshal(b, &e)
		return res, fmt.Errorf("%s", e["error"])
	}
	return res, nil
}

type UpdateRequestData struct {
	Ctx   context.Context
	Id    string
	Index string
	Body  []byte
}

// update the document
func UpdateDoc(c *elasticsearch.Client, u *UpdateRequestData) (*esapi.Response, error) {
	updateReq := esapi.UpdateRequest{
		Index:      u.Index,
		DocumentID: u.Id,
		Body:       bytes.NewReader(u.Body),
		Refresh:    "true",
	}
	res, err := updateReq.Do(u.Ctx, c)
	if err != nil {
		return res, err
	}
	defer func() { _ = res.Body.Close() }()
	if res.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("no document:%s", err)
	}
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return res, fmt.Errorf("update failed: %s", res.Status())
	}
	return res, nil
}

type DeleteRequestData struct {
	Ctx   context.Context
	Id    string
	Index string
}

// delete  a doc by id
func DeleteDoc(c *elasticsearch.Client, d *DeleteRequestData) (*esapi.Response, error) {
	deleteReq := esapi.DeleteRequest{
		Index:      d.Index,
		DocumentID: d.Id,
	}
	return deleteReq.Do(d.Ctx, c)
}

// search data from index
type SearchRequestData struct {
	Ctx        context.Context
	Query      interface{}
	PageSize   int
	CurrentPag int
	SortField  []string
	Index      []string
}

func (s *SearchRequestData) BuildDSL() map[string]interface{} {
	searchSource := make(map[string]interface{})
	searchSource["query"] = s.Query
	if len(s.SortField) > 0 {
		sort := make([]interface{}, len(s.SortField))
		for i, field := range s.SortField {
			tmpMap := make(map[string]interface{})
			tmpMap["order"] = "asc"
			sortSlice := []interface{}{tmpMap, "field", field}
			sort[i] = append(sortSlice)
		}
		searchSource["sort"] = sort
	}
	searchSource["from"] = (s.CurrentPag - 1) * s.PageSize
	searchSource["size"] = s.PageSize
	return searchSource
}

func SearchDocs(c *elasticsearch.Client, d *SearchRequestData) (*esapi.Response, error) {
	searchSource, _ := json.Marshal(d.BuildDSL())
	searchReq := esapi.SearchRequest{
		Index:      d.Index,
		Body:       bytes.NewReader(searchSource),
		Pretty:     true, // pretty print request and response JSON
		Human:      true, // return human readable values for
		ErrorTrace: true,
	}
	return searchReq.Do(d.Ctx, c)
}


