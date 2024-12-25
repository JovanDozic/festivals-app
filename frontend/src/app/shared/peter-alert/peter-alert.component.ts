import { Component, inject } from '@angular/core';
import {
  MatDialogRef,
  MAT_DIALOG_DATA,
  MatDialogModule,
} from '@angular/material/dialog';
import { MatButtonModule } from '@angular/material/button';
import {
  MatDialogTitle,
  MatDialogContent,
  MatDialogActions,
} from '@angular/material/dialog';

export interface ConfirmationDialogData {
  title: string;
  message: string;
  confirmButtonText: string;
  cancelButtonText: string;
}

@Component({
  selector: 'app-peter-alert',
  templateUrl: './peter-alert.component.html',
  imports: [
    MatButtonModule,
    MatDialogTitle,
    MatDialogModule,
    MatDialogContent,
    MatDialogActions,
  ],
})
export class PeterAlertComponent {
  readonly dialogRef = inject(MatDialogRef<PeterAlertComponent>);
  readonly data = inject<ConfirmationDialogData>(MAT_DIALOG_DATA);

  closeDialog(confirm: boolean) {
    this.dialogRef.close({ confirm });
  }
}
