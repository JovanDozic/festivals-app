import { Component, inject, OnInit } from '@angular/core';
import {
  CreateTicketOrderRequest,
  Festival,
  ItemCurrentPrice,
} from '../../../../models/festival/festival.model';
import { ActivatedRoute, Router } from '@angular/router';
import { FestivalService } from '../../../../services/festival/festival.service';
import { SnackbarService } from '../../../../services/snackbar/snackbar.service';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatIconModule } from '@angular/material/icon';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { MatCardModule } from '@angular/material/card';
import { MatChipsModule } from '@angular/material/chips';
import { MatMenuModule } from '@angular/material/menu';
import { MatStepperModule } from '@angular/material/stepper';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { ItemService } from '../../../../services/festival/item.service';
import { UserService } from '../../../../services/user/user.service';
import { UserProfileResponse } from '../../../../models/user/user-responses';
import { AddressResponse } from '../../../../models/common/address-response.model';
import {
  ConfirmationDialogComponent,
  ConfirmationDialogData,
} from '../../../../shared/confirmation-dialog/confirmation-dialog.component';
import { firstValueFrom, Observable, throwError } from 'rxjs';
import { OrderService } from '../../../../services/festival/order.service';
import { StorePaymentDialogComponent } from '../store-payment-dialog/store-payment-dialog.component';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { CountryPickerComponent } from '../../../../shared/country-picker/country-picker.component';

@Component({
  selector: 'app-store-ticket',
  imports: [
    CountryPickerComponent,
    CommonModule,
    ReactiveFormsModule,
    MatButtonModule,
    MatTooltipModule,
    MatIconModule,
    MatDialogModule,
    MatCardModule,
    MatChipsModule,
    MatMenuModule,
    MatInputModule,
    MatFormFieldModule,
    MatStepperModule,
    MatProgressSpinnerModule,
  ],
  templateUrl: './store-ticket.component.html',
  styleUrls: [
    './store-ticket.component.scss',
    '../../../../app.component.scss',
  ],
})
export class StoreTicketComponent implements OnInit {
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private festivalService = inject(FestivalService);
  private itemService = inject(ItemService);
  private userService = inject(UserService);
  private snackbarService = inject(SnackbarService);
  private orderService = inject(OrderService);
  private dialog = inject(MatDialog);
  private fb = inject(FormBuilder);

  userProfile: UserProfileResponse | null = null;
  address: AddressResponse | null = null;
  festival: Festival | null = null;
  isLoading = false;

  personalFormGroup: FormGroup;
  addressFormGroup: FormGroup;
  tickets: ItemCurrentPrice[] = [];
  selectedTicket: ItemCurrentPrice | null = null;

  constructor() {
    this.personalFormGroup = this.fb.group({
      firstNameCtrl: ['', Validators.required],
      lastNameCtrl: ['', Validators.required],
      emailCtrl: ['', [Validators.required, Validators.email]],
      dateOfBirthCtrl: ['', Validators.required],
      phoneCtrl: ['', Validators.required],
    });

    this.addressFormGroup = this.fb.group({
      streetCtrl: ['', Validators.required],
      numberCtrl: ['', Validators.required],
      apartmentSuiteCtrl: [''],
      cityCtrl: ['', Validators.required],
      postalCodeCtrl: ['', Validators.required],
      countryISO3Ctrl: ['', [Validators.required, Validators.maxLength(3)]],
    });
  }

  ngOnInit() {
    this.loadFestival();
    this.loadTickets();
    this.loadUser();
  }

  goBack() {
    this.router.navigate([`festivals/${this.festival?.id}`]);
  }

