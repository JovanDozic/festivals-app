import { Component, inject, OnInit } from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialogModule,
  MatDialogRef,
} from '@angular/material/dialog';
import { CommonModule } from '@angular/common';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { SnackbarService } from '../../../../shared/snackbar/snackbar.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-store-payment-dialog',
  imports: [
    CommonModule,
    MatDialogModule,
    MatIconModule,
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
