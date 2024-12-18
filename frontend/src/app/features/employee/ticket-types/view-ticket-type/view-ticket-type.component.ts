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
  templateUrl: './view-ticket-type.component.html',
  styleUrls: [
    './view-ticket-type.component.scss',
    '../../../../app.component.scss',
  ],
  providers: [provideNativeDateAdapter()],
})
export class ViewTicketTypeComponent implements OnInit {
  private fb = inject(FormBuilder);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<ViewTicketTypeComponent>);
  private itemService = inject(ItemService);
  private data: {
    festivalId: number;
    itemId: number;
  } = inject(MAT_DIALOG_DATA);

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
      availableNumberCtrl: [0, Validators.required],
    });

    this.fixedPriceFormGroup = this.fb.group({
      fixedPriceCtrl: ['', [Validators.required, Validators.min(0)]],
    });

    this.variablePricesFormGroup = this.fb.group({
      variablePricesFormArray: this.fb.array([this.createVariablePriceGroup()]),
    });

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

      this.infoFormGroup.disable();
      this.fixedPriceFormGroup.disable();
      this.variablePricesFormGroup.disable();
      this.variablePricesFormArray.disable();
      this.variablePricesFormArray.controls.forEach((control) => {
        control.disable();
      });
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
          },
          error: (error) => {
            console.log(error);
            this.snackbarService.show('Error getting ticket type');
          },
        });
    }
  }
}
