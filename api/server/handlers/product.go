package handlers

import (
	"Hackathon/api/server/data"
	"Hackathon/api/tools"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("GetProducts called")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		bde_id := r.FormValue("bdeId")
		if bde_id == "" {
			Products := tools.GetProducts()
			if Products != false {
				Response := data.Response{
					Status: "request successful",
					Code:   "200",
					Data:   Products,
				}
				tools.ResponseF(w, Response)
			} else {
				Response := data.Response{
					Status: "error",
					Code:   "400",
				}
				tools.ResponseF(w, Response)
			}
		} else {
			Products := tools.GetProductsByBDE(bde_id)
			if Products != false {
				Response := data.Response{
					Status: "request successful",
					Code:   "200",
					Data:   Products,
				}
				tools.ResponseF(w, Response)
			} else {
				Response := data.Response{
					Status: "error",
					Code:   "400",
				}
				tools.ResponseF(w, Response)

			}
		}

	} else {
		Response := data.Response{
			Status: "method incorrect",
			Code:   "400",
		}
		tools.ResponseF(w, Response)
	}
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("GetProducts called")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		ObjectId := r.FormValue("id")
		if ObjectId != "" {
			Product := tools.GetProductById(ObjectId)
			if Product != false {
				Response := data.Response{
					Status: "request successful",
					Code:   "200",
					Data:   Product,
				}
				tools.ResponseF(w, Response)
			} else {
				Response := data.Response{
					Status: "error",
					Code:   "400",
				}
				tools.ResponseF(w, Response)
			}
		} else {
			Response := data.Response{
				Status: "missing parameter",
				Code:   "400",
			}
			tools.ResponseF(w, Response)
		}

	} else {
		Response := data.Response{
			Status: "method incorrect",
			Code:   "400",
		}
		tools.ResponseF(w, Response)
	}
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("AddInventory called")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		bde_id := r.FormValue("bdeId")
		if name != "" && description != "" && price != "" && bde_id != "" {
			tools.AddProduct(name, description, price, bde_id)
			Response := data.Response{
				Status: "request successful",
				Code:   "200",
			}
			tools.ResponseF(w, Response)
		} else {
			Response := data.Response{
				Status: "missing parameter",
				Code:   "400",
			}
			tools.ResponseF(w, Response)
		}
	} else {
		Response := data.Response{
			Status: "method incorrect",
			Code:   "400",
		}
		tools.ResponseF(w, Response)
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("DeleteProduct called")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "DELETE" {
		ObjectId := r.FormValue("id")
		if ObjectId != "" {
			if tools.DeleteProduct(ObjectId) {
				Response := data.Response{
					Status: "Product deleted successfully",
					Code:   "200",
				}
				tools.ResponseF(w, Response)
			} else {
				Response := data.Response{
					Status: "incorrect id",
					Code:   "400",
				}
				tools.ResponseF(w, Response)
			}

		} else {
			Response := data.Response{
				Status: "missing parameter",
				Code:   "400",
			}
			tools.ResponseF(w, Response)
		}
	} else {
		Response := data.Response{
			Status: "method incorrect",
			Code:   "400",
		}
		tools.ResponseF(w, Response)
	}
}
