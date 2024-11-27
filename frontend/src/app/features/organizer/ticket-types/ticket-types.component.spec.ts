import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TicketTypesComponent } from './ticket-types.component';

describe('TicketTypesComponent', () => {
  let component: TicketTypesComponent;
  let fixture: ComponentFixture<TicketTypesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TicketTypesComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TicketTypesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
