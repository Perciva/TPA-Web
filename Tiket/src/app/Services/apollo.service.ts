import { Hotel } from './../Models/hotel';
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
    try{
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
    }catch(err){
      return null
    }
  }
  selectAllLocation():Observable<Query>{
    return this.apollo.query<Query>({
      query: gql`
      query getallloc{
        getallloc{
          city
          country
          id
          province
        }
      }`
    })
  }
  selectAllHotelFacility():Observable<Query>{
    return this.apollo.query<Query>({
      query: gql`
        query getallhotelfacility{
          getallhotelfacility{
            id
            hotelId
            name
            path
          }
        }`
    })
  }
  selectAllHotelImage():Observable<Query>{
    return this.apollo.query<Query>({
      query: gql`
        query getallhotelimage{
          getallhotelimage{
            id
            hotelId
            path
          }
        }`
    })
  }
  selectAllHotelType():Observable<Query>{
    return this.apollo.query<Query>({
      query: gql`
        query getallhoteltype{
          getallhoteltype{
            id
            hotelId
            path
          }
        }`
    })
  }
  selectAllHotel():Observable<Query>{
    return this.apollo.query<Query>({
      query: gql`
        query allHotel{
          allHotel{
            id
            name
            description
            price
            rating
            address
            facility{
              name
              path
            }
            photo{
              path
            }
            type{
              name
            }
          }
        }`,
        fetchPolicy: 'no-cache'
    })
  }
  
  //search
  searchUserByEmailPhone(arg:string): Observable<Query>{
      return this.apollo.query<Query>(
        {
          query: gql`
            query userbyemailphone($arg:String!){
              userbyemailphone(arg:$arg){
                id
                firstname
                lastname
                password
                email
                phone
              }
            }`,
            variables:{
              arg :arg
            }
        }
      )
  }
  searchLocationByProvince(province:string):Observable<Query>{
    return this.apollo.query<Query>(
      {
        query:gql`
          query getlocbyprovince($province:String!){
            getlocbyprovince(province:$province){
              id
              city
              province
              country
            }
          }`,
        variables:{
          province: province
        }
      }
    )
  }
  searchLocationByCity(city:string):Observable<Query>{
    return this.apollo.query<Query>(
      {
        query:gql`
        query getlocbycity($city:String){
          getlocbycity(city: $city){
            id
          }
        }`,
        variables:{
          city: city
        }
      }
    )
  }

  searchHotelByProvince(province:string):Observable<Query>{
    return this.apollo.query<Query>(
      {
        query:gql`
          query hotelbyprovince($province:String!){
            hotelbyprovince(province:$province){
              id
              name
              description
              rating
            }
          }`,
        variables:{
          province: province
        }
      }
    )
  }
  searchHotelById(id:string):Observable<Query>{
    return this.apollo.query<Query>(
      {
        query:gql`
          query hotelbyid($id:Int!){
            hotelbyid(id:$id){
              name
              description
              lat
              long
              price
              rating
            }
          }`,
        variables:{
          id:id
        }
      }
    )
  }
  getNearbyHotel(lat:any, long:any):Observable<Query>{
    return this.apollo.query<Query>(
      {
        query: gql
          `query getnearbyhotel($lat:Float!, $long:Float!){
            getnearbyhotel(lat:$lat, long:$long){
              id
    					name
              description
              price
              rating
    					photo {
    					  path
    					}
            }
          }`, 
          variables:{
            lat: parseFloat(lat),
            long: parseFloat(long)
          }
      }
    )
  }
  getHotelByLatLong(lat:any, long:any):Observable<Query>{
    return this.apollo.query<Query>(
      {
        query:gql`
          query hotelbylatlong($lat:Float!, $long:Float!){
            hotelbylatlong(lat:$lat, long:$long){
              name
              description
              price
              rating
            }
          }`,
          variables:{
            lat: parseFloat(lat),
            long: parseFloat(long)
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
  createHotel(newHotel: Hotel){
    return this.apollo.mutate({
      mutation:gql`
        mutation createhotel(
          $rating:Float,
          $lat:Float,
          $long:Float,
          $description:String,
          $name:String,
          $address:String,
          $location:String,
          $price:Int
        ){
          createhotel(
            rating: $rating,
            lat:$lat,
            long: $long,
            description:$description,
            name: $name,
            address: $address,
            location:$location,
            price:$price
          ){
            id
          }
        }`,
        variables:{
          "rating": newHotel.rating,
          "lat": newHotel.lat,
          "long": newHotel.long,
          "name": newHotel.name,
          "price": newHotel.price,
          "description": newHotel.desc,
          "location": newHotel.city,
          "address": newHotel.address
        }
    })
  }
  createHotelFacility(hotelId:number, name:string, path:string){
    return this.apollo.mutate({
      mutation: gql`
        mutation createhotelfacility(
          $hotelId: Int,
          $name:String,
          $path:String
          
        ){
          createhotelfacility(
            hotelId:$hotelId,
            name: $name,
            path:$path
            
          ){
            id
          }
        }`,
      variables:{
        "hotelId": hotelId,
        "name": name,
        "path": path
      }
    })
  
  }
  createHotelImage(hotelid:number, path:string){
    console.log(hotelid, path)
    return this.apollo.mutate({
      mutation: gql`
        mutation createhotelimage(
          $hotelId: Int,
          $path:String
          
        ){
          createhotelimage(
            hotelId:$hotelId,
            path:$path
            
          ){
            id
          }
        }`,
        variables:{
          "hotelId": hotelid,
          "path":path
        }
    })
  }
  createHotelType(hotelId:number, name:string, path:string){
    return this.apollo.mutate({
      mutation: gql`
        mutation createhoteltype(
          $hotelId: Int,
          $name:String,
          $path:String
          
        ){
          createhoteltype(
            hotelId:$hotelId,
            name: $name,
            path:$path
            
          ){
            id
          }
        }`,
      variables:{
        "hotelId": hotelId,
        "name": name,
        "path": path
      }
    })
}

//Update

UpdateHotel(rating:number, desc:string, id:number, name:string,price:number){
  return this.apollo.mutate({
    mutation: gql`
    mutation updatehotel(
      $rating: Float,
      $description: String,
      $id: Int,
      $name:String,
      $price: Int
    ){
      updatehotel(
        rating: $rating
        description: $description
        id:$id,
        name: $name,
        price: $price 
      ){
        name
      }
    }`,
    variables:{
      "rating": rating,
      "description": desc,
      "id": id,
      "price": price,
      "name": name
    }
  })
}
DeleteHotel(hotelid:number){
  return this.apollo.mutate({
    mutation: gql`
      mutation deletehotel($hotelId: Int){
        deletehotel(
          id:$hotelId      
        ){
          id
        }
      }`,
    variables:{
      "hotelId" : hotelid
    }
  })
}

}