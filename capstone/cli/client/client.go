package client

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Goal struct {
	ID          int64         `json:"id"`
	Discription string        `json:"discription"`
	Completed   bool          `json:"completed"`
	UserID      sql.NullInt32 `json:"user_id"`
	CreatedAt   time.Time     `json:"created_at"`
}

type Client struct {
	BaseURL    string
	HttpClient *http.Client
}

func NewClient() Client {
	return Client{
		BaseURL: "http://localhost:8000/user",
		HttpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (c *Client) CreateUser(a, b string) {
	jsonBody := []byte(`{"first_name": "` + a + `", "last_name": "` + b + `"}`)
	bodyReader := bytes.NewReader(jsonBody)

	res, err := http.NewRequest("POST", "http://localhost:8000/user", bodyReader)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	r, err := client.Do(res)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	user := &User{}
	err = json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		log.Fatal(err)
	}

	// if user == nil {
	// 	log.Fatalf("uh oh. It's empty. Please try again")
	// }

	fmt.Printf("Successfully created user: %s %s with ID: %d", user.FirstName, user.LastName, user.ID)
}

func (c *Client) DeleteData(id, gID int64, f string) {
	var path string
	if f == "" && gID == 0 {
		path = fmt.Sprintf("http://localhost:8000/user/%d", id)
	} else {
		path = fmt.Sprintf("http://localhost:8000/user/%d/%s/%d", id, f, gID)
	}

	res, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	r, err := client.Do(res)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	if r.Status == "200 OK" {
		fmt.Println("account has been deleted - response status : ", r.Status)
	} else {
		fmt.Println("something went wrong - response status : ", r.Status)
	}

}

func (c *Client) CreateGoal(id int64, d, f string) {
	jsonBody := []byte(`{"discription": "` + d + `"}`)
	bodyReader := bytes.NewReader(jsonBody)

	path := fmt.Sprintf("http://localhost:8000/user/%d/%s", id, f)

	res, err := http.NewRequest("POST", path, bodyReader)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	r, err := client.Do(res)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	goal := &Goal{}
	err = json.NewDecoder(r.Body).Decode(goal)
	if err != nil {
		log.Fatal(err)
	}

	if r.Status == "200 OK" {
		fmt.Printf("successfully created a %s - response status %s: ", f, r.Status)
		fmt.Printf("User ID: %d\nGoal ID: %d\nDetail: %s\nComplete: %v\n", goal.UserID.Int32, goal.ID, goal.Discription, goal.Completed)
	} else {
		fmt.Println("something went wrong - response status : ", r.Status)
	}

}

func (c *Client) UpdateGoals(id, gID int64, s bool, f string) {
	body := fmt.Sprintf(`{"completed": %t}`, s)

	jsonBody := []byte(body)
	bodyReader := bytes.NewReader(jsonBody)

	path := fmt.Sprintf("http://localhost:8000/user/%d/%s/%d", id, f, gID)

	res, err := http.NewRequest("PUT", path, bodyReader)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	r, err := client.Do(res)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	goal := &Goal{}
	err = json.NewDecoder(r.Body).Decode(goal)
	if err != nil {
		log.Fatal(err)
	}

	if r.Status == "200 OK" {
		fmt.Printf("successfully updated a %s - response status %s: \n", f, r.Status)
		fmt.Printf("User ID: %d\nGoal ID: %d\nDetail: %s\nComplete: %v\n", goal.UserID.Int32, goal.ID, goal.Discription, goal.Completed)
	} else {
		fmt.Println("something went wrong - response status : ", r.Status)
	}

}

func (c *Client) GetGoals(id int64, g, f string) {
	var path string
	if f == "" {
		path = fmt.Sprintf("http://localhost:8000/user/%d/%s", id, g)
	} else {
		path = fmt.Sprintf("http://localhost:8000/user/%d/%s/%s", id, g, f)
	}

	res, err := http.NewRequest("GET", path, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	r, err := client.Do(res)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	// respBody, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(respBody))
	var goals []Goal
	err = json.NewDecoder(r.Body).Decode(&goals)
	if err != nil {
		log.Fatal(err)
	}

	if r.Status == "200 OK" {
		fmt.Printf("successfully created a %s - response status %s: \n", f, r.Status)
		fmt.Printf("User ID: %d\n", goals[0].UserID.Int32)
		for _, goal := range goals {
			fmt.Printf("Goal ID: %d | Detail: %s | Complete: %v\n", goal.ID, goal.Discription, goal.Completed)
		}

	} else {
		fmt.Println("something went wrong - response status : ", r.Status)
	}

}
