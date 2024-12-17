import { Component, inject } from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialogModule,
  MatDialogRef,
} from '@angular/material/dialog';
import { FormsModule } from '@angular/forms';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatRadioModule } from '@angular/material/radio';
import { Log } from '../../../../models/common/log.model';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-view-log',
  imports: [
    CommonModule,
    FormsModule,
    MatInputModule,
    MatButtonModule,
    MatCardModule,
    MatIconModule,
    MatDialogModule,
    MatRadioModule,
  ],
  templateUrl: './view-log.component.html',
  styleUrls: ['./view-log.component.scss', '../../../../app.component.scss'],
})
export class ViewLogComponent {
  private dialogRef = inject(MatDialogRef<ViewLogComponent>);

  private data: {
    log: Log;
  } = inject(MAT_DIALOG_DATA);

  log: Log | null = null;

  constructor() {
    console.log(this.data);
    this.log = this.data.log;
  }

  closeDialog() {
    this.dialogRef.close(false);
  }
}
