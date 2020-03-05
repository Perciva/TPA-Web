import { ApolloService } from 'src/app/Services/apollo.service';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-promo',
  templateUrl: './promo.component.html',
  styleUrls: ['./promo.component.scss']
})
export class PromoComponent implements OnInit {
  private id:number
  private promo:any
  private rec:any[]=[]
  constructor(private activatedRoute: ActivatedRoute,
    private apollo: ApolloService) {
    this.activatedRoute.queryParams.subscribe(async p => {
    await this.getParam(p);
});
}

  ngOnInit() {
    this.apollo.getPromoById(this.id).subscribe(async a=>{
      await this.getPromo(a.data.getpromobyid)
    })
    this.apollo.getOtherPromo(this.id).subscribe(async a=>{
      await this.getRec(a.data.getotherpromo)
    })
  }
  getPromo(a:any){
    this.promo = a

  }
  getRec(a:any){
    this.rec = a
    console.log(a)
  }
  getParam(a:any){
    this.id =a.id
  }
  fb() {
    window.open('http://www.facebook.com/sharer.php?u=localhost:4200/Promo?id=' + this.id, 'facebookShare', 'width=626,height=436');
  }

  wa() {
    window.open('https://api.whatsapp.com/send?text=localhost:4200/Promo?id=' + this.id)
  }

  copy() {
    navigator.clipboard.writeText('localhost:4200/Promo?id=' + this.id);
  }

  line() {
    window.open('http://line.me/R/msg/text/?Promo?id=' + this.id);
  }

}
