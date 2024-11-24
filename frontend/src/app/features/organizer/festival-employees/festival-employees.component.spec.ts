import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FestivalEmployeesComponent } from './festival-employees.component';

describe('FestivalEmployeesComponent', () => {
  let component: FestivalEmployeesComponent;
  let fixture: ComponentFixture<FestivalEmployeesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FestivalEmployeesComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FestivalEmployeesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
