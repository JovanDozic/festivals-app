import { Component, inject, ViewChild } from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialog,
  MatDialogModule,
  MatDialogRef,
} from '@angular/material/dialog';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import {
  CreateItemPriceRequest,
  CreateItemRequest,
  IssueBraceletRequest,
} from '../../../../models/festival/festival.model';
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
import { SnackbarService } from '../../../../shared/snackbar/snackbar.service';
import { MatStepper, MatStepperModule } from '@angular/material/stepper';
import { ItemService } from '../../../../services/festival/item.service';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { OrderService } from '../../../../services/festival/order.service';
import { AddressResponse } from '../../../../models/common/address-response.model';
import { ConfirmationDialogComponent } from '../../../../shared/confirmation-dialog/confirmation-dialog.component';

@Component({
  selector: 'app-issue-bracelet',
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
    MatTabsModule,
    MatStepperModule,
    MatSlideToggleModule,
    MatDialogModule,
  ],
  templateUrl: './issue-bracelet.component.html',
  styleUrls: [
    './issue-bracelet.component.scss',
    '../../../../app.component.scss',
  ],
})
export class IssueBraceletComponent {
  private fb = inject(FormBuilder);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<IssueBraceletComponent>);
  private data: {
    festivalId: number;
    orderId: number;
    festivalTicketId: number;
    attendeeUsername: string;
  } = inject(MAT_DIALOG_DATA);
  private orderService = inject(OrderService);
  private dialog = inject(MatDialog);

  @ViewChild('stepper') private stepper: MatStepper | undefined;

  infoFormGroup: FormGroup;
  addressFormGroup: FormGroup;

  constructor() {
    this.infoFormGroup = this.fb.group({
      pinCtrl: [
        '',
        [
          Validators.required,
          Validators.minLength(4),
          Validators.maxLength(20),
          Validators.pattern('^[0-9]+$'),
        ],
      ],
      barcodeNumberCtrl: [
        '',
        [
          Validators.required,
          Validators.minLength(4),
          Validators.maxLength(20),
        ],
      ],
    });

    this.addressFormGroup = this.fb.group({
      streetCtrl: ['', Validators.required],
      numberCtrl: ['', Validators.required],
      apartmentSuiteCtrl: [''],
      cityCtrl: ['', Validators.required],
      postalCodeCtrl: ['', Validators.required],
      countryNiceNameCtrl: ['', Validators.required],
    });
  }

  closeDialog() {
    this.dialogRef.close(false);
  }

  done() {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        title: 'Are you sure?',
        message: `Are you sure you want to complete Bracelet issuing? Don't forget to print the shipping label.`,
        confirmButtonText: 'Yes',
        cancelButtonText: 'No',
      },
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.dialogRef.close(true);
      }
    });
  }

  issueBracelet() {
    if (
      this.infoFormGroup.valid &&
      this.data &&
      this.data.orderId &&
      this.data.festivalTicketId &&
      this.data.attendeeUsername
    ) {
      const request: IssueBraceletRequest = {
        orderId: this.data.orderId,
        pin: this.infoFormGroup.value.pinCtrl,
        barcodeNumber: this.infoFormGroup.value.barcodeNumberCtrl,
        festivalTicketId: this.data.festivalTicketId,
        attendeeUsername: this.data.attendeeUsername,
      };

      this.orderService.issueBracelet(request).subscribe({
        next: (response) => {
          this.stepper?.next();
          this.snackbarService.show('Bracelet created');
          console.log(response);
          console.log(response.shippingAddress);
          if (response.shippingAddress) {
            this.fillAddressForm(response.shippingAddress);
          }
        },
      });
    }
  }

  fillAddressForm(address: AddressResponse) {
    this.addressFormGroup.setValue({
      streetCtrl: address.street,
      numberCtrl: address.number,
      apartmentSuiteCtrl: address.apartmentSuite,
      cityCtrl: address.city,
      postalCodeCtrl: address.postalCode,
      countryNiceNameCtrl: address.niceName,
    });

    this.addressFormGroup.disable();
  }

  printShippingLabel() {
    if (this.data.orderId) {
      this.snackbarService.show('Printing shipping label...');
      this.orderService.printShippingLabel(this.data.orderId).subscribe({
        next: (blob) => {
          const url = URL.createObjectURL(blob);
          window.open(url, '_blank');
          this.snackbarService.show('Shipping label opened successfully');
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Error printing shipping label');
        },
      });
    }
  }
}
