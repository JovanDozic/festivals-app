import { Component, inject, OnInit } from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialog,
  MatDialogActions,
  MatDialogContent,
  MatDialogRef,
  MatDialogTitle,
} from '@angular/material/dialog';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatStepperModule } from '@angular/material/stepper';
import { FestivalService } from '../../../services/festival/festival.service';
import {
  Festival,
  UpdateFestivalRequest,
} from '../../../models/festival/festival.model';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatIconModule } from '@angular/material/icon';
import {
  MatOptionModule,
  MAT_DATE_LOCALE,
  provideNativeDateAdapter,
} from '@angular/material/core';
import { MatTabsModule } from '@angular/material/tabs';
import { SnackbarService } from '../../../services/snackbar/snackbar.service';
import {
  ConfirmationDialogComponent,
  ConfirmationDialogData,
} from '../../../shared/confirmation-dialog/confirmation-dialog.component';
import { forkJoin, map, Observable, of } from 'rxjs';
import { ImageService } from '../../../services/image/image.service';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { NgxMatSelectSearchModule } from 'ngx-mat-select-search';
import * as countries from 'i18n-iso-countries';
import enLocale from 'i18n-iso-countries/langs/en.json';
import { MatSelectModule } from '@angular/material/select';
import { CountryPickerComponent } from '../../../shared/country-picker/country-picker.component';

interface ImagePreview {
  id?: number;
  file?: File;
  previewUrl: string | ArrayBuffer | null;
  isNew: boolean;
}

