package main

import (
	"github.com/graphql-go/graphql"
	"time"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"events": &graphql.Field{
				Name: "Events",
				Type: graphql.NewList(eventType),
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: nil,
						Description:  "the title of the event",
					},
					"place": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: nil,
						Description:  "the name of the provider of the event",
					},
					"timestampLT": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: nil,
						Description:  "Event has to be before this timestamp",
					},
					"timestampGT": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: (time.Now().Unix()),
						Description:  "Event has to be after this timestamp",
					},
					"genre": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: nil,
						Description:  "The genre of event NOT IMPLEMENTED",
					},
					"category": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: nil,
						Description:  "the category of the event",
					},
					"area": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: nil,
						Description:  "the area in the city of event",
					},
					"zip": &graphql.ArgumentConfig{
						Type:         graphql.String,
						DefaultValue: nil,
						Description:  "The zip code of the location of the event",
					},
					"priceLT": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: nil,
						Description:  "The max price of the event",
					},
					"priceGT": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: nil,
						Description:  "The min price of the event",
					},
				},
				Resolve:     queryResolver,
				Description: "Get events by arguments",
			},
		},
	})

var addressType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Address",
		Fields: graphql.Fields{
			"street": &graphql.Field{
				Name:        "StreetName",
				Description: "The name of the road of the event",
				Type:        graphql.String,
			},
			"no": &graphql.Field{
				Name:        "Number",
				Description: "The number of the house or establishment on the road.",
				Type:        graphql.String,
			},
			"zip": &graphql.Field{
				Name:        "PostNumber",
				Description: "The zip code of the location",
				Type:        graphql.Int,
			},
			"city": &graphql.Field{
				Name:        "City",
				Description: "The name of the city of the event",
				Type:        graphql.String,
			},
			"state": &graphql.Field{
				Name:        "Country",
				Description: "The name of the country, probably Denmark...",
				Type:        graphql.String,
			},
		},
		Description: "The address have more specific information about the location.",
	},
)
var coordinatesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Coordinates",
		Fields: graphql.Fields{
			"latitude": &graphql.Field{
				Name:        "Latitude",
				Description: "I thing the horizontal axis or something...",
				Type:        graphql.Float,
			},
			"longitude": &graphql.Field{
				Name:        "Longitude",
				Description: "I thing the vertical axis or something... depending on the rotation of the earth",
				Type:        graphql.Float,
			},
		},
		Description: "The coordinates are used to pin the location on a map.",
	},
)
var locationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Location",
		Fields: graphql.Fields{
			"address": &graphql.Field{
				Name:        "Address",
				Description: "The name of the road of the event",
				Type:        addressType,
			},
			"area": &graphql.Field{
				Name:        "CityArea",
				Description: "The name of the area, in the city, of the event",
				Type:        graphql.String,
			},
			"place": &graphql.Field{
				Name:        "Place",
				Description: "The name of the place that holds the event.",
				Type:        graphql.String,
			},
			"coordinates": &graphql.Field{
				Name:        "Coordinates",
				Description: "The coordinates of the event!",
				Type:        coordinatesType,
			},
		},
		Description: "The location is holds the information used to locate the event.",
	},
)

var eventType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Event",
		Fields: graphql.Fields{
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"time": &graphql.Field{
				Type: graphql.Int,
			},
			"price": &graphql.Field{
				Type: graphql.Int,
			},
			"genre": &graphql.Field{
				Type: graphql.String,
			},
			"image": &graphql.Field{
				Type: graphql.String,
			},
			"tickets": &graphql.Field{
				Type: graphql.String,
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"link": &graphql.Field{
				Type: graphql.String,
			},
			"category": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"other": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"location": &graphql.Field{
				Type: locationType,
			},
		},
	},
)
