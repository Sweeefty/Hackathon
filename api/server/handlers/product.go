package handlers

import (
	"Hackathon/api/tools"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("GetProducts called")
	if r.Method == "GET" {
		bde_id := r.FormValue("bdeId")
		if bde_id == "" {
			Products := tools.GetProducts()
			if Products != false {
				tools.ResponseF(w, "200", "request successful", Products)

			} else {
				tools.ResponseF(w, "400", "internal error please retry later", nil)

			}
		} else {
			Products := tools.GetProductsByBDE(bde_id)
			if Products != false {
				tools.ResponseF(w, "200", "request successful", Products)

			} else {
				tools.ResponseF(w, "400", "internal error please retry later", nil)

			}
		}

	} else {
		tools.ResponseF(w, "400", "method incorrect", nil)
	}
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("GetProducts called")
	if r.Method == "GET" {
		ObjectId := r.FormValue("id")
		if ObjectId != "" {
			Product := tools.GetProductById(ObjectId)
			if Product != false {
				tools.ResponseF(w, "200", "request successful", Product)
			} else {
				tools.ResponseF(w, "400", "internal error please retry later", nil)
			}
		} else {
			tools.ResponseF(w, "400", "missing parameter(s)", nil)
		}

	} else {
		tools.ResponseF(w, "400", "method incorrect", nil)
	}
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("AddInventory called")
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		bde_id := r.FormValue("bdeId")
		if name != "" && description != "" && price != "" && bde_id != "" {
			tools.AddProduct(name, description, price, bde_id)

			tools.ResponseF(w, "200", "request successful", nil)
		} else {
			tools.ResponseF(w, "400", "missing parameter(s)", nil)
		}
	} else {
		tools.ResponseF(w, "400", "method incorrect", nil)
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	tools.WriteLog("DeleteProduct called")
	if r.Method == "DELETE" {
		ObjectId := r.FormValue("id")
		if ObjectId != "" {
			if tools.DeleteProduct(ObjectId) {
				tools.ResponseF(w, "200", "request successful", nil)
			} else {
				tools.ResponseF(w, "400", "incorrect id", nil)
			}

		} else {
			tools.ResponseF(w, "400", "missing parameter(s)", nil)
		}
	} else {
		tools.ResponseF(w, "400", "method incorrect", nil)
	}
}
