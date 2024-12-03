import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StorePaymentDialogComponent } from './store-payment-dialog.component';

describe('StorePaymentDialogComponent', () => {
  let component: StorePaymentDialogComponent;
  let fixture: ComponentFixture<StorePaymentDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [StorePaymentDialogComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(StorePaymentDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
