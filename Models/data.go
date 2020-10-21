package Models

type Data struct {
	Device string `json:"devices"`
	Date   string `json:"date"`
	Value  string `json:"value"`
}

type Location struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}
