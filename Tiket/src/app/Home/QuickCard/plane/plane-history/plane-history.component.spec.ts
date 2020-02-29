import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PlaneHistoryComponent } from './plane-history.component';

describe('PlaneHistoryComponent', () => {
  let component: PlaneHistoryComponent;
  let fixture: ComponentFixture<PlaneHistoryComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PlaneHistoryComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PlaneHistoryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
