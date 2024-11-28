import { Component, inject, OnInit, ViewChild } from '@angular/core';
import {
  MAT_DIALOG_DATA,
  MatDialogContent,
  MatDialogModule,
  MatDialogRef,
  MatDialogTitle,
} from '@angular/material/dialog';
import {
  FormArray,
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { FestivalService } from '../../../services/festival/festival.service';
import {
  CreateItemPriceRequest,
  CreateItemRequest,
  Item,
  PriceListItem,
  VariablePrice,
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
import { UserService } from '../../../services/user/user.service';
import { SnackbarService } from '../../../shared/snackbar/snackbar.service';
import { MatStepper, MatStepperModule } from '@angular/material/stepper';
import { ItemService } from '../../../services/festival/item.service';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { forkJoin } from 'rxjs';

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
    MatStepperModule,
    MatSlideToggleModule,
    MatDialogModule,
  ],
  templateUrl: './view-edit-ticket-type.component.html',
  styleUrls: [
    './view-edit-ticket-type.component.scss',
    '../../../app.component.scss',
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

  isEditing: boolean = true;

  infoFormGroup: FormGroup;
  fixedPriceFormGroup: FormGroup;
  variablePricesFormGroup: FormGroup;

  isFixedPrice: boolean = true;

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
      'variablePricesFormArray'
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
    const festivalId = this.data.festivalId;
    const ticketTypeId = this.data.itemId;
    if (festivalId && ticketTypeId) {
      this.itemService.getTicketType(festivalId, ticketTypeId).subscribe({
        next: (ticketType) => {
          this.ticketType = ticketType;
          console.log('Ticket type: ', ticketType);
          this.loadForms();
          this.toggleIsEditing();
        },
        error: (error) => {
          console.log('Error fetching ticket type: ', error);
          this.snackbarService.show('Error getting ticket type');
        },
      });
    } else {
      console.log(
        'No festival id or ticket type id:',
        festivalId,
        ticketTypeId
      );
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

      console.log('Request: ', request);

      this.itemService.updateItem(this.data.festivalId, request).subscribe({
        next: () => {
          this.snackbarService.show('Ticket type updated');
          this.dialogRef.close(true);
        },
        error: (error) => {
          console.log('Error updating ticket type: ', error);
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
}
