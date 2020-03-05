import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { InsertBlogComponent } from './insert-blog.component';

describe('InsertBlogComponent', () => {
  let component: InsertBlogComponent;
  let fixture: ComponentFixture<InsertBlogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ InsertBlogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(InsertBlogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
