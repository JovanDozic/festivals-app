import { Component, inject, ViewChild } from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialogModule,
  MatDialogRef,
} from '@angular/material/dialog';
import {
  FormArray,
  FormBuilder,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  ValidationErrors,
  Validators,
} from '@angular/forms';
import {
  CreateItemPriceRequest,
  CreateItemRequest,
  VariablePrice,
} from '../../../../models/festival/festival.model';
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
import { SnackbarService } from '../../../../shared/snackbar/snackbar.service';
import { MatStepper, MatStepperModule } from '@angular/material/stepper';
import { ItemService } from '../../../../services/festival/item.service';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { forkJoin } from 'rxjs';
import { MatSelectModule } from '@angular/material/select';

interface Category {
  value: string;
  viewValue: string;
}

@Component({
  selector: 'app-create-package-addon-chooser',
  imports: [
    FormsModule,
    CommonModule,
    ReactiveFormsModule,
    MatInputModule,
    MatFormFieldModule,
    MatButtonModule,
    MatCardModule,
    MatDatepickerModule,
    MatGridListModule,
    MatIconModule,
    MatTabsModule,
    MatStepperModule,
    MatSlideToggleModule,
    MatDialogModule,
    MatSelectModule,
  ],
  templateUrl: './create-package-addon-chooser.component.html',
  styleUrls: [
    './create-package-addon-chooser.component.scss',
    '../../../../app.component.scss',
  ],
})
export class CreatePackageAddonChooserComponent {
  private dialogRef = inject(MatDialogRef<CreatePackageAddonChooserComponent>);
  private data: { festivalId: number } = inject(MAT_DIALOG_DATA);

  selectedCategory: Category | null = null;
  categories: Category[] = [
    { value: 'GENERAL', viewValue: 'General' },
    { value: 'TRANSPORT', viewValue: 'Transport' },
    { value: 'CAMP', viewValue: 'Camp' },
  ];

  closeDialog() {
    this.dialogRef.close(false);
  }

  choose() {
    this.dialogRef.close(this.selectedCategory);
  }
}
