import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PackageAddonsComponent } from './package-addons.component';

describe('PackageAddonsComponent', () => {
  let component: PackageAddonsComponent;
  let fixture: ComponentFixture<PackageAddonsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PackageAddonsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PackageAddonsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
