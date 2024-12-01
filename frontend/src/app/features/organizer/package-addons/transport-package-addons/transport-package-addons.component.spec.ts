import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TransportPackageAddonsComponent } from './transport-package-addons.component';

describe('TransportPackageAddonsComponent', () => {
  let component: TransportPackageAddonsComponent;
  let fixture: ComponentFixture<TransportPackageAddonsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TransportPackageAddonsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TransportPackageAddonsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
