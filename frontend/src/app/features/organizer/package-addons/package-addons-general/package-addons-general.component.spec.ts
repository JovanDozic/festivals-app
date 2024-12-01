import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PackageAddonsGeneralComponent } from './package-addons-general.component';

describe('PackageAddonsGeneralComponent', () => {
  let component: PackageAddonsGeneralComponent;
  let fixture: ComponentFixture<PackageAddonsGeneralComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PackageAddonsGeneralComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PackageAddonsGeneralComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
