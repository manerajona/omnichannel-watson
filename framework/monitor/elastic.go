package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	elastic "github.com/olivere/elastic/v7"
	config "github.com/olivere/elastic/v7/config"
)

type ElasticSession struct {
	Client *elastic.Client
}

// NewEslasticSession initialize a client with config.
func NewEslasticSession(cfg *config.Config) (*ElasticSession, error) {
	client, err := elastic.NewClientFromConfig(cfg)
	return &ElasticSession{client}, err
}

// Get the json.RawMessage if the document exists.
func (s *ElasticSession) Get(indexName string, docID string) (json json.RawMessage) {
	res, err := s.Client.Get().Index(indexName).Id(docID).Do(context.TODO())
	if err != nil {
		return
	}
	return res.Source
}

// Index a document with the specified docID.
func (s *ElasticSession) Index(indexName string, docID string, object interface{}) (err error) {
	res, err := s.Client.Index().Index(indexName).Id(docID).BodyJson(&object).Do(context.TODO())
	if res == nil || res.Index == "" {
		err = fmt.Errorf("[%s] Fail indexing document", res.Status)
	}
	return
}
