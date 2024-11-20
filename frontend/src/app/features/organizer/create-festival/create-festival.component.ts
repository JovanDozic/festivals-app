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
  ],
  providers: [provideNativeDateAdapter()],
})
export class CreateFestivalComponent {
  private fb = inject(FormBuilder);
  private http = inject(HttpClient);
  private snackbar = inject(MatSnackBar);
  private festivalService = inject(FestivalService);

  @ViewChild('stepper') private stepper: MatStepper | undefined;

  isLinear = true;
  festivalId: string | null = null;
  selectedFile: File | null = null;

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

  imagesFormGroup = this.fb.group({});

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
        startDate: this.basicInfoFormGroup.get('startDateCtrl')?.value ?? '',
        endDate: this.basicInfoFormGroup.get('endDateCtrl')?.value ?? '',
        capacity:
          Number(this.basicInfoFormGroup.get('capacityCtrl')?.value) ?? 0,
        address: {
          street:
            this.addressFormGroup.get('streetCtrl')?.value ?? 'Temp Street',
          number: this.addressFormGroup.get('numberCtrl')?.value ?? '1',
          apartmentSuite:
            this.addressFormGroup.get('apartmentSuiteCtrl')?.value ?? '',
          city: this.addressFormGroup.get('cityCtrl')?.value ?? 'Koralovo',
          postalCode:
            this.addressFormGroup.get('postalCodeCtrl')?.value ?? '99999',
          countryISO3:
            this.addressFormGroup.get('countryISO3Ctrl')?.value ?? 'SRB',
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

  onFileSelected(event: Event) {
    const target = event.target as HTMLInputElement;
    if (target.files && target.files.length > 0) {
      this.selectedFile = target.files[0];
    }
  }

  onDragOver(event: DragEvent) {
    event.preventDefault();
    event.stopPropagation();
  }

  onDrop(event: DragEvent) {
    event.preventDefault();
    event.stopPropagation();

    if (event.dataTransfer && event.dataTransfer.files.length > 0) {
      this.selectedFile = event.dataTransfer.files[0];
    }
  }

  uploadFestivalImage() {
    if (this.selectedFile && this.festivalId) {
      const formData = new FormData();
      formData.append('image', this.selectedFile);

      this.http
        .post(
          `http://localhost:4000/festival/${this.festivalId}/image`,
          formData
        )
        .subscribe({
          next: () => {
            this.snackbar.open('Image uploaded successfully!', 'Close', {
              duration: 2000,
            });
            setTimeout(() => {
              window.location.reload();
            }, 1000);
          },
          error: (error) => {
            console.error('Error uploading image:', error);
            this.snackbar.open('Error uploading image', 'Close', {
              duration: 2000,
            });
          },
        });
    }
  }
}
