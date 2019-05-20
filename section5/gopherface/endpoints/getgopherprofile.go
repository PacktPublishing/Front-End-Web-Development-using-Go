package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/EngineerKamesh/gofullstack/volume3/section5/gopherface/common"
)

func GetGopherProfileEndpoint(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var gopherUUID string
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}
		err = json.Unmarshal(reqBody, &gopherUUID)
		if err != nil {
			log.Print("Encountered error when attempting to unmarshal JSON: ", err)
		}

		u, err := env.DB.GetGopherProfile(gopherUUID)

		if err != nil {
			log.Print(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)

	})
}
