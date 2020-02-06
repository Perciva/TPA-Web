import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Observable } from 'rxjs';
import gql from 'graphql-tag';
import {Query} from '../Models/query'
import {User} from '../Models/user'
@Injectable({
  providedIn: 'root'
})
export class ApolloService {

  constructor(private apollo: Apollo) { }
  
  
  //select all
  selectAllUser():Observable<Query>{
    return this.apollo.query<Query>({
      query: gql`
        query users{
          users{
            firstname
            lastname
            password
            email
            phone
          }
        }`
    })
  }

  //search
  searchUserByEmailPhone(arg:string): Observable<Query>{
    return this.apollo.query<Query>(
      {
        query: gql`
          query userbyemailphone($arg:String){
            userbyemailphone(arg:$arg){
              firstname
              lastname
              password
              email
              phone
            }
          }`,
          variables:{
            "arg":arg
          }
      }
    )
  }

  //CRUD
  createUser(newUser:User){
    return this.apollo.mutate({
      mutation:gql`
        mutation createuser(
          $firstname:String,
          $lastname:String,
          $email:String,
          $phone:String,
          $password:String
        ){
          createuser(
            phone: $phone
            firstname: $firstname
            lastname: $lastname
            password: $password
            email: $email
          ){
            id
          }
        }`,
        variables:{
          "firstname" : newUser.getFirstname(),
          "lastname" : newUser.getLastname(),
          "email" : newUser.getEmail(),
          "phone" : newUser.getPhone(),
          "password" : newUser.getPassword()
        }
    })
  }







}

