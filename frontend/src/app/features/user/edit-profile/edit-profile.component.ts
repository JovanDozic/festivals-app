import { Component, inject } from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialogActions,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle,
} from '@angular/material/dialog';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatStepperModule } from '@angular/material/stepper';
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
  MAT_DATE_LOCALE,
  provideNativeDateAdapter,
} from '@angular/material/core';
import { MatTabsModule } from '@angular/material/tabs';
import { SnackbarService } from '../../../services/snackbar/snackbar.service';
import { forkJoin } from 'rxjs';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { NgxMatSelectSearchModule } from 'ngx-mat-select-search';
import * as countries from 'i18n-iso-countries';
import enLocale from 'i18n-iso-countries/langs/en.json';
import { MatSelectModule } from '@angular/material/select';
import { CountryPickerComponent } from '../../../shared/country-picker/country-picker.component';
import { UserService } from '../../../services/user/user.service';
import { CreateUpdateUserProfileRequest } from '../../../models/user/user-requests';
import { UpdateAddressRequest } from '../../../models/common/create-address-request.model';
import { UserProfileResponse } from '../../../models/user/user-responses';

@Component({
  selector: 'app-edit-profile',
  imports: [
    CountryPickerComponent,
    CommonModule,
    ReactiveFormsModule,
    MatInputModule,
    MatStepperModule,
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
    MatProgressSpinnerModule,
    NgxMatSelectSearchModule,
    MatOptionModule,
    MatSelectModule,
  ],
  templateUrl: './edit-profile.component.html',
  styleUrls: ['./edit-profile.component.scss', '../../../app.component.scss'],
  providers: [
    provideNativeDateAdapter(),
    { provide: MAT_DATE_LOCALE, useValue: 'en-GB' },
  ],
})
export class EditProfileComponent {
  private fb = inject(FormBuilder);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<EditProfileComponent>);
  private data: UserProfileResponse = inject(MAT_DIALOG_DATA);
  private userService = inject(UserService);

  isLinear = true;

  infoFormGroup: FormGroup;
  addressFormGroup: FormGroup;

  constructor() {
    countries.registerLocale(enLocale);

    this.infoFormGroup = this.fb.group({
      usernameCtrl: [this.data.username, Validators.required],
      emailCtrl: [this.data.email, Validators.required],
      firstNameCtrl: [this.data.firstName, Validators.required],
      lastNameCtrl: [this.data.lastName, Validators.required],
      dateOfBirthCtrl: [new Date(this.data.dateOfBirth), Validators.required],
      phoneNumberCtrl: [this.data.phoneNumber, Validators.required],
    });

    this.addressFormGroup = this.fb.group({
      streetCtrl: [this.data.address?.street, Validators.required],
      numberCtrl: [this.data.address?.number, Validators.required],
      apartmentSuiteCtrl: [this.data.address?.apartmentSuite],
      cityCtrl: [this.data.address?.city, Validators.required],
      postalCodeCtrl: [this.data.address?.postalCode, Validators.required],
      countryISO3Ctrl: [
        this.data.address?.countryISO3,
        [Validators.required, Validators.maxLength(3)],
      ],
    });

    this.infoFormGroup.get('usernameCtrl')?.disable();
  }

  closeDialog() {
    this.dialogRef.close(false);
  }

  saveChanges() {
    if (this.infoFormGroup.invalid || this.addressFormGroup.invalid) {
      this.snackbarService.show('Please complete all required fields.');
      return;
    }

    const requestProfile: CreateUpdateUserProfileRequest = {
      firstName: this.infoFormGroup.get('firstNameCtrl')?.value,
      lastName: this.infoFormGroup.get('lastNameCtrl')?.value,
      dateOfBirth: this.infoFormGroup.get('dateOfBirthCtrl')?.value,
      phoneNumber: this.infoFormGroup.get('phoneNumberCtrl')?.value,
    };

    const emailRequest = this.infoFormGroup.get('emailCtrl')?.value;

    const requestAddress: UpdateAddressRequest = {
      street: this.addressFormGroup.get('streetCtrl')?.value,
      number: this.addressFormGroup.get('numberCtrl')?.value,
      apartmentSuite: this.addressFormGroup.get('apartmentSuiteCtrl')?.value,
      city: this.addressFormGroup.get('cityCtrl')?.value,
      postalCode: this.addressFormGroup.get('postalCodeCtrl')?.value,
      countryISO3: this.addressFormGroup.get('countryISO3Ctrl')?.value,
    };

    let addressCall = null;

    if (this.data.address !== null) {
      addressCall = this.userService.updateUserAddress(requestAddress);
    } else {
      addressCall = this.userService.createAddress(requestAddress);
    }

    forkJoin({
      profile: this.userService.updateUserProfile(requestProfile),
      address: addressCall,
      email: this.userService.updateUserEmail(emailRequest),
    }).subscribe({
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
