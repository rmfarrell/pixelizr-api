package main

import (
  "encoding/json"
  "github.com/apex/go-apex"
  "log"
  // "github.com/apex/log"
  // "github.com/apex/log/handlers/text"
)

// type s3Notification struct {
//   Records []struct {
//     S3 struct {
//       Bucket struct {
//         Arn string `json:"arn"`
//         Name string `json:"name"`
//       } `json:"bucket"`
//       Object struct {
//         Key string `json:"key"`
//         Size int `json:"size"`
//       } `json:"object"`
//     } `json:"s3"`
//   } `json:"Records"`
// }

type pixelizrRequest struct {
  MediaType string `json:"type"`
  Resolution string `json:"output_resolution"`
  Pattern string `json:"pattern"`
  Media string `json:"media"` //TODO
}

func main() {

  // log.SetHandler(text.New(os.Stderr))

  apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
    var m pixelizrRequest

    if err := json.Unmarshal(event, &m); err != nil {
      log.Print(err)
      return nil, err
    }
    // log.Infof(m)
    // log.Infof(m.Records[0].S3.Object.Key)
    // log.Infof(string(event.Records[0].s3.object.key))

    return m, nil
  })
}