import { Component, inject, OnInit } from '@angular/core';
import {
  MAT_DIALOG_DATA,
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
import { provideNativeDateAdapter } from '@angular/material/core';
import { MatTabsModule } from '@angular/material/tabs';
import { SnackbarService } from '../../../shared/snackbar/snackbar.service';

@Component({
  selector: 'app-edit-festival',
  templateUrl: './edit-festival.component.html',
  styleUrls: ['./edit-festival.component.scss', '../../../app.component.scss'],
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
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions,
    MatTabsModule,
  ],
  providers: [provideNativeDateAdapter()],
})
export class EditFestivalComponent implements OnInit {
  private fb = inject(FormBuilder);
  private festivalService = inject(FestivalService);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<EditFestivalComponent>);
  private data: Festival = inject(MAT_DIALOG_DATA);

  isLinear = true;

  basicInfoFormGroup: FormGroup;
  addressFormGroup: FormGroup;

  constructor() {
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
      countryISO3Ctrl: ['', Validators.required],
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
    }
  }

  saveChanges() {
    if (this.basicInfoFormGroup.invalid || this.addressFormGroup.invalid) {
      this.snackbarService.show('Please complete all required fields.');
      return;
    }

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
        this.snackbarService.show('Festival updated successfully!');
        this.dialogRef.close(true);
      },
      error: (error) => {
        console.error('Error updating festival:', error);
        this.snackbarService.show('Error updating festival. Please try again.');
      },
    });
  }

  closeDialog() {
    this.dialogRef.close(false);
  }

  private formatDate(date: Date): string {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    return `${year}-${month}-${day}`;
  }
}
