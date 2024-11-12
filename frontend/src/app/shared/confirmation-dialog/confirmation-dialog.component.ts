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
  selector: 'confirmation-dialog',
  templateUrl: './confirmation-dialog.component.html',
  standalone: true,
  imports: [
    MatButtonModule,
    MatDialogTitle,
    MatDialogModule,
    MatDialogContent,
    MatDialogActions,
  ],
})
export class ConfirmationDialog {
  readonly dialogRef = inject(MatDialogRef<ConfirmationDialog>);
  readonly data = inject<ConfirmationDialogData>(MAT_DIALOG_DATA);

  closeDialog(confirm: boolean) {
    this.dialogRef.close({ confirm });
  }
}
