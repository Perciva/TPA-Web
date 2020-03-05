import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { InsertTrainComponent } from './insert-train.component';

describe('InsertTrainComponent', () => {
  let component: InsertTrainComponent;
  let fixture: ComponentFixture<InsertTrainComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ InsertTrainComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(InsertTrainComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
