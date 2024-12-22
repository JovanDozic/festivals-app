import { Component, inject, ViewChild } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { MatStepper, MatStepperModule } from '@angular/material/stepper';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { FestivalService } from '../../../services/festival/festival.service';
import { CreateFestivalRequest } from '../../../models/festival/festival.model';
import { forkJoin } from 'rxjs';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatIconModule } from '@angular/material/icon';
import { MatDialog } from '@angular/material/dialog';
import {
  ConfirmationDialogComponent,
  ConfirmationDialogData,
} from '../../../shared/confirmation-dialog/confirmation-dialog.component';
import { Router } from '@angular/router';
import { SnackbarService } from '../../../services/snackbar/snackbar.service';
import { ImageService } from '../../../services/image/image.service';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { CountryPickerComponent } from '../../../shared/country-picker/country-picker.component';
import {
  MAT_DATE_LOCALE,
  provideNativeDateAdapter,
} from '@angular/material/core';

interface ImagePreview {
  file: File;
  previewUrl: string | ArrayBuffer | null;
}

@Component({
  selector: 'app-create-festival',
  templateUrl: './create-festival.component.html',
  styleUrls: ['./create-festival.component.scss'],
  imports: [
    CountryPickerComponent,
    CommonModule,
    ReactiveFormsModule,
    MatInputModule,
    MatStepperModule,
    MatFormFieldModule,
    MatButtonModule,
    MatCardModule,
    MatDatepickerModule,
    MatGridListModule,
    MatIconModule,
    MatProgressSpinnerModule,
  ],
  providers: [
    provideNativeDateAdapter(),
    { provide: MAT_DATE_LOCALE, useValue: 'en-GB' },
  ],
})
export class CreateFestivalComponent {
  private fb = inject(FormBuilder);
  private http = inject(HttpClient);
  private snackbarService = inject(SnackbarService);
  private festivalService = inject(FestivalService);
  private dialog = inject(MatDialog);
  private imageService = inject(ImageService);
  router: Router = inject(Router);

  @ViewChild('stepper') private stepper: MatStepper | undefined;

  festivalId: string | null = null;
  images: ImagePreview[] = [];
  isUploading = false;

  basicInfoFormGroup = this.fb.group({
    nameCtrl: ['', Validators.required],
    descriptionCtrl: ['', Validators.required],
    startDateCtrl: ['', Validators.required],
    endDateCtrl: ['', Validators.required],
    capacityCtrl: ['', [Validators.required, Validators.min(1)]],
  });

  addressFormGroup = this.fb.group({
    streetCtrl: ['', Validators.required],
    numberCtrl: ['', Validators.required],
    apartmentSuiteCtrl: [''],
    cityCtrl: ['', Validators.required],
    postalCodeCtrl: ['', Validators.required],
    countryISO3Ctrl: ['', [Validators.required, Validators.maxLength(3)]],
  });

  goBack() {
    this.router.navigate(['organizer/my-festivals']);
  }

  createFestivalBasicInfo() {
    if (this.basicInfoFormGroup.valid) {
      this.stepper?.next();
    }
  }

  addFestivalAddress() {
    if (this.addressFormGroup.valid) {
      const festival: CreateFestivalRequest = {
        name: this.basicInfoFormGroup.get('nameCtrl')?.value ?? '',
        description:
          this.basicInfoFormGroup.get('descriptionCtrl')?.value ?? '',
        startDate: formatDate(
          new Date(this.basicInfoFormGroup.get('startDateCtrl')?.value ?? ''),
        ),
        endDate: formatDate(
          new Date(this.basicInfoFormGroup.get('endDateCtrl')?.value ?? ''),
        ),
        capacity: Number(this.basicInfoFormGroup.get('capacityCtrl')?.value),
        address: {
          street: this.addressFormGroup.get('streetCtrl')?.value ?? '',
          number: this.addressFormGroup.get('numberCtrl')?.value ?? '',
          apartmentSuite:
            this.addressFormGroup.get('apartmentSuiteCtrl')?.value ?? '',
          city: this.addressFormGroup.get('cityCtrl')?.value ?? '',
          postalCode: this.addressFormGroup.get('postalCodeCtrl')?.value ?? '',
          countryISO3:
            this.addressFormGroup.get('countryISO3Ctrl')?.value ?? '',
        },
      };

      this.festivalService.createFestival(festival).subscribe({
        next: (response) => {
          this.festivalId = response.id?.toString() ?? null;
          this.snackbarService.show('Basic info saved successfully!');
          this.stepper?.next();
        },
        error: (error) => {
          console.error('Error creating festival:', error);
          this.snackbarService.show('Error creating festival');
        },
      });
    }
  }

  onFileSelected(event: Event) {
    const fileInput = event.target as HTMLInputElement;

    if (fileInput.files && fileInput.files.length > 0) {
      Array.from(fileInput.files).forEach((file) => {
        const reader = new FileReader();

        reader.onload = () => {
          this.images.push({
            file: file,
            previewUrl: reader.result,
          });
        };

        reader.readAsDataURL(file);
      });
    }
  }

  removeImage(image: ImagePreview) {
    const dialogRef = this.dialog.open(ConfirmationDialogComponent, {
      data: {
        title: 'Remove Image',
        message: `Are you sure?`,
        confirmButtonText: 'Remove',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.images = this.images.filter((img) => img !== image);
      }
    });
  }

  uploadFestivalImages() {
    if (this.images.length === 0) {
      this.snackbarService.show('Add at least one image');
      return;
    }
    if (this.festivalId && this.images.length > 0) {
      this.isUploading = true;

      const uploadObservables = this.images.map((image) =>
        this.imageService.uploadImageAndGetURL(image.file),
      );

      forkJoin(uploadObservables).subscribe({
        next: (responses) => {
          const imageUrls = responses.map((response) => response.imageURL);

          const addImageObservables = imageUrls.map((imageUrl) =>
            this.http.post(
              `http://localhost:4000/festival/${this.festivalId}/image`,
              { imageUrl },
            ),
          );

          forkJoin(addImageObservables).subscribe({
            next: () => {
              this.isUploading = false;

              this.snackbarService.show('Festival created successfully!');
              this.stepper?.next();
              this.router.navigate([
                'organizer/my-festivals/' + this.festivalId,
              ]);
            },
            error: (error) => {
              this.isUploading = false;

              console.error('Error adding images to festival:', error);
              this.snackbarService.show('Error adding images to festival');
            },
          });
        },
        error: (error) => {
          this.isUploading = false;
          console.error('Error uploading images:', error);
          this.snackbarService.show('Error uploading images');
        },
      });
    }
  }
}

function formatDate(date: Date): string {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
}