@Component({
  selector: 'app-edit-festival',
  templateUrl: './edit-festival.component.html',
  styleUrls: ['./edit-festival.component.scss', '../../../app.component.scss'],
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
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions,
    MatTabsModule,
    MatProgressSpinnerModule,
    NgxMatSelectSearchModule,
    MatOptionModule,
    MatSelectModule,
  ],
  providers: [
    provideNativeDateAdapter(),
    { provide: MAT_DATE_LOCALE, useValue: 'en-GB' },
  ],
})
export class EditFestivalComponent implements OnInit {
  private fb = inject(FormBuilder);
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<EditFestivalComponent>);
  private data: Festival = inject(MAT_DIALOG_DATA);
  private dialog = inject(MatDialog);
  private imageService = inject(ImageService);

  isLinear = true;

  basicInfoFormGroup: FormGroup;
  addressFormGroup: FormGroup;

  images: ImagePreview[] = [];
  imagesToDelete: number[] = [];
  isUploading = false;

  constructor() {
    countries.registerLocale(enLocale);

    this.basicInfoFormGroup = this.fb.group({
      nameCtrl: ['', Validators.required],
      descriptionCtrl: ['', Validators.required],
      startDateCtrl: ['', Validators.required],
      endDateCtrl: ['', Validators.required],
      capacityCtrl: ['', [Validators.required, Validators.min(1)]],
    });

    this.addressFormGroup = this.fb.group({
      streetCtrl: ['', Validators.required],
      numberCtrl: ['', Validators.required],
      apartmentSuiteCtrl: [''],
      cityCtrl: ['', Validators.required],
      postalCodeCtrl: ['', Validators.required],
      countryISO3Ctrl: [
        this.data.address?.countryISO3,
        [Validators.required, Validators.maxLength(3)],
      ],
    });
  }

  ngOnInit() {
    if (this.data) {
      this.basicInfoFormGroup.patchValue({
        nameCtrl: this.data.name,
        descriptionCtrl: this.data.description,
        startDateCtrl: new Date(this.data.startDate),
        endDateCtrl: new Date(this.data.endDate),
        capacityCtrl: this.data.capacity,
      });

      this.addressFormGroup.patchValue({
        streetCtrl: this.data.address?.street,
        numberCtrl: this.data.address?.number,
        apartmentSuiteCtrl: this.data.address?.apartmentSuite,
        cityCtrl: this.data.address?.city,
        postalCodeCtrl: this.data.address?.postalCode,
        countryISO3Ctrl: this.data.address?.countryISO3,
      });

      this.images = this.data.images.map((image) => ({
        id: image.id,
        previewUrl: image.url,
        isNew: false,
      }));
    }
  }

  saveChanges() {
    if (this.basicInfoFormGroup.invalid || this.addressFormGroup.invalid) {
      this.snackbarService.show('Please complete all required fields.');
      return;
    }

    this.isUploading = true;

    const updatedFestival: UpdateFestivalRequest = {
      id: this.data.id,
      name: this.basicInfoFormGroup.get('nameCtrl')?.value,
      description: this.basicInfoFormGroup.get('descriptionCtrl')?.value,
      startDate: this.formatDate(
        this.basicInfoFormGroup.get('startDateCtrl')?.value,
      ),
      endDate: this.formatDate(
        this.basicInfoFormGroup.get('endDateCtrl')?.value,
      ),
      capacity: Number(this.basicInfoFormGroup.get('capacityCtrl')?.value),
      address: {
        street: this.addressFormGroup.get('streetCtrl')?.value,
        number: this.addressFormGroup.get('numberCtrl')?.value,
        apartmentSuite: this.addressFormGroup.get('apartmentSuiteCtrl')?.value,
        city: this.addressFormGroup.get('cityCtrl')?.value,
        postalCode: this.addressFormGroup.get('postalCodeCtrl')?.value,
        countryISO3: this.addressFormGroup.get('countryISO3Ctrl')?.value,
      },
    };

    this.festivalService.updateFestival(updatedFestival).subscribe({
      next: () => {
        this.handleImages();
      },
      error: (error) => {
        this.isUploading = false;
        console.error('Error updating festival:', error);
        this.snackbarService.show('Error updating festival. Please try again.');
      },
    });
  }

  private handleImages() {
    const deleteObservables = this.imagesToDelete.map((imageId) =>
      this.festivalService.deleteFestivalImage(this.data.id, imageId),
    );

    const newImages = this.images.filter((image) => image.isNew && image.file);

    // delete images marked for deletion
    const deleteImages$ =
      deleteObservables.length > 0 ? forkJoin(deleteObservables) : of(null);

    (deleteImages$ as Observable<null>).subscribe({
      next: () => {
        // upload new ones
        const uploadObservables = newImages.map((image) =>
          this.imageService
            .uploadImageAndGetURL(image.file!)
            .pipe(map((response) => response.imageURL)),
        );

        const uploadImages$ =
          uploadObservables.length > 0 ? forkJoin(uploadObservables) : of([]);

        uploadImages$.subscribe({
          next: (imageUrls) => {
            // add new images to the festival
            const addImageObservables = (imageUrls as string[]).map(
              (imageUrl) =>
                this.festivalService.addFestivalImage(this.data.id, imageUrl),
            );

            const addImages$ =
              addImageObservables.length > 0
                ? forkJoin(addImageObservables)
                : of(null);

            (addImages$ as Observable<null>).subscribe({
              next: () => {
                this.isUploading = false;
                this.snackbarService.show('Festival updated successfully!');
                this.dialogRef.close(true);
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
      },
      error: (error) => {
        this.isUploading = false;
        console.error('Error deleting images:', error);
        this.snackbarService.show('Error deleting images');
      },
    });
  }

  closeDialog() {
    this.dialogRef.close(false);
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
            isNew: true,
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
        message: `Are you sure you want to remove this image?`,
        confirmButtonText: 'Remove',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        if (image.isNew) {
          // just remove from array
          this.images = this.images.filter((img) => img !== image);
        } else if (image.id) {
          // add to imagesToDelete array so it can be removed from the festival
          this.imagesToDelete.push(image.id);
          // remove from array
          this.images = this.images.filter((img) => img !== image);
        }
      }
    });
  }

  private formatDate(date: Date): string {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    return `${year}-${month}-${day}`;
  }
}
