import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MainPlaneComponent } from './main-plane.component';

describe('MainPlaneComponent', () => {
  let component: MainPlaneComponent;
  let fixture: ComponentFixture<MainPlaneComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MainPlaneComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MainPlaneComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
