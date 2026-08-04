package repositories

import (
	"api/interfaces"
	"encoding/json"
	"os"
)

func LoadData() interfaces.Data {
	file, err := os.ReadFile("data.json")
	if err != nil {
		return interfaces.Data{}
	}

	var data interfaces.Data

	err = json.Unmarshal(file, &data)
	if err != nil {
		return interfaces.Data{}
	}

	return data
}

func storeAuthors(Authors []interfaces.Author) {
	data := interfaces.Data{
		Blogs:   ReadAllBlogs(),
		Authors: Authors,
	}

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return
	}

	err = os.WriteFile("data.json", file, 0644)
	if err != nil {
		return
	}
}

func storeBlogs(Blogs []interfaces.Blog) {
	data := interfaces.Data{
		Blogs:   Blogs,
		Authors: ReadAllAuthors(),
	}

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return
	}

	err = os.WriteFile("data.json", file, 0644)
	if err != nil {
		return
	}
}
