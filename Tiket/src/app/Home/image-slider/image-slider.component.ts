import { animate,trigger, state, style, transition} from '@angular/animations';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-image-slider',
  templateUrl: './image-slider.component.html',
  styleUrls: ['./image-slider.component.scss'],
  animations:[
    trigger('slider', [
      state('sliderStart', style({
        left: 0,
      })),
      state('sliderMid', style({
        left: '-150%',
      })),
      state('sliderEnd', style({
        left: '150%',
      })),
      transition('siderStart <=> sliderMid', animate('200ms')),
      transition('sliderStart <=> sliderEnd', animate('200ms'))
    ])
  ],
})
export class ImageSliderComponent implements OnInit {

  private idx:number = 0;
  private imageUrls:string[]= 
  [
    "../../../assets/Nav/image-slider/1.jpg",
    "../../../assets/Nav/image-slider/2.jpg",
    "../../../assets/Nav/image-slider/3.jpg",
    "../../../assets/Nav/image-slider/4.jpg",
    "../../../assets/Nav/image-slider/5.jpg",
    "../../../assets/Nav/image-slider/7.jpg",
    "../../../assets/Nav/image-slider/8.jpg",
    "../../../assets/Nav/image-slider/9.jpg",
    "../../../assets/Nav/image-slider/10.jpg",
  ]
  private curr:string;

  constructor() { }

  ngOnInit() {
    this.curr = 'sliderStart';
    this.idx = 0;
    this.slide();
  }
  checkIndex(){
    if(this.idx > 8){
      this.idx = 0;
    }
    else if(this.idx<0 ){
      this.idx = 8;
    }
  }
  // slideNext():void{
  //   setInterval( () => {
  //     this.curr = "sliderMid"

  //     setTimeout( () => {
  //       this.curr="sliderEnd";
  //       this.idx++;
  //       this.checkIndex();
  //     },1000);

  //     setTimeout(()=> {
  //       this.curr = "sliderStart";
  //     }, 1000);
  //   })
  // }
  // slidePrev(){
  //   setInterval( () => {
  //     this.curr = "sliderEnd"

  //     setTimeout( () => {
  //       this.curr="sliderMid";
  //       this.idx++;
  //       this.checkIndex();
  //     },1000);

  //     setTimeout(()=> {
  //       this.curr = "sliderStart";
  //     }, 1000);
  //   })
  // }
  slide(): void {
    setInterval(() => {

      this.curr = 'sliderMid';

      setTimeout(() => {
        this.curr = 'sliderEnd';
        this.idx++;
        this.checkIndex();
      }, 800);

      setTimeout(() => {
        this.curr = 'sliderStart';
      }, 800);
    }, 9000);
  }
}
