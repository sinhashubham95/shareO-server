package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/sinhashubham95/shareO-server/external"
)

func getPaginatedImageList(p graphql.ResolveParams) (interface{}, error) {
	searchType := p.Args["type"].(string)
	startIndex := p.Args["start"].(int)
	return external.GetImageList(searchType, startIndex)
}
