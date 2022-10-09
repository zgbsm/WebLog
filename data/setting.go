package data

type Conf struct {
	Web      int `yaml:"web"`
	Listener int `yaml:"listener"`
}

var Config Conf

func ErrHandle(err error) {
	if err != nil {
		panic(err)
	}
}
