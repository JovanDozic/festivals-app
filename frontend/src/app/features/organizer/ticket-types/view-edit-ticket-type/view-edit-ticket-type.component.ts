import { Component, inject, OnInit } from '@angular/core';
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
  Item,
  PriceListItem,
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
import { SnackbarService } from '../../../../services/snackbar/snackbar.service';
import { ItemService } from '../../../../services/festival/item.service';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';

@Component({
  selector: 'app-view-edit-ticket-type',
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
    MatSlideToggleModule,
    MatDialogModule,
  ],
  templateUrl: './view-edit-ticket-type.component.html',
  styleUrls: [
    './view-edit-ticket-type.component.scss',
    '../../../../app.component.scss',
  ],
  providers: [provideNativeDateAdapter()],
})
export class ViewEditTicketTypeComponent implements OnInit {
  private fb = inject(FormBuilder);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<ViewEditTicketTypeComponent>);
  private itemService = inject(ItemService);
  private data: {
    festivalId: number;
    itemId: number;
  } = inject(MAT_DIALOG_DATA);

  isEditing = true;

  infoFormGroup: FormGroup;
  fixedPriceFormGroup: FormGroup;
  variablePricesFormGroup: FormGroup;

  isFixedPrice = true;

  ticketType: Item | null = null;

  ngOnInit(): void {
    this.loadTicketType();
  }

  closeDialog() {
    this.dialogRef.close(false);
  }

  constructor() {
    this.infoFormGroup = this.fb.group({
      nameCtrl: ['', Validators.required],
      descriptionCtrl: ['', Validators.required],
      availableNumberCtrl: [
        0,
        [
          Validators.required,
          (control: any) => {
            const value = control.value;
            const availableNumber = this.ticketType?.availableNumber ?? 0;
            const remainingNumber = this.ticketType?.remainingNumber ?? 0;
            return value < availableNumber - remainingNumber
              ? { lessThanSold: true }
              : null;
          },
        ],
      ],
    });

    this.fixedPriceFormGroup = this.fb.group({
      fixedPriceCtrl: ['', [Validators.required, Validators.min(0)]],
    });

    this.variablePricesFormGroup = this.fb.group(
      {
        variablePricesFormArray: this.fb.array([
          this.createVariablePriceGroup(),
        ]),
      },
      { validators: this.validateVariablePrices.bind(this) },
    );

    this.infoFormGroup.disable();
    this.fixedPriceFormGroup.disable();
    this.variablePricesFormGroup.disable();
    this.variablePricesFormArray.disable();
    this.variablePricesFormArray.controls.forEach((control) => {
      control.disable();
    });
  }

  get variablePricesFormArray(): FormArray {
    return this.variablePricesFormGroup.get(
      'variablePricesFormArray',
    ) as FormArray;
  }

  private createVariablePriceGroup(): FormGroup {
    return this.fb.group({
      idCtrl: [0],
      isFixed: [false],
      priceCtrl: ['', [Validators.required, Validators.min(0)]],
      dateFromCtrl: [null, Validators.required],
      dateToCtrl: [null, Validators.required],
    });
  }

  toggleIsEditing() {
    this.isEditing = !this.isEditing;
    if (!this.isEditing) {
      this.infoFormGroup.disable();
      this.fixedPriceFormGroup.disable();
      this.variablePricesFormGroup.disable();
      this.variablePricesFormArray.disable();
      this.variablePricesFormArray.controls.forEach((control) => {
        control.disable();
      });
    } else {
      this.infoFormGroup.enable();
      this.fixedPriceFormGroup.enable();
      this.variablePricesFormGroup.enable();
      this.variablePricesFormArray.enable();
      this.variablePricesFormArray.controls.forEach((control) => {
        control.enable();
      });
    }
  }

  loadForms() {
    if (this.ticketType) {
      this.infoFormGroup.setValue({
        nameCtrl: this.ticketType.name,
        descriptionCtrl: this.ticketType.description,
        availableNumberCtrl: this.ticketType.availableNumber,
      });

      this.isFixedPrice = this.ticketType.priceListItems[0].isFixed;

      if (this.isFixedPrice) {
        this.fixedPriceFormGroup.setValue({
          fixedPriceCtrl: this.ticketType.priceListItems[0].price,
        });
      } else {
        this.variablePricesFormArray.clear();
        this.ticketType.priceListItems.forEach((priceListItem) => {
          const variablePriceGroup = this.createVariablePriceGroup();
          variablePriceGroup.setValue({
            idCtrl: priceListItem.id,
            isFixed: priceListItem.isFixed,
            priceCtrl: priceListItem.price,
            dateFromCtrl: new Date(priceListItem.dateFrom ?? ''),
            dateToCtrl: new Date(priceListItem.dateTo ?? ''),
          });
          this.variablePricesFormArray.push(variablePriceGroup);
        });
      }
    }
  }

  loadTicketType() {
    if (this.data.festivalId && this.data.itemId) {
      this.itemService
        .getTicketType(this.data.festivalId, this.data.itemId)
        .subscribe({
          next: (ticketType) => {
            this.ticketType = ticketType;
            this.loadForms();
            this.toggleIsEditing();
          },
          error: (error) => {
            console.log(error);
            this.snackbarService.show('Error getting ticket type');
          },
        });
    }
  }

  saveChanges() {
    if (this.infoFormGroup.valid) {
      const request: Item = {
        id: this.data.itemId,
        name: this.infoFormGroup.get('nameCtrl')?.value,
        description: this.infoFormGroup.get('descriptionCtrl')?.value,
        availableNumber: this.infoFormGroup.get('availableNumberCtrl')?.value,
        type: 'TICKET_TYPE',
        remainingNumber: this.ticketType?.remainingNumber ?? 0,
        priceListItems: [],
      };

      if (this.isFixedPrice) {
        const fixedPriceRequest: PriceListItem = {
          id: this.ticketType?.priceListItems[0].id ?? 0,
          isFixed: true,
          price: this.fixedPriceFormGroup.get('fixedPriceCtrl')?.value,
          dateFrom: this.ticketType?.priceListItems[0].dateFrom,
          dateTo: this.ticketType?.priceListItems[0].dateTo,
        };
        request.priceListItems.push(fixedPriceRequest);
      } else {
        this.variablePricesFormArray.controls.forEach((control) => {
          const variablePriceRequest: PriceListItem = {
            id: control.get('idCtrl')?.value,
            isFixed: control.get('isFixed')?.value,
            price: control.get('priceCtrl')?.value,
            dateFrom: this.formatDate(control.get('dateFromCtrl')?.value),
            dateTo: this.formatDate(control.get('dateToCtrl')?.value),
          };
          request.priceListItems.push(variablePriceRequest);
        });
      }

      this.itemService.updateItem(this.data.festivalId, request).subscribe({
        next: () => {
          this.snackbarService.show('Ticket type updated');
          this.dialogRef.close(true);
        },
        error: (error) => {
          console.log(error);
          this.snackbarService.show('Error updating ticket type');
        },
      });
    }
  }

  private formatDate(date: Date): string {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    return `${year}-${month}-${day}`;
  }

  private validateVariablePrices(
    formGroup: FormGroup,
  ): ValidationErrors | null {
    const variablePricesFormArray = formGroup.get(
      'variablePricesFormArray',
    ) as FormArray;

    const dateRanges = variablePricesFormArray.controls.map((control) => {
      const dateFrom: Date = control.get('dateFromCtrl')?.value;
      const dateTo: Date = control.get('dateToCtrl')?.value;
      return {
        dateFrom,
        dateTo,
        control,
      };
    });

    let hasErrors = false;

    // Clear previous errors but preserve 'required' errors
    variablePricesFormArray.controls.forEach((control) => {
      const dateFromCtrl = control.get('dateFromCtrl');
      const dateToCtrl = control.get('dateToCtrl');

      if (dateFromCtrl?.errors && !dateFromCtrl.errors['required']) {
        delete dateFromCtrl.errors['dateOrder'];
        delete dateFromCtrl.errors['overlap'];
        delete dateFromCtrl.errors['gap'];
        if (Object.keys(dateFromCtrl.errors).length === 0) {
          dateFromCtrl.setErrors(null);
        }
      }

      if (dateToCtrl?.errors && !dateToCtrl.errors['required']) {
        delete dateToCtrl.errors['dateOrder'];
        delete dateToCtrl.errors['overlap'];
        delete dateToCtrl.errors['gap'];
        if (Object.keys(dateToCtrl.errors).length === 0) {
          dateToCtrl.setErrors(null);
        }
      }
    });

    // Validate each date range
    dateRanges.forEach((currentRange) => {
      const dateFromCtrl = currentRange.control.get('dateFromCtrl');
      const dateToCtrl = currentRange.control.get('dateToCtrl');
      const currentDateFrom = currentRange.dateFrom;
      const currentDateTo = currentRange.dateTo;

      // Check for required dates
      if (!currentDateFrom) {
        dateFromCtrl?.setErrors({ ...dateFromCtrl.errors, required: true });
        hasErrors = true;
      }
      if (!currentDateTo) {
        dateToCtrl?.setErrors({ ...dateToCtrl.errors, required: true });
        hasErrors = true;
      }

      // Proceed only if both dates are present
      if (currentDateFrom && currentDateTo) {
        // Check that dateFrom <= dateTo
        if (currentDateFrom > currentDateTo) {
          dateToCtrl?.setErrors({ ...dateToCtrl.errors, dateOrder: true });
          hasErrors = true;
        }
      }
    });

    // Proceed only if all dateFrom and dateTo are present
    if (hasErrors) {
      return { invalidDateRanges: true };
    }

    // Sort date ranges by dateFrom
    dateRanges.sort((a, b) => a.dateFrom.getTime() - b.dateFrom.getTime());

    // Check for overlaps and gaps
    for (let i = 0; i < dateRanges.length; i++) {
      const currentRange = dateRanges[i];
      const currentDateFrom = currentRange.dateFrom;
      const dateFromCtrl = currentRange.control.get('dateFromCtrl');

      if (i > 0) {
        const previousRange = dateRanges[i - 1];
        const previousDateTo = previousRange.dateTo;
        const previousDateToCtrl = previousRange.control.get('dateToCtrl');

        // Check for overlaps
        if (currentDateFrom.getTime() <= previousDateTo.getTime()) {
          // Overlap detected
          dateFromCtrl?.setErrors({ ...dateFromCtrl.errors, overlap: true });
          previousDateToCtrl?.setErrors({
            ...previousDateToCtrl.errors,
            overlap: true,
          });
          hasErrors = true;
        } else {
          // Check for gaps
          const expectedDateFrom = new Date(previousDateTo);
          expectedDateFrom.setDate(expectedDateFrom.getDate() + 1);
          if (currentDateFrom.getTime() !== expectedDateFrom.getTime()) {
            // Gap detected
            dateFromCtrl?.setErrors({ ...dateFromCtrl.errors, gap: true });
            previousDateToCtrl?.setErrors({
              ...previousDateToCtrl.errors,
              gap: true,
            });
            hasErrors = true;
          }
        }
      }
    }

    return hasErrors ? { invalidDateRanges: true } : null;
  }
}
