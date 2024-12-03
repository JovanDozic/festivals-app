import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StorePackageComponent } from './store-package.component';

describe('StorePackageComponent', () => {
  let component: StorePackageComponent;
  let fixture: ComponentFixture<StorePackageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [StorePackageComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(StorePackageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
