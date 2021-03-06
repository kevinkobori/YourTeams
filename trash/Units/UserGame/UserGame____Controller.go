package UserGame

import (
	"be/Entities/UserEntity"
	"be/HyperText"
	"be/Interface"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

//-------------------------------- ADM --------------------------------//
func (c *GameEntityController) Insert(w http.ResponseWriter, r *http.Request) {
	var age GameEntity
	Interface.PutDocValidateNameExist(DOCNAME)
	if err := HyperText.BodyValidate(r, &age); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	uev := UserEntity.UserEntityVerify{}
	uev.Username = age.Username
	uev.Password = age.Password
	ue, boolean := UserEntity.Verify(uev)
	if boolean == false {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["wrong-verify"])
		return
	} else if ue. == false {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["not-"])
		return
	}
	if err := c.GameEntityRepository.InsertGameEntity(&age); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, age)
	return
}

//-------------------------------- USR --------------------------------//
func (c *GameEntityController) UpdateSingle(w http.ResponseWriter, r *http.Request) {
	ageus := GameEntityUpdateSingle{}
	Interface.PutDocValidateNameExist(DOCNAME)
	if err := HyperText.BodyValidate(r, &ageus); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := ValidateUpdateSingle(ageus); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	uev := UserEntity.UserEntityVerify{}
	uev.Username = ageus.Username
	uev.Password = ageus.Password
	if _, boolean := UserEntity.Verify(uev); boolean == false {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["wrong-verify"])
		return
	}
	if _, err := Interface.UpdateSingleDB(ageus.Position, ageus.Value, ageus, DOCNAME, ageus.Name); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["error-update"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, HyperText.CustomResponses["success-update"])
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GameEntityController) Update(w http.ResponseWriter, r *http.Request) {
	var age GameEntity
	Interface.PutDocValidateNameExist(DOCNAME)
	age.Name = mux.Vars(r)["name"]
	if err := HyperText.BodyValidate(r, &age); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}
	uev := UserEntity.UserEntityVerify{}
	uev.Username = age.Username
	uev.Password = age.Password
	ue, boolean := UserEntity.Verify(uev)
	if boolean == false {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["wrong-verify"])
		return
	} else if ue. == false {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["not-"])
		return
	}
	age.Name = strings.ToLower(age.Name)
	result, err := Interface.UpdateDB(age, DOCNAME, age.Name)
	if err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["error-update"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, result)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GameEntityController) GetUnit(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	name = strings.ToLower(name)
	result := GameEntity{}
	if err := Interface.FindUnitDB(name, &result, DOCNAME); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["not-found-entity"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, result)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GameEntityController) GetAllWhile(w http.ResponseWriter, r *http.Request) {
	position := mux.Vars(r)["position"]
	value := mux.Vars(r)["value"]
	position = strings.ToLower(position)
	value = strings.ToLower(value)
	entities := GameEntities{}
	if err := Interface.FindAllWhileDB(position, value, &entities, DOCNAME); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["empty-Interface"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, entities)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GameEntityController) GetAll(w http.ResponseWriter, r *http.Request) {
	result := GameEntities{}
	if err := Interface.FindAllDB(&result, DOCNAME); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["empty-Interface"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, result)
	return
}

//-------------------------------- ADM --------------------------------//
func (c *GameEntityController) Delete(w http.ResponseWriter, r *http.Request) {
	name := strings.ToLower(mux.Vars(r)["name"])
	result := GameEntity{}
	if err := Interface.DeleteDB(name, &result, DOCNAME); err != nil {
		HyperText.HttpErrorResponse(w, http.StatusBadRequest, HyperText.CustomResponses["not-found-entity"])
		return
	}
	HyperText.HttpResponse(w, http.StatusOK, result)
	return
}
