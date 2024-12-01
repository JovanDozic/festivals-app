import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreatePackageAddonChooserComponent } from './create-package-addon-chooser.component';

describe('CreatePackageAddonChooserComponent', () => {
  let component: CreatePackageAddonChooserComponent;
  let fixture: ComponentFixture<CreatePackageAddonChooserComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CreatePackageAddonChooserComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreatePackageAddonChooserComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
