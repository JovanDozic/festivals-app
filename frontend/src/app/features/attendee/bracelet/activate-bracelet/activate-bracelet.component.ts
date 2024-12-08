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
import { SnackbarService } from '../../../../shared/snackbar/snackbar.service';
import { MatStepperModule } from '@angular/material/stepper';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { OrderService } from '../../../../services/festival/order.service';
import { ActivateHelpRequestComponent } from '../activate-help-request/activate-help-request.component';

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
  private fb = inject(FormBuilder);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<ActivateBraceletComponent>);
  private data: {
    order: OrderDTO;
  } = inject(MAT_DIALOG_DATA);
  private orderService = inject(OrderService);
  private dialog = inject(MatDialog);

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

  activateBracelet() {
    if (
      this.infoFormGroup.valid &&
      this.data.order &&
      this.data.order.bracelet
    ) {
      const request: ActivateBraceletRequest = {
        braceletId: this.data.order.bracelet.braceletId,
        pin: this.infoFormGroup.get('pinCtrl')?.value,
      };

      this.orderService.activateBracelet(request).subscribe({
        next: () => {
          this.snackbarService.show('Bracelet activated successfully');
          this.dialogRef.close(true);
        },
        error: (error) => {
          console.log(error);
          if (error.status === 403) {
            this.snackbarService.show('Invalid PIN. Please try again.');
            return;
          }
          this.snackbarService.show('Failed to activate bracelet');
        },
      });
    }
  }

  requestHelp() {
    const dialogRef = this.dialog.open(ActivateHelpRequestComponent, {
      data: { order: this.data.order },
      width: '800px',
      height: '565px',
      disableClose: true,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.dialogRef.close(false);
      }
    });
  }
}
