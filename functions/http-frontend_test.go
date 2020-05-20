package functions

import (
	"net/http"
	"strings"
	"testing"
)

func TestHttpFrontend(t *testing.T) {

	//serviceUrl := os.Getenv("BOOTSTRAP_DATA_SERVER")
	//if serviceUrl != "https://europe-west1-bootstrap-data-cloudfunctions.cloudfunctions.net" {
	//	t.Errorf("BOOTSTRAP_DATA_SERVER environment variable expected %q set %q",
	//		"https://europe-west1-bootstrap-data-cloudfunctions.cloudfunctions.net", serviceUrl)
	//}

	serviceUrl := "https://europe-west1-hybrid-cloud-22365.cloudfunctions.net"

	// Send ping request to service
	res, err := http.Post(serviceUrl+"/ping",
		"application/x-www-form-urlencoded",
		strings.NewReader(""))
	if err != nil {
		t.Errorf("failed to ping BOOTSTRAP_DATA_SERVER: %v\n", err)
		return
	}
	err = res.Body.Close()
	if err != nil {
		t.Errorf("cannot close response body\n")
	}
}

//curl -X POST "https://europe-west1-hybrid-cloud-22365.cloudfunctions.net/translate" \
//-d '{ "clientId": "beab10c6-deee-4843-9757-719566214526", "text": "Today is Monday", "sourceLanguage": "en",  "targetLanguage": "fr"}'
