import { Component, inject } from '@angular/core';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatIconModule } from '@angular/material/icon';
import { MatTabsModule } from '@angular/material/tabs';
import { MatStepperModule } from '@angular/material/stepper';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatSelectModule } from '@angular/material/select';
import { MatRadioModule } from '@angular/material/radio';

interface StoreType {
  value: string;
  viewValue: string;
  description: string;
}

@Component({
  selector: 'app-store-chooser',
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
  templateUrl: './store-chooser.component.html',
  styleUrls: [
    './store-chooser.component.scss',
    '../../../../app.component.scss',
  ],
})
export class StoreChooserComponent {
  private dialogRef = inject(MatDialogRef<StoreChooserComponent>);

  categories: StoreType[] = [
    {
      value: 'TICKET',
      viewValue: 'Festival Ticket',
      description: 'Ticket to access the Festival Grounds',
    },
    {
      value: 'PACKAGE',
      viewValue: 'Festival Package',
      description:
        'Choose a custom experience using Package with Travel, Camp and other options with Ticket included',
    },
  ];

  selectedCategory: StoreType | null = null;

  closeDialog() {
    this.dialogRef.close(false);
  }

  choose() {
    this.dialogRef.close(this.selectedCategory);
  }
}
