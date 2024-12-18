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
  ActivateBraceletRequest,
  OrderDTO,
  TopUpBraceletRequest,
} from '../../../../models/festival/festival.model';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatIconModule } from '@angular/material/icon';
import { MatTabsModule } from '@angular/material/tabs';
import { SnackbarService } from '../../../../services/snackbar/snackbar.service';
import { MatStepper, MatStepperModule } from '@angular/material/stepper';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { OrderService } from '../../../../services/festival/order.service';
import { TopUpPaymentDialogComponent } from '../top-up-payment-dialog/top-up-payment-dialog.component';

@Component({
  selector: 'app-top-up-bracelet',
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
  templateUrl: './top-up-bracelet.component.html',
  styleUrls: [
    './top-up-bracelet.component.scss',
    '../../../../app.component.scss',
  ],
})
export class TopUpBraceletComponent {
  private fb = inject(FormBuilder);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<TopUpBraceletComponent>);
  private dialog = inject(MatDialog);
  private data: {
    order: OrderDTO;
  } = inject(MAT_DIALOG_DATA);
  private orderService = inject(OrderService);

  isLoading = false;
  infoFormGroup: FormGroup;

  constructor() {
    this.infoFormGroup = this.fb.group({
      amountCtrl: ['', [Validators.required]],
    });
  }

  closeDialog() {
    this.dialogRef.close(false);
  }

  openPaymentDialog() {
    this.isLoading = true;
    const dialogRef = this.dialog.open(TopUpPaymentDialogComponent, {
      width: '250px',
      height: '250px',
      disableClose: true,
    });

    dialogRef.afterClosed().subscribe(() => {
      this.topUp();
      this.isLoading = false;
    });
  }

  topUp() {
    if (
      this.infoFormGroup.valid &&
      this.data.order &&
      this.data.order.bracelet
    ) {
      const request: TopUpBraceletRequest = {
        braceletId: this.data.order.bracelet.braceletId,
        amount: this.infoFormGroup.get('amountCtrl')?.value,
      };
      this.orderService.topUpBracelet(request).subscribe({
        next: () => {
          this.snackbarService.show('Top up successful');
          this.dialogRef.close(true);
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Failed to top up bracelet');
        },
      });
    }
  }
}
