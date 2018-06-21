package main

import (
  "strings"
)

func main()  {

  user1 := CreateUser("Jesse Pinkman", RoleEmployee)
  printUserAcessDetails(user1)

  user2 := CreateUser("Mike Ehrmantraut", RoleManager)
  printUserAcessDetails(user2)

  user3 := CreateUser("Walter White", RoleSupervisor, RoleDistrictLeader)
  printUserAcessDetails(user3)

  user4 := CreateUser("Gus Fring", RoleManager, RoleCorporateLeader)
  printUserAcessDetails(user4)
}

func printUserAcessDetails(user User) {

  println(user.String())
  user.PrintPermissionAccessDescription(PermissionFeatureA)
  user.PrintPermissionAccessDescription(PermissionFeatureB)
  user.PrintPermissionAccessDescription(PermissionFeatureC)
  user.PrintPermissionAccessDescription(PermissionFeatureD)
  user.PrintPermissionAccessDescription(PermissionFeatureE)
  user.PrintPermissionAccessDescription(PermissionFeatureF)
}


// USER
type User struct {

  Name string
  Roles []Role
}

func CreateUser(name string, roles ...Role) User {

  user := User{name, []Role{}}
  user.Roles = roles
  return user
}

func (user User) HasPermission(requestedPermissionType PermissionType) bool {

  for _, role := range user.Roles {
    for _, permissionType := range role.PermissionsTypes {
      if permissionType == requestedPermissionType {
        return true
      }
    }
  }
  return false
}

func (user User) String() string {

  description := "\n\nName: " + user.Name

  description += "\nRoles: "
  roleTypeStrings := []string{}
  distinctPermissionTypeValues := map[PermissionType]bool{}
  for _, role := range user.Roles {
    roleTypeStrings = append(roleTypeStrings, role.Type.String())
    for _, permissionType := range role.PermissionsTypes {
      distinctPermissionTypeValues[permissionType] = true
    }
  }
  description += strings.Join(roleTypeStrings[:], ", ")

  description += "\nPermissions: "
  permissionTypeStrings := []string{}
  for permissionTypeValue := range distinctPermissionTypeValues {
    permissionTypeStrings = append(permissionTypeStrings, permissionTypeValue.String())
  }
  description += strings.Join(permissionTypeStrings[:], ", ")

  return description
}

func (user User) PrintPermissionAccessDescription(requestedPermissionType PermissionType) {
  println(user.Name, "has permission for", requestedPermissionType.String(), ":", user.HasPermission(requestedPermissionType))
}


// ROLE
type Role struct {
  Type RoleType
  PermissionsTypes []PermissionType
}

var RoleEmployee = Role{RoleTypeEmployee, []PermissionType{PermissionFeatureA}}
var RoleManager = Role{RoleTypeManager, []PermissionType{PermissionFeatureA, PermissionFeatureB, PermissionFeatureC}}
var RoleSupervisor = Role{RoleTypeSupervisor, []PermissionType{PermissionFeatureA, PermissionFeatureB, PermissionFeatureC}}
var RoleDistrictLeader = Role{RoleTypeDistrictLeader, []PermissionType{PermissionFeatureA, PermissionFeatureB, PermissionFeatureC, PermissionFeatureD, PermissionFeatureE}}
var RoleRegionalLeader = Role{RoleTypeRegionalLeader, []PermissionType{PermissionFeatureA, PermissionFeatureB, PermissionFeatureC, PermissionFeatureD}}
var RoleCorporateLeader = Role{RoleTypeCorporateLeader, []PermissionType{PermissionFeatureA, PermissionFeatureB, PermissionFeatureC, PermissionFeatureD, PermissionFeatureE, PermissionFeatureF}}


// ROLE TYPE
type RoleType int
const (
  RoleTypeEmployee        RoleType = iota
  RoleTypeManager         RoleType = iota
  RoleTypeSupervisor      RoleType = iota
  RoleTypeDistrictLeader  RoleType = iota
  RoleTypeRegionalLeader  RoleType = iota
  RoleTypeCorporateLeader RoleType = iota
)

func (roleType RoleType) String() string {
  names := []string{"Employee", "Manager", "Supervisor", "District Leader", "Regional Leader", "Corporate Leader"}
  return names[roleType]
}


// PERMISSION TYPE
type PermissionType int
const (
  PermissionFeatureA  PermissionType = iota
  PermissionFeatureB  PermissionType = iota
  PermissionFeatureC  PermissionType = iota
  PermissionFeatureD  PermissionType = iota
  PermissionFeatureE  PermissionType = iota
  PermissionFeatureF  PermissionType = iota
)

func (permissionType PermissionType) String() string {
  names := []string {"Feature A", "Feature B", "Feature C", "Feature D", "Feature E", "Feature F"}
  return names[permissionType]
}

func permissionTypeByValue(permissionTypeValue int) PermissionType {
  permissionsTypes := []PermissionType{PermissionFeatureA, PermissionFeatureB, PermissionFeatureC, PermissionFeatureD, PermissionFeatureE, PermissionFeatureF}
  return permissionsTypes[permissionTypeValue]
}
