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
  OrderDTO,
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
  selector: 'app-activate-bracelet',
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
  templateUrl: './activate-bracelet.component.html',
  styleUrls: [
    './activate-bracelet.component.scss',
    '../../../../app.component.scss',
  ],
})
export class ActivateBraceletComponent {
  requestHelp() {
    throw new Error('Method not implemented.');
  }
  private fb = inject(FormBuilder);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<ActivateBraceletComponent>);
  private data: {
    order: OrderDTO;
  } = inject(MAT_DIALOG_DATA);
  private orderService = inject(OrderService);
  private dialog = inject(MatDialog);

  @ViewChild('stepper') private stepper: MatStepper | undefined;

  infoFormGroup: FormGroup;

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
      barcodeNumberCtrl: this.data.order.bracelet?.barcodeNumber,
    });

    this.infoFormGroup.get('barcodeNumberCtrl')?.disable();
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
}
