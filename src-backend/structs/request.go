package structs

type Request struct {
  Message Message `json:"message"`
  Method  string  `json:"method"`
}
