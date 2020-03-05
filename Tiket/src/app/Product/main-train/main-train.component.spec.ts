import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MainTrainComponent } from './main-train.component';

describe('MainTrainComponent', () => {
  let component: MainTrainComponent;
  let fixture: ComponentFixture<MainTrainComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MainTrainComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MainTrainComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
