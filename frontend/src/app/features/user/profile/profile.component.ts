import { Component, inject, OnInit } from '@angular/core';
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
import { UserService } from '../../../services/user/user.service';
import { UserProfileResponse } from '../../../models/user/user-profile-response.model';
import { MatDividerModule } from '@angular/material/divider';
import { MatCardModule } from '@angular/material/card';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss'],
  standalone: true,
  imports: [
    CommonModule,
    MatButtonModule,
    MatIconModule,
    MatDialogModule,
    MatDividerModule,
    MatCardModule,
  ],
})
export class ProfileComponent implements OnInit {
  private authService = inject(AuthService);
  private userService = inject(UserService);
  readonly dialog = inject(MatDialog);

  userProfile: UserProfileResponse | null = null;

  ngOnInit() {
    this.userService.getUserProfile().subscribe((response) => {
      this.userProfile = response;
      console.log('profile:', this.userProfile);
    });
  }

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
