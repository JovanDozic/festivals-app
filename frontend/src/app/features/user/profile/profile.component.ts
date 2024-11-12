import { Component, inject } from '@angular/core';
import { AuthService } from '../../../core/auth.service';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import {
  MAT_DIALOG_DATA,
  MatDialog,
  MatDialogActions,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle,
  MatDialogModule,
} from '@angular/material/dialog';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrl: './profile.component.scss',
  standalone: true,
  imports: [MatButtonModule, MatIconModule, MatDialogModule],
})
export class ProfileComponent {
  private authService = inject(AuthService);
  readonly dialog = inject(MatDialog);

  logout() {
    let dialogRef = this.dialog.open(LogoutDialog, {
      data: { confirm: false },
    });

    dialogRef.afterClosed().subscribe((result) => {
      console.log('The dialog was closed', result);
      if (result !== undefined && result.confirm) {
        this.authService.logout();
      }
    });
  }
}

export interface DialogData {
  confirm: boolean;
}

@Component({
  selector: 'logout-dialog',
  templateUrl: 'logout-dialog.html',
  standalone: true,
  imports: [
    MatButtonModule,
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions,
  ],
})
export class LogoutDialog {
  readonly dialogRef = inject(MatDialogRef<LogoutDialog>);
  readonly data = inject<DialogData>(MAT_DIALOG_DATA);

  closeDialog(confirm: boolean) {
    this.dialogRef.close({ confirm });
  }
}
