import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { InsertHotelComponent } from './insert-hotel.component';

describe('InsertHotelComponent', () => {
  let component: InsertHotelComponent;
  let fixture: ComponentFixture<InsertHotelComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ InsertHotelComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(InsertHotelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
