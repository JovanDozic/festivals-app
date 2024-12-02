import { Component, inject, OnInit } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatDatepickerModule } from '@angular/material/datepicker';
import {
  MatDialogTitle,
  MatDialogContent,
  MatDialogActions,
} from '@angular/material/dialog';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { UserService } from '../../../services/user/user.service';
import { SnackbarService } from '../../../shared/snackbar/snackbar.service';
import { ImageService } from '../../../services/image/image.service';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';

@Component({
  selector: 'app-change-profile-photo-dialog',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatButtonModule,
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions,
    MatInputModule,
    MatDatepickerModule,
    MatIconModule,
    MatProgressSpinnerModule,
  ],
  templateUrl: './change-profile-photo-dialog.component.html',
  styleUrls: [
    './change-profile-photo-dialog.component.scss',
    '../../../app.component.scss',
  ],
})
export class ChangeProfilePhotoDialogComponent implements OnInit {
  readonly dialogRef = inject(MatDialogRef<ChangeProfilePhotoDialogComponent>);
  readonly data = inject<any>(MAT_DIALOG_DATA);

  private imageService = inject(ImageService);
  private userService = inject(UserService);
  private snackbarService = inject(SnackbarService);

  selectedFile: File | null = null;
  imagePreviewUrl: string | ArrayBuffer | null = null;
  isUploading = false;

  ngOnInit() {
    this.imagePreviewUrl = this.data.currentImageURL;
  }

  saveChanges() {
    if (this.selectedFile) {
      this.isUploading = true;
      this.imageService.uploadImageAndGetURL(this.selectedFile).subscribe({
        next: (response) => {
          console.log(response);
          this.snackbarService.show('Profile photo uploaded successfully!');
          this.userService.updateUserProfilePhoto(response.imageURL).subscribe({
            next: () => {
              this.isUploading = false;
              this.snackbarService.show('Profile photo updated successfully!');
              this.dialogRef.close(true);
            },
            error: (error) => {
              console.log(error);
              this.isUploading = false;
              this.snackbarService.show('Failed to update profile photo.');
            },
          });
        },
        error: (error) => {
          console.log(error);
          this.isUploading = false;
          this.snackbarService.show('Failed to update profile photo.');
        },
      });
    }
  }

  closeDialog() {
    this.dialogRef.close(false);
  }

  onFileSelected(event: Event) {
    const fileInput = event.target as HTMLInputElement;

    if (fileInput.files && fileInput.files.length > 0) {
      this.selectedFile = fileInput.files[0];

      const reader = new FileReader();

      reader.onload = () => {
        this.imagePreviewUrl = reader.result;
      };

      reader.readAsDataURL(this.selectedFile);
    }
  }
}
