package main

import (
   
    //"encoding/json"
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

func getAppGroups(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, postgresdb.getAppGroupsData())
	fmt.Println("Endpoint Hit: GroupsAPI")
}
func handleRequests() {
	postgresdb.InitDB()
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/api/getAppAuthorization", getAppGroups)
    myRouter.HandleFunc("/api/users", getUsers)
    myRouter.HandleFunc("/api/applications", getApplications)
    myRouter.HandleFunc("/api/groups", getGroups)
    log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {    
    handleRequests()
}