package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
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
			"id": relay.GlobalIDField("Dimension", nil),
			"height": &graphql.Field{
				Name:        "Height",
				Type:        graphql.String,
				Description: "Height of the image",
			},
			"width": &graphql.Field{
				Name:        "Width",
				Type:        graphql.String,
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
			"id": relay.GlobalIDField("Image", nil),
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
			"id": relay.GlobalIDField("Page", nil),
			"pageNumber": &graphql.Field{
				Name:        "Page Number",
				Type:        graphql.Int,
				Description: "Page Number",
			},
			"count": &graphql.Field{
				Name:        "Count",
				Type:        graphql.Int,
				Description: "Page Number",
			},
			"start": &graphql.Field{
				Name:        "Start",
				Type:        graphql.Int,
				Description: "Start index of the current list of items returned",
			},
			"nextStart": &graphql.Field{
				Name:        "Next Start",
				Type:        graphql.Int,
				Description: "Start index of the next list of items returned that will be returned",
			},
			"isLastPage": &graphql.Field{
				Name:        "Is Last Page",
				Type:        graphql.Boolean,
				Description: "This page returned is the last page or there are more items remaining.",
			},
		},
	})
}

func setPaginatedImageList() {
	paginatedImageList = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Paginated Image List",
		Description: "List of images along with the pagination information",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Paginated Image List", nil),
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
				Name: "Image List",
				Type: paginatedImageList,
				Args: graphql.FieldConfigArgument{
					"type": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: "morning",
						Description:  "Morning, afternoon, evening or night.",
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
func GetSchema() graphql.Schema {
	return schema
}
