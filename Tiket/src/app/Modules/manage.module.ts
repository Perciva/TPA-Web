import { ConfirmDialogComponent } from './../Manage/confirm-dialog/confirm-dialog.component';
import { InsertHotelComponent } from './../Manage/manage-hotel/insert-hotel/insert-hotel.component';
import { FormsModule } from '@angular/forms';

import { ManageHotelComponent } from './../Manage/manage-hotel/manage-hotel.component';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatTableModule, MatButtonModule, MatInputModule, MatDividerModule, MatChipsModule, MatIconModule, MatProgressSpinnerModule, MatPaginatorModule } from '@angular/material';


const decs:any = [
  ManageHotelComponent,
  InsertHotelComponent,
  ConfirmDialogComponent
]
const imps:any =[
  MatTableModule,
  MatButtonModule,
  MatInputModule,
  FormsModule,
  MatDividerModule,
  MatChipsModule,
  MatIconModule,
  MatProgressSpinnerModule,
  MatPaginatorModule
]

@NgModule({
  declarations: decs,
  imports: [
    CommonModule,
    imps
  ],
  exports : [
    decs,
    imps

  ],
  entryComponents:[
    InsertHotelComponent,
    ConfirmDialogComponent
  ]

})
export class ManageModule { 
}
