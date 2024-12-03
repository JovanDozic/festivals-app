import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StoreTicketComponent } from './store-ticket.component';

describe('StoreTicketComponent', () => {
  let component: StoreTicketComponent;
  let fixture: ComponentFixture<StoreTicketComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [StoreTicketComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(StoreTicketComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
