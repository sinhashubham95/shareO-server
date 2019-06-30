package graphql

import (
	"github.com/graphql-go/graphql"
	"log"
)

var image *graphql.Object
var dimension *graphql.Object
var page *graphql.Object
var paginatedImageList *graphql.Object
var query *graphql.Object
var schema graphql.Schema

func init() {
	var err error
	setDimension()
	setImage()
	setPage()
	setPaginatedImageList()
	setQuery()
	schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: query,
	})
	if err != nil {
		log.Fatal("Error initializing schema.")
	}
}

func setDimension() {
	dimension = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Dimension",
		Description: "Dimension Information including height and width.",
		Fields: graphql.Fields{
			"height": &graphql.Field{
				Name:        "Height",
				Type:        graphql.Int,
				Description: "Height of the image",
			},
			"width": &graphql.Field{
				Name:        "Width",
				Type:        graphql.Int,
				Description: "Width of the image",
			},
		},
	})
}

func setImage() {
	image = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Image",
		Description: "Image Information",
		Fields: graphql.Fields{
			"title": &graphql.Field{
				Name:        "Title",
				Type:        graphql.String,
				Description: "Title of the image",
			},
			"link": &graphql.Field{
				Name:        "Link",
				Type:        graphql.String,
				Description: "Link to the image",
			},
			"type": &graphql.Field{
				Name:        "Type",
				Type:        graphql.String,
				Description: "Type of the image like jpg, gif, etc.",
			},
			"dimension": &graphql.Field{
				Name:        "Dimension",
				Type:        dimension,
				Description: "Dimensions of the image",
			},
		},
	})
}

func setPage() {
	page = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Page",
		Description: "Page Information",
		Fields: graphql.Fields{
			"count": &graphql.Field{
				Name:        "Count",
				Type:        graphql.Int,
				Description: "Page Number",
			},
			"start": &graphql.Field{
				Name:        "StartIndex",
				Type:        graphql.Int,
				Description: "Start index of the current list of items returned",
			},
			"nextStart": &graphql.Field{
				Name:        "NextStartIndex",
				Type:        graphql.Int,
				Description: "Start index of the next list of items returned that will be returned",
			},
			"isLastPage": &graphql.Field{
				Name:        "IsLastPage",
				Type:        graphql.Boolean,
				Description: "This page returned is the last page or there are more items remaining.",
			},
		},
	})
}

func setPaginatedImageList() {
	paginatedImageList = graphql.NewObject(graphql.ObjectConfig{
		Name:        "PaginatedImageList",
		Description: "List of images along with the pagination information",
		Fields: graphql.Fields{
			"data": &graphql.Field{
				Name:        "Data",
				Type:        graphql.NewList(image),
				Description: "List of images",
			},
			"page": &graphql.Field{
				Name:        "Page",
				Type:        page,
				Description: "Page information",
			},
		},
	})
}

func setQuery() {
	query = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Query",
		Description: "Queries that the server will serve.",
		Fields: graphql.Fields{
			"imageList": &graphql.Field{
				Name: "ImageList",
				Type: paginatedImageList,
				Args: graphql.FieldConfigArgument{
					"type": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "morning",
						Description:  "morning, afternoon, evening or night.",
					},
					"start": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 1,
						Description:  "Starting index of the image list being returned.",
					},
				},
				Description: "List of images with pagination",
				Resolve:     getPaginatedImageList,
			},
		},
	})
}

// GetSchema is used to get the schema
func GetSchema() *graphql.Schema {
	return &schema
}