  loadFestival() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.festivalService.getFestival(Number(id)).subscribe({
        next: (festival) => {
          this.festival = festival;
          this.isLoading = false;
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Error getting festival');
          this.festival = null;
          this.isLoading = true;
        },
      });
    }
  }

  loadTickets() {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.itemService.getTicketTypes(Number(id)).subscribe({
        next: (tickets) => {
          this.tickets = tickets;
          this.isLoading = false;
          if (this.tickets.length === 0) {
            this.snackbarService.show('Store has no tickets available yet');
            this.router.navigate([`festivals/${this.festival?.id}`]);
          }
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Error getting tickets');
          this.tickets = [];
          this.isLoading = true;
        },
      });
    }
  }

  loadUser() {
    this.userService.getUserProfile().subscribe({
      next: (userProfile) => {
        this.userProfile = userProfile;
        if (userProfile.address) this.address = userProfile.address;
        this.isLoading = false;
        this.fillPersonalFormGroup();
        this.fillAddressFormGroup();
      },
      error: (error) => {
        console.log(error);
        this.snackbarService.show('Error getting user information');
        this.userProfile = null;
        this.isLoading = true;
      },
    });
  }

  selectTicket(ticket: ItemCurrentPrice) {
    if (ticket.remainingNumber === 0) {
      return;
    }
    this.selectedTicket = ticket;
  }

  fillPersonalFormGroup() {
    if (this.userProfile) {
      this.personalFormGroup.setValue({
        firstNameCtrl: this.userProfile.firstName,
        lastNameCtrl: this.userProfile.lastName,
        emailCtrl: this.userProfile.email,
        dateOfBirthCtrl: this.userProfile.dateOfBirth,
        phoneCtrl: this.userProfile.phoneNumber,
      });
    }
  }

  fillAddressFormGroup() {
    if (this.address) {
      this.addressFormGroup.setValue({
        streetCtrl: this.address.street,
        numberCtrl: this.address.number,
        apartmentSuiteCtrl: this.address.apartmentSuite,
        cityCtrl: this.address.city,
        postalCodeCtrl: this.address.postalCode,
        countryISO3Ctrl: this.address.countryISO3,
      });
    }
  }

  saveFormToCurrent() {
    if (this.userProfile) {
      this.userProfile.firstName =
        this.personalFormGroup.get('firstNameCtrl')?.value;
      this.userProfile.lastName =
        this.personalFormGroup.get('lastNameCtrl')?.value;
      this.userProfile.email = this.personalFormGroup.get('emailCtrl')?.value;
      this.userProfile.phoneNumber =
        this.personalFormGroup.get('phoneCtrl')?.value;
      this.address = {
        street: this.addressFormGroup.get('streetCtrl')?.value,
        number: this.addressFormGroup.get('numberCtrl')?.value,
        apartmentSuite: this.addressFormGroup.get('apartmentSuiteCtrl')?.value,
        city: this.addressFormGroup.get('cityCtrl')?.value,
        postalCode: this.addressFormGroup.get('postalCodeCtrl')?.value,
        countryISO3: this.addressFormGroup.get('countryISO3Ctrl')?.value,
        country: '',
        countryISO2: '',
      };
    }
  }

  completeOrder() {
    if (this.isLoading) return;
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        title: 'Complete Order',
        message:
          'Please make sure all information is correct. Are you sure you want to complete the Order?',
        confirmButtonText: 'Confirm',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.sendOrder();
      }
    });
  }

  async sendOrder() {
    try {
      this.isLoading = true;
      await firstValueFrom(this.updateProfile());
      await firstValueFrom(this.updateEmail());
      await firstValueFrom(this.updateAddress());
      const orderResponse = await firstValueFrom(this.createOrder());

      const orderId = orderResponse.orderId;

      if (orderId) {
        this.openPaymentDialog(orderId);
      } else {
        throw new Error('Order ID is missing in the response');
      }
    } catch (error) {
      console.error('Error completing order:', error);
      this.snackbarService.show('Error completing order');
      this.isLoading = false;
    }
  }

  updateProfile(): Observable<void> {
    if (this.personalFormGroup.valid) {
      return this.userService.updateUserProfile({
        firstName: this.personalFormGroup.get('firstNameCtrl')?.value,
        lastName: this.personalFormGroup.get('lastNameCtrl')?.value,
        dateOfBirth: new Date(
          this.personalFormGroup.get('dateOfBirthCtrl')?.value,
        ),
        phoneNumber: this.personalFormGroup.get('phoneCtrl')?.value,
      });
    }
    return throwError(() => new Error('Personal form is invalid'));
  }

  updateEmail(): Observable<void> {
    if (this.personalFormGroup.valid) {
      return this.userService.updateUserEmail(
        this.personalFormGroup.get('emailCtrl')?.value,
      );
    }
    return throwError(() => new Error('Email form is invalid'));
  }

  updateAddress(): Observable<void> {
    if (this.addressFormGroup.valid) {
      return this.userService.updateUserAddress({
        street: this.addressFormGroup.get('streetCtrl')?.value,
        number: this.addressFormGroup.get('numberCtrl')?.value,
        apartmentSuite: this.addressFormGroup.get('apartmentSuiteCtrl')?.value,
        city: this.addressFormGroup.get('cityCtrl')?.value,
        postalCode: this.addressFormGroup.get('postalCodeCtrl')?.value,
        countryISO3: this.addressFormGroup.get('countryISO3Ctrl')?.value,
      });
    }
    return throwError(() => new Error('Address form is invalid'));
  }

  createOrder(): Observable<{
    orderId: number;
  }> {
    if (this.selectedTicket && this.festival && this.festival.id) {
      const request: CreateTicketOrderRequest = {
        ticketTypeId: this.selectedTicket.itemId,
        totalPrice: this.selectedTicket.price,
      };

      return this.orderService.createTicketOrder(this.festival.id, request);
    }
    return throwError(
      () =>
        new Error('No ticket selected or festival or festival ID is missing'),
    );
  }

  openPaymentDialog(orderId: number) {
    const dialogRef = this.dialog.open(StorePaymentDialogComponent, {
      data: { orderId: orderId },
      width: '250px',
      height: '250px',
      disableClose: true,
    });

    dialogRef.afterClosed().subscribe(() => {
      this.isLoading = false;
    });
  }
}
