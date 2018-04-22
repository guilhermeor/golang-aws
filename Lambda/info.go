package main

//Request stores the information needed to complete the requisition
type Request struct {
	ID    float32 `json:"id"`
	Value string  `json:"value"`
}

//Response stores the information needed to respond to the request
type Response struct {
	Message string `json:"message"`
	OK      bool   `json:"ok"`
}
