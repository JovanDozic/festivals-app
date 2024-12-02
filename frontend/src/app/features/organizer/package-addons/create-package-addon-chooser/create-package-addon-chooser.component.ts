import { Component, inject, OnInit, ViewChild } from '@angular/core';
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
import { MatRadioModule } from '@angular/material/radio';

interface Category {
  value: string;
  viewValue: string;
  description: string;
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
    MatRadioModule,
  ],
  templateUrl: './create-package-addon-chooser.component.html',
  styleUrls: [
    './create-package-addon-chooser.component.scss',
    '../../../../app.component.scss',
  ],
})
export class CreatePackageAddonChooserComponent implements OnInit {
  private dialogRef = inject(MatDialogRef<CreatePackageAddonChooserComponent>);
  private data: { festivalId: number } = inject(MAT_DIALOG_DATA);

  categories: Category[] = [
    {
      value: 'GENERAL',
      viewValue: 'General',
      description: 'Enchase The Festival Experience',
    },
    {
      value: 'TRANSPORT',
      viewValue: 'Travel',
      description: 'Help Attendees arrive to The Festival Grounds',
    },
    {
      value: 'CAMP',
      viewValue: 'Camp',
      description: 'Provide Attendees with a place to stay',
    },
  ];

  selectedCategory: any | null = this.categories[2].value;

  ngOnInit(): void {}

  closeDialog() {
    this.dialogRef.close(false);
  }

  choose() {
    this.dialogRef.close(this.selectedCategory);
  }
}
