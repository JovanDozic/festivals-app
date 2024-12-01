import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateCampPackageAddonComponent } from './create-camp-package-addon.component';

describe('CreateCampPackageAddonComponent', () => {
  let component: CreateCampPackageAddonComponent;
  let fixture: ComponentFixture<CreateCampPackageAddonComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CreateCampPackageAddonComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateCampPackageAddonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
