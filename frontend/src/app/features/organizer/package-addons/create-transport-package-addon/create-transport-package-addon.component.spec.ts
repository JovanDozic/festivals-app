import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateTransportPackageAddonComponent } from './create-transport-package-addon.component';

describe('CreateTransportPackageAddonComponent', () => {
  let component: CreateTransportPackageAddonComponent;
  let fixture: ComponentFixture<CreateTransportPackageAddonComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CreateTransportPackageAddonComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateTransportPackageAddonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
