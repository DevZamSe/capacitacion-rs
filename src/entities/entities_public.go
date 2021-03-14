package entities

type ResponseHello struct {
	Message string `json:"mensaje"`
	Nombre  string `json:"name"`
}

type ResponseGeneral struct {
	Result string `json:"result"`
}

type RequestName struct {
	Pass   string `json:"pass"`
	Nombre string `json:"name"`
}
