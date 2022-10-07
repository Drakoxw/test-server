package interfaces

type GettBancosAve struct {
	// Controller string `json:"id"`
	Type string `json:"tipo"`
}

type ResponseAveGen struct {
	Status  string
	Message string
	Data    []any
}
