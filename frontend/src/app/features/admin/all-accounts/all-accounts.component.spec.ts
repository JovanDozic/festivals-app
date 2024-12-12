import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AllAccountsComponent } from './all-accounts.component';

describe('AllAccountsComponent', () => {
  let component: AllAccountsComponent;
  let fixture: ComponentFixture<AllAccountsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AllAccountsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AllAccountsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
