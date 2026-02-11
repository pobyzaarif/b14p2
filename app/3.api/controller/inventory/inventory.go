package inventory

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pobyzaarif/b14p2/service/inventory"
)

type Controller struct {
	inventoryService inventory.Service
}

func NewController(inventoryService inventory.Service) *Controller {
	return &Controller{
		inventoryService: inventoryService,
	}
}

type Inventory struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
}

func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	invs, err := c.inventoryService.GetAll(1, 10)
	newInvs := make([]Inventory, 0)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	for _, v := range invs {
		newInvs = append(newInvs, Inventory{
			Code:        v.Code,
			Name:        v.Name,
			Description: v.Description,
			Stock:       v.Stock,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(newInvs)
}
