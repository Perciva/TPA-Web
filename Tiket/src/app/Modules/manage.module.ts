import { UpdateEventComponent } from './../Manage/manage-event/update-event/update-event.component';
import { ManageEventComponent } from './../Manage/manage-event/manage-event.component';
import { InsertEventComponent } from './../Manage/manage-event/insert-event/insert-event.component';
import { TextEditorComponent } from './../Manage/manage-blog/text-editor/text-editor.component';
import { InsertBlogComponent } from './../Manage/manage-blog/insert-blog/insert-blog.component';
import { ManageBlogComponent } from './../Manage/manage-blog/manage-blog.component';
import { InsertFlightComponent } from './../Manage/manage-flight/insert-flight/insert-flight.component';
import { ManageFlightComponent } from './../Manage/manage-flight/manage-flight.component';
import { ManageTrainComponent } from './../Manage/manage-train/manage-train.component';
import { ConfirmDialogComponent } from './../Manage/confirm-dialog/confirm-dialog.component';
import { InsertHotelComponent } from './../Manage/manage-hotel/insert-hotel/insert-hotel.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { ManageHotelComponent } from './../Manage/manage-hotel/manage-hotel.component';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatTableModule, MatButtonModule, MatInputModule, MatDividerModule, MatChipsModule, MatIconModule, MatProgressSpinnerModule, MatPaginatorModule, MatSortModule, MatDatepickerModule, MatSelectModule } from '@angular/material';
import { InsertTrainComponent } from '../Manage/manage-train/insert-train/insert-train.component';
import {NgxMaterialTimepickerModule} from 'ngx-material-timepicker';

const decs:any = [
  ManageHotelComponent,
  InsertHotelComponent,
  ConfirmDialogComponent,
  ManageTrainComponent,
  ManageFlightComponent,
  InsertFlightComponent,
  InsertTrainComponent,
  ManageBlogComponent,
  InsertBlogComponent,
  TextEditorComponent,
  InsertEventComponent,
  ManageEventComponent,
  UpdateEventComponent,
]
const imps:any =[
  MatTableModule,
  MatButtonModule,
  MatInputModule,
  ReactiveFormsModule,
  FormsModule,
  MatDividerModule,
  MatChipsModule,
  MatIconModule,
  MatProgressSpinnerModule,
  MatPaginatorModule,
  MatSortModule,
  NgxMaterialTimepickerModule,
  MatDatepickerModule,
  MatSelectModule,

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
    InsertFlightComponent,
    InsertTrainComponent,
    ConfirmDialogComponent,
    InsertBlogComponent,
    InsertEventComponent,
    UpdateEventComponent
  ]

})
export class ManageModule { 
}
