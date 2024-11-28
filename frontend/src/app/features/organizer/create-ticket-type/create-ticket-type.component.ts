import { Component, inject, ViewChild } from '@angular/core';
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
  selector: 'app-create-ticket-type',
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
  templateUrl: './create-ticket-type.component.html',
  styleUrls: [
    './create-ticket-type.component.scss',
    '../../../app.component.scss',
  ],
  providers: [provideNativeDateAdapter()],
})
export class CreateTicketTypeComponent {
  private fb = inject(FormBuilder);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<CreateTicketTypeComponent>);
  private data: { festivalId: number } = inject(MAT_DIALOG_DATA);
  private itemService = inject(ItemService);

  @ViewChild('stepper') private stepper: MatStepper | undefined;

  infoFormGroup: FormGroup;
  fixedPriceFormGroup: FormGroup;
  variablePricesFormGroup: FormGroup;

  ticketTypeId: number | null = null;
  isFixedPrice: boolean = false;
  variablePrices: VariablePrice[] = [];

  constructor() {
    this.infoFormGroup = this.fb.group({
      nameCtrl: ['', Validators.required],
      descriptionCtrl: ['', Validators.required],
      availableNumberCtrl: ['', [Validators.required, Validators.min(1)]],
    });

    this.fixedPriceFormGroup = this.fb.group({
      fixedPriceCtrl: ['', [Validators.required, Validators.min(0)]],
    });

    this.variablePricesFormGroup = this.fb.group({
      variablePricesFormArray: this.fb.array([this.createVariablePriceGroup()]),
    });
  }

  get variablePricesFormArray(): FormArray {
    return this.variablePricesFormGroup.get(
      'variablePricesFormArray'
    ) as FormArray;
  }

  private createVariablePriceGroup(): FormGroup {
    return this.fb.group({
      priceCtrl: ['', [Validators.required, Validators.min(0)]],
      dateFromCtrl: ['', Validators.required],
      dateToCtrl: ['', Validators.required],
    });
  }

  toggleIsFixed() {
    this.isFixedPrice = !this.isFixedPrice;
  }

  closeDialog() {
    this.dialogRef.close(false);
  }

  createTicketType() {
    if (this.infoFormGroup.valid && this.data.festivalId) {
      const request: CreateItemRequest = {
        name: this.infoFormGroup.get('nameCtrl')?.value,
        description: this.infoFormGroup.get('descriptionCtrl')?.value,
        availableNumber: this.infoFormGroup.get('availableNumberCtrl')?.value,
        type: 'TICKET_TYPE',
      };

      this.itemService.createItem(this.data.festivalId, request).subscribe({
        next: (response) => {
          console.log('Ticket type created: ', response);
          this.snackbarService.show('Ticket Type created');
          this.ticketTypeId = response;
          this.stepper?.next();
        },
        error: (error) => {
          console.log('Error creating ticket type: ', error);
          this.snackbarService.show('Error creating Ticket Type');
        },
      });
    }
  }

  addVariablePrice() {
    const lastGroup = this.variablePricesFormArray.at(
      this.variablePricesFormArray.length - 1
    ) as FormGroup;

    if (lastGroup.valid) {
      this.variablePricesFormArray.push(this.createVariablePriceGroup());
    } else {
      this.snackbarService.show(
        'Please fill out the last variable price before adding a new one.'
      );
    }
  }

  removeVariablePrice(index: number) {
    if (this.variablePricesFormArray.length > 1) {
      this.variablePricesFormArray.removeAt(index);
    } else {
      this.snackbarService.show(
        'At least one variable price entry is required.'
      );
    }
  }

  done() {
    if (this.ticketTypeId) {
      if (this.isFixedPrice) {
        this.createFixedPrice();
      } else {
        this.createNotFixedPrices();
      }
    }
  }

  createFixedPrice() {
    if (
      this.fixedPriceFormGroup.valid &&
      this.ticketTypeId &&
      this.data.festivalId
    ) {
      const request: CreateItemPriceRequest = {
        itemId: this.ticketTypeId,
        price: this.fixedPriceFormGroup.get('fixedPriceCtrl')?.value,
        isFixed: true,
      };

      this.itemService
        .createItemPrice(this.data.festivalId, request)
        .subscribe({
          next: (response) => {
            console.log('Fixed price created: ', response);
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

  createNotFixedPrices() {
    if (
      this.variablePricesFormArray.valid &&
      this.ticketTypeId &&
      this.data.festivalId
    ) {
      const variablePrices: VariablePrice[] =
        this.variablePricesFormArray.value.map((vp: any) => ({
          price: vp.priceCtrl,
          dateFrom: vp.dateFromCtrl,
          dateTo: vp.dateToCtrl,
        }));

      const requests: CreateItemPriceRequest[] = variablePrices.map((vp) => ({
        itemId: this.ticketTypeId!,
        price: vp.price,
        isFixed: false,
        dateFrom: this.formatDate(vp.dateFrom),
        dateTo: this.formatDate(vp.dateTo),
      }));

      forkJoin(
        requests.map((req) =>
          this.itemService.createItemPrice(this.data.festivalId, req)
        )
      ).subscribe({
        next: (responses) => {
          console.log('Variable prices created: ', responses);
          this.snackbarService.show('Variable Prices created');
          this.dialogRef.close(true);
        },
        error: (error) => {
          console.log('Error creating variable prices: ', error);
          this.snackbarService.show('Error creating Variable Prices');
          this.dialogRef.close(false);
        },
      });
    } else {
      this.snackbarService.show('Please fill out all variable price fields.');
    }
  }

  private formatDate(date: Date): string {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    return `${year}-${month}-${day}`;
  }
}
