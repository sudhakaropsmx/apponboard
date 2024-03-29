package main

import (
   
    "encoding/json"
    "fmt"
    "log"
    //"io/ioutil"
    "net/http"
    "github.com/gorilla/mux"
     postgresdb "github.com/sudhakaropsmx/apponboard/models"
     
)


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Spinnaker Apponboard Compliance!")   
    fmt.Println("Endpoint Hit: homePage")
}
func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,postgresdb.GetUsersData())
	fmt.Println("Endpoint Hit: UserAPI")
}

func getApplications(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, postgresdb.GetApplicationsData())
	fmt.Println("Endpoint Hit: ApplicationAPI")
}

func getGroups(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, postgresdb.GetGroupsData())
	fmt.Println("Endpoint Hit: GroupsAPI")
}
type myAppData struct {
	Application string
}
func getAppGroups(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
       http.Error(w, http.StatusText(405), 405)
      return
   }
	decoder := json.NewDecoder(r.Body)
	var data myAppData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, postgresdb.GetAppGroupsData(data.Application))
	fmt.Println("Endpoint Hit: AppGroupsAPI")
}
type myUserData struct {
	UserName string
}
func getUserGroups(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
       http.Error(w, http.StatusText(405), 405)
      return
   }
	decoder := json.NewDecoder(r.Body)
	var data myUserData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, postgresdb.GetUserGroupsData(data.UserName))
	fmt.Println("Endpoint Hit: UserGroupsAPI")
}
type UserApp struct {
	UserName string
	Application string
}
func GetUserAppAuthorized(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
       http.Error(w, http.StatusText(405), 405)
      return
   }
	decoder := json.NewDecoder(r.Body)

	var data UserApp
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	userdata := postgresdb.GetUserAppGroupsData(data.UserName,data.Application)
	if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
	fmt.Fprintf(w, userdata )
	fmt.Println("Endpoint Hit: GetUserAuthorizedAPI")
}

type myData struct {
	UserName string
	Groups  []string
}
func GetUserAuthorized(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
       http.Error(w, http.StatusText(405), 405)
      return
   }
	decoder := json.NewDecoder(r.Body)

	var data myData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	groups := data.Groups
	
	userdata := postgresdb.GetUserAuthorizedData(data.UserName,groups)
	if err != nil {
        http.Error(w, http.StatusText(500), 500)
        return
    }
	fmt.Fprintf(w, userdata )
	fmt.Println("Endpoint Hit: GetUserAuthorizedAPI")
}

func handleRequests() {
	postgresdb.InitDB()
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/api/getUserAppAuthorized",GetUserAppAuthorized).Methods("POST")
    myRouter.HandleFunc("/api/getAppAuthorized",GetUserAuthorized).Methods("POST")
    myRouter.HandleFunc("/api/getAppGroups", getAppGroups).Methods("POST")
    myRouter.HandleFunc("/api/getUserGroups", getUserGroups).Methods("POST")
    myRouter.HandleFunc("/api/users", getUsers)
    myRouter.HandleFunc("/api/applications", getApplications)
    myRouter.HandleFunc("/api/groups", getGroups)
    log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {    
    handleRequests()
}
