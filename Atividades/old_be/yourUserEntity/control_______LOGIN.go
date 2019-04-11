package yourUserEntity

import (
	"lfkk/be/myHttp"
	"lfkk/be/myJson"
	"lfkk/be/myValidator"
	"log"
	"net/http"
)

func (c *UserEntityController) LoginUserEntity(w http.ResponseWriter, r *http.Request) {
	var uev UserEntityVerify
	myHttp.MyWriteHeaderSet(w, r, myHttp.APPJASON_UTF_8)

	if err := myJson.DecodeJson(r.Body, &uev); err != nil {
		myHttp.MyWriteResponses(w, r, 478)
		log.Println(err)
		return
	}
	myValidator.UserEntityValidator()

	if err := myValidator.GetVarValidate().Struct(uev); err != nil {
		myHttp.MyWriteResponses(w, r, ErrorInsertUserEntity(w, r, uev.Username, err))
		return
	}

	if _, err := c.UserEntityRepository.LoginUserEntityDB(uev); err != nil {
		myHttp.MyWriteResponses(w, r, ErrorInsertUserEntity(w, r, uev.Username, err))
		return
	}

	myHttp.MyWriteResponses(w, r, 219)
	return
}
