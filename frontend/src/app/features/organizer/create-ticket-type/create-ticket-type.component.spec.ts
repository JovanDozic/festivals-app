import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateTicketTypeComponent } from './create-ticket-type.component';

describe('CreateTicketTypeComponent', () => {
  let component: CreateTicketTypeComponent;
  let fixture: ComponentFixture<CreateTicketTypeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CreateTicketTypeComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateTicketTypeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
