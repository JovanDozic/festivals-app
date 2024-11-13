import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegisterAttendeeComponent } from './register-attendee.component';

describe('RegisterAttendeeComponent', () => {
  let component: RegisterAttendeeComponent;
  let fixture: ComponentFixture<RegisterAttendeeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RegisterAttendeeComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RegisterAttendeeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
