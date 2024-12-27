package model

type Customer struct {
	Sid            string `json:"sid" binding:"required"`
	CustomerRoleId string `json:"customer_role_id" binding:"required"`

	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	UserIconUrl string `json:"user_icon_url" binding:"required"`

	LastActiveAt string `json:"last_active_at" binding:"required"`
	CreatedAt    string `json:"created_at" binding:"required"`
	UpdatedAt    string `json:"updated_at" binding:"required"`
	DeletedAt    string `json:"deleted_at"`
}

type RegisterCustomerRequest struct {
	CustomerRoleId string `json:"customer_role_id" binding:"required"`

	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	UserIconUrl string `json:"user_icon_url"`
}

type RegisterCustomerResponse struct {
	Sid string `json:"sid" binding:"required"`
}

type UpdateCustomerRequest struct {
	Sid            string `json:"sid" binding:"required"`
	CustomerRoleId string `json:"customer_role_id" binding:"required"`

	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	UserIconUrl string `json:"user_icon_url" binding:"required"`

	LastActiveAt string `json:"last_active_at" binding:"required"`
	CreatedAt    string `json:"created_at" binding:"required"`
	UpdatedAt    string `json:"updated_at" binding:"required"`
	DeletedAt    string `json:"deleted_at"`
}

type UpdateCustomerResponse struct {
	Sid string `json:"sid"`
}

type GetCustomerRequest struct {
	Sid string `json:"sid"`
}

type GetCustomerResponse struct {
	Sid            string `json:"sid" binding:"required"`
	CustomerRoleId string `json:"customer_role_id" binding:"required"`

	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	UserIconUrl string `json:"user_icon_url" binding:"required"`

	LastActiveAt string `json:"last_active_at" binding:"required"`
	CreatedAt    string `json:"created_at" binding:"required"`
	UpdatedAt    string `json:"updated_at" binding:"required"`
	DeletedAt    string `json:"deleted_at"`
}

type GetManyCustomerRequest struct {
	Metadata       MetadataRequest `json:"metadata" binding:"required"`
	Username       string          `json:"username"`
	CustomerRoleId string          `json:"customer_role_id"`
}

type GetManyCustomerResponse struct {
}
