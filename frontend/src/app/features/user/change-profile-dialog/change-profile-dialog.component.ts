import { Component, inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatDatepickerModule } from '@angular/material/datepicker';
import {
  MatDialogTitle,
  MatDialogContent,
  MatDialogActions,
} from '@angular/material/dialog';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { UserService } from '../../../services/user/user.service';
import { provideNativeDateAdapter } from '@angular/material/core';
import { SnackbarService } from '../../../shared/snackbar/snackbar.service';

@Component({
  selector: 'app-change-profile-dialog',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatButtonModule,
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions,
    MatInputModule,
    MatDatepickerModule,
  ],
  templateUrl: './change-profile-dialog.component.html',
  styleUrls: [
    './change-profile-dialog.component.scss',
    '../../../app.component.scss',
  ],
  providers: [provideNativeDateAdapter()],
})
export class ChangeProfileDialogComponent implements OnInit {
  private userService = inject(UserService);
  readonly dialogRef = inject(MatDialogRef<ChangeProfileDialogComponent>);
  readonly formBuilder = inject(FormBuilder);
  readonly data = inject<any>(MAT_DIALOG_DATA);
  private snackbarService = inject(SnackbarService);

  changeProfileForm: FormGroup = this.formBuilder.group({
    firstName: ['', Validators.required],
    lastName: ['', Validators.required],
    dateOfBirth: [null, Validators.required],
    phoneNumber: [''],
  });

  ngOnInit() {
    if (this.data) {
      const { firstName, lastName, dateOfBirth, phoneNumber } = this.data;
      this.changeProfileForm.patchValue({
        firstName,
        lastName,
        dateOfBirth: dateOfBirth ? new Date(dateOfBirth) : null,
        phoneNumber,
      });
    }
  }

  changeProfile() {
    if (this.changeProfileForm.valid) {
      this.userService
        .updateUserProfile(this.changeProfileForm.value)
        .subscribe({
          next: () => {
            this.dialogRef.close(true);
            this.snackbarService.show('Profile updated successfully');
          },
          error: (error) => {
            console.error(error);
            this.snackbarService.show('Error updating profile');
          },
        });
    }
  }

  closeDialog() {
    this.dialogRef.close(false);
  }
}
