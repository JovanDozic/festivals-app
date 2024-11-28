import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewEditTicketTypeComponent } from './view-edit-ticket-type.component';

describe('ViewEditTicketTypeComponent', () => {
  let component: ViewEditTicketTypeComponent;
  let fixture: ComponentFixture<ViewEditTicketTypeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ViewEditTicketTypeComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ViewEditTicketTypeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
