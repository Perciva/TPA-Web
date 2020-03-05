import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PlaneSearchComponent } from './plane-search.component';

describe('PlaneSearchComponent', () => {
  let component: PlaneSearchComponent;
  let fixture: ComponentFixture<PlaneSearchComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PlaneSearchComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PlaneSearchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
