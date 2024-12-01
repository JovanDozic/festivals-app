import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ChangeProfilePhotoDialogComponent } from './change-profile-photo-dialog.component';

describe('ChangeProfilePhotoDialogComponent', () => {
  let component: ChangeProfilePhotoDialogComponent;
  let fixture: ComponentFixture<ChangeProfilePhotoDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ChangeProfilePhotoDialogComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ChangeProfilePhotoDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
