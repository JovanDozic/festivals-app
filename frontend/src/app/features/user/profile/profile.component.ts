import { Component, inject, OnInit } from '@angular/core';
import { AuthService } from '../../../core/auth.service';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { UserService } from '../../../services/user/user.service';
import { UserProfileResponse } from '../../../models/user/user-profile-response.model';
import { MatDividerModule } from '@angular/material/divider';
import { MatCardModule } from '@angular/material/card';
import { CommonModule } from '@angular/common';
import {
  ConfirmationDialog,
  ConfirmationDialogData,
} from '../../../shared/confirmation-dialog/confirmation-dialog.component';

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
    });
  }

  logout() {
    const dialogRef = this.dialog.open(ConfirmationDialog, {
      data: {
        title: 'Logout',
        message: 'Are you sure you want to log out?',
        confirmButtonText: 'Confirm',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.authService.logout();
      }
    });
  }
}
