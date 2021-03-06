package main

import (
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func queryResolver(p graphql.ResolveParams) (interface{}, error) {
	filter := bson.M{}
	title, titleOk := p.Args["title"].(string)
	if titleOk {
		filter["title"] = primitive.Regex{
			Pattern: title,
			Options: "i",
		}
	}
	price := interval(&p, "priceLT", "priceGT")
	if price != nil {
		filter["price"] = price
	}
	time := interval(&p, "timestampLT", "timestampGT")
	if time != nil {
		filter["time"] = time
	}
	place, placeOk := p.Args["place"].([]interface{})
	if placeOk {
		filter["location.place"] = bson.M{"$in": place}
	}
	category, categoryOk := p.Args["category"].([]interface{})
	if categoryOk {
		filter["category"] = bson.M{"$in": category}
	}
	area, areaOk := p.Args["area"].([]interface{})
	if areaOk {
		filter["location.area"] = bson.M{"$in": area}
	}
	genre, genreOk := p.Args["genre"].([]interface{})
	if genreOk {
		filter["genre"] = bson.M{"$in": genre}
	}
	zip, zipOK := p.Args["zip"].([]interface{})
	if zipOK {
		filter["location.zip"] = bson.M{"$in": zip}
	}

	return getEvents(filter), nil
}

func interval(p *graphql.ResolveParams, lt string, gt string) bson.M {
	a, aOK := p.Args[lt]
	b, bOK := p.Args[gt]

	if aOK && bOK {
		return bson.M{
			"$gte": a,
			"$lte": b,
		}
	} else if aOK {
		return bson.M{
			"$lte": a,
		}
	} else if bOK {
		return bson.M{
			"$gte": b,
		}
	}
	return nil
}
