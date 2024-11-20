import { Component, inject, ViewChild } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { MatStepper, MatStepperModule } from '@angular/material/stepper';
import { MatSnackBar } from '@angular/material/snack-bar';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { provideNativeDateAdapter } from '@angular/material/core';
import { FestivalService } from '../../../services/festival/festival.service';
import {
  CreateFestivalRequest,
  Festival,
} from '../../../models/festival/festival.model';
import { forkJoin } from 'rxjs';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatIconModule } from '@angular/material/icon';
import { MatDialog } from '@angular/material/dialog';
import {
  ConfirmationDialog,
  ConfirmationDialogData,
} from '../../../shared/confirmation-dialog/confirmation-dialog.component';
import { routes } from '../../../app.routes';
import { Router } from '@angular/router';

@Component({
  selector: 'app-create-festival',
  templateUrl: './create-festival.component.html',
  styleUrls: ['./create-festival.component.scss'],
  standalone: true,
  imports: [
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
  ],
  providers: [provideNativeDateAdapter()],
})
export class CreateFestivalComponent {
  private fb = inject(FormBuilder);
  private http = inject(HttpClient);
  private snackbar = inject(MatSnackBar);
  private festivalService = inject(FestivalService);
  private dialog = inject(MatDialog);
  router: Router = inject(Router);

  @ViewChild('stepper') private stepper: MatStepper | undefined;

  isLinear = true;
  festivalId: string | null = null;
  images: string[] = [
    'https://prismic-assets-cdn.tomorrowland.com/ZqJkiB5LeNNTxfzO_1721400232662_dbb1c34f-229a-46b2-81f3-f498bccf476c.jpg_0_10322487413277135069.jpg',
    'https://venga-store.com/cdn/shop/articles/Tomorrowland-festival-outfits.jpg?v=1721132564&width=1920',
    'https://prismic-assets-cdn.tomorrowland.com/ZwqbE4F3NbkBXXNu_1728686140913_7ef44b3b-297c-4a1c-9d05-a7b04824fc23.jpg_0_12063235528760173649.jpg?crop=3000%2C1600%2C0%2C200&width=600&height=320',
    'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTpolaAYzCeS-n-9KirQxBhe4Co9illDu3ffIlISCpVOIkHkiAiYwYWNQVGZWhvkSZlzzs&usqp=CAU',
  ];

  basicInfoFormGroup = this.fb.group({
    nameCtrl: ['', Validators.required],
    descriptionCtrl: ['', Validators.required],
    startDateCtrl: ['', Validators.required],
    endDateCtrl: ['', Validators.required],
    capacityCtrl: ['', [Validators.required, Validators.min(1)]],
  });

  addressFormGroup = this.fb.group({
    streetCtrl: ['Unknown', Validators.required],
    numberCtrl: ['1', Validators.required],
    apartmentSuiteCtrl: [''],
    cityCtrl: ['Koralovo', Validators.required],
    postalCodeCtrl: ['99999', Validators.required],
    countryISO3Ctrl: ['SRB', Validators.required],
  });

  imagesFormGroup = this.fb.group({
    imageUrlCtrl: [''],
  });

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
          new Date(this.basicInfoFormGroup.get('startDateCtrl')?.value ?? '')
        ),
        endDate: formatDate(
          new Date(this.basicInfoFormGroup.get('endDateCtrl')?.value ?? '')
        ),
        capacity:
          Number(this.basicInfoFormGroup.get('capacityCtrl')?.value) ?? 0,
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
          console.log('Festival created:', response);
          this.festivalId = response.id?.toString() ?? null;
          this.snackbar.open('Basic info saved successfully!', 'Close', {
            duration: 2000,
          });
          this.stepper?.next();
        },
        error: (error) => {
          console.error('Error creating festival:', error);
          this.snackbar.open('Error creating festival', 'Close', {
            duration: 2000,
          });
        },
      });
    }
  }

  addImage() {
    if (this.imagesFormGroup.valid) {
      const imageUrl = this.imagesFormGroup.get('imageUrlCtrl')?.value ?? '';
      this.images.push(imageUrl);
      this.imagesFormGroup.reset();
    }
  }

  removeImage(url: string) {
    const dialogRef = this.dialog.open(ConfirmationDialog, {
      data: {
        title: 'Remove Image',
        message: `Are you sure?`,
        confirmButtonText: 'Remove',
        cancelButtonText: 'Cancel',
      } as ConfirmationDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result?.confirm) {
        this.images = this.images.filter((image) => image !== url);
      }
    });
  }

  uploadFestivalImages() {
    if (this.images.length === 0) {
      this.snackbar.open('Add at least one image');
    }
    if (this.festivalId && this.images.length > 0) {
      const uploadObservables = this.images.map((imageUrl) =>
        this.http.post(
          `http://localhost:4000/festival/${this.festivalId}/image`,
          { imageUrl }
        )
      );

      forkJoin(uploadObservables).subscribe({
        next: () => {
          this.snackbar.open('Festival created successfully!');
          this.stepper?.next();
          this.router.navigate(['organizer/my-festivals']);
        },
        error: (error) => {
          console.error('Error uploading images:', error);
          this.snackbar.open('Error uploading images');
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
