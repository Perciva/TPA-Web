import { Hotel } from './../../Models/hotel';
import { ApolloService } from './../../Services/apollo.service';
import { Apollo } from 'apollo-angular';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.scss']
})
export class HomePageComponent implements OnInit {
  private popp:boolean = false;
  constructor(private apollo: ApolloService) { }

  ngOnInit() {
    // let hotel:Hotel = new Hotel()
    // hotel.address = "Manggis kecamatan karangasem"
    // hotel.city="Denpasar"
    // hotel.desc="Villa with 5 bedrooms and seaside view"
    // hotel.lat= -8.502620
    // hotel.long = 115.540718
    // hotel.name = "Villa Blanca"
    // hotel.price = 5000000
    // hotel.rating = 4.7

    // this.apollo.createHotel(hotel).subscribe(({data}) => {
    //   console.log(data)
    //   // this.ref.close()
    // })
  }
  wa(){
    window.open('https://api.whatsapp.com/send?phone=62895360356233')
  }
  pop() {
    scrollTo(<number><unknown>document.documentElement,200);
    document.getElementById('container').classList.add('pop');
    this.popp = true
  }
  unPop(){
    if(this.popp){
      document.getElementById('container').classList.remove('pop');
      this.popp = false
    }
  }

  scrollTo(element: HTMLElement, to: number) {
    const start = element.scrollTop;
    const change = to - start;
    let currentTime = 0;
    const increment = 20;

    const animateScroll = () => {
      currentTime += increment;
      const val = this.easeInOutQuad(currentTime, start, change, 1000000);
      element.scrollTop = val;
      if (currentTime < 1000000) {
        setTimeout(animateScroll, 10);
      }
    };
    animateScroll();
  }

  easeInOutQuad = (t, b, c, d) => {
    t /= d / 2;
    if (t < 1) {
      return c / 2 * t * t + b;
    }
    t--;
    return -c / 2 * (t * (t - 2) - 1) + b;
  }
  
}
