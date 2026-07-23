package vo

type PropertyVO struct {
	Name  string `json:"name" jsonapi:"attr,name"`
	Kind  string `json:"kind" jsonapi:"attr,kind"`
	Value string `json:"value" jsonapi:"attr,value"`
}
