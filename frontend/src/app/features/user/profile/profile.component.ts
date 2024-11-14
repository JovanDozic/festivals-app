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
import { ChangePasswordDialogComponent } from '../change-password-dialog/change-password-dialog.component';
import { MatTooltipModule } from '@angular/material/tooltip';
import { ChangeProfileDialogComponent } from '../change-profile-dialog/change-profile-dialog.component';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss', '../../../app.component.scss'],
  standalone: true,
  imports: [
    CommonModule,
    MatButtonModule,
    MatTooltipModule,
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
    this.getUserProfile();
  }

  getUserProfile() {
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

  changePassword() {
    const dialogRef = this.dialog.open(ChangePasswordDialogComponent);

    dialogRef.afterClosed().subscribe((success) => {
      if (success) {
        console.log('Password changed successfully');
      }
    });
  }

  changeProfile() {
    const dialogRef = this.dialog.open(ChangeProfileDialogComponent, {
      data: {
        firstName: this.userProfile?.firstName,
        lastName: this.userProfile?.lastName,
        dateOfBirth: this.userProfile?.dateOfBirth,
        phoneNumber: this.userProfile?.phoneNumber,
      },
    });

    dialogRef.afterClosed().subscribe((success) => {
      if (success) {
        this.getUserProfile();
        console.log('Profile changed successfully');
      }
    });
  }
}
