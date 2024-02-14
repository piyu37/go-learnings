package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Address  *Address
	Phone    string
	Website  string
	Company  *Company
}

type Address struct {
	Street  string
	Suite   string
	City    string
	Zipcode string
	Geo     *Geo
}

type Geo struct {
	Lat, Lng string
}

type Company struct {
	Name        string
	CatchPhrase string
	BS          string
}

func fetchData(w http.ResponseWriter, r *http.Request) {

}

func server() {
	http.HandleFunc("/user", fetchData)
	panic(http.ListenAndServe("localhost:9767", http.HandlerFunc(fetchData)))
}

func client() {
	req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/users",
		nil)
	if err != nil {
		fmt.Println(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)

		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))

	var users []*User

	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)
}

func main() {
	client()
}
