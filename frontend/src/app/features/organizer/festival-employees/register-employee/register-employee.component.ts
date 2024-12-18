import { Component, inject } from '@angular/core';
import {
  MAT_DIALOG_DATA,
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
import { FestivalService } from '../../../../services/festival/festival.service';
import { CreateStaffRequest } from '../../../../models/festival/festival.model';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatIconModule } from '@angular/material/icon';
import { MatTabsModule } from '@angular/material/tabs';
import { UserService } from '../../../../services/user/user.service';
import { CreateProfileRequest } from '../../../../models/user/user-requests';
import { SnackbarService } from '../../../../services/snackbar/snackbar.service';
import {
  DateAdapter,
  MAT_DATE_FORMATS,
  MAT_DATE_LOCALE,
} from '@angular/material/core';
import { CustomDateAdapter } from '../../../../shared/date-formats/date-adapter';
import { CUSTOM_DATE_FORMATS } from '../../../../shared/date-formats/date-formats';

@Component({
  selector: 'app-register-employee',
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
  templateUrl: './register-employee.component.html',
  styleUrls: [
    './register-employee.component.scss',
    '../../../../app.component.scss',
  ],
  providers: [
    { provide: DateAdapter, useClass: CustomDateAdapter },
    { provide: MAT_DATE_FORMATS, useValue: CUSTOM_DATE_FORMATS },
    { provide: MAT_DATE_LOCALE, useValue: 'en-GB' },
  ],
})
export class RegisterEmployeeComponent {
  private fb = inject(FormBuilder);
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<RegisterEmployeeComponent>);
  private data: { festivalId: number } = inject(MAT_DIALOG_DATA);
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

      this.userService.registerEmployee(request).subscribe({
        next: () => {
          this.snackbarService.show('Employee registered');
          this.dialogRef.close(true);
        },
        error: (error) => {
          console.error('Error registering employee:', error);
          this.snackbarService.show('Failed to register employee');
        },
      });
    }
  }

  registerToFestival() {
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

      this.userService.registerEmployee(request).subscribe({
        next: (employee) => {
          this.festivalService
            .employEmployee(this.data.festivalId, employee.userId)
            .subscribe(() => {
              this.snackbarService.show('Employee registered');
              this.dialogRef.close(true);
            });
        },
        error: (error) => {
          console.error('Error registering employee:', error);
          this.snackbarService.show('Failed to register employee');
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
