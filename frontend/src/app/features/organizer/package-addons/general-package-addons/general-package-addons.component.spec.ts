import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GeneralPackageAddonsComponent } from './general-package-addons.component';

describe('GeneralPackageAddonsComponent', () => {
  let component: GeneralPackageAddonsComponent;
  let fixture: ComponentFixture<GeneralPackageAddonsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GeneralPackageAddonsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GeneralPackageAddonsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
