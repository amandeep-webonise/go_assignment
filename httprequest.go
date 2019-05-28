package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
)

//Global storage of 'Result' type struct
var list []Result
var requestField string
var requestFieldValue string

// Main or Parent Struct Define
type ResponseBody struct {
	Results []Result `json:"results"`
	Info    Info     `json:"info"`
}

//Struct contain user information
type Result struct {
	Gender   string `json:"gender"`
	Name     Name   `json:"name"`
	Location struct {
		Street string `json:"street"`
		City   string `json:"city"`
		State  string `json:"state"`
		//Postcode    int    `json:"postcode,string"`
		Postcode    interface{} `json:"postcode"`
		Coordinates struct {
			Latitude  string `json:"latitude"`
			Longitude string `json:"longitude"`
		}
		Timezone struct {
			Offset      string `json:"offset"`
			Description string `json:"description"`
		}
	}
	Email string `json:"email"`
	Login struct {
		Uuid     string `json:"uuid"`
		Username string `json:"username"`
		Password string `json:"password"`
		Salt     string `json:"salt"`
		Md5      string `json:"md5"`
		Sha1     string `json:"sha1"`
		Sha256   string `json:"sha256"`
	}
	Dob struct {
		Date string `json:"date"`
		Age  int    `json:"age"`
	}
	Registered struct {
		Date string `json:"date"`
		Age  int    `json:"age"`
	}
	Phone string `json:"phone"`
	Cell  string `json:"cell"`
	Id    struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	Picture struct {
		Large     string `json:"large"`
		Medium    string `json:"medium"`
		Thumbnail string `json:"thumbnail"`
	}
	Nat string `json:"nat"`
}

type Info struct {
	Seed    string `json:"seed"`
	Results int    `json:"results"`
	Page    int    `json:"page"`
	Version string `json:"version"`
}
type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

//Function make request to a URL to get user information
func ResquestURL() Result {
	resp, err := http.Get("https://randomuser.me/api/")
	if err != nil {
		fmt.Println("There is some error while making request to API")

	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Not able to get userinfo")
	}

	result := &ResponseBody{}
	err4 := json.Unmarshal([]byte(body), result)

	if err4 != nil {
		fmt.Println("There was an error:", err4)
	}

	return result.Results[0]
}

//New Search
//API function that perform search on user info
func searchRecord(w http.ResponseWriter, r *http.Request) {

	for k, v := range r.URL.Query() {
		fmt.Printf("%s: %s\n", k, v[0])
		requestField = k
		requestFieldValue = v[0]
	}

	resultlist := make([]Result, 0)
	for _, elem := range list {
		found, _ := Parse(elem)
		//fmt.Println("----FOUND RESULT Here----", found, "-------END FOUND RESULT")

		if found == 1 {
			//fmt.Println("----FOUND RESULT Here----", found, "-------END FOUND RESULT")
			resultlist = append(resultlist, elem)
		} else {
			fmt.Println(elem)
		}

	}

	js, err := json.Marshal(resultlist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if len(resultlist) > 0 {
		w.Write(js)
	} else {
		json.NewEncoder(w).Encode(map[string]string{"message": "No record Found"})
	}
}

//End New Search

//Save record api
func saveRecord(w http.ResponseWriter, r *http.Request) {
	requestResult := ResquestURL()
	//API response append to global temp storage
	list = append(list, requestResult)
	js, err := json.Marshal(list)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//var buffer bytes.Buffer
	//buffer.WriteString({Response: "success", Message: "Welcome to awesome go service"})
	//json.NewEncoder(w).Encode(buffer.String())
	w.Write(js)
	//json.NewEncoder(w).Encode(map[string]string{"message": "Record save sucessfully!!"})
}

//API request handler
func handleRequests() {
	http.HandleFunc("/search", searchRecord)
	http.HandleFunc("/save", saveRecord)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	fmt.Println("**** Request Start ****")
	handleRequests()
}

func doParse(data reflect.Value, status int) int {

	t := data.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i).Name
		fv := data.Field(i)
		if fv.Type().Kind() == reflect.Struct {
			fmt.Println("Key {}", f)
			status = doParse(data.Field(i), status)
			continue
		}

		if strings.ToLower(f) == strings.ToLower(requestField) {
			if strings.ToLower(fv.String()) == strings.ToLower(requestFieldValue) {
				status = 1
				break
			}

		}
	}
	return status

}

func Parse(v interface{}) (int, error) {
	ptrRef := reflect.ValueOf(v)
	if ptrRef.Kind() != reflect.Struct {
		return 0, errors.New("not a struct ")
	}
	return doParse(ptrRef, 0), nil
}
