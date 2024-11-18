import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MyFestivalsComponent } from './my-festivals.component';

describe('MyFestivalsComponent', () => {
  let component: MyFestivalsComponent;
  let fixture: ComponentFixture<MyFestivalsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [MyFestivalsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MyFestivalsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
