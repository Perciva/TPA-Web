import { Router } from '@angular/router';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-car',
  templateUrl: './car.component.html',
  styleUrls: ['./car.component.scss']
})
export class CarComponent implements OnInit {
  private supir:boolean;
  constructor(private router: Router) { }

  ngOnInit() {
    this.supir=true 
  }
  search(){
    var bdate:Date = new Date((<HTMLInputElement>document.getElementById('berangkat')).value);
    var pdate:Date = new Date((<HTMLInputElement>document.getElementById('pulang')).value);
    var from:string = (<HTMLInputElement>document.getElementById('from')).value
    var jumlah:number = parseInt((<HTMLInputElement>document.getElementById('jumlah')).value)
  

    
   
    if(from == ""){
      alert("Mohon Isi kolom Kolom dari")
    }
    else if(new Date(bdate).toString() == "Invalid Date"|| new Date(pdate).toString() == "Invalid Date"){
      alert("Tanggal Check In dan Check Out Invalid");
    }
    else if(isNaN(jumlah) ){
      alert("Mohon isi Keterangan penumpang dan Jenis Ruangan");
    }
    else{
      this.router.navigate(["/Car/Search"], {
        queryParams:{
          'from': from,
          'berangkat': bdate.toString(),
          'pulang': pdate.toString(),
          'people': jumlah,
          'room': this.supir
        }
      })

    }
  }  

}
