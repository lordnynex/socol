package collector

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
)

func Pinterest() Platform {
  return Platform{
    enabled:  true,
    name:     "pinterest",
    statsUrl: "http://api.pinterest.com/v1/urls/count.json?callback=call&url=%s",
    format:   "jsonp",
    parseWith: func(r *http.Response) (Stat, error) {
      stat := Stat{
        data: map[string]interface{}{"count": 0},
      }

      body, error := ioutil.ReadAll(r.Body)
      if error != nil {
        return stat, error
      }

      jsonBody, error := parseJSONP(body)
      if error != nil {
        return stat, error
      }

      var jsonBlob map[string]interface{}
      if err := json.Unmarshal([]byte(jsonBody), &jsonBlob); err != nil {
        return stat, err
      }

      return Stat{
        data: map[string]interface{}{"count": jsonBlob["count"]},
      }, nil
    }}
}
