package models

import (
	"fmt"
	"encoding/json"
  _ "github.com/lib/pq"
  "strings"
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
type AppGroups struct {
	Application       string
	GroupName     string
}
type UserGroups struct {
	UserName       string
	GroupName     string
}
type UserAppGroups struct {
	UserName string
	Application string
	Groups   []string
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
type repositoriesAppGroups struct {
	AppGroupRepositories []AppGroups
}
type repositoriesUserGroups struct {
	UserGroupRepositories []UserGroups
}
func GetUserAppGroupsData(username string, application string) string {
	 
	repos := UserAppGroups{}
	repos.Application = application
	repos.UserName = username
	sql := "SELECT groups.group_name FROM  "+
	       " application,user_table,user_groups,groups,application_groups "+
	       " WHERE user_groups.user_id = user_table.user_id AND application_groups .app_id =application.app_id "+
	       " AND user_groups.group_id = groups.group_id  AND application_groups.group_id = groups.group_id  "+
	       " AND user_table.username = $1 AND application.application_name = $2"
   // fmt.Printf(sql)   
    rows, err := db.Query(sql,&username,&application)
	if err != nil {
		fmt.Println("Query Error :",err)
		
	}
	defer rows.Close()
	groups := make([]string, 0)
	for rows.Next() {
		var group string
		err := rows.Scan(&group)
		if err != nil {
			fmt.Println("Error encountered fetch data",err)
		
		}
		groups = append(groups, group)
	 }	
	 repos.Groups = groups
	 out, err := json.Marshal(repos)
	 if err != nil {
	  panic(err)
	 }
     fmt.Println(string(out))
     
     return string(out)
}	
func GetAppGroupsData(application string) string {
	
	repos := repositoriesAppGroups{}
	sql := "SELECT application.application_name,groups.group_name "+  
       " FROM application,application_groups,groups WHERE application_groups.app_id = application.app_id "+
       " AND application_groups.group_id = groups.group_id AND application.application_name = $1"
   // fmt.Printf(sql)   
    rows, err := db.Query(sql,&application)
	if err != nil {
		fmt.Println("Query Error :",err)
	}
	defer rows.Close()
	for rows.Next() {
		repo := AppGroups{}
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
func GetUserGroupsData(username string) string {
	
	repos := repositoriesUserGroups{}
	sql := "SELECT user_table.username,groups.group_name "+  
       " FROM user_table,user_groups,groups WHERE user_groups.user_id = user_table.user_id "+
       " AND user_groups.group_id = groups.group_id AND user_table.username = $1"
   // fmt.Printf(sql)   
    rows, err := db.Query(sql,&username)
	if err != nil {
		fmt.Println("Query Error :",err)
	}
	defer rows.Close()
	for rows.Next() {
		repo := UserGroups{}
		err := rows.Scan(
			&repo.UserName,
			&repo.GroupName, 
		)
		if err != nil {
			fmt.Println("Error encountered fetch data",err)           
		}
		repos.UserGroupRepositories = append(repos.UserGroupRepositories, repo)
	 }	
	 out, err := json.Marshal(repos)
	 if err != nil {
	    panic(err)
	 }
     fmt.Println(string(out))
     
     return string(out)
}	

func GetUserAuthorizedData(username string,groups []string) string {
	
	repos := repositoriesUserGroups{}
	group_params := strings.Join(groups, "','")
	sql := "SELECT user_table.username,groups.group_name "+  
       " FROM user_table,user_groups,groups WHERE user_groups.user_id = user_table.user_id "+
       " AND user_groups.group_id = groups.group_id AND user_table.username = $1 AND groups.group_name IN ('%s') "
       fmt.Printf(sql)   
    sqlRaw := fmt.Sprintf(sql, group_params)
	rows, err := db.Query(sqlRaw, &username)    
	if err != nil {
		fmt.Println("Query Error :",err)
	}
	defer rows.Close()
	for rows.Next() {
		repo := UserGroups{}
		err := rows.Scan(
			&repo.UserName,
			&repo.GroupName, 
		)
		if err != nil {
			fmt.Println("Error encountered fetch data",err)
		}
		repos.UserGroupRepositories = append(repos.UserGroupRepositories, repo)
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
		FROM groups` )
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
