import { Component, inject, OnInit } from '@angular/core';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { CommonModule } from '@angular/common';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { SnackbarService } from '../../../../services/snackbar/snackbar.service';

@Component({
  selector: 'app-top-up-payment-dialog',
  imports: [
    CommonModule,
    MatDialogModule,
    MatIconModule,
    MatProgressSpinnerModule,
  ],
  templateUrl: './top-up-payment-dialog.component.html',
  styleUrls: [
    './top-up-payment-dialog.component.scss',
    '../../../../app.component.scss',
  ],
})
export class TopUpPaymentDialogComponent implements OnInit {
  readonly dialogRef = inject(MatDialogRef<TopUpPaymentDialogComponent>);
  private snackbarService = inject(SnackbarService);

  isLoading = true;

  ngOnInit(): void {
    setTimeout(() => {
      this.isLoading = false;

      this.snackbarService.show('Payment successful');
      setTimeout(() => {
        this.dialogRef.close();
      }, 1000);
    }, 1000);
  }
}
