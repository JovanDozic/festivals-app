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
  FormsModule,
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
import {
  MatOptionModule,
  provideNativeDateAdapter,
} from '@angular/material/core';
import { MatTabsModule } from '@angular/material/tabs';
import { UserService } from '../../../services/user/user.service';
import { CreateProfileRequest } from '../../../models/user/user-requests';
import { SnackbarService } from '../../../shared/snackbar/snackbar.service';
import { MatSelectModule } from '@angular/material/select';

export interface Role {
  value: string;
  viewValue: string;
}

@Component({
  selector: 'app-register-admin',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    FormsModule,
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
    MatSelectModule,
    MatOptionModule,
  ],
  templateUrl: './register-user.component.html',
  styleUrls: ['./register-user.component.scss', '../../../app.component.scss'],
  providers: [provideNativeDateAdapter()],
})
export class RegisterUserComponent {
  private fb = inject(FormBuilder);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<RegisterUserComponent>);
  private userService = inject(UserService);

  roles: Role[] = [
    { value: 'ORGANIZER', viewValue: 'Organizer' },
    { value: 'ADMINISTRATOR', viewValue: 'Administrator' },
  ];

  roleFormGroup: FormGroup;
  infoFormGroup: FormGroup;

  constructor() {
    this.roleFormGroup = this.fb.group({
      roleCtrl: ['', Validators.required],
    });

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
    if (this.roleFormGroup.valid) {
      if (this.roleFormGroup.get('roleCtrl')?.value === 'ORGANIZER') {
        this.registerOrganizer();
      } else if (
        this.roleFormGroup.get('roleCtrl')?.value === 'ADMINISTRATOR'
      ) {
        this.registerAdmin();
      }
    }
  }

  registerAdmin() {
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

  registerOrganizer() {
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

      this.userService.registerOrganizer(request).subscribe({
        next: () => {
          this.snackbarService.show('Organizer registered');
          this.dialogRef.close(true);
        },
        error: (error) => {
          console.error('Error registering organizer:', error);
          this.snackbarService.show('Failed to register organizer');
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
