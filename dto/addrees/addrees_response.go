package dtoaddrees

import "synapsis-test/models"

type AddreesResponse struct {
	FullName    string      `json:"name" form:"fullname" validate:"required"`
	Street      string      `json:"street" form:"street" validate:"required"`
	HouseNumber string      `json:"housenumber" form:"housenumber" validate:"required"`
	PostCode    string      `json:"postcode" form:"postcode" validate:"required"`
	City        string      `json:"city" form:"city" validate:"required"`
	Province    string      `json:"province" form:"province" validate:"required"`
	Country     string      `json:"country" form:"country" validate:"required"`
	UserID      int         `json:"user_id " form:"user_id"`
	User        models.User `json:"user" form:"user"`
}
