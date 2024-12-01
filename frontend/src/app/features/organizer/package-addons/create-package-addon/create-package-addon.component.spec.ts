import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreatePackageAddonComponent } from './create-package-addon.component';

describe('CreatePackageAddonComponent', () => {
  let component: CreatePackageAddonComponent;
  let fixture: ComponentFixture<CreatePackageAddonComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CreatePackageAddonComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreatePackageAddonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
