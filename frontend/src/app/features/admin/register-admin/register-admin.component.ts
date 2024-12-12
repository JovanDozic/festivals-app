import { Component, inject } from '@angular/core';
import {
  MatDialogActions,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle,
} from '@angular/material/dialog';
import {
  AbstractControl,
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  ValidationErrors,
  Validators,
} from '@angular/forms';
import { CreateStaffRequest } from '../../../models/festival/festival.model';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatIconModule } from '@angular/material/icon';
import { provideNativeDateAdapter } from '@angular/material/core';
import { MatTabsModule } from '@angular/material/tabs';
import { UserService } from '../../../services/user/user.service';
import { CreateProfileRequest } from '../../../models/user/user-requests';
import { SnackbarService } from '../../../shared/snackbar/snackbar.service';

@Component({
  selector: 'app-register-admin',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatInputModule,
    MatFormFieldModule,
    MatButtonModule,
    MatCardModule,
    MatDatepickerModule,
    MatGridListModule,
    MatIconModule,
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions,
    MatTabsModule,
  ],
  templateUrl: './register-admin.component.html',
  styleUrls: ['./register-admin.component.scss', '../../../app.component.scss'],
  providers: [provideNativeDateAdapter()],
})
export class RegisterAdminComponent {
  private fb = inject(FormBuilder);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<RegisterAdminComponent>);
  private userService = inject(UserService);

  infoFormGroup: FormGroup;

  constructor() {
    this.infoFormGroup = this.fb.group({
      usernameCtrl: ['', Validators.required],
      emailCtrl: ['', Validators.required],
      passwordCtrl: ['', [Validators.required, Validators.minLength(5)]],
      confirmPasswordCtrl: ['', Validators.required],
      firstNameCtrl: ['', Validators.required],
      lastNameCtrl: ['', Validators.required],
      phoneNumberCtrl: ['', Validators.required],
      dateOfBirthCtrl: ['', Validators.required],
    });

    this.infoFormGroup
      .get('confirmPasswordCtrl')
      ?.setValidators([
        Validators.required,
        this.passwordMatchValidator.bind(this),
      ]);
  }

  passwordMatchValidator(control: AbstractControl): ValidationErrors | null {
    return control.value === this.infoFormGroup.get('passwordCtrl')?.value
      ? null
      : { mismatch: true };
  }

  register() {
    if (this.infoFormGroup.valid) {
      const profile: CreateProfileRequest = {
        firstName: this.infoFormGroup.get('firstNameCtrl')?.value,
        lastName: this.infoFormGroup.get('lastNameCtrl')?.value,
        phoneNumber: this.infoFormGroup.get('phoneNumberCtrl')?.value,
        dateOfBirth: this.formatDate(
          this.infoFormGroup.get('dateOfBirthCtrl')?.value,
        ),
      };
      const request: CreateStaffRequest = {
        username: this.infoFormGroup.get('usernameCtrl')?.value,
        password: this.infoFormGroup.get('passwordCtrl')?.value,
        email: this.infoFormGroup.get('emailCtrl')?.value,
        userProfile: profile,
      };

      this.userService.registerAdmin(request).subscribe({
        next: () => {
          this.snackbarService.show('Administrator registered');
          this.dialogRef.close(true);
        },
        error: (error) => {
          console.error('Error registering administrator:', error);
          this.snackbarService.show('Failed to register administrator');
        },
      });
    }
  }

  closeDialog() {
    this.dialogRef.close(false);
  }

  private formatDate(date: Date): string {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    return `${year}-${month}-${day}`;
  }
}
