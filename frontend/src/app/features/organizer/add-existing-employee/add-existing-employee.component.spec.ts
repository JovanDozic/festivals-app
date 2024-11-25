import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddExistingEmployeeComponent } from './add-existing-employee.component';

describe('AddExistingEmployeeComponent', () => {
  let component: AddExistingEmployeeComponent;
  let fixture: ComponentFixture<AddExistingEmployeeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AddExistingEmployeeComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddExistingEmployeeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
