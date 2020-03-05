import { BlogData } from './../Interfaces/blog-interface';
import { FlightData } from './../Interfaces/flight-interface';
import { Hotel } from './../Models/hotel';
import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Observable } from 'rxjs';
import gql from 'graphql-tag';
import {Query} from '../Models/query'
import {User} from '../Models/user'
import { TrainData } from '../Interfaces/train-interface';
import { EventData } from '../Interfaces/event-interface';
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
        query{
          users{
            title
            lastname
            address
            firstname
            id
            kode_pos
            phone
            title
            email
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
  subscribe(email:string):Observable<Query>{
    return this.apollo.query<Query>({
      query:gql`
      query subscribe($mail:String){
        subscribe(email: $mail){
          email
        }
      }`,
      variables:{
        mail:email
      }
    })
  }
  selectUserById(id:number):Observable<Query>{
    return this.apollo.query<Query>({
      query: gql`
      query userbyid($id:Int){
        userbyid(id:$id){
           title
          lastname
          address
          firstname
          id
          kode_pos
          phone
          title
          languange
          email
        }
      }`,
      variables:{
        id:id
      }
    })
  }
  selectEventByCategory(category:string):Observable<Query>{
    return this.apollo.query<Query>({
      query: gql`
      query getevenbycategory($category:String){
        geteventbycategory(category:$category){
          content
          date
          id
          image
          latitude
          location
          longitude
          price
          title
          type
        }
      }`,
      variables:{
        category:category
      }
    })
  }
  selectEventById(id:number):Observable<Query>{
    return this.apollo.query<Query>({
      query: gql`
      query getevenbyid($id:Int){
        geteventbyid(id:$id){
          content
          date
          id
          image
          latitude
          location
          longitude
          price
          title
          type
        }
      }`,
      variables:{
        id:id
      }
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
  selectAllEvent():Observable<Query>{
    return this.apollo.query<Query>({
      query: gql`
      query {
        getallentertainment{
            type
            date
            id
            latitude
            location
            longitude
            image
            price
            title
            content
        }
    }`,
    fetchPolicy: 'no-cache'
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
  selectAllTrain():Observable<Query>{
    return this.apollo.query<Query>({
      query:gql`query{
        getalltrain {
            arrivalTime
            departureTime
            Id
            arrival {
                code
                name
            }
            class
            transit {
                code
                name
            }
            departure {
                code
                name
            }
            name
            code
            price
            seat
        }
    }`,
    fetchPolicy: 'no-cache'
    })
  }
  selectTrainName():Observable<Query>{
    return this.apollo.query<Query>({
      query: gql`
        query{
          gettrainnames{
            name
          }
        }`,
    })
  }
  selectTrainClass():Observable<Query>{
    return this.apollo.query<Query>({
      query: gql`
      query{
        gettrainclass{
          class
        }
      }`,
        
    })
  }
  getFilteredTrain(arrival:string, dest:string, date:string, names:string[],classes:string[] ){
    return this.apollo.query<Query>({
      query: gql`query getfilteredtrain(
        $arrival:String,
        $dept:String,
        $date:String,
        $class:[String],
        $names:[String]
      ){
        getfilteredtrain(
          arrival:$arrival,
          dept:$dept,
          date:$date,
          classes:$class,
          names:$names
        ){
          Id
          arrival{
              code
              id
              locationId
              name
          }
          transit {
              code
              id
              locationId
              name
          }
          departure {
              code
              id
              locationId
              name
          }
          arrivalTime
          departureTime
          name
          price
          seat
          code
          class
        }
      }`,
      variables:{
        arrival: arrival,
        dept:dest,
        date:date,
        names:names,
        class:classes
      }
    })
  }
  selectTrainByLocation(arr: string, dep: string, time: string) {
    return this.apollo.query<any>({
      query: gql`
      query gettrainbyloc(
        $arrival: String, 
        $departure:String, 
        $date: String
        ){
        gettrainbyloc(
          arrival: $arrival, 
          departure:$departure, 
          date:$date) {
            Id
            arrival{
                code
                id
                locationId
                name
            }
            transit {
                code
                id
                locationId
                name
            }
            departure {
                code
                id
                locationId
                name
            }
            arrivalTime
            departureTime
            name
            price
            seat
            code
            class
        }
    }`,
      variables: {
        arrival: arr,
        departure: dep,
        date: time,
      },
      fetchPolicy: 'no-cache',
    });
  }
  selectAllBlog():Observable<Query>{
    return this.apollo.query<Query>({
      query:gql`query{
        getallblog {
          category
          content
          id
          image
          title
          userId
          viewCount
        }
    }`,
    fetchPolicy: 'no-cache'
    })
  }
  selectAllCar():Observable<Query>{
    return this.apollo.query<Query>({
      query:gql`query{
        getallcar {
          id
          carModel{
            baggage
            passenger
            brand
            model
          }
        }
    }`,
    fetchPolicy: 'no-cache'
    })
  }
  selectAllFlight():Observable<Query>{
    return this.apollo.query<Query>({
      query:gql`query{
        getallflight {
          arrivalTime
          company {
              name
              image
          }
          companyId
          departureTime
          facility {
              name
              
          }
          fromAirport {
              code
              id
              location {
                  city
                  id
                  province
                  country
              }
              locationID
              name
          }
          fromAirportId
          Id
          model
          price
          transit {
              code
              id
              location {
                  city
                  id
                  province
                  country
              }
              locationID
              name
          }
          toAirport {
              code
              id
              location {
                  city
                  id
                  province
                  country
              }
              locationID
              name
          }
          toAirportId
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
  searchCarByCity(city:string):Observable<Query>{
    return this.apollo.query<Query>(
      {
        query:gql`
        query getcarbycity($city: String){
          getcarbycity(city: $city) {
              carModel {
                  id
                  baggage
                  image
                  brand
                  model
                  passenger
              }
              price
              id
              location {
                  city
                  id
                  province
                  country
              }
          }
      }`,
        variables:{
          city: city
        }
      }
    )
  }
  searchFilteredHotel(locations:number[], ratings:number[], facilities:string[], min:number, max:number):Observable<Query>{
    return this.apollo.query<Query>({
      query: gql`
      query getfilteredhotel(
        $loc:[Int],
        $rat: [Int],
        $fac: [String],
        $min: Int,
        $max: Int
      ){
        getfilteredhotel(
          locations: $loc,
          ratings: $rat,
          facilities: $fac,
          min: $min,
          max: $max
        ){
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
      variables:{
        loc: locations,
        rat:ratings,
        fac: facilities,
        min: min,
        max:max
      }
    })
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
  searchBlogById(id:number):Observable<Query>{
    return this.apollo.query<Query>(
      {
        query:gql`
        query getblogbyid($id: Int){
          getblogbyid(id: $id) {
              category
              content
              id
              image
              title
              userId
              viewCount
          }
      }`,
        variables:{
          id:id
        }
      }
    )
  }
  searchRecBlog(id:number):Observable<Query>{
    return this.apollo.query<Query>(
      {
        query:gql`
        query getrecblog($id: Int){
          getrecblog(id: $id) {
              category
              content
              id
              image
              title
              userId
              viewCount
          }
        }`,
        variables:{
          id:id
        }
      }
    )
  }
  

  searchHotelById(id:number):Observable<Query>{
    return this.apollo.query<Query>(
      {
        query:gql`
          query hotelbyid($id:Int){
            hotelbyid(id:$id){
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
          lat
          long
          photo{
            path
          }
          type{
            name
          }
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
              lat
              long
              location{
                city
                province
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
  getPromoById(id:number):Observable<Query>{
    return this.apollo.query<Query>(
      {
        query:gql`
          query getPromo($id: Int){
            getpromobyid(id: $id) {
              detail
              id
              image
              name
            }
          }`,
        variables:{
          id:id
        }
      }
    )
  }
  getOtherPromo(id:number):Observable<Query>{
    return this.apollo.query<Query>(
      {
        query:gql`
          query getPromo($id: Int){
            getotherpromo(id: $id) {
              detail
              id
              image
              name
            }
          }`,
        variables:{
          id:id
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
              id
    					name
              description
              price
              rating
    					photo {
    					  path
              }
              lat
              long
              location{
                city
                province
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
  CreateEvent(event: EventData) {
    // console.log(event)
    // return;
    return this.apollo.mutate<any>({
      mutation: gql`
      mutation(
        $category:String,
        $image:String,
        $title:String, 
        $price:Int, 
        $location: String, 
        $latitude: Float,
        $longitude: Float, 
        $date: String,
        $desc:String
        ) {
            createevent(
              category: $category, 
              image: $image, 
              title: $title,
              price: $price, 
              location: $location, 
              latitude: $latitude,
              longitude: $longitude, 
              date: $date,
              content:$desc
            ){
              id
            }
          }`,
      variables: {
        category: event.type,
        date: event.date,
        image: event.image,
        latitude: event.latitude,
        location: event.location,
        longitude: event.longitude,
        price: event.price,
        title: event.title,
        desc: event.description
      }
    });
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
  createFlight(flight: FlightData){
    return this.apollo.mutate({
      mutation:gql`
      mutation createflight(
        $company: String, 
        $model: String,
        $price: Int, 
        $fromAirport: String, 
        $toAirport: String,
        $transitAirport: String,
        $arrival: String, 
        $departure: String
        ){
          createflight(
              company: $company, 
              model: $model,
              price: $price, 
              fromAirport: $fromAirport,
              toAirport: $toAirport,
              transitAirport: $transitAirport,
              arrivalTime: $arrival, 
              departureTime: $departure
              ) {
                  Id
              }
          }`,
        variables: {
          company: flight.companyName,
          model: flight.model,
          price: flight.price,
          fromAirport: flight.from,
          toAirport: flight.to,
          transitAirport: flight.transit,
          arrival: flight.arriveDate,
          departure: flight.departDate,
        }
    })
  }
  createTrain(train: TrainData){
    const nameCode = train.nameCode.split(',', 2);

    const nameConverted = String(nameCode[0]);
    const codeConverted = String(nameCode[1]);
    const seatConverted = train.seat;
    const priceConverted = train.price;
    const arrivalConverted = String(train.arrivalName);
    const arrivalTimeConverted = String(train.arrivalTime);
    const transitConverted = String(train.transit);
    const departureConverted = String(train.departureName);
    const departureTimeConverted = String(train.departureTime);
    return this.apollo.mutate({
      mutation:gql`
      mutation(
        $name: String ,
        $code: String, 
        $seat: Int, 
        $price: Int,
        $arrival: String,
        $arrivalTime: String, 
        $transit: String, 
        $departure: String, 
        $departureTime: String,
        $class:String
        ){
            createtrain(
              name: $name, 
              code: $code, 
              seat: $seat, 
              price: $price, 
              arrival: $arrival,
              arrivalTime: $arrivalTime, 
              transit: $transit, 
              departure: $departure, 
              departureTime: $departureTime,
              class: $class
            ){
                Id
            }
        }`,
        variables: {
          name: nameConverted,
          code: codeConverted,
          seat: seatConverted,
          price: priceConverted,
          arrival: arrivalConverted,
          arrivalTime: arrivalTimeConverted,
          transit: transitConverted,
          departure: departureConverted,
          departureTime: departureTimeConverted,
          class : train.class
        }
    })
  }
  createBlog(blog: BlogData){
    return this.apollo.mutate({
      mutation:gql`
      mutation createblog(
        $user: Int, 
        $image: String,
        $category: String, 
        $title: String, 
        $content: String
        ){
          createblog(
              userId: $user, 
              image: $image,
              category:$category, 
              title: $title, 
              content: $content
              ){
                id
              }
          }`,
          variables: {
            user: blog.userID,
            image: blog.image,
            category: blog.category,
            title: blog.title,
            content: blog.content,
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
updateEvent(event: EventData) {
  return this.apollo.mutate<any>({
    mutation: gql`
    mutation updateEvent(
      $id:Int, 
      $category:String, 
      $image:String,
      $title:String, 
      $price:Int, 
      $location: String, 
      $latitude: Float,
      $longitude: Float, 
      $date: String,
      $desc: String
    ){
        updateevent(
          id: $id, 
          category: $category, 
          image: $image, 
          title: $title,
          price: $price, 
          location: $location, 
          latitude: $latitude,
          longitude: $longitude, 
          date: $date,
          content: $desc
        ){
            id
        }
      }`,
    variables: {
      id: event.id,
      category: event.type,
      date: event.date,
      image: event.image,
      latitude: event.latitude,
      location: event.location,
      longitude: event.longitude,
      price: event.price,
      title: event.title,
      desc : event.description
    }
  });
}
deleteEvent(id: number) {
  return this.apollo.mutate<any>({
    mutation: gql`
      mutation deleteEvent($id: Int!) {
        deleteevent(id: $id) {
          id
        }
      }`,
    variables: {
      id: id,
    }
  });
}
updateUser(address:string,email:string, fname:string, lname :string, id:number, kode_pos:string, phone:
  string,title:string, lang:string){
  return this.apollo.mutate<any>({
    mutation: gql`
    mutation updateUser(
      $firstname :String,
      $lastname:String,
      $email:String,
      $kode_pos:String,
      $title:String,
      $id:Int,
      $address: String,
      $phone:String,
      $languange:String
    ){
     updateuser(
      firstname : $firstname,
      lastname: $lastname,
      email:$email,
      kode_pos: $kode_pos,
      title :  $title,
      id : $id,
      address : $address,
      phone: $phone,
      languange : $languange
    ){
      id
    }
    }`,
    variables:{
      address: address,
      email: email,
      firstname: fname,
      id: id,
      kode_pos: kode_pos,
      lastname: lname,
      phone: phone,
      title: title,
      languange: lang
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
DeleteTrain(trainid:number){
  return this.apollo.mutate({
    mutation: gql`
      mutation deletetrain($trainId: Int){
        deletetrain(
          id:$trainId      
        ){
          Id
        }
      }`,
    variables:{
      "trainId" : trainid
    }
  })
}
DeleteBlog(blogid:number){
  return this.apollo.mutate({
    mutation: gql`
      mutation deleteblog($blogid: Int){
        deleteblog(
          id:$blogid      
        ){
          id
        }
      }`,
    variables:{
      "blogid" : blogid
    }
  })
}
DeleteFlight(flightid:number){
  return this.apollo.mutate({
    mutation: gql`
      mutation deleteflight($flightId: Int){
        deleteflight(
          id:$flightId      
        ){
          Id
        }
      }`,
    variables:{
      "flightId" : flightid
    }
  })
}

UpdateTrain(trainId:number, arrival:string, departure:string,price:number, seat:number){
  return this.apollo.mutate({
    mutation: gql`
    mutation(
      $trainId: Int, 
      $seat: Int, 
      $price: Int,
      $arrival: String, 
      $departure: String
      ){
        updatetrain(
            id: $trainId, 
            seat: $seat, 
            price: $price,
            arrivalTime: $arrival, 
            departureTime: $departure
            ){
              Id       
          }
      }`,
    variables:{
      "trainId": trainId,
      "seat": seat,
      "arrival": arrival,
      "price": price,
      "departure": departure
    }
  })
}
UpdateFlight(flightId:number, from:string,to:string, transit:string, arrival:string ,departure:string,model:string, price:number){
  return this.apollo.mutate({
    mutation: gql`
    mutation(
      $flightId: Int, 
      $from: String, 
      $price: Int,
      $arrival: String, 
      $departure: String,
      $to : String,
      $transit : String,
      $model:String
      ){
        updateflight(
            id: $flightId, 
            fromAirport: $from, 
            toAirport: $to,
            arrivalTime: $arrival, 
            departureTime: $departure,
            model: $model,
            price: $price,
            transitAirport: $transit
            ){
              Id       
          }
      }`,
    variables:{
      "flightId": flightId,
      "from": from,
      "arrival": arrival,
      "price": price,
      "departure": departure,
      "transit": transit,
      "model": model,
      "to": to
    }
  })
}
UpdateBlog(blogId:number, content:string, image:string, category:string){
  return this.apollo.mutate({
    mutation: gql`
    mutation(
      $blogId: Int, 
      $content: String, 
      $image: String,
      $category: String
      ){
        updateblog(
            id: $blogId, 
            content: $content, 
            image: $image,
            category: $category, 
            ){
              id       
          }
      }`,
    variables:{
      "blogId": blogId,
      "content":content,
      "image": image,
      "category": category
    }
  })
}


}
