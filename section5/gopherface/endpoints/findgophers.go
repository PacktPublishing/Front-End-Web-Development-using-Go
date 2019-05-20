package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/EngineerKamesh/gofullstack/volume3/section5/gopherface/common/authenticate"

	"github.com/EngineerKamesh/gofullstack/volume3/section5/gopherface/common"
)

func FindGophersEndpoint(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		gfSession, err := authenticate.SessionStore.Get(r, "gopherface-session")
		if err != nil {
			log.Print(err)
			return
		}
		uuid := gfSession.Values["uuid"].(string)

		var searchTerm string
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}
		err = json.Unmarshal(reqBody, &searchTerm)
		if err != nil {
			log.Print("Encountered error when attempting to unmarshal JSON: ", err)
		}

		gophers, err := env.DB.FindGophers(uuid, searchTerm)

		if err != nil {
			log.Print(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(gophers)

	})
}
