import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AllFestivalsComponent } from './all-festivals.component';

describe('AllFestivalsComponent', () => {
  let component: AllFestivalsComponent;
  let fixture: ComponentFixture<AllFestivalsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AllFestivalsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AllFestivalsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
