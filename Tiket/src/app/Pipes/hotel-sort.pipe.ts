import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'hotelSort'
})
export class HotelSortPipe implements PipeTransform {

  PriceHighLow(hot){
    hot.sort((a:any, b:any) =>{
      if(a.price < b.price)return 1
      else if (a.price > b.price) return -1
      else return 0
    })
    return hot
  }
  PriceLowHigh(hot){
    hot.sort((a:any, b:any) =>{
      if(a.price > b.price)return 1
      else if (a.price < b.price) return -1
      else return 0
    })
    return hot
  }
  RatingHighLow(hot){
    hot.sort((a:any, b:any) =>{
      if(a.rating < b.rating)return 1
      else if (a.rating > b.rating) return -1
      else return 0
    })
    return hot
  }
  RatingLowHigh(hot){
    hot.sort((a:any, b:any) =>{
      if(a.rating > b.rating)return 1
      else if (a.rating < b.rating) return -1
      else return 0
    })
    return hot
  }

  transform(hot: Object[], by:string): any {
    if(by== "Recommended"){
      return hot
    }
    else if(by == "Price Low"){
      return this.PriceLowHigh(hot)
    }
    else if(by == "Price High"){
      return this.PriceHighLow(hot)
    }
    else if(by == "Rating Low"){
      return this.RatingLowHigh(hot)
    }
    else if(by == "Rating High"){
      return this.RatingHighLow(hot)
    }
  }

}
