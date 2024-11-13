import { Component, inject } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { MatStepperModule } from '@angular/material/stepper';
import { MatFormFieldModule } from '@angular/material/form-field';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';

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
  ],
  templateUrl: './register-attendee.component.html',
  styleUrl: './register-attendee.component.scss',
})
export class RegisterAttendeeComponent {
  private fb = inject(FormBuilder);

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
    addressCtrl: ['', Validators.required],
    cityCtrl: ['', Validators.required],
    zipCodeCtrl: ['', Validators.required],
    countryCtrl: ['', Validators.required],
  });

  constructor(private http: HttpClient) {}

  ngOnInit() {}

  createAccount() {
    if (this.accountFormGroup.valid) {
      const accountData = {
        username: this.accountFormGroup.get('usernameCtrl')?.value,
        email: this.accountFormGroup.get('emailCtrl')?.value,
        password: this.accountFormGroup.get('passwordCtrl')?.value,
      };
      // Call backend to create user
      this.http.post('/api/register', accountData).subscribe(
        () => {
          // On success, automatically log in
          if (accountData.username && accountData.password) {
            this.login(accountData.username, accountData.password);
          }
        },
        (error) => {
          // Handle error (e.g., display error message)
          console.error('Account creation failed', error);
        }
      );
    }
  }

  login(username: string, password: string) {
    const loginData = { username, password };
    this.http.post('/api/login', loginData).subscribe(
      (response: any) => {
        // Store the authentication token if necessary
        localStorage.setItem('authToken', response.token);
      },
      (error) => {
        // Handle error
        console.error('Login failed', error);
      }
    );
  }

  savePersonalInfo() {
    if (this.personalFormGroup.valid) {
      const personalData = {
        firstName: this.personalFormGroup.get('firstNameCtrl')?.value,
        lastName: this.personalFormGroup.get('lastNameCtrl')?.value,
        birthday: this.personalFormGroup.get('birthdayCtrl')?.value,
        phone: this.personalFormGroup.get('phoneCtrl')?.value,
      };
      // Call backend to save personal information
      this.http.post('/api/personal-info', personalData).subscribe(
        () => {
          // Personal info saved successfully
        },
        (error) => {
          // Handle error
          console.error('Saving personal info failed', error);
        }
      );
    }
  }

  submitRegistration() {
    if (this.addressFormGroup.valid) {
      const addressData = {
        address: this.addressFormGroup.get('addressCtrl')?.value,
        city: this.addressFormGroup.get('cityCtrl')?.value,
        zipCode: this.addressFormGroup.get('zipCodeCtrl')?.value,
        country: this.addressFormGroup.get('countryCtrl')?.value,
      };
      // Call backend to save address information
      this.http.post('/api/address-info', addressData).subscribe(
        () => {
          // Registration complete
          alert('Registration successful!');
        },
        (error) => {
          // Handle error
          console.error('Saving address info failed', error);
        }
      );
    }
  }
}
