package app

import (
	_ "errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kosha/serviceNow-connector/pkg/httpclient"
)

// listConnectorSpecification godoc
// @Summary Get connector specification details
// @Description Retrieve necessary environment variables
// @Tags specification
// @Accept json
// @Produce json
// @Success 200
// @Router /api/v1/specification/list [get]
func (a *App) listConnectorSpecification(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	respondWithJSON(w, http.StatusOK, map[string]string{

		"USERNAME":      "serviceNow Username",
		"PASSWORD":      "serviceNow Password",
		"INSTANCE_NAME": "serviceNow Instance Name",
	})
}

// // testConnectorSpecification godoc
// // @Summary Test auth against the specification
// // @Description Check if domain account can be verified
// // @Tags specification
// // @Accept  json
// // @Produce  json
// // @Param text body models.Specification false "Enter auth and domain name properties"
// // @Success 200
// // @Router /api/v1/specification/test [get]
// func (a *App) testConnectorSpecification(w http.ResponseWriter, r *http.Request) {

// 	//Allow CORS here By * or specific origin
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "*")

// 	if (*r).Method == "OPTIONS" {
// 		w.WriteHeader(200)
// 		return
// 	}

// 	var s models.Specification
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&s); err != nil {
// 		a.Log.Errorf("Error parsing json payload", err)
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	defer r.Body.Close()

// 	account := httpclient.GetAccount(s.Username, s.Password, r.URL.Query())
// 	if account != nil {
// 		respondWithJSON(w, http.StatusOK, account)
// 	} else {
// 		respondWithError(w, http.StatusBadRequest, "Account not verified")
// 	}

// }

// retrieveAllTableRecords godoc
// @Summary Retrieve all records from a table
// @Tags table
// @Accept  json
// @Produce  json
// @Param tableName path string true "Enter table name"
// @Success 200
// @Router /table/{tableName} [get]
func (a *App) retrieveAllTableRecords(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	vars := mux.Vars(r)
	tableName := vars["tableName"]
	res, err := httpclient.RetrieveAllTableRecords(tableName, a.Cfg.GetURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error in retrieveAllTableRecords", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, res)

}

// createTableRecords godoc
// @Summary Create a records for a table
// @Tags table
// @Accept  json
// @Produce  json
// @Param tableName path string true "Enter table name"
// @Success 200
// @Router /table/{tableName} [post]
func (a *App) createTableRecords(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	vars := mux.Vars(r)
	tableName := vars["tableName"]
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	} else {
		respondWithError(w, http.StatusBadRequest, "Empty Body")
	}
	res, err := httpclient.CreateTableRecords(tableName, a.Cfg.GetURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), bodyBytes, r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error in createTableRecords", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, res)
}

// retrieveSingleTableRecord godoc
// @Summary Retrieve a single record from a table
// @Tags table
// @Accept  json
// @Produce  json
// @Param tableName path string true "Enter table name"
// @Param sysId path string true "Enter table record sys id"
// @Success 200
// @Router /table/{tableName}/{sysId} [get]
func (a *App) retrieveSingleTableRecord(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	vars := mux.Vars(r)
	tableName := vars["tableName"]
	sysId := vars["sysId"]
	res, err := httpclient.RetrieveSingleTableRecord(tableName, sysId, a.Cfg.GetURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error in retrieveSingleTableRecord", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, res)

}

// modifyTableRecord godoc
// @Summary Modify a record from a table
// @Tags table
// @Accept  json
// @Produce  json
// @Param tableName path string true "Enter table name"
// @Param sysId path string true "Enter table record sys id"
// @Success 200
// @Router /table/{tableName}/{sysId} [put]
func (a *App) modifyTableRecord(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	vars := mux.Vars(r)
	tableName := vars["tableName"]
	sysId := vars["sysId"]
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	} else {
		respondWithError(w, http.StatusBadRequest, "Empty Body")
	}
	res, err := httpclient.ModifyTableRecord(tableName, sysId, a.Cfg.GetURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), bodyBytes, r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error in modifyTableRecord", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, res)

}

// deleteTableRecord godoc
// @Summary Delete a record from a table
// @Tags table
// @Accept  json
// @Produce  json
// @Param tableName path string true "Enter table name"
// @Param sysId path string true "Enter table record sys id"
// @Success 200
// @Router /table/{tableName}/{sysId} [delete]
func (a *App) deleteTableRecord(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	vars := mux.Vars(r)
	tableName := vars["tableName"]
	sysId := vars["sysId"]
	res, err := httpclient.DeleteTableRecord(tableName, sysId, a.Cfg.GetURL(), a.Cfg.GetUsername(), a.Cfg.GetPassword(), r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error in deleteTableRecord", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, res)
}
