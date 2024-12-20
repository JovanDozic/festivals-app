import { Component, inject, OnInit } from '@angular/core';
import { AuthService } from '../../../services/auth/auth.service';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { UserService } from '../../../services/user/user.service';
import { UserProfileResponse } from '../../../models/user/user-responses';
import { MatDividerModule } from '@angular/material/divider';
import { MatCardModule } from '@angular/material/card';
import { CommonModule } from '@angular/common';
import {
  ConfirmationDialogComponent,
  ConfirmationDialogData,
} from '../../../shared/confirmation-dialog/confirmation-dialog.component';
import { ChangePasswordDialogComponent } from '../change-password-dialog/change-password-dialog.component';
import { MatTooltipModule } from '@angular/material/tooltip';
import { ChangeProfilePhotoDialogComponent } from '../change-profile-photo-dialog/change-profile-photo-dialog.component';
import { EditProfileComponent } from '../edit-profile/edit-profile.component';
import { SnackbarService } from '../../../services/snackbar/snackbar.service';
import { CreateUpdateUserProfileRequest } from '../../../models/user/user-requests';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss', '../../../app.component.scss'],
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
  private snackbarService = inject(SnackbarService);
  readonly dialog = inject(MatDialog);

  userProfile: UserProfileResponse | null = null;

  ngOnInit() {
    this.getUserProfile();
  }

  getUserProfile() {
    this.userService.getUserProfile().subscribe({
      next: (response) => {
        this.userProfile = response;
      },
      error: (error) => {
        this.snackbarService.show(
          'User profile was not created completely, creating temporary profile...',
        );
        console.error(error);
        if (error.status === 404) {
          const request: CreateUpdateUserProfileRequest = {
            firstName: 'FIRST NAME',
            lastName: 'LAST NAME',
            dateOfBirth: new Date('2002-01-21'),
            phoneNumber: 'PHONE NUMBER',
          };
          this.userService.createUserProfile(request).subscribe({
            next: () => {
              this.getUserProfile();
            },
            error: (error) => {
              this.snackbarService.show(
                "Couldn't create temporary user profile",
              );
              console.error(error);
            },
          });
        }
      },
    });
  }

  logout() {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
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
    this.dialog.open(ChangePasswordDialogComponent);
  }

  changeProfile() {
    const dialogRef = this.dialog.open(EditProfileComponent, {
      data: this.userProfile,
      height: '550px',
      width: '500px',
      disableClose: true,
    });

    dialogRef.afterClosed().subscribe((success) => {
      if (success) {
        this.getUserProfile();
      }
    });
  }

  changeProfilePhoto() {
    const dialogRef = this.dialog.open(ChangeProfilePhotoDialogComponent, {
      data: {
        currentImageURL: this.userProfile?.imageURL,
      },
      width: '400px',
      height: '410px',
    });

    dialogRef.afterClosed().subscribe((success) => {
      if (success) {
        this.getUserProfile();
      }
    });
  }
}
