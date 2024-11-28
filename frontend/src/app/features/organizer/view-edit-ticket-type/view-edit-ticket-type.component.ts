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
  addVariablePrice() {
    throw new Error('Method not implemented.');
  }
  saveChanges() {
    throw new Error('Method not implemented.');
  }
  private fb = inject(FormBuilder);
  private snackbarService = inject(SnackbarService);
  private dialogRef = inject(MatDialogRef<ViewEditTicketTypeComponent>);
  private data: { festivalId: number; itemId: number } =
    inject(MAT_DIALOG_DATA);
  private itemService = inject(ItemService);

  isEditing: boolean = false;

  infoFormGroup: FormGroup;
  fixedPriceFormGroup: FormGroup;
  variablePricesFormGroup: FormGroup;

  isFixedPrice: boolean = true;

  ticketType: Item | null = null;

  ngOnInit(): void {
    this.loadTicketType();
  }

  closeDialog() {
    this.dialogRef.close();
  }

  constructor() {
    // todo: fill in information for selected ticket type
    // todo: fetch ticket type information from server
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

  loadForms() {
    if (this.ticketType) {
      console.log('Ticket type: ', this.ticketType);
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
            priceCtrl: priceListItem.price,
            dateFromCtrl: new Date(priceListItem.dateFrom),
            dateToCtrl: new Date(priceListItem.dateTo),
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
}
