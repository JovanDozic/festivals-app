import { Component, inject } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import {
  MatDialogTitle,
  MatDialogContent,
  MatDialogActions,
} from '@angular/material/dialog';
import { AuthService } from '../../../core/auth.service';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { SnackbarService } from '../../../shared/snackbar/snackbar.service';

@Component({
  selector: 'change-password-dialog',
  templateUrl: './change-password-dialog.component.html',
  styleUrls: ['./change-password-dialog.component.scss'],
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatButtonModule,
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions,
    MatInputModule,
  ],
})
export class ChangePasswordDialogComponent {
  private authService = inject(AuthService);
  readonly dialogRef = inject(MatDialogRef<ChangePasswordDialogComponent>);
  readonly formBuilder = inject(FormBuilder);
  readonly snackbarService = inject(SnackbarService);

  changePasswordForm: FormGroup = this.formBuilder.group({
    oldPassword: ['', Validators.required],
    newPassword: ['', [Validators.required, Validators.minLength(5)]],
    confirmPassword: ['', Validators.required],
  });

  constructor() {
    this.changePasswordForm
      .get('confirmPassword')
      ?.setValidators([
        Validators.required,
        this.passwordMatchValidator.bind(this),
      ]);
  }

  passwordMatchValidator(control: any) {
    return control.value === this.changePasswordForm.get('newPassword')?.value
      ? null
      : { mismatch: true };
  }

  changePassword() {
    if (this.changePasswordForm.valid) {
      const { oldPassword, newPassword } = this.changePasswordForm.value;
      this.authService.changePassword(oldPassword, newPassword).subscribe({
        next: () => {
          this.dialogRef.close(true);
          this.snackbarService.show('Password changed successfully');
        },
        error: (error) => {
          console.error('Error changing password:', error);
          if (error.status === 401) {
            this.snackbarService.show('Old password is not correct');
          } else {
            this.snackbarService.show('Error changing password');
          }
        },
      });
    }
  }

  closeDialog() {
    this.dialogRef.close(false);
  }
}
