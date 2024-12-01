import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateGeneralPackageAddonComponent } from './create-general-package-addon.component';

describe('CreateGeneralPackageAddonComponent', () => {
  let component: CreateGeneralPackageAddonComponent;
  let fixture: ComponentFixture<CreateGeneralPackageAddonComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CreateGeneralPackageAddonComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateGeneralPackageAddonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
