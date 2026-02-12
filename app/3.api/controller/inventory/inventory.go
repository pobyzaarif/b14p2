package inventory

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pobyzaarif/b14p2/app/3.api/controller/common"
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

// ShowAccount godoc
// @Summary      Show list of inventory
// @Description  desc show list of inventory
// @Tags         Inventory
// @Accept       json
// @Produce      json
// @Param        page    query      int  false  "Page"
// @Param        limit   query      int  flase  "Limit"
// @Success      200  	 {array}    Inventory
// @Failure      400     {object}   common.ErrorResponse
// @Failure      500     {object}   common.ErrorResponse
// @Router       /inventory [get]
func (c *Controller) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	paramPage := r.URL.Query().Get("page")
	page, _ := strconv.Atoi(paramPage)

	paramLimit := r.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(paramLimit)

	invs, err := c.inventoryService.GetAll(page, limit)
	newInvs := make([]Inventory, 0)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		_ = json.NewEncoder(w).Encode(common.NewErrorMessage(err.Error()))
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

	_ = json.NewEncoder(w).Encode(newInvs)
}
