package models

import (
	"fmt"
	"encoding/json"
  _ "github.com/lib/pq"
)

// repository contains the details of a repository
type Applications struct {
	App_Id      int
	Application string
}
type Groups struct {
	Group_Id       int
	GroupName      string
}
type Users struct {
	Use_Id       int
	UserName     string
}
type AppGroup struct {
	Application       string
	GroupName     string
}

type repositoriesApplication struct {
	AppRepositories []Applications
}
type repositoriesUser struct {
	UserRepositories []Users
}
type repositoriesGroup struct {
	GroupRepositories []Groups
}
type repositoriesAppGroup struct {
	AppGroupRepositories []AppGroup
}
func GetAppGroupsData(application string) {
	repos := repositoriesAppGroup{}
    rows, err := db.Query(`SELECT application.application_name,groups.group_name  
       FROM application,application_groups,groups WHERE application_groups.app_id = application.app_id 
        AND application_groups.group_id = groups.group_id AND application.application_name=?
    `,application)
	if err != nil {
		fmt.Println("Query Error :",err)
	}
	defer rows.Close()
	for rows.Next() {
		repo := AppGroup{}
		err := rows.Scan(
			&repo.Application,
			&repo.GroupName, 
		)
		if err != nil {
			fmt.Println("Error encountered fetch data",err)
		}
		repos.AppGroupRepositories = append(repos.AppGroupRepositories, repo)
	 }	
	 out, err := json.Marshal(repos)
	 if err != nil {
	    panic(err)
	 }
     fmt.Println(string(out))
     
     return string(out)
}	
func  GetUsersData() string{
	
   repos := repositoriesUser{}
   rows, err := db.Query(`
		SELECT
			user_id,
			username
		FROM user_table` )
	if err != nil {
		fmt.Println("Query Error :",err)
	}
	defer rows.Close()
	for rows.Next() {
		repo := Users{}
		err := rows.Scan(
			&repo.Use_Id,
			&repo.UserName, 
		)
		if err != nil {
			fmt.Println("Error encountered fetch data",err)
		}
		repos.UserRepositories = append(repos.UserRepositories, repo)
	 }	
	 out, err := json.Marshal(repos)
	 if err != nil {
	    panic(err)
	 }
     fmt.Println(string(out))
     
     return string(out)
}
func  GetApplicationsData() string{
  repos := repositoriesApplication{}
  rows, err := db.Query(`
		SELECT
			app_id,
			application_name
		FROM application` )
	if err != nil {
		fmt.Println("Query Error :",err)
	}
	defer rows.Close()
	for rows.Next() {
		repo := Applications{}
		err := rows.Scan(
			&repo.App_Id,
			&repo.Application, 
		)
		if err != nil {
			fmt.Println("Error encountered fetch data",err)
		}
		repos.AppRepositories = append(repos.AppRepositories, repo)
	}	  
  out, err := json.Marshal(repos)
  if err != nil {
    panic(err)
  }
  fmt.Println(string(out))
  
  return string(out)
}
func  GetGroupsData() string{
	
 repos := repositoriesGroup{}
  rows, err := db.Query(`
		SELECT
			group_id,
			group_name
		FROM group` )
	if err != nil {
		fmt.Println("Query Error :",err)
	}
	defer rows.Close()
	for rows.Next() {
		repo := Groups{}
		err := rows.Scan(
			&repo.Group_Id,
			&repo.GroupName, 
		)
		if err != nil {
			fmt.Println("Error encountered fetch data",err)
		}
		repos.GroupRepositories = append(repos.GroupRepositories, repo)
	}	
  
  out, err := json.Marshal(repos)
  if err != nil {
    panic(err)
  }
  fmt.Println(string(out))
  return string(out)
}
