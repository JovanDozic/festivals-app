import { Component, inject, ViewChild } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { MatStepper, MatStepperModule } from '@angular/material/stepper';
import { MatFormFieldModule } from '@angular/material/form-field';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { AuthService } from '../../../core/auth.service';
import { SnackbarService } from '../../../shared/snackbar/snackbar.service';
import { UserService } from '../../../services/user/user.service';
import { CreateUpdateUserProfileRequest } from '../../../models/user/user-profile-request.model';
import { CreateAddressRequest } from '../../../models/common/create-address-request.model';
import { MatCardModule } from '@angular/material/card';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { provideNativeDateAdapter } from '@angular/material/core';

@Component({
  selector: 'app-register-attendee',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatInputModule,
    MatStepperModule,
    MatFormFieldModule,
    MatButtonModule,
    MatCardModule,
    MatDatepickerModule,
  ],
  templateUrl: './register-attendee.component.html',
  styleUrls: [
    './register-attendee.component.scss',
    '../../../app.component.scss',
  ],
  providers: [provideNativeDateAdapter()],
})
export class RegisterAttendeeComponent {
  private fb = inject(FormBuilder);
  readonly authService = inject(AuthService);
  readonly userService = inject(UserService);
  readonly snackbarService = inject(SnackbarService);

  @ViewChild('stepper') private stepper: MatStepper | undefined;

  isLinear = true;
  accountFormGroup = this.fb.group({
    usernameCtrl: ['', Validators.required],
    emailCtrl: ['', [Validators.required, Validators.email]],
    passwordCtrl: ['', Validators.required],
  });
  personalFormGroup = this.fb.group({
    firstNameCtrl: ['', Validators.required],
    lastNameCtrl: ['', Validators.required],
    birthdayCtrl: ['', Validators.required],
    phoneCtrl: ['', Validators.required],
  });
  addressFormGroup = this.fb.group({
    streetCtrl: ['', Validators.required],
    numberCtrl: ['', Validators.required],
    apartmentSuiteCtrl: [''],
    cityCtrl: ['', Validators.required],
    postalCodeCtrl: ['', Validators.required],
    countryISO3Ctrl: ['', Validators.required],
  });

  constructor(private http: HttpClient) {}

  ngOnInit() {}

  createAccount() {
    if (this.accountFormGroup.valid) {
      const accountData = {
        username: this.accountFormGroup.get('usernameCtrl')?.value ?? '',
        email: this.accountFormGroup.get('emailCtrl')?.value ?? '',
        password: this.accountFormGroup.get('passwordCtrl')?.value ?? '',
      };
      this.authService.registerAttendee(accountData).subscribe({
        next: () => {
          this.snackbarService.show(
            'Account created successfully, logging in...'
          );

          this.authService
            .login(
              {
                username: accountData.username,
                password: accountData.password,
              },
              false
            )
            .subscribe({
              next: () => {
                setTimeout(() => {
                  this.snackbarService.show('Logged in successfully');
                  this.stepper?.next();
                }, 1000);
              },
              error: (error) => {
                console.error('Error logging in:', error);
                this.snackbarService.show('Error logging in');
              },
            });
        },
        error: (error) => {
          console.error('Error creating an account:', error);
          if (error.status === 409) {
            this.snackbarService.show('Username or email already in use');
          } else {
            this.snackbarService.show('Error creating an account');
          }
        },
      });
    }
  }

  createUserProfile() {
    if (this.personalFormGroup.valid) {
      const personalData: CreateUpdateUserProfileRequest = {
        firstName: this.personalFormGroup.get('firstNameCtrl')?.value ?? '',
        lastName: this.personalFormGroup.get('lastNameCtrl')?.value ?? '',
        dateOfBirth: new Date(
          this.personalFormGroup.get('birthdayCtrl')?.value ?? ''
        ),
        phoneNumber: this.personalFormGroup.get('phoneCtrl')?.value ?? '',
      };
      this.userService.createUserProfile(personalData).subscribe({
        next: () => {
          this.snackbarService.show('Personal information saved successfully');
          this.stepper?.next();
        },
        error: (error) => {
          console.error('Error saving personal information:', error);
          this.snackbarService.show('Error saving personal information');
        },
      });
    }
  }

  createAddress() {
    if (this.addressFormGroup.valid) {
      const addressData: CreateAddressRequest = {
        street: this.addressFormGroup.get('streetCtrl')?.value ?? '',
        number: this.addressFormGroup.get('numberCtrl')?.value ?? '',
        apartmentSuite: this.addressFormGroup.get('apartmentSuiteCtrl')?.value,
        city: this.addressFormGroup.get('cityCtrl')?.value ?? '',
        postalCode: this.addressFormGroup.get('postalCodeCtrl')?.value ?? '',
        countryISO3: this.addressFormGroup.get('countryISO3Ctrl')?.value ?? '',
      };
      this.userService.createAddress(addressData).subscribe({
        next: () => {
          this.snackbarService.show('Account created successfully');
          this.stepper?.next();
          setTimeout(() => {
            window.location.reload();
          }, 1000);
        },
        error: (error) => {
          console.error('Error saving address:', error);
          this.snackbarService.show('Error saving address');
        },
      });
    }
  }
}
