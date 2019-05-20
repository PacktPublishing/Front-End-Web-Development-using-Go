package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/EngineerKamesh/gofullstack/volume3/section5/gopherface/common/authenticate"

	"github.com/EngineerKamesh/gofullstack/volume3/section5/gopherface/common"
)

func FollowGopherEndpoint(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		gfSession, err := authenticate.SessionStore.Get(r, "gopherface-session")
		if err != nil {
			log.Print(err)
			return
		}
		uuid := gfSession.Values["uuid"].(string)

		var gopherUUID string
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}
		err = json.Unmarshal(reqBody, &gopherUUID)
		if err != nil {
			log.Print("Encountered error when attempting to unmarshal JSON: ", err)
		}

		err = env.DB.FollowGopher(uuid, gopherUUID)

		if err != nil {
			log.Print(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)

	})
}
