import { Component, inject } from '@angular/core';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { FormsModule } from '@angular/forms';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
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
    MatInputModule,
    MatButtonModule,
    MatCardModule,
    MatIconModule,
    MatDialogModule,
    MatRadioModule,
  ],
  templateUrl: './create-package-addon-chooser.component.html',
  styleUrls: [
    './create-package-addon-chooser.component.scss',
    '../../../../app.component.scss',
  ],
})
export class CreatePackageAddonChooserComponent {
  private dialogRef = inject(MatDialogRef<CreatePackageAddonChooserComponent>);

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

  selectedCategory: Category | null = null;

  closeDialog() {
    this.dialogRef.close(false);
  }

  choose() {
    this.dialogRef.close(this.selectedCategory);
  }
}
