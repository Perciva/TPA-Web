import { state } from '@angular/animations';
import { ApolloService } from 'src/app/Services/apollo.service';
import {
  Component,
  ChangeDetectionStrategy,
  ViewChild,
  TemplateRef,
  OnInit
} from '@angular/core';
import {
  startOfDay,
  endOfDay,
  subDays,
  addDays,
  endOfMonth,
  isSameDay,
  isSameMonth,
  addHours
} from 'date-fns';
import { Subject } from 'rxjs';
import {
  CalendarEvent,
  CalendarEventAction,
  CalendarEventTimesChangedEvent,
  CalendarView
} from 'angular-calendar';
import { DatePipe } from '@angular/common';
import { cpus } from 'os';

const colors: any = {
  red: {
    primary: '#ad2121',
    secondary: '#FAE3E3'
  },
  blue: {
    primary: '#1e90ff',
    secondary: '#D1E8FF'
  },
  yellow: {
    primary: '#e3bc08',
    secondary: '#FDF1BA'
  }
};

@Component({
  selector: 'app-calendar',
  templateUrl: './calendar.component.html',
  styleUrls: ['./calendar.component.scss'],
  providers : [DatePipe]
})
export class CalendarComponent  implements OnInit{
  private planes:any[]
  constructor(private apollo:ApolloService, private datePipe: DatePipe) {}

  ngOnInit(): void {
    this.apollo.selectAllFlight().subscribe(async a=>{
      await this.getFlight(a.data.getallflight)
    })
    
  }
  getFlight(a:any){
    this.planes = a
    for(let f of this.planes){
      let flag=true
      for(let e of this.events){
        if(this.datePipe.transform(e.start, "yyyy-MM-dd") == this.datePipe.transform(f.arrivalTime, 'yyyy-MM-dd')){
          if(parseInt(f.price) < parseInt(e.title)){
            e.title = f.price
          }
          flag=false
        }
      }
      if(flag){
        this.events.push({
          start: startOfDay(new Date(f.arrivalTime)),
          title: f.price,
          color: colors.red
        })
      }
    }
  }
 

  @ViewChild('modalContent', { static: true }) modalContent: TemplateRef<any>;

  view: CalendarView = CalendarView.Month;

  CalendarView = CalendarView;

  viewDate: Date = new Date();

  modalData: {
    action: string;
    event: CalendarEvent;
  };

  refresh: Subject<any> = new Subject();

  events: CalendarEvent[] = [
    {
      start: startOfDay(new Date()),
      title: 'An event with no end date',
      color: colors.yellow
    }
  ];

  activeDayIsOpen: boolean = true;

  

  dayClicked({ date, events }: { date: Date; events: CalendarEvent[] }): void {
    if (isSameMonth(date, this.viewDate)) {
      if (
        (isSameDay(this.viewDate, date) && this.activeDayIsOpen === true) ||
        events.length === 0
      ) {
        this.activeDayIsOpen = false;
      } else {
        this.activeDayIsOpen = true;
      }
      this.viewDate = date;
    }
  }

  eventTimesChanged({
    event,
    newStart,
    newEnd
  }: CalendarEventTimesChangedEvent): void {
    this.events = this.events.map(iEvent => {
      if (iEvent === event) {
        return {
          ...event,
          start: newStart,
          end: newEnd
        };
      }
      return iEvent;
    });
    this.handleEvent('Dropped or resized', event);
  }

  handleEvent(action: string, event: CalendarEvent): void {
    this.modalData = { event, action };
    // this.modal.open(this.modalContent, { size: 'lg' });
  }

  setView(view: CalendarView) {
    this.view = view;
  }

  closeOpenMonthViewDay() {
    this.activeDayIsOpen = false;
  }
}


