package models

type (
	Warehouse struct {
		WarehouseId string    `json:"warehouse_id,omitempty"  bson:"warehouse_id,omitempty"`
		Name        string    `json:"name,omitempty"  bson:"name,omitempty"`
		Products    []Product `json:"products,omitempty"  bson:"products,omitempty"`
	}
	RotaWarehouseRequest struct {
		FromWarehouseId string    `json:"from_warehouse_id,omitempty" binding:"required"`
		ToWarehouseId   string    `json:"to_warehouse_id,omitempty" binding:"required"`
		Products        []Product `json:"products,omitempty" binding:"required"`
	}

	InsertOrderParams struct {
		Warehouse
		Total int `bson:"total"`
	}
)
