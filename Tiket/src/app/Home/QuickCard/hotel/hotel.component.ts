import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-hotel',
  templateUrl: './hotel.component.html',
  styleUrls: ['./hotel.component.scss']
})
export class HotelComponent implements OnInit {

  constructor(private router: Router) { }

  ngOnInit() {
  }

  search(){
    var bdate:Date = new Date((<HTMLInputElement>document.getElementById('berangkat')).value);
    var pdate:Date = new Date((<HTMLInputElement>document.getElementById('pulang')).value);
    var from:string = (<HTMLInputElement>document.getElementById('from')).value
    var temp: String[] = ((<HTMLInputElement>document.getElementById('pclass')).value).split(' ')
    if(temp.length !=2){
      alert("Mohon isi keterangan jumlah penumpang dan kelas kabin")
      return;
    }
    var people:number = parseInt(temp[0].toString());
    var room:string = temp[1].toString();

    
   
    if(from == ""){
      alert("Mohon Isi kolom Kolom dari")
    }
    else if(new Date(bdate).toString() == "Invalid Date"|| new Date(pdate).toString() == "Invalid Date"){
      alert("Tanggal Check In dan Check Out Invalid");
    }
    else if(isNaN(people) || room == ""){
      alert("Mohon isi Keterangan penumpang dan Jenis Ruangan");
    }
    else{
      this.router.navigate(["/Hotel/Search"], {
        queryParams:{
          'from': from,
          'berangkat': bdate.toString(),
          'pulang': pdate.toString(),
          'people': people,
          'room': room
        }
      })

    }
  }  
  
  

}
