import { Component, inject, OnInit } from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialog,
  MatDialogActions,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle,
} from '@angular/material/dialog';
import { ReactiveFormsModule } from '@angular/forms';
import { MatStepperModule } from '@angular/material/stepper';

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

import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { SnackbarService } from '../../../../shared/snackbar/snackbar.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-store-payment-dialog',
  imports: [
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
  ],
  templateUrl: './store-payment-dialog.component.html',
  styleUrls: [
    './store-payment-dialog.component.scss',
    '../../../../app.component.scss',
  ],
})
export class StorePaymentDialogComponent implements OnInit {
  readonly dialogRef = inject(MatDialogRef<StorePaymentDialogComponent>);
  readonly data = inject(MAT_DIALOG_DATA);
  private router = inject(Router);
  private snackbarService = inject(SnackbarService);

  isLoading = true;

  ngOnInit(): void {
    setTimeout(() => {
      this.isLoading = false;

      this.snackbarService.show('Order created successfully');
      setTimeout(() => {
        // todo: adjust the route with order ID
        this.dialogRef.close();
        this.router.navigate([`/festivals/${this.data.festivalId}`]);
      }, 1000);
    }, 1000);
  }
}
