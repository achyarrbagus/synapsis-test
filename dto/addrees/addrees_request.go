package dtoaddrees

import "synapsis-test/models"

type AddreesRequest struct {
	FullName    string      `json:"full_name" form:"fullname" validate:"required"`
	Street      string      `json:"street" form:"street" validate:"required"`
	HouseNumber string      `json:"house_number" form:"housenumber" validate:"required"`
	PostCode    string      `json:"post_code" form:"postcode" validate:"required"`
	Phone       string      `json:"phone"  validate:"required"`
	City        string      `json:"city" form:"city" validate:"required"`
	Province    string      `json:"province" form:"province" validate:"required"`
	Country     string      `json:"country" form:"country" validate:"required"`
	UserID      int         `json:"user_id " form:"user_id"`
	User        models.User `json:"user" form:"user"`
}

type UpdateAddreesRequest struct {
	FullName    string      `json:"full_name" form:"fullname"`
	Street      string      `json:"street" form:"street"`
	HouseNumber string      `json:"house_number" form:"housenumber"`
	PostCode    string      `json:"post_code" form:"postcode"`
	City        string      `json:"city" form:"city"`
	Phone       string      `json:"phone"`
	Province    string      `json:"province" form:"province"`
	Country     string      `json:"country" form:"country"`
	UserID      int         `json:"user_id " form:"user_id"`
	User        models.User `json:"user" form:"user"`
}
