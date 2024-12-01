import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CampPackageAddonsComponent } from './camp-package-addons.component';

describe('CampPackageAddonsComponent', () => {
  let component: CampPackageAddonsComponent;
  let fixture: ComponentFixture<CampPackageAddonsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CampPackageAddonsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CampPackageAddonsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
