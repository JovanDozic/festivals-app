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

@Component({
  selector: 'app-create-package-addon',
  imports: [
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
  ],
  templateUrl: './create-package-addon.component.html',
  styleUrls: [
    './create-package-addon.component.scss',
    '../../../../app.component.scss',
  ],
  providers: [provideNativeDateAdapter()],
})
export class CreatePackageAddonComponent {
  private fb = inject(FormBuilder);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<CreatePackageAddonComponent>);
  private data: { festivalId: number; category: string } =
    inject(MAT_DIALOG_DATA);
  private itemService = inject(ItemService);

  @ViewChild('stepper') private stepper: MatStepper | undefined;

  category: string = '';

  infoFormGroup: FormGroup;
  fixedPriceFormGroup: FormGroup;

  itemId: number | null = null;
  isFixedPrice: boolean = true;

  constructor() {
    this.category = this.data?.category;

    this.infoFormGroup = this.fb.group({
      nameCtrl: ['', Validators.required],
      descriptionCtrl: ['', Validators.required],
      availableNumberCtrl: ['', [Validators.required, Validators.min(1)]],
    });

    this.fixedPriceFormGroup = this.fb.group({
      fixedPriceCtrl: ['', [Validators.required, Validators.min(0)]],
    });
  }

  closeDialog() {
    this.dialogRef.close(false);
  }

  done() {
    if (this.itemId) {
      this.createFixedPrice();
    }
  }

  createPackageAddon() {
    if (this.infoFormGroup.valid && this.data.festivalId) {
      const request: CreateItemRequest = {
        name: this.infoFormGroup.get('nameCtrl')?.value,
        description: this.infoFormGroup.get('descriptionCtrl')?.value,
        availableNumber: this.infoFormGroup.get('availableNumberCtrl')?.value,
        type: 'PACKAGE_ADDON',
      };

      this.itemService
        .createPackageAddon(this.data.festivalId, request, this.category)
        .subscribe({
          next: (response) => {
            this.snackbarService.show('Package Addon created');
            this.itemId = response;
            this.stepper?.next();
          },
          error: (error) => {
            console.log('Error creating Package Addon: ', error);
            this.snackbarService.show('Error creating Package Addon');
          },
        });
    }
  }

  createFixedPrice() {
    if (this.fixedPriceFormGroup.valid && this.itemId && this.data.festivalId) {
      const request: CreateItemPriceRequest = {
        itemId: this.itemId,
        price: this.fixedPriceFormGroup.get('fixedPriceCtrl')?.value,
        isFixed: true,
      };

      this.itemService
        .createItemPrice(this.data.festivalId, request)
        .subscribe({
          next: (response) => {
            this.snackbarService.show('Fixed Price created');
            this.dialogRef.close(true);
          },
          error: (error) => {
            console.log('Error creating fixed price: ', error);
            this.snackbarService.show('Error creating Fixed Price');
            this.dialogRef.close(false);
          },
        });
    }
  }
}
